package immutable

import (
	"context"
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/immutable/imx-core-sdk-golang/api"
	"github.com/immutable/imx-core-sdk-golang/internal/convert"
)

type TokenDeposit interface {
	Deposit(ctx context.Context, c *Client, l1signer L1Signer, overrides *bind.TransactOpts) (*types.Transaction, error)
}

// ETHDeposit implements TokenDeposit. Used to deposit Eth Tokens.
type ETHDeposit struct {
	Amount string
}

// ERC20Deposit implements TokenDeposit. Used to deposit ERC20 Tokens.
type ERC20Deposit struct {
	Amount       string
	TokenAddress string
}

// ERC721Deposit implements TokenDeposit. Used to deposit ERC721 Tokens.
type ERC721Deposit struct {
	TokenID      string
	TokenAddress string
}

// NewETHDeposit instantiates a new ETHDeposit object with the given amount.
func NewETHDeposit(amount string) *ETHDeposit {
	this := ETHDeposit{}
	this.Amount = amount
	return &this
}

// NewERC20Deposit instantiates a new ERC20Deposit object with given amount and tokenAddress.
func NewERC20Deposit(amount, tokenAddress string) *ERC20Deposit {
	this := ERC20Deposit{}
	this.Amount = amount
	this.TokenAddress = strings.ToLower(tokenAddress)
	return &this
}

// NewERC721Deposit instantiates a new ERC721Deposit object with given tokenID and tokenAddress.
func NewERC721Deposit(tokenID, tokenAddress string) *ERC721Deposit {
	this := ERC721Deposit{}
	this.TokenID = tokenID
	this.TokenAddress = tokenAddress
	return &this
}

func (c *Client) getSignableDeposit(
	ctx context.Context,
	request *api.GetSignableDepositRequest) (*api.GetSignableDepositResponse, error) {
	signableDepositResponse, httpResp, err := c.GetSignableDeposit(ctx).GetSignableDepositRequest(*request).Execute()
	if err != nil {
		return nil, fmt.Errorf("error when calling `Deposits.GetSignableDeposit`: %v, HTTP response body: %v", err, httpResp.Body)
	}
	return signableDepositResponse, nil
}

// Deposit performs the deposit workflow on the ETHDeposit.
func (d *ETHDeposit) Deposit(ctx context.Context, c *Client, l1signer L1Signer, overrides *bind.TransactOpts) (*types.Transaction, error) {
	amount, err := convert.ToDenomination(d.Amount, convert.EtherDecimals)
	if err != nil {
		return nil, fmt.Errorf("error when parsing deposit amount: %v", err)
	}
	signableDepositRequest := newSignableDepositRequestForEth(amount.String(), l1signer.GetAddress())
	signableDeposit, err := c.getSignableDeposit(ctx, signableDepositRequest)
	if err != nil {
		return nil, err
	}

	assetType, err := c.encodeETHAsset(ctx)
	if err != nil {
		return nil, err
	}

	starkKeyHex := signableDeposit.StarkKey
	starkKey, err := hexutil.DecodeBig(starkKeyHex)
	if err != nil {
		return nil, fmt.Errorf("error converting StarkKey to bigint: %s", starkKeyHex)
	}
	isRegistered, _ := c.RegistrationContract.IsRegistered(&bind.CallOpts{Context: ctx}, starkKey)
	// Note: if we reach here, it means we are registered off-chain.
	// Above call will return an error user is not registered but this is for on-chain
	// we should swallow this error to allow the register and deposit flow to execute.

	if isRegistered {
		return c.depositEth(ctx, l1signer, starkKey, big.NewInt(int64(signableDeposit.VaultId)), assetType, amount, overrides)
	}
	return c.registerAndDepositEth(ctx, l1signer, starkKeyHex, starkKey, big.NewInt(int64(signableDeposit.VaultId)), assetType, amount, overrides)
}

func newSignableDepositRequestForEth(amount, user string) *api.GetSignableDepositRequest {
	return &api.GetSignableDepositRequest{
		Amount: amount,
		Token:  SignableETHToken(),
		User:   user,
	}
}

func (c *Client) registerAndDepositEth(
	ctx context.Context,
	l1signer L1Signer,
	starkKeyHex string,
	starkKey *big.Int,
	vaultID *big.Int,
	assetType *big.Int,
	amount *big.Int,
	overrides *bind.TransactOpts,
) (*types.Transaction, error) {
	etherKey := l1signer.GetAddress()
	signableRegistration, err := c.GetSignableRegistrationOnchain(ctx, etherKey, starkKeyHex)
	if err != nil {
		return nil, err
	}

	opts := c.buildTransactOpts(ctx, l1signer, overrides)

	operatorSignature, err := convert.HexToByteArray(signableRegistration.OperatorSignature)
	if err != nil {
		return nil, err
	}
	opts.Value = amount
	tnx, err := c.CoreContract.RegisterAndDepositEth(
		opts,
		common.HexToAddress(etherKey),
		starkKey,
		operatorSignature,
		assetType,
		vaultID,
	)
	if err != nil {
		return nil, err
	}
	return tnx, nil
}

