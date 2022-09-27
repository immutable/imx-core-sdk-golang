package common

import (
	"context"
	"encoding/json"
	"log"
	"math/big"
	"os"
	"runtime"
	"strings"

	"github.com/immutable/imx-core-sdk-golang/imx"
	"github.com/immutable/imx-core-sdk-golang/imx/api"
	"github.com/immutable/imx-core-sdk-golang/imx/signers/ethereum"
	"github.com/immutable/imx-core-sdk-golang/imx/signers/stark"
	"github.com/joho/godotenv"
)

func CommonInitialise(configFilePath string) (context.Context, map[string]string, *imx.Client, imx.L1Signer, imx.L2Signer) {
	var envs map[string]string
	envs, err := godotenv.Read(configFilePath)
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	alchemyAPIKey := envs["ALCHEMY_API_KEY"]
	signerPrivateKey := envs["OWNER_ACCOUNT_PRIVATE_KEY"]
	enableDebugLogging := strings.EqualFold(envs["DEBUG_LOGGING"], "true")

	apiConfiguration := api.NewConfiguration()
	// Enable debug logging.
	if enableDebugLogging {
		apiConfiguration.Debug = true
	}

	// Using context value to switch/specify the server before sending request.
	// If nothing is specified, the default server will be used which will be first one in the open api spec list.
	ctx := context.TODO()

	cfg := imx.Config{
		APIConfig:     apiConfiguration,
		AlchemyAPIKey: alchemyAPIKey,
		Environment:   imx.Sandbox,
	}

	c, err := imx.NewClient(&cfg)
	if err != nil {
		log.Panicf("error on NewClient: %v\n", err)
	}
	defer c.EthClient.Close()

	l1signer, err := ethereum.NewSigner(signerPrivateKey, cfg.ChainID)
	if err != nil {
		log.Panicf("error in creating L1Signer: %v\n", err)
	}

	var starkPrivateKey *big.Int
	privateStarkKeyStr := envs["STARK_PRIVATE_KEY"]
	if privateStarkKeyStr != "" {
		var ok bool
		starkPrivateKey, ok = new(big.Int).SetString(privateStarkKeyStr, 10)
		if !ok {
			log.Panicf("error in converting stark private key value from string to big.Int")
		}
	} else {
		starkPrivateKey, err = stark.GenerateKey()
		if err != nil {
			log.Panicf("error in Generating Stark Private Key: %v\n", err)
		}
	}

	l2signer, err := stark.NewSigner(starkPrivateKey)
	if err != nil {
		log.Panicf("error in creating StarkSigner: %v\n", err)
	}

	return ctx, envs, c, l1signer, l2signer
}

func PrettyStruct(data interface{}) (string, error) {
	val, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return "", err
	}
	return string(val), nil
}

func GetCurrentFunctionName() string {
	counter, _, _, success := runtime.Caller(1)

	if !success {
		println("functionName: runtime.Caller: failed")
		os.Exit(1)
	}

	return runtime.FuncForPC(counter).Name()
}
