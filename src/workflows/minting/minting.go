package minting

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"immutable.com/imx-core-sdk-golang/api/client"
	"immutable.com/imx-core-sdk-golang/api/client/mints"
	"immutable.com/imx-core-sdk-golang/api/models"
	"immutable.com/imx-core-sdk-golang/signers"
)

func MintTokensWorkflow(ctx context.Context,
	api *client.ImmutableXAPI,
	l1signer signers.L1Signer,
	request []*models.MintRequest) (*models.MintTokensResponse, error) {

	for _, eachMintRequst := range request {
		mintRequestInBytes, err := eachMintRequst.MarshalBinary()
		if err != nil {
			return nil, fmt.Errorf("error marshaling MintRequest: %v", err)
		}
		authSignatureInBytes, err := l1signer.SignMessage(string(mintRequestInBytes))
		if err != nil {
			return nil, fmt.Errorf("error generating signature from MintRequest data: %v", err)
		}
		authSignatureEncodedInHex := hexutil.Encode(authSignatureInBytes)
		eachMintRequst.AuthSignature = &authSignatureEncodedInHex
		if err != nil {
			return nil, fmt.Errorf("error marshaling MintRequest: %v", err)
		}
	}

	mintTokenParams := mints.NewMintTokensParamsWithContext(ctx)
	mintTokenParams.SetMintTokensRequestV2(request)
	mintTokensOk, err := api.Mints.MintTokens(mintTokenParams)
	if err != nil {
		return nil, fmt.Errorf("error when calling `api.Mints.MintTokens`: %v", err)
	}
	return mintTokensOk.GetPayload(), nil
}
