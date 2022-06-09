// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// RegistrationMetaData contains all meta data concerning the Registration contract.
var RegistrationMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractCore\",\"name\":\"_imx\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"imx\",\"outputs\":[{\"internalType\":\"contractCore\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"}],\"name\":\"isRegistered\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ethKey\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"registerAndDepositNft\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ethKey\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"}],\"name\":\"registerAndWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ethKey\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"registerAndWithdrawNft\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ethKey\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"registerAndWithdrawNftTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ethKey\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"registerAndWithdrawTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ethKey\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"mintingBlob\",\"type\":\"bytes\"}],\"name\":\"regsiterAndWithdrawAndMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// RegistrationABI is the input ABI used to generate the binding from.
// Deprecated: Use RegistrationMetaData.ABI instead.
var RegistrationABI = RegistrationMetaData.ABI

// Registration is an auto generated Go binding around an Ethereum contract.
type Registration struct {
	RegistrationCaller     // Read-only binding to the contract
	RegistrationTransactor // Write-only binding to the contract
	RegistrationFilterer   // Log filterer for contract events
}

// RegistrationCaller is an auto generated read-only Go binding around an Ethereum contract.
type RegistrationCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistrationTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RegistrationTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistrationFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RegistrationFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistrationSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RegistrationSession struct {
	Contract     *Registration     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RegistrationCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RegistrationCallerSession struct {
	Contract *RegistrationCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// RegistrationTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RegistrationTransactorSession struct {
	Contract     *RegistrationTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// RegistrationRaw is an auto generated low-level Go binding around an Ethereum contract.
type RegistrationRaw struct {
	Contract *Registration // Generic contract binding to access the raw methods on
}

// RegistrationCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RegistrationCallerRaw struct {
	Contract *RegistrationCaller // Generic read-only contract binding to access the raw methods on
}

// RegistrationTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RegistrationTransactorRaw struct {
	Contract *RegistrationTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRegistration creates a new instance of Registration, bound to a specific deployed contract.
func NewRegistration(address common.Address, backend bind.ContractBackend) (*Registration, error) {
	contract, err := bindRegistration(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Registration{RegistrationCaller: RegistrationCaller{contract: contract}, RegistrationTransactor: RegistrationTransactor{contract: contract}, RegistrationFilterer: RegistrationFilterer{contract: contract}}, nil
}

// NewRegistrationCaller creates a new read-only instance of Registration, bound to a specific deployed contract.
func NewRegistrationCaller(address common.Address, caller bind.ContractCaller) (*RegistrationCaller, error) {
	contract, err := bindRegistration(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RegistrationCaller{contract: contract}, nil
}

// NewRegistrationTransactor creates a new write-only instance of Registration, bound to a specific deployed contract.
func NewRegistrationTransactor(address common.Address, transactor bind.ContractTransactor) (*RegistrationTransactor, error) {
	contract, err := bindRegistration(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RegistrationTransactor{contract: contract}, nil
}

// NewRegistrationFilterer creates a new log filterer instance of Registration, bound to a specific deployed contract.
func NewRegistrationFilterer(address common.Address, filterer bind.ContractFilterer) (*RegistrationFilterer, error) {
	contract, err := bindRegistration(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RegistrationFilterer{contract: contract}, nil
}

// bindRegistration binds a generic wrapper to an already deployed contract.
func bindRegistration(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RegistrationABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Registration *RegistrationRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Registration.Contract.RegistrationCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Registration *RegistrationRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registration.Contract.RegistrationTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Registration *RegistrationRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Registration.Contract.RegistrationTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Registration *RegistrationCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Registration.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Registration *RegistrationTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registration.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Registration *RegistrationTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Registration.Contract.contract.Transact(opts, method, params...)
}

// Imx is a free data retrieval call binding the contract method 0x0f08025f.
//
// Solidity: function imx() view returns(address)
func (_Registration *RegistrationCaller) Imx(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Registration.contract.Call(opts, &out, "imx")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Imx is a free data retrieval call binding the contract method 0x0f08025f.
//
// Solidity: function imx() view returns(address)
func (_Registration *RegistrationSession) Imx() (common.Address, error) {
	return _Registration.Contract.Imx(&_Registration.CallOpts)
}

// Imx is a free data retrieval call binding the contract method 0x0f08025f.
//
// Solidity: function imx() view returns(address)
func (_Registration *RegistrationCallerSession) Imx() (common.Address, error) {
	return _Registration.Contract.Imx(&_Registration.CallOpts)
}

// IsRegistered is a free data retrieval call binding the contract method 0x579a6988.
//
// Solidity: function isRegistered(uint256 starkKey) view returns(bool)
func (_Registration *RegistrationCaller) IsRegistered(opts *bind.CallOpts, starkKey *big.Int) (bool, error) {
	var out []interface{}
	err := _Registration.contract.Call(opts, &out, "isRegistered", starkKey)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRegistered is a free data retrieval call binding the contract method 0x579a6988.
//
// Solidity: function isRegistered(uint256 starkKey) view returns(bool)
func (_Registration *RegistrationSession) IsRegistered(starkKey *big.Int) (bool, error) {
	return _Registration.Contract.IsRegistered(&_Registration.CallOpts, starkKey)
}

// IsRegistered is a free data retrieval call binding the contract method 0x579a6988.
//
// Solidity: function isRegistered(uint256 starkKey) view returns(bool)
func (_Registration *RegistrationCallerSession) IsRegistered(starkKey *big.Int) (bool, error) {
	return _Registration.Contract.IsRegistered(&_Registration.CallOpts, starkKey)
}

// RegisterAndDepositNft is a paid mutator transaction binding the contract method 0x4280d50a.
//
// Solidity: function registerAndDepositNft(address ethKey, uint256 starkKey, bytes signature, uint256 assetType, uint256 vaultId, uint256 tokenId) returns()
func (_Registration *RegistrationTransactor) RegisterAndDepositNft(opts *bind.TransactOpts, ethKey common.Address, starkKey *big.Int, signature []byte, assetType *big.Int, vaultId *big.Int, tokenId *big.Int) (*types.Transaction, error) {
	return _Registration.contract.Transact(opts, "registerAndDepositNft", ethKey, starkKey, signature, assetType, vaultId, tokenId)
}

// RegisterAndDepositNft is a paid mutator transaction binding the contract method 0x4280d50a.
//
// Solidity: function registerAndDepositNft(address ethKey, uint256 starkKey, bytes signature, uint256 assetType, uint256 vaultId, uint256 tokenId) returns()
func (_Registration *RegistrationSession) RegisterAndDepositNft(ethKey common.Address, starkKey *big.Int, signature []byte, assetType *big.Int, vaultId *big.Int, tokenId *big.Int) (*types.Transaction, error) {
	return _Registration.Contract.RegisterAndDepositNft(&_Registration.TransactOpts, ethKey, starkKey, signature, assetType, vaultId, tokenId)
}

// RegisterAndDepositNft is a paid mutator transaction binding the contract method 0x4280d50a.
//
// Solidity: function registerAndDepositNft(address ethKey, uint256 starkKey, bytes signature, uint256 assetType, uint256 vaultId, uint256 tokenId) returns()
func (_Registration *RegistrationTransactorSession) RegisterAndDepositNft(ethKey common.Address, starkKey *big.Int, signature []byte, assetType *big.Int, vaultId *big.Int, tokenId *big.Int) (*types.Transaction, error) {
	return _Registration.Contract.RegisterAndDepositNft(&_Registration.TransactOpts, ethKey, starkKey, signature, assetType, vaultId, tokenId)
}

// RegisterAndWithdraw is a paid mutator transaction binding the contract method 0xea864adf.
//
// Solidity: function registerAndWithdraw(address ethKey, uint256 starkKey, bytes signature, uint256 assetType) returns()
func (_Registration *RegistrationTransactor) RegisterAndWithdraw(opts *bind.TransactOpts, ethKey common.Address, starkKey *big.Int, signature []byte, assetType *big.Int) (*types.Transaction, error) {
	return _Registration.contract.Transact(opts, "registerAndWithdraw", ethKey, starkKey, signature, assetType)
}

// RegisterAndWithdraw is a paid mutator transaction binding the contract method 0xea864adf.
//
// Solidity: function registerAndWithdraw(address ethKey, uint256 starkKey, bytes signature, uint256 assetType) returns()
func (_Registration *RegistrationSession) RegisterAndWithdraw(ethKey common.Address, starkKey *big.Int, signature []byte, assetType *big.Int) (*types.Transaction, error) {
	return _Registration.Contract.RegisterAndWithdraw(&_Registration.TransactOpts, ethKey, starkKey, signature, assetType)
}

// RegisterAndWithdraw is a paid mutator transaction binding the contract method 0xea864adf.
//
// Solidity: function registerAndWithdraw(address ethKey, uint256 starkKey, bytes signature, uint256 assetType) returns()
func (_Registration *RegistrationTransactorSession) RegisterAndWithdraw(ethKey common.Address, starkKey *big.Int, signature []byte, assetType *big.Int) (*types.Transaction, error) {
	return _Registration.Contract.RegisterAndWithdraw(&_Registration.TransactOpts, ethKey, starkKey, signature, assetType)
}

// RegisterAndWithdrawNft is a paid mutator transaction binding the contract method 0x352eb84c.
//
// Solidity: function registerAndWithdrawNft(address ethKey, uint256 starkKey, bytes signature, uint256 assetType, uint256 tokenId) returns()
func (_Registration *RegistrationTransactor) RegisterAndWithdrawNft(opts *bind.TransactOpts, ethKey common.Address, starkKey *big.Int, signature []byte, assetType *big.Int, tokenId *big.Int) (*types.Transaction, error) {
	return _Registration.contract.Transact(opts, "registerAndWithdrawNft", ethKey, starkKey, signature, assetType, tokenId)
}

// RegisterAndWithdrawNft is a paid mutator transaction binding the contract method 0x352eb84c.
//
// Solidity: function registerAndWithdrawNft(address ethKey, uint256 starkKey, bytes signature, uint256 assetType, uint256 tokenId) returns()
func (_Registration *RegistrationSession) RegisterAndWithdrawNft(ethKey common.Address, starkKey *big.Int, signature []byte, assetType *big.Int, tokenId *big.Int) (*types.Transaction, error) {
	return _Registration.Contract.RegisterAndWithdrawNft(&_Registration.TransactOpts, ethKey, starkKey, signature, assetType, tokenId)
}

// RegisterAndWithdrawNft is a paid mutator transaction binding the contract method 0x352eb84c.
//
// Solidity: function registerAndWithdrawNft(address ethKey, uint256 starkKey, bytes signature, uint256 assetType, uint256 tokenId) returns()
func (_Registration *RegistrationTransactorSession) RegisterAndWithdrawNft(ethKey common.Address, starkKey *big.Int, signature []byte, assetType *big.Int, tokenId *big.Int) (*types.Transaction, error) {
	return _Registration.Contract.RegisterAndWithdrawNft(&_Registration.TransactOpts, ethKey, starkKey, signature, assetType, tokenId)
}

// RegisterAndWithdrawNftTo is a paid mutator transaction binding the contract method 0x1259cc6c.
//
// Solidity: function registerAndWithdrawNftTo(address ethKey, uint256 starkKey, bytes signature, uint256 assetType, uint256 tokenId, address recipient) returns()
func (_Registration *RegistrationTransactor) RegisterAndWithdrawNftTo(opts *bind.TransactOpts, ethKey common.Address, starkKey *big.Int, signature []byte, assetType *big.Int, tokenId *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Registration.contract.Transact(opts, "registerAndWithdrawNftTo", ethKey, starkKey, signature, assetType, tokenId, recipient)
}

// RegisterAndWithdrawNftTo is a paid mutator transaction binding the contract method 0x1259cc6c.
//
// Solidity: function registerAndWithdrawNftTo(address ethKey, uint256 starkKey, bytes signature, uint256 assetType, uint256 tokenId, address recipient) returns()
func (_Registration *RegistrationSession) RegisterAndWithdrawNftTo(ethKey common.Address, starkKey *big.Int, signature []byte, assetType *big.Int, tokenId *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Registration.Contract.RegisterAndWithdrawNftTo(&_Registration.TransactOpts, ethKey, starkKey, signature, assetType, tokenId, recipient)
}

// RegisterAndWithdrawNftTo is a paid mutator transaction binding the contract method 0x1259cc6c.
//
// Solidity: function registerAndWithdrawNftTo(address ethKey, uint256 starkKey, bytes signature, uint256 assetType, uint256 tokenId, address recipient) returns()
func (_Registration *RegistrationTransactorSession) RegisterAndWithdrawNftTo(ethKey common.Address, starkKey *big.Int, signature []byte, assetType *big.Int, tokenId *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Registration.Contract.RegisterAndWithdrawNftTo(&_Registration.TransactOpts, ethKey, starkKey, signature, assetType, tokenId, recipient)
}

// RegisterAndWithdrawTo is a paid mutator transaction binding the contract method 0x4627d598.
//
// Solidity: function registerAndWithdrawTo(address ethKey, uint256 starkKey, bytes signature, uint256 assetType, address recipient) returns()
func (_Registration *RegistrationTransactor) RegisterAndWithdrawTo(opts *bind.TransactOpts, ethKey common.Address, starkKey *big.Int, signature []byte, assetType *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Registration.contract.Transact(opts, "registerAndWithdrawTo", ethKey, starkKey, signature, assetType, recipient)
}

// RegisterAndWithdrawTo is a paid mutator transaction binding the contract method 0x4627d598.
//
// Solidity: function registerAndWithdrawTo(address ethKey, uint256 starkKey, bytes signature, uint256 assetType, address recipient) returns()
func (_Registration *RegistrationSession) RegisterAndWithdrawTo(ethKey common.Address, starkKey *big.Int, signature []byte, assetType *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Registration.Contract.RegisterAndWithdrawTo(&_Registration.TransactOpts, ethKey, starkKey, signature, assetType, recipient)
}

// RegisterAndWithdrawTo is a paid mutator transaction binding the contract method 0x4627d598.
//
// Solidity: function registerAndWithdrawTo(address ethKey, uint256 starkKey, bytes signature, uint256 assetType, address recipient) returns()
func (_Registration *RegistrationTransactorSession) RegisterAndWithdrawTo(ethKey common.Address, starkKey *big.Int, signature []byte, assetType *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Registration.Contract.RegisterAndWithdrawTo(&_Registration.TransactOpts, ethKey, starkKey, signature, assetType, recipient)
}

// RegsiterAndWithdrawAndMint is a paid mutator transaction binding the contract method 0x0a9c3beb.
//
// Solidity: function regsiterAndWithdrawAndMint(address ethKey, uint256 starkKey, bytes signature, uint256 assetType, bytes mintingBlob) returns()
func (_Registration *RegistrationTransactor) RegsiterAndWithdrawAndMint(opts *bind.TransactOpts, ethKey common.Address, starkKey *big.Int, signature []byte, assetType *big.Int, mintingBlob []byte) (*types.Transaction, error) {
	return _Registration.contract.Transact(opts, "regsiterAndWithdrawAndMint", ethKey, starkKey, signature, assetType, mintingBlob)
}

// RegsiterAndWithdrawAndMint is a paid mutator transaction binding the contract method 0x0a9c3beb.
//
// Solidity: function regsiterAndWithdrawAndMint(address ethKey, uint256 starkKey, bytes signature, uint256 assetType, bytes mintingBlob) returns()
func (_Registration *RegistrationSession) RegsiterAndWithdrawAndMint(ethKey common.Address, starkKey *big.Int, signature []byte, assetType *big.Int, mintingBlob []byte) (*types.Transaction, error) {
	return _Registration.Contract.RegsiterAndWithdrawAndMint(&_Registration.TransactOpts, ethKey, starkKey, signature, assetType, mintingBlob)
}

// RegsiterAndWithdrawAndMint is a paid mutator transaction binding the contract method 0x0a9c3beb.
//
// Solidity: function regsiterAndWithdrawAndMint(address ethKey, uint256 starkKey, bytes signature, uint256 assetType, bytes mintingBlob) returns()
func (_Registration *RegistrationTransactorSession) RegsiterAndWithdrawAndMint(ethKey common.Address, starkKey *big.Int, signature []byte, assetType *big.Int, mintingBlob []byte) (*types.Transaction, error) {
	return _Registration.Contract.RegsiterAndWithdrawAndMint(&_Registration.TransactOpts, ethKey, starkKey, signature, assetType, mintingBlob)
}
