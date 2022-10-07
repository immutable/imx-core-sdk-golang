package imx

import (
	"fmt"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

// CreateSignatures accepts signable data and signers and returns signatures.
// signableMessage is signed by l1signer producing ethSignature.
// payloadHash is signed by l2signer producing starkSignature.
func createSignatures(signableMessage, payloadHash *string, l1signer L1Signer, l2signer L2Signer) (ethSignature, starkSignature string, err error) {
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

// generateIMXAuthorisationHeaders gets the Unix epoch timestamp and its signature signed with given ETH key.
// It is required for performing some of the administrative level tasks.
func generateIMXAuthorisationHeaders(l1signer L1Signer) (timestamp, signature string, err error) {
	timestamp = strconv.FormatInt(time.Now().Unix(), 10)
	signedTimestamp, err := l1signer.SignMessage(timestamp)
	if err != nil {
		return "", "", err
	}
	signature = hexutil.Encode(signedTimestamp)
	return timestamp, signature, nil
}
