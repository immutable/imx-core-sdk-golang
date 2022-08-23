package burn

import (
	"context"
	"encoding/json"
	"log"
	"strconv"

	"github.com/immutable/imx-core-sdk-golang/examples/workflows/utils"
	"github.com/immutable/imx-core-sdk-golang/generated/api"
	"github.com/immutable/imx-core-sdk-golang/signers"
	"github.com/immutable/imx-core-sdk-golang/tokens"
	burnWorkflow "github.com/immutable/imx-core-sdk-golang/workflows/burn"
)

func DemoBurnWorkflow(ctx context.Context, apiClient *api.APIClient, l1signer signers.L1Signer, l2signer signers.L2Signer) {
	log.Println("-------------------------------------------------------")
	log.Printf("Running %s", utils.GetCurrentFunctionName())

	signableToken := tokens.NewSignableTokenEth()
	burnRequest := burnWorkflow.GetSignableBurnRequest{
		Amount: "100000000",
		Sender: l1signer.GetAddress(),
		Token:  signableToken,
	}

	/*
		signableToken := tokens.NewSignableTokenERC20(18, "0x73f99ca65b1a0aef2d4591b1b543d789860851bf")
		ftcToBurn, _ := converters.ToDenomination(".025", 18)
		burnRequest := burnWorkflow.GetSignableBurnRequest{
			 Amount: ftcToBurn.String(),
			 Sender: l1signer.GetAddress(),
			 Token:  signableToken,
		}
	*/

	/*
		signableToken := tokens.NewSignableTokenERC721("token id to burn", "token contract address")
		burnRequest := burnWorkflow.GetSignableBurnRequest{
			 Amount: "1",
			 Sender: l1signer.GetAddress(),
			 Token:  signableToken,
		}
	*/

	burnResponse, err := burnWorkflow.Burn(ctx, apiClient, l1signer, l2signer, burnRequest)
	if err != nil {
		log.Panicf("error calling burn workflow: %v", err)
	}
	val, _ := json.MarshalIndent(burnResponse, "", "  ")
	log.Printf("Burn response:\n%s\n", val)

	transferID := strconv.FormatInt(int64(burnResponse.TransferId), 10) // GetBurn method requires transferID as a string
	transfer, err := burnWorkflow.GetBurn(ctx, apiClient, transferID)
	if err != nil {
		log.Panicf("error calling GetBurn workflow: %v", err)
	}
	val, _ = json.MarshalIndent(transfer, "", "  ")
	log.Printf("GetBurn response:\n%s\n", val)

	log.Printf("Running %s completed.", utils.GetCurrentFunctionName())
	log.Println("-------------------------------------------------------")
}
