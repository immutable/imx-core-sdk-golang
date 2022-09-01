package imx

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/immutable/imx-core-sdk-golang/api"
)

// BurnAddress 0x00 where tokens are sent to be removed from circulation
const BurnAddress = "0x0000000000000000000000000000000000000000"

type GetSignableBurnRequest struct {
	// Amount of the token to burn
	Amount string
	// Ethereum address of the user
	Sender string
	// Token to burn
	Token api.SignableToken
}

// Burn permanently removes tokens from circulation.
// To burn an asset, you need to transfer the asset to the burn address.
// We have added burn functions in the SDK for clarity, and these are just wrappers for standard transfer functions to BurnAddress.
// Using the transfer functions with the burn address as a recipient will also work just the same.
func (c *Client) Burn(
	ctx context.Context,
	l1signer L1Signer,
	l2signer L2Signer,
	request GetSignableBurnRequest) (*api.CreateTransferResponseV1, error) {
	burnAddress := BurnAddress
	transferRequest := api.GetSignableTransferRequestV1{
		Amount:   request.Amount,
		Sender:   request.Sender,
		Token:    request.Token,
		Receiver: burnAddress,
	}
	return c.CreateTransfer(ctx, l1signer, l2signer, transferRequest)
}

// GetBurn returns the status of the transfer given the transferID.
func (c *Client) GetBurn(ctx context.Context, transferID string) (*api.Transfer, error) {
	response, httpResponse, err := c.TransfersApi.GetTransfer(ctx, transferID).Execute()
	if err != nil {
		return nil, fmt.Errorf("error when calling `TransfersApi.GetTransfer`: %v, HTTP response body: %v", err, httpResponse.Body)
	}
	return response, nil
}

// CreateTrade submits a matched order to the CreateTrade endpoint.
// https://docs.x.immutable.com/reference#/operations/createTrade
func (c *Client) CreateTrade(
	ctx context.Context,
	l1signer L1Signer,
	l2signer L2Signer,
	request api.GetSignableTradeRequest,
) (*api.CreateTradeResponse, error) {
	ethAddress := l1signer.GetAddress()
	request.User = ethAddress
	signableTrade, httpResponse, err := c.GetSignableTrade(ctx).GetSignableTradeRequest(request).Execute()
	if err != nil {
		return nil, fmt.Errorf("error when calling `TradesApi.GetSignableTrade`: %v, HTTP response body: %v", err, httpResponse.Body)
	}

	ethSignature, starkSignature, err := createSignatures(&signableTrade.SignableMessage, &signableTrade.PayloadHash, l1signer, l2signer)
	if err != nil {
		return nil, err
	}

	includeFees := true
	createTradeResponse, httpResponse, err := c.TradesApi.CreateTrade(ctx).
		CreateTradeRequest(api.CreateTradeRequestV1{
			AmountBuy:           signableTrade.AmountBuy,
			AmountSell:          signableTrade.AmountSell,
			AssetIdBuy:          signableTrade.AssetIdBuy,
			AssetIdSell:         signableTrade.AssetIdSell,
			ExpirationTimestamp: signableTrade.ExpirationTimestamp,
			FeeInfo:             signableTrade.FeeInfo,
			Fees:                request.Fees,
			IncludeFees:         &includeFees,
			Nonce:               signableTrade.Nonce,
			OrderId:             request.OrderId,
			StarkKey:            signableTrade.StarkKey,
			StarkSignature:      starkSignature,
			VaultIdBuy:          signableTrade.VaultIdBuy,
			VaultIdSell:         signableTrade.VaultIdSell,
		}).
		XImxEthAddress(ethAddress).XImxEthSignature(ethSignature).Execute()
	if err != nil {
		return nil, fmt.Errorf("error when calling `TradesApi.CreateTrade`: %v, HTTP response body: %v", err, httpResponse.Body)
	}
	return createTradeResponse, nil
}

