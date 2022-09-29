package imx

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/immutable/imx-core-sdk-golang/imx/api"
	"github.com/immutable/imx-core-sdk-golang/imx/internal/convert"
)

/*
PrepareWithdrawal submits a withdrawal request for ETH, ERC20 and ERC721 tokens to be included in the generation and submission of the next batch.
Upon batch confirmation (on-chain state update), the asset is available to be withdrawn by the initial owner/originator of the asset.

@param ctx context.Context - for cancellation, deadlines, tracing, etc or context.Background().
@param l1Signer Ethereum signer to sign message.
@param l2signer Stark signer to sign the payload hash.
@param request The request struct with all the params.
@return CreateWithdrawalResponse
*/
func (c *Client) PrepareWithdrawal(
	ctx context.Context,
	l1signer L1Signer,
	l2signer L2Signer,
	request api.GetSignableWithdrawalRequest,
) (*api.CreateWithdrawalResponse, error) {
	ethAddress := l1signer.GetAddress()
	request.User = ethAddress
	signableResponse, httpResponse, err := c.withdrawalsAPI.GetSignableWithdrawal(ctx).GetSignableWithdrawalRequest(request).Execute()
	if err != nil {
		return nil, NewIMXError(httpResponse, err)
	}

	ethSignature, starkSignature, err := createSignatures(&signableResponse.SignableMessage, &signableResponse.PayloadHash, l1signer, l2signer)
	if err != nil {
		return nil, err
	}

	withdrawalRequest := api.CreateWithdrawalRequest{
		Amount:         request.Amount,
		AssetId:        signableResponse.AssetId,
		Nonce:          signableResponse.Nonce,
		StarkKey:       signableResponse.StarkKey,
		StarkSignature: starkSignature,
		VaultId:        signableResponse.VaultId,
	}
	apiCreateWithdrawalRequest := c.withdrawalsAPI.CreateWithdrawal(ctx).XImxEthAddress(ethAddress).XImxEthSignature(ethSignature)
	withdrawalResponse, httpResponse, err := apiCreateWithdrawalRequest.CreateWithdrawalRequest(withdrawalRequest).Execute()
	if err != nil {
		return nil, NewIMXError(httpResponse, err)
	}
	return withdrawalResponse, nil
}

type TokenWithdrawal interface {
	CompleteWithdrawal(ctx context.Context, c *Client, l1signer L1Signer, starkPublicKey string, overrides *bind.TransactOpts) (*types.Transaction, error)
}

// ERC20Withdrawal implements TokenWithdrawal. Used for withdrawal of ERC20 Tokens.
type ERC20Withdrawal struct {
	TokenAddress string
}

// ERC721Withdrawal implements TokenWithdrawal. Used for withdrawal of ERC721 Tokens.
type ERC721Withdrawal struct {
	TokenID      string
	TokenAddress string
}

// NewERC20Withdrawal instantiates a new ERC20Withdrawal object with the given tokenAddress.
func NewERC20Withdrawal(tokenAddress string) *ERC20Withdrawal {
	this := ERC20Withdrawal{}
	this.TokenAddress = tokenAddress
	return &this
}

// NewERC721Withdrawal instantiates a new ERC721Withdrawal object with the given tokenId and tokenAddress.
func NewERC721Withdrawal(tokenID, tokenAddress string) *ERC721Withdrawal {
	this := ERC721Withdrawal{}
	this.TokenID = tokenID
	this.TokenAddress = tokenAddress
	return &this
}

/*
CompleteEthWithdrawal performs the complete withdrawal workflow for ETH

@param ctx context.Context - for cancellation, deadlines, tracing, etc or context.Background().
@param l1Signer Ethereum signer to sign message.
@param starkKeyHex Stark key string in hex decimal format.
@param overrides Optional transaction params that overrides the default values.
@return Transaction
*/
func (c *Client) CompleteEthWithdrawal(
	ctx context.Context,
	l1signer L1Signer,
	starkKeyHex string,
	overrides *bind.TransactOpts,
) (*types.Transaction, error) {
	assetType, err := c.encodeETHAsset(ctx)
	if err != nil {
		return nil, err
	}
	return c.completeFungiblesWithdrawal(ctx, l1signer, starkKeyHex, assetType, overrides)
}

/*
CompleteWithdrawal performs the complete withdrawal workflow for ERC20 tokens.

@param ctx context.Context - for cancellation, deadlines, tracing, etc or context.Background().
@param l1Signer Ethereum signer to sign message.
@param starkKeyHex Stark key string in hex decimal format.
@param overrides Optional transaction params that overrides the default values.
@return Transaction
*/
func (w *ERC20Withdrawal) CompleteWithdrawal(
	ctx context.Context,
	c *Client,
	l1signer L1Signer,
	starkKeyHex string,
	overrides *bind.TransactOpts,
) (*types.Transaction, error) {
	assetType, err := c.encodeERC20Asset(ctx, w.TokenAddress)
	if err != nil {
		return nil, err
	}
	return c.completeFungiblesWithdrawal(ctx, l1signer, starkKeyHex, assetType, overrides)
}

