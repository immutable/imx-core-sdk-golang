package imx

import (
	"context"
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/immutable/imx-core-sdk-golang/imx/api"
	"github.com/immutable/imx-core-sdk-golang/imx/internal/convert"
)

/*
Wei type used to specify the token amounts.

For example,
wei  = Wei(1)
gwei = Wei(1000000000)
eth  = Wei(1000000000000000000)
*/
type Wei = uint64

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
func NewETHDeposit(amount Wei) *ETHDeposit {
	this := ETHDeposit{}
	this.Amount = strconv.FormatUint(amount, 10)
	return &this
}

// NewERC20Deposit instantiates a new ERC20Deposit object with given amount and tokenAddress.
func NewERC20Deposit(unDecimalisedAmount uint64, tokenAddress string) *ERC20Deposit {
	this := ERC20Deposit{}
	this.Amount = strconv.FormatUint(unDecimalisedAmount, 10)
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

/*
GetDeposit Gets details of a deposit with the given ID

@param ctx context.Context - for cancellation, deadlines, tracing, etc or context.Background().
@param id Deposit Id
@return Deposit
*/
func (c *Client) GetDeposit(ctx context.Context, id string) (*api.Deposit, error) {
	response, httpResponse, err := c.depositsAPI.GetDeposit(ctx, id).Execute()
	if err != nil {
		return nil, NewAPIError(httpResponse, err)
	}
	return response, nil
}

/*
NewListDepositsRequest Creates a new ApiListDepositsRequest object with required params.

@param ctx context.Context - for cancellation, deadlines, tracing, etc or context.Background().
@return ApiListDepositsRequest
*/
func (c *Client) NewListDepositsRequest(ctx context.Context) api.ApiListDepositsRequest {
	return c.depositsAPI.ListDeposits(ctx)
}

/*
ListDeposits Gets a list of deposits

@param req the api request struct with all params populated to make the request
@return ListDepositsResponse
*/
func (c *Client) ListDeposits(req *api.ApiListDepositsRequest) (*api.ListDepositsResponse, error) {
	response, httpResponse, err := c.depositsAPI.ListDepositsExecute(*req)
	if err != nil {
		return nil, NewAPIError(httpResponse, err)
	}
	return response, nil
}

/*
Deposit performs the deposit workflow on the ETHDeposit.

@param ctx context.Context - for cancellation, deadlines, tracing, etc or context.Background().
@param c Client object from interface.go used to make API calls.
@param l1Signer Ethereum signer to sign message.
@param overrides Optional transaction params that overrides the default values.
@return Transaction
*/
func (d *ETHDeposit) Deposit(ctx context.Context, c *Client, l1signer L1Signer, overrides *bind.TransactOpts) (*types.Transaction, error) {
	signableDepositRequest := newSignableDepositRequestForEth(d.Amount, l1signer.GetAddress())
	signableDeposit, err := c.getSignableDeposit(ctx, signableDepositRequest)
	if err != nil {
		return nil, err
	}

	assetType, err := c.encodeETHAsset(ctx)
	if err != nil {
		return nil, err
	}

	starkKey, ok := new(big.Int).SetString(signableDeposit.StarkKey, 0)
	if !ok {
		return nil, fmt.Errorf("error converting StarkKey to bigint")
	}

	isRegistered, _ := c.registrationContract.IsRegistered(&bind.CallOpts{Context: ctx}, starkKey)
	// Note: if we reach here, it means we are registered off-chain.
	// Above call will return an error user is not registered but this is for on-chain
	// we should swallow this error to allow the register and deposit flow to execute.

	amount, ok := new(big.Int).SetString(d.Amount, 10)
	if !ok {
		return nil, fmt.Errorf("error converting amount to bigint: %v", d.Amount)
	}
	if isRegistered {
		return c.depositEth(ctx, l1signer, starkKey, big.NewInt(int64(signableDeposit.VaultId)), assetType, amount, overrides)
	}
	return c.registerAndDepositEth(ctx, l1signer, signableDeposit.StarkKey, starkKey, big.NewInt(int64(signableDeposit.VaultId)), assetType, amount, overrides)
}

func (c *Client) getSignableDeposit(
	ctx context.Context,
	request *api.GetSignableDepositRequest) (*api.GetSignableDepositResponse, error) {
	signableDepositResponse, httpResponse, err := c.depositsAPI.GetSignableDeposit(ctx).GetSignableDepositRequest(*request).Execute()
	if err != nil {
		return nil, NewAPIError(httpResponse, err)
	}
	return signableDepositResponse, nil
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
	signableRegistration, err := c.getSignableRegistrationOnchain(ctx, etherKey, starkKeyHex)
	if err != nil {
		return nil, err
	}

	opts := c.buildTransactOpts(ctx, l1signer, overrides)

	operatorSignature, err := convert.HexToByteArray(signableRegistration.OperatorSignature)
	if err != nil {
		return nil, err
	}
	opts.Value = amount
	tnx, err := c.coreContract.RegisterAndDepositEth(
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
	tnx, err := c.coreContract.Deposit(opts, starkPublicKey, assetType, vaultID)
	if err != nil {
		return nil, err
	}
	return tnx, nil
}

/*
Deposit performs the deposit workflow on the ERC721Deposit.

@param ctx context.Context - for cancellation, deadlines, tracing, etc or context.Background().
@param c Client object from interface.go used to make API calls.
@param l1Signer Ethereum signer to sign message.
@param overrides Optional transaction params that overrides the default values.
@return Transaction
*/
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

	starkKey, ok := new(big.Int).SetString(signableDeposit.StarkKey, 0)
	if !ok {
		return nil, fmt.Errorf("error converting StarkKey to bigint")
	}

	isRegistered, _ := c.registrationContract.IsRegistered(&bind.CallOpts{Context: ctx}, starkKey)
	// Note: if we reach here, it means we are registered off-chain.
	// Above call will return an error user is not registered but this is for on-chain
	// we should swallow this error to allow the register and deposit flow to execute.

	if isRegistered {
		return c.depositERC721(ctx, l1signer, starkKey, big.NewInt(int64(signableDeposit.VaultId)), assetType, tokenID, overrides)
	}
	// Passing starkKey string in hex format to register method because it may be padded and converting back from Int loses the padding
	return c.registerAndDepositERC721(ctx, l1signer, signableDeposit.StarkKey, starkKey, big.NewInt(int64(signableDeposit.VaultId)), assetType, tokenID, overrides)
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
	tnx, err := c.coreContract.DepositNft(opts, starkPublicKey, assetType, vaultID, tokenID)
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
	signableRegistration, err := c.getSignableRegistrationOnchain(ctx, etherKey, starkKeyHex)
	if err != nil {
		return nil, err
	}

	opts := c.buildTransactOpts(ctx, l1signer, overrides)

	operatorSignature, err := convert.HexToByteArray(signableRegistration.OperatorSignature)
	if err != nil {
		return nil, err
	}
	tnx, err := c.registrationContract.RegisterAndDepositNft(
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

/*
Deposit performs the deposit workflow on the ERC20Deposit.

@param ctx context.Context - for cancellation, deadlines, tracing, etc or context.Background().
@param c Client object from interface.go used to make API calls.
@param l1Signer Ethereum signer to sign message.
@param overrides Optional transaction params that overrides the default values.
@return Transaction
*/
func (d *ERC20Deposit) Deposit(ctx context.Context, c *Client, l1signer L1Signer, overrides *bind.TransactOpts) (*types.Transaction, error) {
	// Get decimals for this specific ERC20
	token, httpResponse, err := c.tokensAPI.GetToken(ctx, d.TokenAddress).Execute()
	if err != nil {
		return nil, NewAPIError(httpResponse, err)
	}

	// Approve whether an amount of token from an account can be spent by a third-party account
	opts := c.buildTransactOpts(ctx, l1signer, overrides)
	ierc20Contract, err := c.newIERC20Contract(ctx, d.TokenAddress)
	if err != nil {
		return nil, err
	}

	amount, ok := new(big.Int).SetString(d.Amount, 10)
	if !ok {
		return nil, fmt.Errorf("error converting amount to bigint: %v", d.Amount)
	}

	if err != nil {
		return nil, err
	}
	_, err = ierc20Contract.Approve(opts, common.HexToAddress(c.Environment.CoreContractAddress), amount)
	if err != nil {
		return nil, err
	}

	decimals, err := strconv.Atoi(token.Decimals)
	if err != nil {
		return nil, fmt.Errorf("error parsing token decimals: %v", err)
	}
	// Get signable deposit details
	signableDepositRequest := newSignableDepositRequestForERC20(d.Amount, d.TokenAddress, l1signer.GetAddress(), decimals)
	signableDeposit, err := c.getSignableDeposit(ctx, signableDepositRequest)
	if err != nil {
		return nil, err
	}

	// Perform encoding on asset details to get an assetType (required for stark contract request)
	assetType, err := c.encodeERC20Asset(ctx, d.TokenAddress)
	if err != nil {
		return nil, err
	}

	starkKey, ok := new(big.Int).SetString(signableDeposit.StarkKey, 0)
	if !ok {
		return nil, fmt.Errorf("error converting StarkKey to bigint")
	}

	isRegistered, _ := c.registrationContract.IsRegistered(&bind.CallOpts{Context: ctx}, starkKey)
	// Note: if we reach here, it means we are registered off-chain.
	// Above call will return an error user is not registered but this is for on-chain
	// we should swallow this error to allow the register and deposit flow to execute.

	depositResponseAmount, success := new(big.Int).SetString(signableDeposit.Amount, 10)
	if !success {
		return nil, fmt.Errorf("error converting string value '%s' to bigint", signableDeposit.Amount)
	}

	if isRegistered {
		return c.depositERC20(ctx, l1signer, starkKey, big.NewInt(int64(signableDeposit.VaultId)), assetType, depositResponseAmount, overrides)
	}
	// Passing starkKey string in hex format to register method because it may be padded and converting back from Int loses the padding
	return c.registerAndDepositERC20(ctx, l1signer, signableDeposit.StarkKey, starkKey, big.NewInt(int64(signableDeposit.VaultId)), assetType, depositResponseAmount, overrides)
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
	tnx, err := c.coreContract.DepositERC20(opts, starkKey, assetType, vaultID, quantizedAmount)
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
	signableRegistration, err := c.getSignableRegistrationOnchain(ctx, etherKey, starkKeyHex)
	if err != nil {
		return nil, err
	}

	opts := c.buildTransactOpts(ctx, l1signer, overrides)

	operatorSignature, err := convert.HexToByteArray(signableRegistration.OperatorSignature)
	if err != nil {
		return nil, err
	}

	tnx, err := c.coreContract.RegisterAndDepositERC20(
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