// Note: 	These structs need to be in the order defined as below and should match the
//   		implementation in imx-engine to pass the validation step while minting tokens.

type MintFee struct {
	Recipient  string  `json:"recipient"`
	Percentage float32 `json:"percentage" validate:"max=100,gt=0"`
}

type MintableTokenData struct {
	ID        string    `json:"id"`
	Blueprint *string   `json:"blueprint"`
	Royalties []MintFee `json:"royalties,omitempty"` // token-level overridable fees (optional)
}

type User struct {
	User   string              `json:"ether_key"`
	Tokens []MintableTokenData `json:"tokens"`
}

type UnsignedMintRequest struct {
	ContractAddress string    `json:"contract_address" validate:"required,eth_addr"`
	Royalties       []MintFee `json:"royalties,omitempty" validate:"dive"` // contract-level (optional)
	Users           []User    `json:"users" validate:"required,dive,min=1"`
	AuthSignature   string    `json:"auth_signature" validate:"required"`
}

// MintTokensWorkflow assists in minting tokens to the given imx user.
func (c *Client) Mint(
	ctx context.Context,
	l1signer L1Signer,
	unsignedMintRequest []UnsignedMintRequest,
) (*api.MintTokensResponse, error) {
	mintRequest := make([]api.MintRequest, len(unsignedMintRequest))
	for requestIndex, eachMintRequest := range unsignedMintRequest {
		mintRequestInBytes, err := json.Marshal(eachMintRequest)
		if err != nil {
			return nil, fmt.Errorf("error in json marshaling eachMintRequest: %v", err)
		}
		requestHash := crypto.Keccak256Hash(mintRequestInBytes)
		authSignatureInBytes, err := l1signer.SignMessage(requestHash.String())
		if err != nil {
			return nil, fmt.Errorf("error in signing eachMintRequest message: %v", err)
		}

		mintFees := make([]api.MintFee, len(eachMintRequest.Royalties))
		for index, eachMintFee := range eachMintRequest.Royalties {
			mintFees[index] = api.MintFee{
				Percentage: eachMintFee.Percentage,
				Recipient:  eachMintFee.Recipient,
			}
		}

		mintToUsers := make([]api.MintUser, len(eachMintRequest.Users))
		for userIndex, eachMintUser := range eachMintRequest.Users {
			mintTokens := make([]api.MintTokenDataV2, len(eachMintUser.Tokens))
			for tokenIndex, eachMintToken := range eachMintUser.Tokens {
				mintFeesPerToken := make([]api.MintFee, len(eachMintToken.Royalties))
				for royaltyIndex, eachMintFeePerToken := range eachMintToken.Royalties {
					mintFeesPerToken[royaltyIndex] = api.MintFee{
						Percentage: eachMintFeePerToken.Percentage,
						Recipient:  eachMintFeePerToken.Recipient,
					}
				}

				mintTokens[tokenIndex] = api.MintTokenDataV2{
					Blueprint: eachMintToken.Blueprint,
					Id:        eachMintToken.ID,
					Royalties: mintFeesPerToken,
				}
			}
			mintToUsers[userIndex] = api.MintUser{
				Tokens: mintTokens,
				User:   eachMintUser.User,
			}
		}

		authSignatureEncodedInHex := hexutil.Encode(authSignatureInBytes)
		mintRequest[requestIndex] = api.MintRequest{
			AuthSignature:   authSignatureEncodedInHex,
			ContractAddress: eachMintRequest.ContractAddress,
			Royalties:       mintFees,
			Users:           mintToUsers,
		}
	}

	mintTokensResponse, httpResp, err := c.MintTokens(ctx).MintTokensRequestV2(mintRequest).Execute()
	if err != nil {
		return nil, fmt.Errorf("error when calling `api.Mints.MintTokens`: %v, HTTP response body: %v", err, httpResp.Body)
	}
	return mintTokensResponse, nil
}

