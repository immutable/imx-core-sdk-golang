package imx

import (
	"context"
	"errors"
	"fmt"
	"regexp"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/immutable/imx-core-sdk-golang/contracts"
)

const (
	ContractInvalidError = "contract_invalid_error"
	EthClientError       = "ethereum_client_error"
)

func (c *Client) newIERC20Contract(ctx context.Context, address string) (*contracts.IERC20, error) {
	if err := c.validateContract(ctx, address); err != nil {
		return nil, err
	}
	return contracts.NewIERC20(common.HexToAddress(address), c.EthClient)
}

func (c *Client) newIERC721Contract(ctx context.Context, address string) (*contracts.IERC721, error) {
	if err := c.validateContract(ctx, address); err != nil {
		return nil, err
	}
	return contracts.NewIERC721(common.HexToAddress(address), c.EthClient)
}

func (c *Client) attachRegistrationContract(ctx context.Context) error {
	if err := c.validateContract(ctx, c.Environment.RegistrationContractAddress); err != nil {
		return err
	}
	client, err := contracts.NewRegistration(common.HexToAddress(c.Environment.CoreContractAddress), c.EthClient)
	if err != nil {
		return err
	}
	c.RegistrationContract = client
	return nil
}

func (c *Client) attachCoreContract(ctx context.Context) error {
	if err := c.validateContract(ctx, c.Environment.CoreContractAddress); err != nil {
		return err
	}
	client, err := contracts.NewCore(common.HexToAddress(c.Environment.CoreContractAddress), c.EthClient)
	if err != nil {
		return err
	}
	c.CoreContract = client
	return nil
}

func (c *Client) validateContract(ctx context.Context, address string) error {
	if err := validateEthereumAddress(address); err != nil {
		return err
	}

	if err := c.isValidContract(ctx, address); err != nil {
		return err
	}
	return nil
}

func validateEthereumAddress(address string) error {
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	if !re.MatchString(address) {
		return fmt.Errorf("invalid ethereum address")
	}
	return nil
}

func (c *Client) buildTransactOpts(ctx context.Context, l1signer L1Signer, overrides *bind.TransactOpts) *bind.TransactOpts {
	keyAddr := common.HexToAddress(l1signer.GetAddress())
	if overrides == nil {
		overrides = new(bind.TransactOpts)
	}
	overrides.From = keyAddr
	overrides.Context = ctx
	overrides.Signer = func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
		if address != keyAddr {
			return nil, bind.ErrNotAuthorized
		}
		return l1signer.SignTx(tx)
	}
	return overrides
}

// isValidContract validates whether the given address is a contract
func (c *Client) isValidContract(ctx context.Context, address string) error {
	bytecode, err := c.EthClient.CodeAt(ctx, common.HexToAddress(address), nil) // nil: latest head block
	if err != nil {
		return errors.New(EthClientError)
	}
	if len(bytecode) == 0 {
		return errors.New(ContractInvalidError)
	}
	return nil
}
