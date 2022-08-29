package ethereumutil

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"regexp"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/immutable/imx-core-sdk-golang/config"
	"github.com/immutable/imx-core-sdk-golang/generated/contracts"
	"github.com/immutable/imx-core-sdk-golang/signers"
)

type EtherClient interface {
	NewIERC20Contract(ctx context.Context, address string) (*contracts.IERC20, error)
	NewIERC721Contract(ctx context.Context, address string) (*contracts.IERC721, error)
}

// Client to interact with ethereum blockchain
type Client struct {
	Client                      *ethclient.Client
	RegistrationContract        *contracts.Registration
	CoreContract                *contracts.Core
	GasLimit                    uint64
	GasPrice                    *big.Int
	GasMultiplier               float64
	MaxGasPrice                 *big.Int
	EgsApiKey                   string
	EGSSpeed                    string
	RegistrationContractAddress ethcommon.Address
	StarkContractAddress        ethcommon.Address
}

// NewEthereumClientAndAttachContracts is a factory method to create a new ethereum client and attaches registration and core contracts to the client.
func NewEthereumClientAndAttachContracts(
	ctx context.Context,
	cfg *config.Config,
	gasParams GasParams) (*Client, error) {
	ethClient, err := NewEthereumClient(cfg.EthereumClientEndpoint, gasParams)
	if err != nil {
		return nil, err
	}
	if err := ethClient.attachRegistrationContract(ctx, cfg.RegistrationContractAddress); err != nil {
		return nil, err
	}
	if err := ethClient.attachCoreContract(ctx, cfg.StarkContractAddress); err != nil {
		return nil, err
	}
	return ethClient, nil
}

// NewEthereumClient returns a new ethereum client
func NewEthereumClient(url string, params GasParams) (*Client, error) {
	client, err := ethclient.Dial(url)
	if err != nil {
		return nil, err
	}

	ethClient := &Client{
		Client:        client,
		GasLimit:      params.GasLimit,
		GasPrice:      params.GasPrice,
		GasMultiplier: params.GasMultiplier,
		MaxGasPrice:   params.MaxGasPrice,
		EgsApiKey:     params.EgsApiKey,
		EGSSpeed:      params.EGSSpeed,
	}
	return ethClient, nil
}

func (e *Client) NewIERC20Contract(ctx context.Context, address string) (*contracts.IERC20, error) {
	if err := e.validateContract(ctx, address); err != nil {
		return nil, err
	}

	return contracts.NewIERC20(ethcommon.HexToAddress(address), e.Client)
}

func (e *Client) NewIERC721Contract(ctx context.Context, address string) (*contracts.IERC721, error) {
	if err := e.validateContract(ctx, address); err != nil {
		return nil, err
	}

	return contracts.NewIERC721(ethcommon.HexToAddress(address), e.Client)
}

func (e *Client) attachRegistrationContract(ctx context.Context, address string) error {
	if err := e.validateContract(ctx, address); err != nil {
		return err
	}

	e.RegistrationContractAddress = ethcommon.HexToAddress(address)
	client, err := contracts.NewRegistration(e.RegistrationContractAddress, e.Client)

	if err != nil {
		return err
	}
	e.RegistrationContract = client
	return nil
}

func (e *Client) attachCoreContract(ctx context.Context, address string) error {
	if err := e.validateContract(ctx, address); err != nil {
		return err
	}

	e.StarkContractAddress = ethcommon.HexToAddress(address)
	client, err := contracts.NewCore(e.StarkContractAddress, e.Client)

	if err != nil {
		return err
	}
	e.CoreContract = client
	return nil
}

func (e *Client) validateContract(ctx context.Context, address string) error {
	if err := ValidateEthereumAddress(address); err != nil {
		return err
	}

	if err := e.isValidContract(ctx, address); err != nil {
		return err
	}
	return nil
}

func ValidateEthereumAddress(address string) error {
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	if !re.MatchString(address) {
		return fmt.Errorf("invalid ethereum address")
	}
	return nil
}

func (e *Client) BuildTransactOpts(ctx context.Context, l1signer signers.L1Signer) (*bind.TransactOpts, error) {
	keyAddr := ethcommon.HexToAddress(l1signer.GetAddress())

	auth := &bind.TransactOpts{
		From: keyAddr,
		Signer: func(address ethcommon.Address, tx *types.Transaction) (*types.Transaction, error) {
			if address != keyAddr {
				return nil, bind.ErrNotAuthorized
			}
			return l1signer.SignTx(tx)
		},
		Context: ctx,
	}

	if err := e.setGasPriceAndLimit(ctx, auth); err != nil {
		return nil, err
	}

	return auth, nil
}

func (e *Client) setGasPriceAndLimit(ctx context.Context, auth *bind.TransactOpts) error {
	auth.Value = big.NewInt(0)
	auth.GasLimit = e.GasLimit

	suggestedGasPrice, err := e.estimateGasPrice(ctx)
	if err != nil {
		return err
	}

	auth.GasPrice = suggestedGasPrice
	return nil
}

// estimateGasPrice it will return the gas price from eth gas station or the suggested gas price from the go-ethereum lib
func (e *Client) estimateGasPrice(ctx context.Context) (*big.Int, error) {
	var gasPrice *big.Int

	ctx, cancel := context.WithTimeout(ctx, time.Second*Timeout)
	defer cancel()

	if e.EgsApiKey != "" {
		price, err := fetchGasPrice(e.EgsApiKey, e.EGSSpeed)
		if err != nil {
			// log? - ignore for now and use node price
		} else {
			gasPrice = price
		}
	}

	if gasPrice == nil {
		// use recommended gas price from node
		nodePriceEstimate, err := e.Client.SuggestGasPrice(ctx)
		if err != nil {
			return nil, err
		} else {
			gasPrice = multiplyGasPrice(nodePriceEstimate, new(big.Float).SetFloat64(e.GasMultiplier))
		}
	}
	if gasPrice.Cmp(e.MaxGasPrice) == 1 {
		return e.MaxGasPrice, nil
	}
	return gasPrice, nil // gasPrice in wei
}

func multiplyGasPrice(gasEstimate *big.Int, gasMultiplier *big.Float) *big.Int {
	gasPrice := big.NewInt(0)
	gasEstimateFloat := big.NewFloat(0).SetInt(gasEstimate)
	result := gasEstimateFloat.Mul(gasEstimateFloat, gasMultiplier)
	result.Int(gasPrice)
	return gasPrice
}

const (
	ContractInvalidError = "contract_invalid_error"
	EthClientError       = "ethereum_client_error"
)

// isValidContract validates whether the given address is a contract
func (e *Client) isValidContract(ctx context.Context, address string) error {
	addr := ethcommon.HexToAddress(address)

	bytecode, err := e.Client.CodeAt(ctx, addr, nil) // nil: latest head block

	if err != nil {
		return errors.New(EthClientError)
	}
	if len(bytecode) == 0 {
		return errors.New(ContractInvalidError)
	}
	return nil
}