func (c *Client) depositEth(
	ctx context.Context,
	l1signer L1Signer,
	starkPublicKey *big.Int,
	vaultID *big.Int,
	assetType *big.Int,
	amount *big.Int,
	overrides *bind.TransactOpts,
) (*types.Transaction, error) {
	opts := c.buildTransactOpts(ctx, l1signer, overrides)
	opts.Value = amount
	tnx, err := c.CoreContract.Deposit(opts, starkPublicKey, assetType, vaultID)
	if err != nil {
		return nil, err
	}
	return tnx, nil
}

// Deposit performs the deposit workflow on the ERC721Deposit.
func (d *ERC721Deposit) Deposit(ctx context.Context, c *Client, l1signer L1Signer, overrides *bind.TransactOpts) (*types.Transaction, error) {
	// Approve whether an amount of token from an account can be spent by a third-party account
	opts := c.buildTransactOpts(ctx, l1signer, overrides)
	tokenID, ok := new(big.Int).SetString(d.TokenID, 10)
	if !ok {
		return nil, fmt.Errorf("error converting tokenID to bigint: %v", d.TokenID)
	}
	ierc721Contract, err := c.newIERC721Contract(ctx, d.TokenAddress)
	if err != nil {
		return nil, err
	}
	_, err = ierc721Contract.Approve(opts, common.HexToAddress(c.Environment.CoreContractAddress), tokenID)
	if err != nil {
		return nil, err
	}

	// Get signable deposit details
	signableDepositRequest := newSignableDepositRequestForERC721("1", d.TokenID, d.TokenAddress, l1signer.GetAddress())
	signableDeposit, err := c.getSignableDeposit(ctx, signableDepositRequest)
	if err != nil {
		return nil, err
	}

	// Perform encoding on asset details to get an assetType (required for stark contract request)
	assetType, err := c.encodeERC721Asset(ctx, d.TokenID, d.TokenAddress)
	if err != nil {
		return nil, err
	}

	// Passing starkKeyHex to register method because it may be padded and converting back from Int loses the padding
	starkKeyHex := signableDeposit.StarkKey
	starkKey, err := hexutil.DecodeBig(starkKeyHex)
	if err != nil {
		return nil, fmt.Errorf("error converting StarkKey to bigint: %s", starkKeyHex)
	}

	isRegistered, _ := c.RegistrationContract.IsRegistered(&bind.CallOpts{Context: ctx}, starkKey)
	// Note: if we reach here, it means we are registered off-chain.
	// Above call will return an error user is not registered but this is for on-chain
	// we should swallow this error to allow the register and deposit flow to execute.

	if isRegistered {
		return c.depositERC721(ctx, l1signer, starkKey, big.NewInt(int64(signableDeposit.VaultId)), assetType, tokenID, overrides)
	}
	return c.registerAndDepositERC721(ctx, l1signer, starkKeyHex, starkKey, big.NewInt(int64(signableDeposit.VaultId)), assetType, tokenID, overrides)
}

func newSignableDepositRequestForERC721(amount, tokenID, tokenAddress, user string) *api.GetSignableDepositRequest {
	return &api.GetSignableDepositRequest{
		Amount: amount,
		Token:  SignableERC721Token(tokenID, tokenAddress),
		User:   user,
	}
}

func (c *Client) depositERC721(
	ctx context.Context,
	l1signer L1Signer,
	starkPublicKey *big.Int,
	vaultID *big.Int,
	assetType *big.Int,
	tokenID *big.Int,
	overrides *bind.TransactOpts,
) (*types.Transaction, error) {
	opts := c.buildTransactOpts(ctx, l1signer, overrides)
	tnx, err := c.CoreContract.DepositNft(opts, starkPublicKey, assetType, vaultID, tokenID)
	if err != nil {
		return nil, err
	}
	return tnx, nil
}

func (c *Client) registerAndDepositERC721(
	ctx context.Context,
	l1signer L1Signer,
	starkKeyHex string,
	starkKey *big.Int,
	vaultID *big.Int,
	assetType *big.Int,
	tokenID *big.Int,
	overrides *bind.TransactOpts,
) (*types.Transaction, error) {
	etherKey := l1signer.GetAddress()
	signableRegistration, err := c.GetSignableRegistrationOnchain(ctx, etherKey, starkKeyHex)
	if err != nil {
		return nil, err
	}

	opts := c.buildTransactOpts(ctx, l1signer, overrides)

	operatorSignature, err := convert.HexToByteArray(signableRegistration.OperatorSignature)
	if err != nil {
		return nil, err
	}
	tnx, err := c.RegistrationContract.RegisterAndDepositNft(
		opts,
		common.HexToAddress(etherKey),
		starkKey,
		operatorSignature,
		assetType,
		vaultID,
		tokenID,
	)
	if err != nil {
		return nil, err
	}
	return tnx, nil
}

