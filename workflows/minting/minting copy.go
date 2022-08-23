package minting

// import (
// 	"context"
// 	"crypto/ecdsa"
// 	"encoding/hex"
// 	"encoding/json"
// 	"fmt"

// 	"github.com/ethereum/go-ethereum/crypto"

// 	"github.com/ethereum/go-ethereum/common/hexutil"
// 	"github.com/immutable/imx-core-sdk-golang/generated/api/client"
// 	"github.com/immutable/imx-core-sdk-golang/generated/api/client/mints"
// 	"github.com/immutable/imx-core-sdk-golang/generated/api/models"
// 	"github.com/immutable/imx-core-sdk-golang/signers"
// )

// // // UnsignedMintRequest is MintRequest with out the AuthSignature.
// // type UnsignedMintRequest struct {

// // 	// minting contract
// // 	// Required: true
// // 	ContractAddress *string `json:"contract_address"`

// // 	// Global contract-level royalty fees
// // 	Royalties []*MintFee `json:"royalties"`

// // 	// Users to mint to
// // 	// Required: true
// // 	Users []*MintUser `json:"users"`
// // }

// func prefixMessageDigest(hashBytes []byte) []byte {
// 	ext := []byte("\x19Ethereum Signed Message:\n")
// 	len := []byte(fmt.Sprintf("%d", len(hashBytes)))
// 	return append(ext, append(len, hashBytes...)...)
// }

// func HashWithPrefix(data []byte) []byte {
// 	hash := crypto.Keccak256Hash(data)
// 	hashBytes := []byte(hash.String())
// 	return crypto.Keccak256Hash(prefixMessageDigest(hashBytes)).Bytes()
// }

// func encodeMintBody(mint models.MintRequest) ([]byte, error) {
// 	out, err := json.Marshal(mint)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return out, nil
// }

// func Sign(prv *ecdsa.PrivateKey, mint models.MintRequest) (string, error) {
// 	data, err := encodeMintBody(mint)
// 	if err != nil {
// 		return "", err
// 	}
// 	hash := crypto.Keccak256Hash(data).Bytes()
// 	sig, err := crypto.Sign(hash, prv)
// 	if err != nil {
// 		return "", err
// 	}
// 	return hexutil.Encode(sig), nil
// }

// func GetMintSigner(mint models.MintRequest) (string, error) {
// 	emptyStr := ""
// 	authSigBytes, err := hexutil.Decode(*mint.AuthSignature)
// 	if err != nil {
// 		return emptyStr, err
// 	}

// 	mint.AuthSignature = &emptyStr
// 	data, err := encodeMintBody(mint)
// 	if err != nil {
// 		return emptyStr, err
// 	}
// 	hash := crypto.Keccak256Hash(data).Bytes()
// 	sigPublicKey, err := crypto.Ecrecover(hash, authSigBytes)
// 	if err != nil {
// 		return emptyStr, err
// 	}
// 	return hexutil.Encode(sigPublicKey), nil
// }

// var signerPrivateKey string = "4b56f5ce0ebaf6665a2712b4b0cbcafe42e7a0661091d67776547870a3e3d1a5"

// // Encode encodes b as a hex string with 0x prefix.
// func Encode(b []byte) string {
// 	enc := make([]byte, len(b)*2+2)
// 	copy(enc, "0x")
// 	hex.Encode(enc[2:], b)
// 	return string(enc)
// }

// func MintTokensWorkflow(ctx context.Context,
// 	api *client.ImmutableXAPI,
// 	l1signer signers.L1Signer,
// 	request []*models.MintRequest) (*models.MintTokensResponse, error) {

// 	for _, eachMintRequst := range request {
// 		mintRequestInBytes, err := eachMintRequst.MarshalBinary()
// 		if err != nil {
// 			return nil, fmt.Errorf("error marshaling MintRequest: %v", err)
// 		}
// 		authSignatureInBytes, err := l1signer.SignMessage(string(mintRequestInBytes))
// 		if err != nil {
// 			return nil, fmt.Errorf("error generating signature from MintRequest data: %v", err)
// 		}
// 		authSignatureEncodedInHex := hexutil.Encode(authSignatureInBytes)
// 		eachMintRequst.AuthSignature = &authSignatureEncodedInHex
// 		if err != nil {
// 			return nil, fmt.Errorf("error marshaling MintRequest: %v", err)
// 		}
// 	}
// 	//privateKey, err := coreutils.TrimHexPrefix(signerPrivateKey)
// 	//if err != nil {
// 	//	return nil, err
// 	//}
// 	//privateKeyInEcdsa, err := crypto.HexToECDSA(privateKey)
// 	//if err != nil {
// 	//	return nil, err
// 	//}
// 	//
// 	//for _, eachMintRequst := range request {
// 	//	sig, err := Sign(privateKeyInEcdsa, *eachMintRequst)
// 	//	if err != nil {
// 	//		return nil, err
// 	//	}
// 	//	eachMintRequst.AuthSignature = &sig
// 	//
// 	//	signer, err := GetMintSigner(*eachMintRequst)
// 	//	log.Info(signer)
// 	//	publicKey := l1signer.GetAddress()
// 	//	log.Info(publicKey)
// 	//}
// 	mintTokenParams := mints.NewMintTokensParamsWithContext(ctx)
// 	mintTokenParams.SetMintTokensRequestV2(request)
// 	mintTokensOk, err := api.Mints.MintTokens(mintTokenParams)
// 	if err != nil {
// 		return nil, fmt.Errorf("error when calling `api.Mints.MintTokens`: %v", err)
// 	}
// 	return mintTokensOk.GetPayload(), nil
// }