// CreateOrder will list the given asset for sale on the marketplace.
func (c *Client) CreateOrder(ctx context.Context,
	l1signer L1Signer,
	l2signer L2Signer,
	request *api.GetSignableOrderRequest) (*api.CreateOrderResponse, error) {
	ethAddress := l1signer.GetAddress()
	request.User = ethAddress
	signableOrder, httpResponse, err := c.GetSignableOrder(ctx).GetSignableOrderRequestV3(*request).Execute()
	if err != nil {
		return nil, fmt.Errorf("error when calling `OrdersApi.GetSignableOrder`: %v, HTTP response body: %v", err, httpResponse.Body)
	}

	ethSignature, starkSignature, err := createSignatures(&signableOrder.SignableMessage, &signableOrder.PayloadHash, l1signer, l2signer)
	if err != nil {
		return nil, err
	}

	includeFees := true
	createOrderResponse, httpResponse, err := c.OrdersApi.CreateOrder(ctx).
		CreateOrderRequest(api.CreateOrderRequest{
			AmountBuy: signableOrder.AmountBuy, // The amount (listing price) should be in Wei for Eth tokens,
			// see https://docs.starkware.co/starkex-v4/starkex-deep-dive/starkex-specific-concepts and https://eth-converter.com/
			AmountSell:          signableOrder.AmountSell,  // Quantity to list for sale
			AssetIdBuy:          signableOrder.AssetIdBuy,  // Token with which this asset can be bought with. For example Eth token can be used to buy this asset.
			AssetIdSell:         signableOrder.AssetIdSell, // Token of the asset being listed for sale. For example, ERC720 is a NFT Token and ERC20 is a fungible token.
			ExpirationTimestamp: signableOrder.ExpirationTimestamp,
			Fees:                request.Fees,
			IncludeFees:         &includeFees,
			Nonce:               signableOrder.Nonce,
			StarkKey:            signableOrder.StarkKey,
			StarkSignature:      starkSignature,
			VaultIdBuy:          signableOrder.VaultIdBuy,
			VaultIdSell:         signableOrder.VaultIdSell,
		}).XImxEthAddress(ethAddress).XImxEthSignature(ethSignature).Execute()
	if err != nil {
		return nil, fmt.Errorf("error when calling `OrdersApi.CreateOrder`: %v, HTTP response body: %v", err, httpResponse.Body)
	}
	return createOrderResponse, nil
}

// CancelOrder will remove the listed asset on marketplace from sale.
func (c *Client) CancelOrder(ctx context.Context,
	l1signer L1Signer,
	l2signer L2Signer,
	request api.GetSignableCancelOrderRequest,
) (*api.CancelOrderResponse, error) {
	signableCancelOrder, httpResponse, err := c.GetSignableCancelOrder(ctx).GetSignableCancelOrderRequest(request).Execute()
	if err != nil {
		return nil, fmt.Errorf("error when calling `OrdersApi.GetSignableCancelOrder`: %v, HTTP response body: %v", err, httpResponse.Body)
	}

	ethSignature, starkSignature, err := createSignatures(&signableCancelOrder.SignableMessage, &signableCancelOrder.PayloadHash, l1signer, l2signer)
	if err != nil {
		return nil, err
	}

	ethAddress := l1signer.GetAddress()
	orderID := strconv.FormatInt(int64(request.OrderId), 10)
	cancelOrderResponse, httpResponse, err := c.OrdersApi.CancelOrder(ctx, orderID).
		CancelOrderRequest(api.CancelOrderRequest{
			OrderId:        request.OrderId,
			StarkSignature: starkSignature,
		}).XImxEthAddress(ethAddress).XImxEthSignature(ethSignature).Execute()
	if err != nil {
		return nil, fmt.Errorf("error when calling `OrdersApi.CancelOrder`: %v, HTTP response body: %v", err, httpResponse.Body)
	}
	return cancelOrderResponse, nil
}