func (c *Client) completeFungiblesWithdrawal(
	ctx context.Context,
	l1signer L1Signer,
	starkKeyHex string,
	assetType *big.Int,
	overrides *bind.TransactOpts,
) (*types.Transaction, error) {
	starkKey, ok := new(big.Int).SetString(starkKeyHex, 0)
	if !ok {
		return nil, fmt.Errorf("error converting starkKeyHex to bigint")
	}

	isRegistered, _ := c.registrationContract.IsRegistered(&bind.CallOpts{Context: ctx}, starkKey)

	if isRegistered {
		return c.withdrawFungibles(ctx, l1signer, starkKey, assetType, overrides)
	}
	return c.registerAndWithdrawFungibles(ctx, l1signer, starkKey, assetType, overrides)
}

func (c *Client) withdrawFungibles(
	ctx context.Context,
	l1signer L1Signer,
	starkKey, assetType *big.Int,
	overrides *bind.TransactOpts,
) (*types.Transaction, error) {
	opts := c.buildTransactOpts(ctx, l1signer, overrides)
	tnx, err := c.coreContract.Withdraw(opts, starkKey, assetType)
	if err != nil {
		return nil, err
	}
	return tnx, nil
}

func (c *Client) registerAndWithdrawFungibles(
	ctx context.Context,
	l1signer L1Signer,
	starkKey *big.Int,
	assetType *big.Int,
	overrides *bind.TransactOpts,
) (*types.Transaction, error) {
	etherKey := l1signer.GetAddress()
	signableRegistration, err := c.getSignableRegistrationOnchain(ctx, etherKey, hexutil.EncodeBig(starkKey))
	if err != nil {
		return nil, err
	}

	opts := c.buildTransactOpts(ctx, l1signer, overrides)

	operatorSignature, err := convert.HexToByteArray(signableRegistration.OperatorSignature)
	if err != nil {
		return nil, err
	}
	transaction, err := c.registrationContract.RegisterAndWithdraw(opts, common.HexToAddress(etherKey), starkKey, operatorSignature, assetType)
	if err != nil {
		return nil, err
	}
	return transaction, nil
}

func (w *ERC721Withdrawal) withdrawMintedNft(
	ctx context.Context,
	c *Client,
	l1signer L1Signer,
	starkKeyHex string,
	overrides *bind.TransactOpts,
) (*types.Transaction, error) {
	assetType, err := c.encodeERC721Asset(ctx, w.TokenID, w.TokenAddress)
	if err != nil {
		return nil, err
	}

	starkKey, ok := new(big.Int).SetString(starkKeyHex, 0)
	if !ok {
		return nil, fmt.Errorf("error converting starkKeyHex to bigint")
	}

	isRegistered, err := c.registrationContract.IsRegistered(&bind.CallOpts{Context: ctx}, starkKey)
	if err != nil {
		return nil, fmt.Errorf("error when calling 'ethClient.registrationContract.IsRegistered': %v", err)
	}

	tokenID, ok := new(big.Int).SetString(w.TokenID, 10)
	if !ok {
		return nil, fmt.Errorf("error converting tokenID to bigint: %v", w.TokenID)
	}

	if isRegistered {
		return c.withdrawMintedNft(ctx, l1signer, starkKey, assetType, tokenID, overrides)
	}
	return c.registerAndWithdrawMintedNft(ctx, l1signer, starkKey, assetType, tokenID, overrides)
}

func (c *Client) withdrawMintedNft(
	ctx context.Context,
	l1signer L1Signer,
	starkKey, assetType, tokenID *big.Int,
	overrides *bind.TransactOpts,
) (*types.Transaction, error) {
	opts := c.buildTransactOpts(ctx, l1signer, overrides)
	tnx, err := c.coreContract.WithdrawNft(opts, starkKey, assetType, tokenID)
	if err != nil {
		return nil, err
	}
	return tnx, nil
}

func (c *Client) registerAndWithdrawMintedNft(
	ctx context.Context,
	l1signer L1Signer,
	starkKey *big.Int,
	assetType *big.Int,
	tokenID *big.Int,
	overrides *bind.TransactOpts,
) (*types.Transaction, error) {
	etherKey := l1signer.GetAddress()
	signableRegistration, err := c.getSignableRegistrationOnchain(ctx, etherKey, hexutil.EncodeBig(starkKey))
	if err != nil {
		return nil, err
	}

	opts := c.buildTransactOpts(ctx, l1signer, overrides)

	operatorSignature, err := convert.HexToByteArray(signableRegistration.OperatorSignature)
	if err != nil {
		return nil, err
	}
	tnx, err := c.registrationContract.RegisterAndWithdrawNft(opts, common.HexToAddress(etherKey), starkKey, operatorSignature, assetType, tokenID)
	if err != nil {
		return nil, err
	}
	return tnx, nil
}

