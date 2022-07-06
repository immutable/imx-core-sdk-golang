package utils

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"immutable.com/imx-core-sdk-golang/api/models"
	"immutable.com/imx-core-sdk-golang/signers"
	"immutable.com/imx-core-sdk-golang/workflows/types"
)

// CreateSignatures accepts signable data and signers and returns signatures.
// signableMessage is signed by l1signer producing ethSignature.
// payloadHash is signed by l2signer producing starkSignature.
func CreateSignatures(signableMessage, payloadHash *string, l1signer signers.L1Signer, l2signer signers.L2Signer) (ethSignature, starkSignature string, err error) {
	ethSignatureBytes, err := l1signer.SignMessage(*signableMessage)
	if err != nil {
		return "", "", fmt.Errorf("error generating EthSignature from SignableMessage: %v", err)
	}
	ethSignature = hexutil.Encode(ethSignatureBytes)

	if starkSignature, err = l2signer.SignMessage(*payloadHash); err != nil {
		return "", "", fmt.Errorf("error generating StarkSignature from PayloadHash: %v", err)
	}
	return ethSignature, starkSignature, nil
}

func NewSignableTokenEth() *models.SignableToken {
	return &models.SignableToken{
		Data: map[string]interface{}{
			"decimals": 18,
		},
		Type: string(types.ETHType),
	}
}

func NewSignableTokenERC20(decimals, tokenAddress string) *models.SignableToken {
	return &models.SignableToken{
		Data: map[string]interface{}{
			"decimals":      decimals,
			"token_address": tokenAddress,
		},
		Type: string(types.ERC20Type),
	}
}

func NewSignableTokenERC721(tokenId, tokenAddress string) *models.SignableToken {
	return &models.SignableToken{
		Data: map[string]interface{}{
			"token_id":      tokenId,
			"token_address": tokenAddress,
		},
		Type: string(types.ERC721Type),
	}
}
