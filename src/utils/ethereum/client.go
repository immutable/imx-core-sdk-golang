package ethereum

import (
	"context"
	"fmt"
	"immutable.com/imx-core-sdk-golang/generated/contracts"
	"immutable.com/imx-core-sdk-golang/signers"
	"math/big"
	"regexp"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	EthereumDefaultGasLimit uint64  = 30000        // Ethereum default gas limit in gas units
	EthereumDefaultGasPrice uint64  = 160000000000 // Ethereum default gas price in wei
	EthereumEGSAPIKey       string  = ""           // Ethereum eth gas station (EGS) API key
	EthereumEGSSpeed        string  = "fast"       // Ethereum eth gas station (EGS) speed
	EthereumGasMultiplier   float64 = 1.5          // Ethereum default gas multiplier
	EthereumMaxGasPrice     uint64  = 300000000000 // Ethereum max gas price in wei
)

var DefaultGasParams = NewDefaultGasParams()

func NewDefaultGasParams() GasParams {
	return GasParams{
		GasLimit:      EthereumDefaultGasLimit,
		GasPrice:      big.NewInt(int64(EthereumDefaultGasPrice)),
		GasMultiplier: EthereumGasMultiplier,
		MaxGasPrice:   big.NewInt(int64(EthereumMaxGasPrice)),
		EGSAPIKey:     EthereumEGSAPIKey,
		EGSSpeed:      EthereumEGSSpeed,
	}
}

type GasParams struct {
	GasLimit      uint64
	GasPrice      *big.Int
	GasMultiplier float64
	MaxGasPrice   *big.Int
	EGSAPIKey     string
	EGSSpeed      string
}

type EtherClient interface {
	AttachRegistrationContract(ctx context.Context, address string) error
	AttachCoreContract(ctx context.Context, address string) error
	NewIERC20Contract(ctx context.Context, address string) (*contracts.IERC20, error)
	NewIERC721Contract(ctx context.Context, address string) (*contracts.IERC721, error)
	IsValidContract(ctx context.Context, address string) error
	EstimateGasPrice(ctx context.Context) (*big.Int, error)
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
	EGSAPIKey                   string
	EGSSpeed                    string
	RegistrationContractAddress ethcommon.Address
	StarkContractAddress        ethcommon.Address
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
		EGSAPIKey:     params.EGSAPIKey,
		EGSSpeed:      params.EGSSpeed,
	}
	return ethClient, nil
}

func (e *Client) AttachRegistrationContract(ctx context.Context, address string) error {
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

func (e *Client) AttachCoreContract(ctx context.Context, address string) error {
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

func (e *Client) validateContract(ctx context.Context, address string) error {
	if err := ValidateEthereumAddress(address); err != nil {
		return err
	}

	if err := e.IsValidContract(ctx, address); err != nil {
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
	keyAddr := l1signer.GetAddress()

	auth := &bind.TransactOpts{
		From: keyAddr,
		Signer: func(address ethcommon.Address, tx *types.Transaction) (*types.Transaction, error) {
			if address != keyAddr {
				return nil, bind.ErrNotAuthorized
			}
			return l1signer.SignTx(tx)
		},
		Context: context.Background(),
	}

	if err := e.setGasPriceAndLimit(ctx, auth); err != nil {
		return nil, err
	}

	return auth, nil
}

func (e *Client) setGasPriceAndLimit(ctx context.Context, auth *bind.TransactOpts) error {
	auth.Value = big.NewInt(0)
	auth.GasLimit = e.GasLimit

	suggestedGasPrice, err := e.EstimateGasPrice(ctx)
	if err != nil {
		return err
	}

	auth.GasPrice = suggestedGasPrice
	return nil
}

// EstimateGasPrice it will return the gas price from ethgasstation or the suggested gas price from the go-ethereum lib
func (e *Client) EstimateGasPrice(ctx context.Context) (*big.Int, error) {
	var gasPrice *big.Int

	ctx, cancel := context.WithTimeout(ctx, time.Second*Timeout)
	defer cancel()

	if e.EGSAPIKey != "" {
		price, err := FetchGasPrice(e.EGSAPIKey, e.EGSSpeed)
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

type ContractValidationError struct {
	Reason string
}

func (err ContractValidationError) Error() string {
	return err.Reason
}

// IsValidContract validates whether the given address is a contract
func (e *Client) IsValidContract(ctx context.Context, address string) error {
	addr := ethcommon.HexToAddress(address)

	bytecode, err := e.Client.CodeAt(ctx, addr, nil) // nil: latest head block

	if err != nil {
		return ContractValidationError{EthClientError}
	}
	if len(bytecode) <= 0 {
		return ContractValidationError{ContractInvalidError}
	}
	return nil
}