/*
CompleteWithdrawal performs the completion step of the withdrawal process for ERC721 token.

@param ctx context.Context - for cancellation, deadlines, tracing, etc or context.Background().
@param l1Signer Ethereum signer to sign message.
@param starkKeyHex Stark key string in hex decimal format.
@param overrides Optional transaction params that overrides the default values.
@return Transaction
*/
func (w *ERC721Withdrawal) CompleteWithdrawal(
	ctx context.Context,
	c *Client,
	l1signer L1Signer,
	starkKeyHex string,
	overrides *bind.TransactOpts,
) (*types.Transaction, error) {
	mintableTokenResponse, httpResponse, err := c.mintsAPI.GetMintableTokenDetailsByClientTokenId(ctx, w.TokenAddress, w.TokenID).Execute()
	if err != nil {
		if httpResponse.StatusCode == 404 {
			// Token is already minted on L1
			return w.withdrawMintedNft(ctx, c, l1signer, starkKeyHex, overrides)
		}
		return nil, NewIMXError(httpResponse, err)
	}

	blueprint := mintableTokenResponse.Blueprint
	mintingBlob := getMintingBlob(w.TokenID, blueprint)
	assetType, err := c.encodeMintableERC721Asset(ctx, w.TokenID, w.TokenAddress, blueprint)
	if err != nil {
		return nil, err
	}

	starkKey, ok := new(big.Int).SetString(starkKeyHex, 0)
	if !ok {
		return nil, fmt.Errorf("error converting starkKeyHex to bigint")
	}

	isRegistered, _ := c.registrationContract.IsRegistered(&bind.CallOpts{Context: ctx}, starkKey)
	// Note: if we reach here, it means we are registered off-chain.
	// Above call will return an error user is not registered but this is for on-chain
	// we should swallow this error to allow the register and withdraw flow to execute.

	if isRegistered {
		return c.withdrawAndMintNft(ctx, l1signer, starkKey, assetType, mintingBlob, overrides)
	}
	return c.registerAndWithdrawAndMintNft(ctx, l1signer, starkKey, assetType, mintingBlob, overrides)
}

func (c *Client) withdrawAndMintNft(
	ctx context.Context,
	l1signer L1Signer,
	starkKey, assetType *big.Int,
	mintingBlob []byte,
	overrides *bind.TransactOpts,
) (*types.Transaction, error) {
	opts := c.buildTransactOpts(ctx, l1signer, overrides)
	tnx, err := c.coreContract.WithdrawAndMint(opts, starkKey, assetType, mintingBlob)
	if err != nil {
		return nil, err
	}
	return tnx, nil
}

func (c *Client) registerAndWithdrawAndMintNft(
	ctx context.Context,
	l1signer L1Signer,
	starkKey *big.Int,
	assetType *big.Int,
	mintingBlob []byte,
	overrides *bind.TransactOpts,
) (*types.Transaction, error) {
	etherKey := l1signer.GetAddress()
	signableRegistration, err := c.getSignableRegistrationOnchain(ctx, etherKey, hexutil.EncodeBig(starkKey))
	if err != nil {
		return nil, err
	}

	opts := c.buildTransactOpts(ctx, l1signer, overrides)

	operatorSignature, err := convert.HexToByteArray(signableRegistration.OperatorSignature)
	if err != nil {
		return nil, err
	}

	tnx, err := c.registrationContract.RegsiterAndWithdrawAndMint(opts, common.HexToAddress(etherKey), starkKey, operatorSignature, assetType, mintingBlob)
	if err != nil {
		return nil, err
	}
	return tnx, nil
}

func getMintingBlob(tokenID, blueprint string) []byte {
	return []byte(fmt.Sprintf("{%s}:{%s}", tokenID, blueprint))
}

/*
GetWithdrawal Get details of a withdrawal with the given ID

@param ctx context.Context - for cancellation, deadlines, tracing, etc or context.Background().
@param id Withdrawal ID
@return Withdrawal
*/
func (c *Client) GetWithdrawal(ctx context.Context, id string) (*api.Withdrawal, error) {
	response, httpResponse, err := c.withdrawalsAPI.GetWithdrawal(ctx, id).Execute()
	if err != nil {
		return nil, NewIMXError(httpResponse, err)
	}
	return response, nil
}

/*
NewListWithdrawalsRequest Creates a new ApiListWithdrawalsRequest object with required params.

@param ctx context.Context - for cancellation, deadlines, tracing, etc or context.Background().
@return ApiListWithdrawalsRequest
*/
func (c *Client) NewListWithdrawalsRequest(ctx context.Context) api.ApiListWithdrawalsRequest {
	return c.withdrawalsAPI.ListWithdrawals(ctx)
}

/*
ListWithdrawals Gets a list of withdrawals

@param req the api request struct with all params populated to make the request
@return ListWithdrawalsResponse
*/
func (c *Client) ListWithdrawals(req *api.ApiListWithdrawalsRequest) (*api.ListWithdrawalsResponse, error) {
	response, httpResponse, err := c.withdrawalsAPI.ListWithdrawalsExecute(*req)
	if err != nil {
		return nil, NewIMXError(httpResponse, err)
	}
	return response, nil
}