// Deposit performs the deposit workflow on the ERC20Deposit.
func (d *ERC20Deposit) Deposit(ctx context.Context, c *Client, l1signer L1Signer, overrides *bind.TransactOpts) (*types.Transaction, error) {
	// Get decimals for this specific ERC20
	token, httpResp, err := c.GetToken(ctx, d.TokenAddress).Execute()
	if err != nil {
		return nil, fmt.Errorf("error when calling `Tokens.GetToken`: %v, http reponse body: %v", err, httpResp.Body)
	}

	decimals, err := strconv.Atoi(token.Decimals)
	if err != nil {
		return nil, fmt.Errorf("error parsing token decimals: %v", err)
	}
	amount, err := convert.ToDenomination(d.Amount, decimals)
	if err != nil {
		return nil, err
	}

	// Approve whether an amount of token from an account can be spent by a third-party account
	opts := c.buildTransactOpts(ctx, l1signer, overrides)
	ierc20Contract, err := c.newIERC20Contract(ctx, d.TokenAddress)
	if err != nil {
		return nil, err
	}
	_, err = ierc20Contract.Approve(opts, common.HexToAddress(c.Environment.CoreContractAddress), amount)
	if err != nil {
		return nil, err
	}

	// Get signable deposit details
	signableDepositRequest := newSignableDepositRequestForERC20(amount.String(), d.TokenAddress, l1signer.GetAddress(), decimals)
	signableDepositResponse, err := c.getSignableDeposit(ctx, signableDepositRequest)
	if err != nil {
		return nil, err
	}

	// Perform encoding on asset details to get an assetType (required for stark contract request)
	assetType, err := c.encodeERC20Asset(ctx, d.TokenAddress)
	if err != nil {
		return nil, err
	}

	starkKeyHex := signableDepositResponse.StarkKey
	starkKey, err := hexutil.DecodeBig(starkKeyHex)
	if err != nil {
		return nil, fmt.Errorf("error converting StarkKey to bigint: %s", starkKeyHex)
	}

	isRegistered, _ := c.RegistrationContract.IsRegistered(&bind.CallOpts{Context: ctx}, starkKey)
	// Note: if we reach here, it means we are registered off-chain.
	// Above call will return an error user is not registered but this is for on-chain
	// we should swallow this error to allow the register and deposit flow to execute.

	quantizedAmount, success := new(big.Int).SetString(signableDepositResponse.Amount, 10)
	if !success {
		return nil, fmt.Errorf("error converting string value '%s' to bigint", signableDepositResponse.Amount)
	}

	if isRegistered {
		return c.depositERC20(ctx, l1signer, starkKey, big.NewInt(int64(signableDepositResponse.VaultId)), assetType, quantizedAmount, overrides)
	}
	return c.registerAndDepositERC20(ctx, l1signer, starkKeyHex, starkKey, big.NewInt(int64(signableDepositResponse.VaultId)), assetType, quantizedAmount, overrides)
}

func newSignableDepositRequestForERC20(amount, tokenAddress, user string, decimals int) *api.GetSignableDepositRequest {
	return &api.GetSignableDepositRequest{
		Amount: amount,
		Token:  SignableERC20Token(decimals, tokenAddress),
		User:   user,
	}
}

func (c *Client) depositERC20(
	ctx context.Context,
	l1signer L1Signer,
	starkKey *big.Int,
	vaultID *big.Int,
	assetType *big.Int,
	quantizedAmount *big.Int,
	overrides *bind.TransactOpts,
) (*types.Transaction, error) {
	opts := c.buildTransactOpts(ctx, l1signer, overrides)
	tnx, err := c.CoreContract.DepositERC20(opts, starkKey, assetType, vaultID, quantizedAmount)
	if err != nil {
		return nil, err
	}
	return tnx, nil
}

func (c *Client) registerAndDepositERC20(
	ctx context.Context,
	l1signer L1Signer,
	starkKeyHex string,
	starkKey *big.Int,
	vaultID *big.Int,
	assetType *big.Int,
	quantizedAmount *big.Int,
	overrides *bind.TransactOpts,
) (*types.Transaction, error) {
	etherKey := l1signer.GetAddress()
	signableRegistration, err := c.GetSignableRegistrationOnchain(ctx, etherKey, starkKeyHex)
	if err != nil {
		return nil, err
	}

	opts := c.buildTransactOpts(ctx, l1signer, overrides)

	operatorSignature, err := convert.HexToByteArray(signableRegistration.OperatorSignature)
	if err != nil {
		return nil, err
	}

	tnx, err := c.CoreContract.RegisterAndDepositERC20(
		opts,
		common.HexToAddress(etherKey),
		starkKey,
		operatorSignature,
		assetType,
		vaultID,
		quantizedAmount,
	)
	if err != nil {
		return nil, err
	}
	return tnx, nil
}
