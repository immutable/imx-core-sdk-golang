package ethereum

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"immutable.com/imx-core-sdk/contracts"
	"immutable.com/imx-core-sdk/utils"
	"math/big"
	"regexp"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type ClientParams struct {
	URL           string
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
	AttachIERC20Contract(ctx context.Context, address string) error
	AttachIERC721Contract(ctx context.Context, address string) error
	LoadPrivateKey(privateKey string) error
	IsValidContract(ctx context.Context, address string) error
	EstimateGasPrice(ctx context.Context) (*big.Int, error)
}

// Client to interact with ethereum blockchain
type Client struct {
	Client                      *ethclient.Client
	RegistrationContract        *contracts.Registration
	CoreContract                *contracts.Core
	IERC20Contract              *contracts.IERC20
	IERC721Contract             *contracts.IERC721
	PrivateKey                  *ecdsa.PrivateKey
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
func NewEthereumClient(params ClientParams) (*Client, error) {
	client, err := ethclient.Dial(params.URL)
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

func (e *Client) AttachIERC20Contract(ctx context.Context, address string) error {
	if err := e.validateContract(ctx, address); err != nil {
		return err
	}

	client, err := contracts.NewIERC20(ethcommon.HexToAddress(address), e.Client)

	if err != nil {
		return err
	}
	e.IERC20Contract = client
	return nil
}

func (e *Client) AttachIERC721Contract(ctx context.Context, address string) error {
	if err := e.validateContract(ctx, address); err != nil {
		return err
	}

	client, err := contracts.NewIERC721(ethcommon.HexToAddress(address), e.Client)

	if err != nil {
		return err
	}
	e.IERC721Contract = client
	return nil
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

// LoadPrivateKey loads a new private key for the ethereum client
func (e *Client) LoadPrivateKey(privateKey string) error {
	privateKey, err := utils.RemoveHex(privateKey)
	if err != nil {
		return err
	}
	privKey, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return err
	}
	e.PrivateKey = privKey
	return nil
}

func (e *Client) BuildTransactOpts(ctx context.Context) (*bind.TransactOpts, error) {
	chainID, err := e.Client.ChainID(ctx)
	if err != nil {
		return nil, err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(e.PrivateKey, chainID)
	if err != nil {
		return nil, err
	}

	if err := e.setGasPriceAndLimit(ctx, auth); err != nil {
		return nil, err
	}

	return auth, nil
}

func (e *Client) GetAddress() ethcommon.Address {
	return crypto.PubkeyToAddress(e.PrivateKey.PublicKey)
}

func (e *Client) SignTransaction(ctx context.Context, tx *types.Transaction) (*types.Transaction, error) {
	chainID, err := e.Client.ChainID(ctx)
	if err != nil {
		return nil, err
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), e.PrivateKey)
	if err != nil {
		return nil, err
	}

	return signedTx, nil
}

func (e *Client) Signer(ctx context.Context) (signer func(ethcommon.Address, *types.Transaction) (*types.Transaction, error)) {
	signer = func(address ethcommon.Address, transaction *types.Transaction) (*types.Transaction, error) {
		return e.SignTransaction(ctx, transaction)
	}

	return signer
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
