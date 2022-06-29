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

// CoreMetaData contains all meta data concerning the Core contract.
var CoreMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"depositorEthKey\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonQuantizedAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quantizedAmount\",\"type\":\"uint256\"}],\"name\":\"LogDeposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"}],\"name\":\"LogDepositCancel\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonQuantizedAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quantizedAmount\",\"type\":\"uint256\"}],\"name\":\"LogDepositCancelReclaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"}],\"name\":\"LogDepositNftCancelReclaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"}],\"name\":\"LogFullWithdrawalRequest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ownerKey\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonQuantizedAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quantizedAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"}],\"name\":\"LogMintWithdrawalPerformed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ownerKey\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quantizedAmount\",\"type\":\"uint256\"}],\"name\":\"LogMintableWithdrawalAllowed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"depositorEthKey\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"}],\"name\":\"LogNftDeposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ownerKey\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"}],\"name\":\"LogNftWithdrawalAllowed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ownerKey\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"LogNftWithdrawalPerformed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sequenceNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"batchId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vaultRoot\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"orderRoot\",\"type\":\"uint256\"}],\"name\":\"LogRootUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"stateTransitionFact\",\"type\":\"bytes32\"}],\"name\":\"LogStateTransitionFact\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"ethKey\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"quantizedAmountChange\",\"type\":\"int256\"}],\"name\":\"LogVaultBalanceChangeApplied\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ownerKey\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonQuantizedAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quantizedAmount\",\"type\":\"uint256\"}],\"name\":\"LogWithdrawalAllowed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ownerKey\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonQuantizedAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quantizedAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"LogWithdrawalPerformed\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"announceAvailabilityVerifierRemovalIntent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"announceVerifierRemovalIntent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantizedAmount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"}],\"name\":\"depositCancel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantizedAmount\",\"type\":\"uint256\"}],\"name\":\"depositERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"}],\"name\":\"depositEth\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"depositNft\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"depositNftReclaim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"}],\"name\":\"depositReclaim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantizedAmount\",\"type\":\"uint256\"}],\"name\":\"escape\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"}],\"name\":\"freezeRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"}],\"name\":\"fullWithdrawalRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"}],\"name\":\"getAssetInfo\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"assetInfo\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"}],\"name\":\"getCancellationRequest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"request\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"}],\"name\":\"getDepositBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"}],\"name\":\"getEthKey\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"ethKey\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"}],\"name\":\"getFullWithdrawalRequest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"res\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLastBatchId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"batchId\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOrderRoot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"root\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOrderTreeHeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"}],\"name\":\"getQuantizedDepositBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"presumedAssetType\",\"type\":\"uint256\"}],\"name\":\"getQuantum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"quantum\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRegisteredAvailabilityVerifiers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRegisteredVerifiers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSequenceNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"seq\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVaultRoot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"root\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVaultTreeHeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"ownerKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"}],\"name\":\"getWithdrawalBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isAvailabilityVerifier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isFrozen\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isTokenAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isUserAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isVerifier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mainAcceptGovernance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mainCancelNomination\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"mainIsGovernor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"mainNominateNewGovernor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"mainRemoveGovernor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ethKey\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantizedAmount\",\"type\":\"uint256\"}],\"name\":\"registerAndDepositERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ethKey\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"}],\"name\":\"registerAndDepositEth\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"registerAvailabilityVerifier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"registerOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"registerToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"registerTokenAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"registerUser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"registerUserAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"registerVerifier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"removeAvailabilityVerifier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"removeVerifier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unFreeze\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"unregisterOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"unregisterTokenAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"unregisterUserAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"publicInput\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"applicationData\",\"type\":\"uint256[]\"}],\"name\":\"updateState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"ownerKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"ownerKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"mintingBlob\",\"type\":\"bytes\"}],\"name\":\"withdrawAndMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"ownerKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"withdrawNft\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"withdrawNftTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"withdrawTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// CoreABI is the input ABI used to generate the binding from.
// Deprecated: Use CoreMetaData.ABI instead.
var CoreABI = CoreMetaData.ABI

// Core is an auto generated Go binding around an Ethereum contract.
type Core struct {
	CoreCaller     // Read-only binding to the contract
	CoreTransactor // Write-only binding to the contract
	CoreFilterer   // Log filterer for contract events
}

// CoreCaller is an auto generated read-only Go binding around an Ethereum contract.
type CoreCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CoreTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CoreTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CoreFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CoreFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CoreSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CoreSession struct {
	Contract     *Core             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CoreCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CoreCallerSession struct {
	Contract *CoreCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// CoreTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CoreTransactorSession struct {
	Contract     *CoreTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CoreRaw is an auto generated low-level Go binding around an Ethereum contract.
type CoreRaw struct {
	Contract *Core // Generic contract binding to access the raw methods on
}

// CoreCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CoreCallerRaw struct {
	Contract *CoreCaller // Generic read-only contract binding to access the raw methods on
}

// CoreTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CoreTransactorRaw struct {
	Contract *CoreTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCore creates a new instance of Core, bound to a specific deployed contract.
func NewCore(address common.Address, backend bind.ContractBackend) (*Core, error) {
	contract, err := bindCore(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Core{CoreCaller: CoreCaller{contract: contract}, CoreTransactor: CoreTransactor{contract: contract}, CoreFilterer: CoreFilterer{contract: contract}}, nil
}

// NewCoreCaller creates a new read-only instance of Core, bound to a specific deployed contract.
func NewCoreCaller(address common.Address, caller bind.ContractCaller) (*CoreCaller, error) {
	contract, err := bindCore(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CoreCaller{contract: contract}, nil
}

// NewCoreTransactor creates a new write-only instance of Core, bound to a specific deployed contract.
func NewCoreTransactor(address common.Address, transactor bind.ContractTransactor) (*CoreTransactor, error) {
	contract, err := bindCore(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CoreTransactor{contract: contract}, nil
}

// NewCoreFilterer creates a new log filterer instance of Core, bound to a specific deployed contract.
func NewCoreFilterer(address common.Address, filterer bind.ContractFilterer) (*CoreFilterer, error) {
	contract, err := bindCore(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CoreFilterer{contract: contract}, nil
}

// bindCore binds a generic wrapper to an already deployed contract.
func bindCore(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CoreABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Core *CoreRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Core.Contract.CoreCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Core *CoreRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Core.Contract.CoreTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Core *CoreRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Core.Contract.CoreTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Core *CoreCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Core.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Core *CoreTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Core.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Core *CoreTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Core.Contract.contract.Transact(opts, method, params...)
}

// GetAssetInfo is a free data retrieval call binding the contract method 0xf637d950.
//
// Solidity: function getAssetInfo(uint256 assetType) view returns(bytes assetInfo)
func (_Core *CoreCaller) GetAssetInfo(opts *bind.CallOpts, assetType *big.Int) ([]byte, error) {
	var out []interface{}
	err := _Core.contract.Call(opts, &out, "getAssetInfo", assetType)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetAssetInfo is a free data retrieval call binding the contract method 0xf637d950.
//
// Solidity: function getAssetInfo(uint256 assetType) view returns(bytes assetInfo)
func (_Core *CoreSession) GetAssetInfo(assetType *big.Int) ([]byte, error) {
	return _Core.Contract.GetAssetInfo(&_Core.CallOpts, assetType)
}

// GetAssetInfo is a free data retrieval call binding the contract method 0xf637d950.
//
// Solidity: function getAssetInfo(uint256 assetType) view returns(bytes assetInfo)
func (_Core *CoreCallerSession) GetAssetInfo(assetType *big.Int) ([]byte, error) {
	return _Core.Contract.GetAssetInfo(&_Core.CallOpts, assetType)
}

// GetCancellationRequest is a free data retrieval call binding the contract method 0x333ac20b.
//
// Solidity: function getCancellationRequest(uint256 starkKey, uint256 assetId, uint256 vaultId) view returns(uint256 request)
func (_Core *CoreCaller) GetCancellationRequest(opts *bind.CallOpts, starkKey *big.Int, assetId *big.Int, vaultId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Core.contract.Call(opts, &out, "getCancellationRequest", starkKey, assetId, vaultId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCancellationRequest is a free data retrieval call binding the contract method 0x333ac20b.
//
// Solidity: function getCancellationRequest(uint256 starkKey, uint256 assetId, uint256 vaultId) view returns(uint256 request)
func (_Core *CoreSession) GetCancellationRequest(starkKey *big.Int, assetId *big.Int, vaultId *big.Int) (*big.Int, error) {
	return _Core.Contract.GetCancellationRequest(&_Core.CallOpts, starkKey, assetId, vaultId)
}

// GetCancellationRequest is a free data retrieval call binding the contract method 0x333ac20b.
//
// Solidity: function getCancellationRequest(uint256 starkKey, uint256 assetId, uint256 vaultId) view returns(uint256 request)
func (_Core *CoreCallerSession) GetCancellationRequest(starkKey *big.Int, assetId *big.Int, vaultId *big.Int) (*big.Int, error) {
	return _Core.Contract.GetCancellationRequest(&_Core.CallOpts, starkKey, assetId, vaultId)
}

// GetDepositBalance is a free data retrieval call binding the contract method 0xabf98fe1.
//
// Solidity: function getDepositBalance(uint256 starkKey, uint256 assetId, uint256 vaultId) view returns(uint256 balance)
func (_Core *CoreCaller) GetDepositBalance(opts *bind.CallOpts, starkKey *big.Int, assetId *big.Int, vaultId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Core.contract.Call(opts, &out, "getDepositBalance", starkKey, assetId, vaultId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDepositBalance is a free data retrieval call binding the contract method 0xabf98fe1.
//
// Solidity: function getDepositBalance(uint256 starkKey, uint256 assetId, uint256 vaultId) view returns(uint256 balance)
func (_Core *CoreSession) GetDepositBalance(starkKey *big.Int, assetId *big.Int, vaultId *big.Int) (*big.Int, error) {
	return _Core.Contract.GetDepositBalance(&_Core.CallOpts, starkKey, assetId, vaultId)
}

// GetDepositBalance is a free data retrieval call binding the contract method 0xabf98fe1.
//
// Solidity: function getDepositBalance(uint256 starkKey, uint256 assetId, uint256 vaultId) view returns(uint256 balance)
func (_Core *CoreCallerSession) GetDepositBalance(starkKey *big.Int, assetId *big.Int, vaultId *big.Int) (*big.Int, error) {
	return _Core.Contract.GetDepositBalance(&_Core.CallOpts, starkKey, assetId, vaultId)
}

// GetEthKey is a free data retrieval call binding the contract method 0x1dbd1da7.
//
// Solidity: function getEthKey(uint256 starkKey) view returns(address ethKey)
func (_Core *CoreCaller) GetEthKey(opts *bind.CallOpts, starkKey *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Core.contract.Call(opts, &out, "getEthKey", starkKey)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetEthKey is a free data retrieval call binding the contract method 0x1dbd1da7.
//
// Solidity: function getEthKey(uint256 starkKey) view returns(address ethKey)
func (_Core *CoreSession) GetEthKey(starkKey *big.Int) (common.Address, error) {
	return _Core.Contract.GetEthKey(&_Core.CallOpts, starkKey)
}

// GetEthKey is a free data retrieval call binding the contract method 0x1dbd1da7.
//
// Solidity: function getEthKey(uint256 starkKey) view returns(address ethKey)
func (_Core *CoreCallerSession) GetEthKey(starkKey *big.Int) (common.Address, error) {
	return _Core.Contract.GetEthKey(&_Core.CallOpts, starkKey)
}

// GetFullWithdrawalRequest is a free data retrieval call binding the contract method 0x296e2f37.
//
// Solidity: function getFullWithdrawalRequest(uint256 starkKey, uint256 vaultId) view returns(uint256 res)
func (_Core *CoreCaller) GetFullWithdrawalRequest(opts *bind.CallOpts, starkKey *big.Int, vaultId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Core.contract.Call(opts, &out, "getFullWithdrawalRequest", starkKey, vaultId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFullWithdrawalRequest is a free data retrieval call binding the contract method 0x296e2f37.
//
// Solidity: function getFullWithdrawalRequest(uint256 starkKey, uint256 vaultId) view returns(uint256 res)
func (_Core *CoreSession) GetFullWithdrawalRequest(starkKey *big.Int, vaultId *big.Int) (*big.Int, error) {
	return _Core.Contract.GetFullWithdrawalRequest(&_Core.CallOpts, starkKey, vaultId)
}

// GetFullWithdrawalRequest is a free data retrieval call binding the contract method 0x296e2f37.
//
// Solidity: function getFullWithdrawalRequest(uint256 starkKey, uint256 vaultId) view returns(uint256 res)
func (_Core *CoreCallerSession) GetFullWithdrawalRequest(starkKey *big.Int, vaultId *big.Int) (*big.Int, error) {
	return _Core.Contract.GetFullWithdrawalRequest(&_Core.CallOpts, starkKey, vaultId)
}

// GetLastBatchId is a free data retrieval call binding the contract method 0x515535e8.
//
// Solidity: function getLastBatchId() view returns(uint256 batchId)
func (_Core *CoreCaller) GetLastBatchId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Core.contract.Call(opts, &out, "getLastBatchId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLastBatchId is a free data retrieval call binding the contract method 0x515535e8.
//
// Solidity: function getLastBatchId() view returns(uint256 batchId)
func (_Core *CoreSession) GetLastBatchId() (*big.Int, error) {
	return _Core.Contract.GetLastBatchId(&_Core.CallOpts)
}

// GetLastBatchId is a free data retrieval call binding the contract method 0x515535e8.
//
// Solidity: function getLastBatchId() view returns(uint256 batchId)
func (_Core *CoreCallerSession) GetLastBatchId() (*big.Int, error) {
	return _Core.Contract.GetLastBatchId(&_Core.CallOpts)
}

// GetOrderRoot is a free data retrieval call binding the contract method 0x0dded952.
//
// Solidity: function getOrderRoot() view returns(uint256 root)
func (_Core *CoreCaller) GetOrderRoot(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Core.contract.Call(opts, &out, "getOrderRoot")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetOrderRoot is a free data retrieval call binding the contract method 0x0dded952.
//
// Solidity: function getOrderRoot() view returns(uint256 root)
func (_Core *CoreSession) GetOrderRoot() (*big.Int, error) {
	return _Core.Contract.GetOrderRoot(&_Core.CallOpts)
}

// GetOrderRoot is a free data retrieval call binding the contract method 0x0dded952.
//
// Solidity: function getOrderRoot() view returns(uint256 root)
func (_Core *CoreCallerSession) GetOrderRoot() (*big.Int, error) {
	return _Core.Contract.GetOrderRoot(&_Core.CallOpts)
}

// GetOrderTreeHeight is a free data retrieval call binding the contract method 0x7e9da4c5.
//
// Solidity: function getOrderTreeHeight() view returns(uint256 height)
func (_Core *CoreCaller) GetOrderTreeHeight(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Core.contract.Call(opts, &out, "getOrderTreeHeight")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetOrderTreeHeight is a free data retrieval call binding the contract method 0x7e9da4c5.
//
// Solidity: function getOrderTreeHeight() view returns(uint256 height)
func (_Core *CoreSession) GetOrderTreeHeight() (*big.Int, error) {
	return _Core.Contract.GetOrderTreeHeight(&_Core.CallOpts)
}

// GetOrderTreeHeight is a free data retrieval call binding the contract method 0x7e9da4c5.
//
// Solidity: function getOrderTreeHeight() view returns(uint256 height)
func (_Core *CoreCallerSession) GetOrderTreeHeight() (*big.Int, error) {
	return _Core.Contract.GetOrderTreeHeight(&_Core.CallOpts)
}

// GetQuantizedDepositBalance is a free data retrieval call binding the contract method 0x4e8912da.
//
// Solidity: function getQuantizedDepositBalance(uint256 starkKey, uint256 assetId, uint256 vaultId) view returns(uint256 balance)
func (_Core *CoreCaller) GetQuantizedDepositBalance(opts *bind.CallOpts, starkKey *big.Int, assetId *big.Int, vaultId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Core.contract.Call(opts, &out, "getQuantizedDepositBalance", starkKey, assetId, vaultId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetQuantizedDepositBalance is a free data retrieval call binding the contract method 0x4e8912da.
//
// Solidity: function getQuantizedDepositBalance(uint256 starkKey, uint256 assetId, uint256 vaultId) view returns(uint256 balance)
func (_Core *CoreSession) GetQuantizedDepositBalance(starkKey *big.Int, assetId *big.Int, vaultId *big.Int) (*big.Int, error) {
	return _Core.Contract.GetQuantizedDepositBalance(&_Core.CallOpts, starkKey, assetId, vaultId)
}

// GetQuantizedDepositBalance is a free data retrieval call binding the contract method 0x4e8912da.
//
// Solidity: function getQuantizedDepositBalance(uint256 starkKey, uint256 assetId, uint256 vaultId) view returns(uint256 balance)
func (_Core *CoreCallerSession) GetQuantizedDepositBalance(starkKey *big.Int, assetId *big.Int, vaultId *big.Int) (*big.Int, error) {
	return _Core.Contract.GetQuantizedDepositBalance(&_Core.CallOpts, starkKey, assetId, vaultId)
}

// GetQuantum is a free data retrieval call binding the contract method 0xdd7202d8.
//
// Solidity: function getQuantum(uint256 presumedAssetType) view returns(uint256 quantum)
func (_Core *CoreCaller) GetQuantum(opts *bind.CallOpts, presumedAssetType *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Core.contract.Call(opts, &out, "getQuantum", presumedAssetType)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetQuantum is a free data retrieval call binding the contract method 0xdd7202d8.
//
// Solidity: function getQuantum(uint256 presumedAssetType) view returns(uint256 quantum)
func (_Core *CoreSession) GetQuantum(presumedAssetType *big.Int) (*big.Int, error) {
	return _Core.Contract.GetQuantum(&_Core.CallOpts, presumedAssetType)
}

// GetQuantum is a free data retrieval call binding the contract method 0xdd7202d8.
//
// Solidity: function getQuantum(uint256 presumedAssetType) view returns(uint256 quantum)
func (_Core *CoreCallerSession) GetQuantum(presumedAssetType *big.Int) (*big.Int, error) {
	return _Core.Contract.GetQuantum(&_Core.CallOpts, presumedAssetType)
}

// GetSequenceNumber is a free data retrieval call binding the contract method 0x42af35fd.
//
// Solidity: function getSequenceNumber() view returns(uint256 seq)
func (_Core *CoreCaller) GetSequenceNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Core.contract.Call(opts, &out, "getSequenceNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSequenceNumber is a free data retrieval call binding the contract method 0x42af35fd.
//
// Solidity: function getSequenceNumber() view returns(uint256 seq)
func (_Core *CoreSession) GetSequenceNumber() (*big.Int, error) {
	return _Core.Contract.GetSequenceNumber(&_Core.CallOpts)
}

// GetSequenceNumber is a free data retrieval call binding the contract method 0x42af35fd.
//
// Solidity: function getSequenceNumber() view returns(uint256 seq)
func (_Core *CoreCallerSession) GetSequenceNumber() (*big.Int, error) {
	return _Core.Contract.GetSequenceNumber(&_Core.CallOpts)
}

// GetVaultRoot is a free data retrieval call binding the contract method 0x64da5dfe.
//
// Solidity: function getVaultRoot() view returns(uint256 root)
func (_Core *CoreCaller) GetVaultRoot(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Core.contract.Call(opts, &out, "getVaultRoot")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVaultRoot is a free data retrieval call binding the contract method 0x64da5dfe.
//
// Solidity: function getVaultRoot() view returns(uint256 root)
func (_Core *CoreSession) GetVaultRoot() (*big.Int, error) {
	return _Core.Contract.GetVaultRoot(&_Core.CallOpts)
}

// GetVaultRoot is a free data retrieval call binding the contract method 0x64da5dfe.
//
// Solidity: function getVaultRoot() view returns(uint256 root)
func (_Core *CoreCallerSession) GetVaultRoot() (*big.Int, error) {
	return _Core.Contract.GetVaultRoot(&_Core.CallOpts)
}

// GetVaultTreeHeight is a free data retrieval call binding the contract method 0xf288a3ff.
//
// Solidity: function getVaultTreeHeight() view returns(uint256 height)
func (_Core *CoreCaller) GetVaultTreeHeight(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Core.contract.Call(opts, &out, "getVaultTreeHeight")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVaultTreeHeight is a free data retrieval call binding the contract method 0xf288a3ff.
//
// Solidity: function getVaultTreeHeight() view returns(uint256 height)
func (_Core *CoreSession) GetVaultTreeHeight() (*big.Int, error) {
	return _Core.Contract.GetVaultTreeHeight(&_Core.CallOpts)
}

// GetVaultTreeHeight is a free data retrieval call binding the contract method 0xf288a3ff.
//
// Solidity: function getVaultTreeHeight() view returns(uint256 height)
func (_Core *CoreCallerSession) GetVaultTreeHeight() (*big.Int, error) {
	return _Core.Contract.GetVaultTreeHeight(&_Core.CallOpts)
}

// GetWithdrawalBalance is a free data retrieval call binding the contract method 0xec3161b0.
//
// Solidity: function getWithdrawalBalance(uint256 ownerKey, uint256 assetId) view returns(uint256 balance)
func (_Core *CoreCaller) GetWithdrawalBalance(opts *bind.CallOpts, ownerKey *big.Int, assetId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Core.contract.Call(opts, &out, "getWithdrawalBalance", ownerKey, assetId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetWithdrawalBalance is a free data retrieval call binding the contract method 0xec3161b0.
//
// Solidity: function getWithdrawalBalance(uint256 ownerKey, uint256 assetId) view returns(uint256 balance)
func (_Core *CoreSession) GetWithdrawalBalance(ownerKey *big.Int, assetId *big.Int) (*big.Int, error) {
	return _Core.Contract.GetWithdrawalBalance(&_Core.CallOpts, ownerKey, assetId)
}

// GetWithdrawalBalance is a free data retrieval call binding the contract method 0xec3161b0.
//
// Solidity: function getWithdrawalBalance(uint256 ownerKey, uint256 assetId) view returns(uint256 balance)
func (_Core *CoreCallerSession) GetWithdrawalBalance(ownerKey *big.Int, assetId *big.Int) (*big.Int, error) {
	return _Core.Contract.GetWithdrawalBalance(&_Core.CallOpts, ownerKey, assetId)
}

// AnnounceAvailabilityVerifierRemovalIntent is a paid mutator transaction binding the contract method 0x1d078bbb.
//
// Solidity: function announceAvailabilityVerifierRemovalIntent(address ) returns()
func (_Core *CoreTransactor) AnnounceAvailabilityVerifierRemovalIntent(opts *bind.TransactOpts, arg0 common.Address) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "announceAvailabilityVerifierRemovalIntent", arg0)
}

// AnnounceAvailabilityVerifierRemovalIntent is a paid mutator transaction binding the contract method 0x1d078bbb.
//
// Solidity: function announceAvailabilityVerifierRemovalIntent(address ) returns()
func (_Core *CoreSession) AnnounceAvailabilityVerifierRemovalIntent(arg0 common.Address) (*types.Transaction, error) {
	return _Core.Contract.AnnounceAvailabilityVerifierRemovalIntent(&_Core.TransactOpts, arg0)
}

// AnnounceAvailabilityVerifierRemovalIntent is a paid mutator transaction binding the contract method 0x1d078bbb.
//
// Solidity: function announceAvailabilityVerifierRemovalIntent(address ) returns()
func (_Core *CoreTransactorSession) AnnounceAvailabilityVerifierRemovalIntent(arg0 common.Address) (*types.Transaction, error) {
	return _Core.Contract.AnnounceAvailabilityVerifierRemovalIntent(&_Core.TransactOpts, arg0)
}

// AnnounceVerifierRemovalIntent is a paid mutator transaction binding the contract method 0x418573b7.
//
// Solidity: function announceVerifierRemovalIntent(address ) returns()
func (_Core *CoreTransactor) AnnounceVerifierRemovalIntent(opts *bind.TransactOpts, arg0 common.Address) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "announceVerifierRemovalIntent", arg0)
}

// AnnounceVerifierRemovalIntent is a paid mutator transaction binding the contract method 0x418573b7.
//
// Solidity: function announceVerifierRemovalIntent(address ) returns()
func (_Core *CoreSession) AnnounceVerifierRemovalIntent(arg0 common.Address) (*types.Transaction, error) {
	return _Core.Contract.AnnounceVerifierRemovalIntent(&_Core.TransactOpts, arg0)
}

// AnnounceVerifierRemovalIntent is a paid mutator transaction binding the contract method 0x418573b7.
//
// Solidity: function announceVerifierRemovalIntent(address ) returns()
func (_Core *CoreTransactorSession) AnnounceVerifierRemovalIntent(arg0 common.Address) (*types.Transaction, error) {
	return _Core.Contract.AnnounceVerifierRemovalIntent(&_Core.TransactOpts, arg0)
}

// Deposit is a paid mutator transaction binding the contract method 0x00aeef8a.
//
// Solidity: function deposit(uint256 starkKey, uint256 assetType, uint256 vaultId) payable returns()
func (_Core *CoreTransactor) Deposit(opts *bind.TransactOpts, starkKey *big.Int, assetType *big.Int, vaultId *big.Int) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "deposit", starkKey, assetType, vaultId)
}

// Deposit is a paid mutator transaction binding the contract method 0x00aeef8a.
//
// Solidity: function deposit(uint256 starkKey, uint256 assetType, uint256 vaultId) payable returns()
func (_Core *CoreSession) Deposit(starkKey *big.Int, assetType *big.Int, vaultId *big.Int) (*types.Transaction, error) {
	return _Core.Contract.Deposit(&_Core.TransactOpts, starkKey, assetType, vaultId)
}

// Deposit is a paid mutator transaction binding the contract method 0x00aeef8a.
//
// Solidity: function deposit(uint256 starkKey, uint256 assetType, uint256 vaultId) payable returns()
func (_Core *CoreTransactorSession) Deposit(starkKey *big.Int, assetType *big.Int, vaultId *big.Int) (*types.Transaction, error) {
	return _Core.Contract.Deposit(&_Core.TransactOpts, starkKey, assetType, vaultId)
}

// Deposit0 is a paid mutator transaction binding the contract method 0x2505c3d9.
//
// Solidity: function deposit(uint256 starkKey, uint256 assetType, uint256 vaultId, uint256 quantizedAmount) returns()
func (_Core *CoreTransactor) Deposit0(opts *bind.TransactOpts, starkKey *big.Int, assetType *big.Int, vaultId *big.Int, quantizedAmount *big.Int) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "deposit0", starkKey, assetType, vaultId, quantizedAmount)
}

// Deposit0 is a paid mutator transaction binding the contract method 0x2505c3d9.
//
// Solidity: function deposit(uint256 starkKey, uint256 assetType, uint256 vaultId, uint256 quantizedAmount) returns()
func (_Core *CoreSession) Deposit0(starkKey *big.Int, assetType *big.Int, vaultId *big.Int, quantizedAmount *big.Int) (*types.Transaction, error) {
	return _Core.Contract.Deposit0(&_Core.TransactOpts, starkKey, assetType, vaultId, quantizedAmount)
}

// Deposit0 is a paid mutator transaction binding the contract method 0x2505c3d9.
//
// Solidity: function deposit(uint256 starkKey, uint256 assetType, uint256 vaultId, uint256 quantizedAmount) returns()
func (_Core *CoreTransactorSession) Deposit0(starkKey *big.Int, assetType *big.Int, vaultId *big.Int, quantizedAmount *big.Int) (*types.Transaction, error) {
	return _Core.Contract.Deposit0(&_Core.TransactOpts, starkKey, assetType, vaultId, quantizedAmount)
}

// DepositCancel is a paid mutator transaction binding the contract method 0x7df7dc04.
//
// Solidity: function depositCancel(uint256 starkKey, uint256 assetId, uint256 vaultId) returns()
func (_Core *CoreTransactor) DepositCancel(opts *bind.TransactOpts, starkKey *big.Int, assetId *big.Int, vaultId *big.Int) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "depositCancel", starkKey, assetId, vaultId)
}

// DepositCancel is a paid mutator transaction binding the contract method 0x7df7dc04.
//
// Solidity: function depositCancel(uint256 starkKey, uint256 assetId, uint256 vaultId) returns()
func (_Core *CoreSession) DepositCancel(starkKey *big.Int, assetId *big.Int, vaultId *big.Int) (*types.Transaction, error) {
	return _Core.Contract.DepositCancel(&_Core.TransactOpts, starkKey, assetId, vaultId)
}

// DepositCancel is a paid mutator transaction binding the contract method 0x7df7dc04.
//
// Solidity: function depositCancel(uint256 starkKey, uint256 assetId, uint256 vaultId) returns()
func (_Core *CoreTransactorSession) DepositCancel(starkKey *big.Int, assetId *big.Int, vaultId *big.Int) (*types.Transaction, error) {
	return _Core.Contract.DepositCancel(&_Core.TransactOpts, starkKey, assetId, vaultId)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0x9ed17084.
//
// Solidity: function depositERC20(uint256 starkKey, uint256 assetType, uint256 vaultId, uint256 quantizedAmount) returns()
func (_Core *CoreTransactor) DepositERC20(opts *bind.TransactOpts, starkKey *big.Int, assetType *big.Int, vaultId *big.Int, quantizedAmount *big.Int) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "depositERC20", starkKey, assetType, vaultId, quantizedAmount)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0x9ed17084.
//
// Solidity: function depositERC20(uint256 starkKey, uint256 assetType, uint256 vaultId, uint256 quantizedAmount) returns()
func (_Core *CoreSession) DepositERC20(starkKey *big.Int, assetType *big.Int, vaultId *big.Int, quantizedAmount *big.Int) (*types.Transaction, error) {
	return _Core.Contract.DepositERC20(&_Core.TransactOpts, starkKey, assetType, vaultId, quantizedAmount)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0x9ed17084.
//
// Solidity: function depositERC20(uint256 starkKey, uint256 assetType, uint256 vaultId, uint256 quantizedAmount) returns()
func (_Core *CoreTransactorSession) DepositERC20(starkKey *big.Int, assetType *big.Int, vaultId *big.Int, quantizedAmount *big.Int) (*types.Transaction, error) {
	return _Core.Contract.DepositERC20(&_Core.TransactOpts, starkKey, assetType, vaultId, quantizedAmount)
}

// DepositEth is a paid mutator transaction binding the contract method 0x6ce5d957.
//
// Solidity: function depositEth(uint256 starkKey, uint256 assetType, uint256 vaultId) payable returns()
func (_Core *CoreTransactor) DepositEth(opts *bind.TransactOpts, starkKey *big.Int, assetType *big.Int, vaultId *big.Int) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "depositEth", starkKey, assetType, vaultId)
}

// DepositEth is a paid mutator transaction binding the contract method 0x6ce5d957.
//
// Solidity: function depositEth(uint256 starkKey, uint256 assetType, uint256 vaultId) payable returns()
func (_Core *CoreSession) DepositEth(starkKey *big.Int, assetType *big.Int, vaultId *big.Int) (*types.Transaction, error) {
	return _Core.Contract.DepositEth(&_Core.TransactOpts, starkKey, assetType, vaultId)
}

// DepositEth is a paid mutator transaction binding the contract method 0x6ce5d957.
//
// Solidity: function depositEth(uint256 starkKey, uint256 assetType, uint256 vaultId) payable returns()
func (_Core *CoreTransactorSession) DepositEth(starkKey *big.Int, assetType *big.Int, vaultId *big.Int) (*types.Transaction, error) {
	return _Core.Contract.DepositEth(&_Core.TransactOpts, starkKey, assetType, vaultId)
}

// DepositNft is a paid mutator transaction binding the contract method 0xae1cdde6.
//
// Solidity: function depositNft(uint256 starkKey, uint256 assetType, uint256 vaultId, uint256 tokenId) returns()
func (_Core *CoreTransactor) DepositNft(opts *bind.TransactOpts, starkKey *big.Int, assetType *big.Int, vaultId *big.Int, tokenId *big.Int) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "depositNft", starkKey, assetType, vaultId, tokenId)
}

// DepositNft is a paid mutator transaction binding the contract method 0xae1cdde6.
//
// Solidity: function depositNft(uint256 starkKey, uint256 assetType, uint256 vaultId, uint256 tokenId) returns()
func (_Core *CoreSession) DepositNft(starkKey *big.Int, assetType *big.Int, vaultId *big.Int, tokenId *big.Int) (*types.Transaction, error) {
	return _Core.Contract.DepositNft(&_Core.TransactOpts, starkKey, assetType, vaultId, tokenId)
}

// DepositNft is a paid mutator transaction binding the contract method 0xae1cdde6.
//
// Solidity: function depositNft(uint256 starkKey, uint256 assetType, uint256 vaultId, uint256 tokenId) returns()
func (_Core *CoreTransactorSession) DepositNft(starkKey *big.Int, assetType *big.Int, vaultId *big.Int, tokenId *big.Int) (*types.Transaction, error) {
	return _Core.Contract.DepositNft(&_Core.TransactOpts, starkKey, assetType, vaultId, tokenId)
}

// DepositNftReclaim is a paid mutator transaction binding the contract method 0xfcb05822.
//
// Solidity: function depositNftReclaim(uint256 starkKey, uint256 assetType, uint256 vaultId, uint256 tokenId) returns()
func (_Core *CoreTransactor) DepositNftReclaim(opts *bind.TransactOpts, starkKey *big.Int, assetType *big.Int, vaultId *big.Int, tokenId *big.Int) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "depositNftReclaim", starkKey, assetType, vaultId, tokenId)
}

// DepositNftReclaim is a paid mutator transaction binding the contract method 0xfcb05822.
//
// Solidity: function depositNftReclaim(uint256 starkKey, uint256 assetType, uint256 vaultId, uint256 tokenId) returns()
func (_Core *CoreSession) DepositNftReclaim(starkKey *big.Int, assetType *big.Int, vaultId *big.Int, tokenId *big.Int) (*types.Transaction, error) {
	return _Core.Contract.DepositNftReclaim(&_Core.TransactOpts, starkKey, assetType, vaultId, tokenId)
}

// DepositNftReclaim is a paid mutator transaction binding the contract method 0xfcb05822.
//
// Solidity: function depositNftReclaim(uint256 starkKey, uint256 assetType, uint256 vaultId, uint256 tokenId) returns()
func (_Core *CoreTransactorSession) DepositNftReclaim(starkKey *big.Int, assetType *big.Int, vaultId *big.Int, tokenId *big.Int) (*types.Transaction, error) {
	return _Core.Contract.DepositNftReclaim(&_Core.TransactOpts, starkKey, assetType, vaultId, tokenId)
}

// DepositReclaim is a paid mutator transaction binding the contract method 0xae873816.
//
// Solidity: function depositReclaim(uint256 starkKey, uint256 assetId, uint256 vaultId) returns()
func (_Core *CoreTransactor) DepositReclaim(opts *bind.TransactOpts, starkKey *big.Int, assetId *big.Int, vaultId *big.Int) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "depositReclaim", starkKey, assetId, vaultId)
}

// DepositReclaim is a paid mutator transaction binding the contract method 0xae873816.
//
// Solidity: function depositReclaim(uint256 starkKey, uint256 assetId, uint256 vaultId) returns()
func (_Core *CoreSession) DepositReclaim(starkKey *big.Int, assetId *big.Int, vaultId *big.Int) (*types.Transaction, error) {
	return _Core.Contract.DepositReclaim(&_Core.TransactOpts, starkKey, assetId, vaultId)
}

// DepositReclaim is a paid mutator transaction binding the contract method 0xae873816.
//
// Solidity: function depositReclaim(uint256 starkKey, uint256 assetId, uint256 vaultId) returns()
func (_Core *CoreTransactorSession) DepositReclaim(starkKey *big.Int, assetId *big.Int, vaultId *big.Int) (*types.Transaction, error) {
	return _Core.Contract.DepositReclaim(&_Core.TransactOpts, starkKey, assetId, vaultId)
}

// Escape is a paid mutator transaction binding the contract method 0x9e3adac4.
//
// Solidity: function escape(uint256 starkKey, uint256 vaultId, uint256 assetId, uint256 quantizedAmount) returns()
func (_Core *CoreTransactor) Escape(opts *bind.TransactOpts, starkKey *big.Int, vaultId *big.Int, assetId *big.Int, quantizedAmount *big.Int) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "escape", starkKey, vaultId, assetId, quantizedAmount)
}

// Escape is a paid mutator transaction binding the contract method 0x9e3adac4.
//
// Solidity: function escape(uint256 starkKey, uint256 vaultId, uint256 assetId, uint256 quantizedAmount) returns()
func (_Core *CoreSession) Escape(starkKey *big.Int, vaultId *big.Int, assetId *big.Int, quantizedAmount *big.Int) (*types.Transaction, error) {
	return _Core.Contract.Escape(&_Core.TransactOpts, starkKey, vaultId, assetId, quantizedAmount)
}

// Escape is a paid mutator transaction binding the contract method 0x9e3adac4.
//
// Solidity: function escape(uint256 starkKey, uint256 vaultId, uint256 assetId, uint256 quantizedAmount) returns()
func (_Core *CoreTransactorSession) Escape(starkKey *big.Int, vaultId *big.Int, assetId *big.Int, quantizedAmount *big.Int) (*types.Transaction, error) {
	return _Core.Contract.Escape(&_Core.TransactOpts, starkKey, vaultId, assetId, quantizedAmount)
}

// FreezeRequest is a paid mutator transaction binding the contract method 0x93c1e466.
//
// Solidity: function freezeRequest(uint256 starkKey, uint256 vaultId) returns()
func (_Core *CoreTransactor) FreezeRequest(opts *bind.TransactOpts, starkKey *big.Int, vaultId *big.Int) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "freezeRequest", starkKey, vaultId)
}

// FreezeRequest is a paid mutator transaction binding the contract method 0x93c1e466.
//
// Solidity: function freezeRequest(uint256 starkKey, uint256 vaultId) returns()
func (_Core *CoreSession) FreezeRequest(starkKey *big.Int, vaultId *big.Int) (*types.Transaction, error) {
	return _Core.Contract.FreezeRequest(&_Core.TransactOpts, starkKey, vaultId)
}

// FreezeRequest is a paid mutator transaction binding the contract method 0x93c1e466.
//
// Solidity: function freezeRequest(uint256 starkKey, uint256 vaultId) returns()
func (_Core *CoreTransactorSession) FreezeRequest(starkKey *big.Int, vaultId *big.Int) (*types.Transaction, error) {
	return _Core.Contract.FreezeRequest(&_Core.TransactOpts, starkKey, vaultId)
}

// FullWithdrawalRequest is a paid mutator transaction binding the contract method 0xa93310c4.
//
// Solidity: function fullWithdrawalRequest(uint256 starkKey, uint256 vaultId) returns()
func (_Core *CoreTransactor) FullWithdrawalRequest(opts *bind.TransactOpts, starkKey *big.Int, vaultId *big.Int) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "fullWithdrawalRequest", starkKey, vaultId)
}

// FullWithdrawalRequest is a paid mutator transaction binding the contract method 0xa93310c4.
//
// Solidity: function fullWithdrawalRequest(uint256 starkKey, uint256 vaultId) returns()
func (_Core *CoreSession) FullWithdrawalRequest(starkKey *big.Int, vaultId *big.Int) (*types.Transaction, error) {
	return _Core.Contract.FullWithdrawalRequest(&_Core.TransactOpts, starkKey, vaultId)
}

// FullWithdrawalRequest is a paid mutator transaction binding the contract method 0xa93310c4.
//
// Solidity: function fullWithdrawalRequest(uint256 starkKey, uint256 vaultId) returns()
func (_Core *CoreTransactorSession) FullWithdrawalRequest(starkKey *big.Int, vaultId *big.Int) (*types.Transaction, error) {
	return _Core.Contract.FullWithdrawalRequest(&_Core.TransactOpts, starkKey, vaultId)
}

// GetRegisteredAvailabilityVerifiers is a paid mutator transaction binding the contract method 0x1ac347f2.
//
// Solidity: function getRegisteredAvailabilityVerifiers() returns()
func (_Core *CoreTransactor) GetRegisteredAvailabilityVerifiers(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "getRegisteredAvailabilityVerifiers")
}

// GetRegisteredAvailabilityVerifiers is a paid mutator transaction binding the contract method 0x1ac347f2.
//
// Solidity: function getRegisteredAvailabilityVerifiers() returns()
func (_Core *CoreSession) GetRegisteredAvailabilityVerifiers() (*types.Transaction, error) {
	return _Core.Contract.GetRegisteredAvailabilityVerifiers(&_Core.TransactOpts)
}

// GetRegisteredAvailabilityVerifiers is a paid mutator transaction binding the contract method 0x1ac347f2.
//
// Solidity: function getRegisteredAvailabilityVerifiers() returns()
func (_Core *CoreTransactorSession) GetRegisteredAvailabilityVerifiers() (*types.Transaction, error) {
	return _Core.Contract.GetRegisteredAvailabilityVerifiers(&_Core.TransactOpts)
}

// GetRegisteredVerifiers is a paid mutator transaction binding the contract method 0x4eab9ed3.
//
// Solidity: function getRegisteredVerifiers() returns()
func (_Core *CoreTransactor) GetRegisteredVerifiers(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "getRegisteredVerifiers")
}

// GetRegisteredVerifiers is a paid mutator transaction binding the contract method 0x4eab9ed3.
//
// Solidity: function getRegisteredVerifiers() returns()
func (_Core *CoreSession) GetRegisteredVerifiers() (*types.Transaction, error) {
	return _Core.Contract.GetRegisteredVerifiers(&_Core.TransactOpts)
}

// GetRegisteredVerifiers is a paid mutator transaction binding the contract method 0x4eab9ed3.
//
// Solidity: function getRegisteredVerifiers() returns()
func (_Core *CoreTransactorSession) GetRegisteredVerifiers() (*types.Transaction, error) {
	return _Core.Contract.GetRegisteredVerifiers(&_Core.TransactOpts)
}

// IsAvailabilityVerifier is a paid mutator transaction binding the contract method 0xbd1279ae.
//
// Solidity: function isAvailabilityVerifier(address ) returns()
func (_Core *CoreTransactor) IsAvailabilityVerifier(opts *bind.TransactOpts, arg0 common.Address) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "isAvailabilityVerifier", arg0)
}

// IsAvailabilityVerifier is a paid mutator transaction binding the contract method 0xbd1279ae.
//
// Solidity: function isAvailabilityVerifier(address ) returns()
func (_Core *CoreSession) IsAvailabilityVerifier(arg0 common.Address) (*types.Transaction, error) {
	return _Core.Contract.IsAvailabilityVerifier(&_Core.TransactOpts, arg0)
}

// IsAvailabilityVerifier is a paid mutator transaction binding the contract method 0xbd1279ae.
//
// Solidity: function isAvailabilityVerifier(address ) returns()
func (_Core *CoreTransactorSession) IsAvailabilityVerifier(arg0 common.Address) (*types.Transaction, error) {
	return _Core.Contract.IsAvailabilityVerifier(&_Core.TransactOpts, arg0)
}

// IsFrozen is a paid mutator transaction binding the contract method 0x33eeb147.
//
// Solidity: function isFrozen() returns()
func (_Core *CoreTransactor) IsFrozen(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "isFrozen")
}

// IsFrozen is a paid mutator transaction binding the contract method 0x33eeb147.
//
// Solidity: function isFrozen() returns()
func (_Core *CoreSession) IsFrozen() (*types.Transaction, error) {
	return _Core.Contract.IsFrozen(&_Core.TransactOpts)
}

// IsFrozen is a paid mutator transaction binding the contract method 0x33eeb147.
//
// Solidity: function isFrozen() returns()
func (_Core *CoreTransactorSession) IsFrozen() (*types.Transaction, error) {
	return _Core.Contract.IsFrozen(&_Core.TransactOpts)
}

// IsOperator is a paid mutator transaction binding the contract method 0x6d70f7ae.
//
// Solidity: function isOperator(address ) returns()
func (_Core *CoreTransactor) IsOperator(opts *bind.TransactOpts, arg0 common.Address) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "isOperator", arg0)
}

// IsOperator is a paid mutator transaction binding the contract method 0x6d70f7ae.
//
// Solidity: function isOperator(address ) returns()
func (_Core *CoreSession) IsOperator(arg0 common.Address) (*types.Transaction, error) {
	return _Core.Contract.IsOperator(&_Core.TransactOpts, arg0)
}

// IsOperator is a paid mutator transaction binding the contract method 0x6d70f7ae.
//
// Solidity: function isOperator(address ) returns()
func (_Core *CoreTransactorSession) IsOperator(arg0 common.Address) (*types.Transaction, error) {
	return _Core.Contract.IsOperator(&_Core.TransactOpts, arg0)
}

// IsTokenAdmin is a paid mutator transaction binding the contract method 0xa2bdde3d.
//
// Solidity: function isTokenAdmin(address ) returns()
func (_Core *CoreTransactor) IsTokenAdmin(opts *bind.TransactOpts, arg0 common.Address) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "isTokenAdmin", arg0)
}

// IsTokenAdmin is a paid mutator transaction binding the contract method 0xa2bdde3d.
//
// Solidity: function isTokenAdmin(address ) returns()
func (_Core *CoreSession) IsTokenAdmin(arg0 common.Address) (*types.Transaction, error) {
	return _Core.Contract.IsTokenAdmin(&_Core.TransactOpts, arg0)
}

// IsTokenAdmin is a paid mutator transaction binding the contract method 0xa2bdde3d.
//
// Solidity: function isTokenAdmin(address ) returns()
func (_Core *CoreTransactorSession) IsTokenAdmin(arg0 common.Address) (*types.Transaction, error) {
	return _Core.Contract.IsTokenAdmin(&_Core.TransactOpts, arg0)
}

// IsUserAdmin is a paid mutator transaction binding the contract method 0x74d523a8.
//
// Solidity: function isUserAdmin(address ) returns()
func (_Core *CoreTransactor) IsUserAdmin(opts *bind.TransactOpts, arg0 common.Address) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "isUserAdmin", arg0)
}

// IsUserAdmin is a paid mutator transaction binding the contract method 0x74d523a8.
//
// Solidity: function isUserAdmin(address ) returns()
func (_Core *CoreSession) IsUserAdmin(arg0 common.Address) (*types.Transaction, error) {
	return _Core.Contract.IsUserAdmin(&_Core.TransactOpts, arg0)
}

// IsUserAdmin is a paid mutator transaction binding the contract method 0x74d523a8.
//
// Solidity: function isUserAdmin(address ) returns()
func (_Core *CoreTransactorSession) IsUserAdmin(arg0 common.Address) (*types.Transaction, error) {
	return _Core.Contract.IsUserAdmin(&_Core.TransactOpts, arg0)
}

// IsVerifier is a paid mutator transaction binding the contract method 0x33105218.
//
// Solidity: function isVerifier(address ) returns()
func (_Core *CoreTransactor) IsVerifier(opts *bind.TransactOpts, arg0 common.Address) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "isVerifier", arg0)
}

// IsVerifier is a paid mutator transaction binding the contract method 0x33105218.
//
// Solidity: function isVerifier(address ) returns()
func (_Core *CoreSession) IsVerifier(arg0 common.Address) (*types.Transaction, error) {
	return _Core.Contract.IsVerifier(&_Core.TransactOpts, arg0)
}

// IsVerifier is a paid mutator transaction binding the contract method 0x33105218.
//
// Solidity: function isVerifier(address ) returns()
func (_Core *CoreTransactorSession) IsVerifier(arg0 common.Address) (*types.Transaction, error) {
	return _Core.Contract.IsVerifier(&_Core.TransactOpts, arg0)
}

// MainAcceptGovernance is a paid mutator transaction binding the contract method 0x28700a15.
//
// Solidity: function mainAcceptGovernance() returns()
func (_Core *CoreTransactor) MainAcceptGovernance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "mainAcceptGovernance")
}

// MainAcceptGovernance is a paid mutator transaction binding the contract method 0x28700a15.
//
// Solidity: function mainAcceptGovernance() returns()
func (_Core *CoreSession) MainAcceptGovernance() (*types.Transaction, error) {
	return _Core.Contract.MainAcceptGovernance(&_Core.TransactOpts)
}

// MainAcceptGovernance is a paid mutator transaction binding the contract method 0x28700a15.
//
// Solidity: function mainAcceptGovernance() returns()
func (_Core *CoreTransactorSession) MainAcceptGovernance() (*types.Transaction, error) {
	return _Core.Contract.MainAcceptGovernance(&_Core.TransactOpts)
}

// MainCancelNomination is a paid mutator transaction binding the contract method 0x72eb3688.
//
// Solidity: function mainCancelNomination() returns()
func (_Core *CoreTransactor) MainCancelNomination(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "mainCancelNomination")
}

// MainCancelNomination is a paid mutator transaction binding the contract method 0x72eb3688.
//
// Solidity: function mainCancelNomination() returns()
func (_Core *CoreSession) MainCancelNomination() (*types.Transaction, error) {
	return _Core.Contract.MainCancelNomination(&_Core.TransactOpts)
}

// MainCancelNomination is a paid mutator transaction binding the contract method 0x72eb3688.
//
// Solidity: function mainCancelNomination() returns()
func (_Core *CoreTransactorSession) MainCancelNomination() (*types.Transaction, error) {
	return _Core.Contract.MainCancelNomination(&_Core.TransactOpts)
}

// MainIsGovernor is a paid mutator transaction binding the contract method 0x45f5cd97.
//
// Solidity: function mainIsGovernor(address ) returns()
func (_Core *CoreTransactor) MainIsGovernor(opts *bind.TransactOpts, arg0 common.Address) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "mainIsGovernor", arg0)
}

// MainIsGovernor is a paid mutator transaction binding the contract method 0x45f5cd97.
//
// Solidity: function mainIsGovernor(address ) returns()
func (_Core *CoreSession) MainIsGovernor(arg0 common.Address) (*types.Transaction, error) {
	return _Core.Contract.MainIsGovernor(&_Core.TransactOpts, arg0)
}

// MainIsGovernor is a paid mutator transaction binding the contract method 0x45f5cd97.
//
// Solidity: function mainIsGovernor(address ) returns()
func (_Core *CoreTransactorSession) MainIsGovernor(arg0 common.Address) (*types.Transaction, error) {
	return _Core.Contract.MainIsGovernor(&_Core.TransactOpts, arg0)
}

// MainNominateNewGovernor is a paid mutator transaction binding the contract method 0x8c4bce1c.
//
// Solidity: function mainNominateNewGovernor(address ) returns()
func (_Core *CoreTransactor) MainNominateNewGovernor(opts *bind.TransactOpts, arg0 common.Address) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "mainNominateNewGovernor", arg0)
}

// MainNominateNewGovernor is a paid mutator transaction binding the contract method 0x8c4bce1c.
//
// Solidity: function mainNominateNewGovernor(address ) returns()
func (_Core *CoreSession) MainNominateNewGovernor(arg0 common.Address) (*types.Transaction, error) {
	return _Core.Contract.MainNominateNewGovernor(&_Core.TransactOpts, arg0)
}

// MainNominateNewGovernor is a paid mutator transaction binding the contract method 0x8c4bce1c.
//
// Solidity: function mainNominateNewGovernor(address ) returns()
func (_Core *CoreTransactorSession) MainNominateNewGovernor(arg0 common.Address) (*types.Transaction, error) {
	return _Core.Contract.MainNominateNewGovernor(&_Core.TransactOpts, arg0)
}

// MainRemoveGovernor is a paid mutator transaction binding the contract method 0xa1cc921e.
//
// Solidity: function mainRemoveGovernor(address ) returns()
func (_Core *CoreTransactor) MainRemoveGovernor(opts *bind.TransactOpts, arg0 common.Address) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "mainRemoveGovernor", arg0)
}

// MainRemoveGovernor is a paid mutator transaction binding the contract method 0xa1cc921e.
//
// Solidity: function mainRemoveGovernor(address ) returns()
func (_Core *CoreSession) MainRemoveGovernor(arg0 common.Address) (*types.Transaction, error) {
	return _Core.Contract.MainRemoveGovernor(&_Core.TransactOpts, arg0)
}

// MainRemoveGovernor is a paid mutator transaction binding the contract method 0xa1cc921e.
//
// Solidity: function mainRemoveGovernor(address ) returns()
func (_Core *CoreTransactorSession) MainRemoveGovernor(arg0 common.Address) (*types.Transaction, error) {
	return _Core.Contract.MainRemoveGovernor(&_Core.TransactOpts, arg0)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns()
func (_Core *CoreTransactor) OnERC721Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "onERC721Received", arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns()
func (_Core *CoreSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _Core.Contract.OnERC721Received(&_Core.TransactOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns()
func (_Core *CoreTransactorSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _Core.Contract.OnERC721Received(&_Core.TransactOpts, arg0, arg1, arg2, arg3)
}

// RegisterAndDepositERC20 is a paid mutator transaction binding the contract method 0xd5280589.
//
// Solidity: function registerAndDepositERC20(address ethKey, uint256 starkKey, bytes signature, uint256 assetType, uint256 vaultId, uint256 quantizedAmount) returns()
func (_Core *CoreTransactor) RegisterAndDepositERC20(opts *bind.TransactOpts, ethKey common.Address, starkKey *big.Int, signature []byte, assetType *big.Int, vaultId *big.Int, quantizedAmount *big.Int) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "registerAndDepositERC20", ethKey, starkKey, signature, assetType, vaultId, quantizedAmount)
}

// RegisterAndDepositERC20 is a paid mutator transaction binding the contract method 0xd5280589.
//
// Solidity: function registerAndDepositERC20(address ethKey, uint256 starkKey, bytes signature, uint256 assetType, uint256 vaultId, uint256 quantizedAmount) returns()
func (_Core *CoreSession) RegisterAndDepositERC20(ethKey common.Address, starkKey *big.Int, signature []byte, assetType *big.Int, vaultId *big.Int, quantizedAmount *big.Int) (*types.Transaction, error) {
	return _Core.Contract.RegisterAndDepositERC20(&_Core.TransactOpts, ethKey, starkKey, signature, assetType, vaultId, quantizedAmount)
}

// RegisterAndDepositERC20 is a paid mutator transaction binding the contract method 0xd5280589.
//
// Solidity: function registerAndDepositERC20(address ethKey, uint256 starkKey, bytes signature, uint256 assetType, uint256 vaultId, uint256 quantizedAmount) returns()
func (_Core *CoreTransactorSession) RegisterAndDepositERC20(ethKey common.Address, starkKey *big.Int, signature []byte, assetType *big.Int, vaultId *big.Int, quantizedAmount *big.Int) (*types.Transaction, error) {
	return _Core.Contract.RegisterAndDepositERC20(&_Core.TransactOpts, ethKey, starkKey, signature, assetType, vaultId, quantizedAmount)
}

// RegisterAndDepositEth is a paid mutator transaction binding the contract method 0x3ccfc8ed.
//
// Solidity: function registerAndDepositEth(address ethKey, uint256 starkKey, bytes signature, uint256 assetType, uint256 vaultId) payable returns()
func (_Core *CoreTransactor) RegisterAndDepositEth(opts *bind.TransactOpts, ethKey common.Address, starkKey *big.Int, signature []byte, assetType *big.Int, vaultId *big.Int) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "registerAndDepositEth", ethKey, starkKey, signature, assetType, vaultId)
}

// RegisterAndDepositEth is a paid mutator transaction binding the contract method 0x3ccfc8ed.
//
// Solidity: function registerAndDepositEth(address ethKey, uint256 starkKey, bytes signature, uint256 assetType, uint256 vaultId) payable returns()
func (_Core *CoreSession) RegisterAndDepositEth(ethKey common.Address, starkKey *big.Int, signature []byte, assetType *big.Int, vaultId *big.Int) (*types.Transaction, error) {
	return _Core.Contract.RegisterAndDepositEth(&_Core.TransactOpts, ethKey, starkKey, signature, assetType, vaultId)
}

// RegisterAndDepositEth is a paid mutator transaction binding the contract method 0x3ccfc8ed.
//
// Solidity: function registerAndDepositEth(address ethKey, uint256 starkKey, bytes signature, uint256 assetType, uint256 vaultId) payable returns()
func (_Core *CoreTransactorSession) RegisterAndDepositEth(ethKey common.Address, starkKey *big.Int, signature []byte, assetType *big.Int, vaultId *big.Int) (*types.Transaction, error) {
	return _Core.Contract.RegisterAndDepositEth(&_Core.TransactOpts, ethKey, starkKey, signature, assetType, vaultId)
}

// RegisterAvailabilityVerifier is a paid mutator transaction binding the contract method 0xbdb75785.
//
// Solidity: function registerAvailabilityVerifier(address , string ) returns()
func (_Core *CoreTransactor) RegisterAvailabilityVerifier(opts *bind.TransactOpts, arg0 common.Address, arg1 string) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "registerAvailabilityVerifier", arg0, arg1)
}

// RegisterAvailabilityVerifier is a paid mutator transaction binding the contract method 0xbdb75785.
//
// Solidity: function registerAvailabilityVerifier(address , string ) returns()
func (_Core *CoreSession) RegisterAvailabilityVerifier(arg0 common.Address, arg1 string) (*types.Transaction, error) {
	return _Core.Contract.RegisterAvailabilityVerifier(&_Core.TransactOpts, arg0, arg1)
}

// RegisterAvailabilityVerifier is a paid mutator transaction binding the contract method 0xbdb75785.
//
// Solidity: function registerAvailabilityVerifier(address , string ) returns()
func (_Core *CoreTransactorSession) RegisterAvailabilityVerifier(arg0 common.Address, arg1 string) (*types.Transaction, error) {
	return _Core.Contract.RegisterAvailabilityVerifier(&_Core.TransactOpts, arg0, arg1)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0x3682a450.
//
// Solidity: function registerOperator(address ) returns()
func (_Core *CoreTransactor) RegisterOperator(opts *bind.TransactOpts, arg0 common.Address) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "registerOperator", arg0)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0x3682a450.
//
// Solidity: function registerOperator(address ) returns()
func (_Core *CoreSession) RegisterOperator(arg0 common.Address) (*types.Transaction, error) {
	return _Core.Contract.RegisterOperator(&_Core.TransactOpts, arg0)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0x3682a450.
//
// Solidity: function registerOperator(address ) returns()
func (_Core *CoreTransactorSession) RegisterOperator(arg0 common.Address) (*types.Transaction, error) {
	return _Core.Contract.RegisterOperator(&_Core.TransactOpts, arg0)
}

// RegisterToken is a paid mutator transaction binding the contract method 0xc8b1031a.
//
// Solidity: function registerToken(uint256 , bytes ) returns()
func (_Core *CoreTransactor) RegisterToken(opts *bind.TransactOpts, arg0 *big.Int, arg1 []byte) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "registerToken", arg0, arg1)
}

// RegisterToken is a paid mutator transaction binding the contract method 0xc8b1031a.
//
// Solidity: function registerToken(uint256 , bytes ) returns()
func (_Core *CoreSession) RegisterToken(arg0 *big.Int, arg1 []byte) (*types.Transaction, error) {
	return _Core.Contract.RegisterToken(&_Core.TransactOpts, arg0, arg1)
}

// RegisterToken is a paid mutator transaction binding the contract method 0xc8b1031a.
//
// Solidity: function registerToken(uint256 , bytes ) returns()
func (_Core *CoreTransactorSession) RegisterToken(arg0 *big.Int, arg1 []byte) (*types.Transaction, error) {
	return _Core.Contract.RegisterToken(&_Core.TransactOpts, arg0, arg1)
}

// RegisterTokenAdmin is a paid mutator transaction binding the contract method 0x0b3a2d21.
//
// Solidity: function registerTokenAdmin(address ) returns()
func (_Core *CoreTransactor) RegisterTokenAdmin(opts *bind.TransactOpts, arg0 common.Address) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "registerTokenAdmin", arg0)
}

// RegisterTokenAdmin is a paid mutator transaction binding the contract method 0x0b3a2d21.
//
// Solidity: function registerTokenAdmin(address ) returns()
func (_Core *CoreSession) RegisterTokenAdmin(arg0 common.Address) (*types.Transaction, error) {
	return _Core.Contract.RegisterTokenAdmin(&_Core.TransactOpts, arg0)
}

// RegisterTokenAdmin is a paid mutator transaction binding the contract method 0x0b3a2d21.
//
// Solidity: function registerTokenAdmin(address ) returns()
func (_Core *CoreTransactorSession) RegisterTokenAdmin(arg0 common.Address) (*types.Transaction, error) {
	return _Core.Contract.RegisterTokenAdmin(&_Core.TransactOpts, arg0)
}

// RegisterUser is a paid mutator transaction binding the contract method 0xdd2414d4.
//
// Solidity: function registerUser(address , uint256 , bytes ) returns()
func (_Core *CoreTransactor) RegisterUser(opts *bind.TransactOpts, arg0 common.Address, arg1 *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "registerUser", arg0, arg1, arg2)
}

// RegisterUser is a paid mutator transaction binding the contract method 0xdd2414d4.
//
// Solidity: function registerUser(address , uint256 , bytes ) returns()
func (_Core *CoreSession) RegisterUser(arg0 common.Address, arg1 *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _Core.Contract.RegisterUser(&_Core.TransactOpts, arg0, arg1, arg2)
}

// RegisterUser is a paid mutator transaction binding the contract method 0xdd2414d4.
//
// Solidity: function registerUser(address , uint256 , bytes ) returns()
func (_Core *CoreTransactorSession) RegisterUser(arg0 common.Address, arg1 *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _Core.Contract.RegisterUser(&_Core.TransactOpts, arg0, arg1, arg2)
}

// RegisterUserAdmin is a paid mutator transaction binding the contract method 0xf3e0c3fb.
//
// Solidity: function registerUserAdmin(address ) returns()
func (_Core *CoreTransactor) RegisterUserAdmin(opts *bind.TransactOpts, arg0 common.Address) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "registerUserAdmin", arg0)
}

// RegisterUserAdmin is a paid mutator transaction binding the contract method 0xf3e0c3fb.
//
// Solidity: function registerUserAdmin(address ) returns()
func (_Core *CoreSession) RegisterUserAdmin(arg0 common.Address) (*types.Transaction, error) {
	return _Core.Contract.RegisterUserAdmin(&_Core.TransactOpts, arg0)
}

// RegisterUserAdmin is a paid mutator transaction binding the contract method 0xf3e0c3fb.
//
// Solidity: function registerUserAdmin(address ) returns()
func (_Core *CoreTransactorSession) RegisterUserAdmin(arg0 common.Address) (*types.Transaction, error) {
	return _Core.Contract.RegisterUserAdmin(&_Core.TransactOpts, arg0)
}

// RegisterVerifier is a paid mutator transaction binding the contract method 0x3776fe2a.
//
// Solidity: function registerVerifier(address , string ) returns()
func (_Core *CoreTransactor) RegisterVerifier(opts *bind.TransactOpts, arg0 common.Address, arg1 string) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "registerVerifier", arg0, arg1)
}

// RegisterVerifier is a paid mutator transaction binding the contract method 0x3776fe2a.
//
// Solidity: function registerVerifier(address , string ) returns()
func (_Core *CoreSession) RegisterVerifier(arg0 common.Address, arg1 string) (*types.Transaction, error) {
	return _Core.Contract.RegisterVerifier(&_Core.TransactOpts, arg0, arg1)
}

// RegisterVerifier is a paid mutator transaction binding the contract method 0x3776fe2a.
//
// Solidity: function registerVerifier(address , string ) returns()
func (_Core *CoreTransactorSession) RegisterVerifier(arg0 common.Address, arg1 string) (*types.Transaction, error) {
	return _Core.Contract.RegisterVerifier(&_Core.TransactOpts, arg0, arg1)
}

// RemoveAvailabilityVerifier is a paid mutator transaction binding the contract method 0xb1e640bf.
//
// Solidity: function removeAvailabilityVerifier(address ) returns()
func (_Core *CoreTransactor) RemoveAvailabilityVerifier(opts *bind.TransactOpts, arg0 common.Address) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "removeAvailabilityVerifier", arg0)
}

// RemoveAvailabilityVerifier is a paid mutator transaction binding the contract method 0xb1e640bf.
//
// Solidity: function removeAvailabilityVerifier(address ) returns()
func (_Core *CoreSession) RemoveAvailabilityVerifier(arg0 common.Address) (*types.Transaction, error) {
	return _Core.Contract.RemoveAvailabilityVerifier(&_Core.TransactOpts, arg0)
}

// RemoveAvailabilityVerifier is a paid mutator transaction binding the contract method 0xb1e640bf.
//
// Solidity: function removeAvailabilityVerifier(address ) returns()
func (_Core *CoreTransactorSession) RemoveAvailabilityVerifier(arg0 common.Address) (*types.Transaction, error) {
	return _Core.Contract.RemoveAvailabilityVerifier(&_Core.TransactOpts, arg0)
}

// RemoveVerifier is a paid mutator transaction binding the contract method 0xca2dfd0a.
//
// Solidity: function removeVerifier(address ) returns()
func (_Core *CoreTransactor) RemoveVerifier(opts *bind.TransactOpts, arg0 common.Address) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "removeVerifier", arg0)
}

// RemoveVerifier is a paid mutator transaction binding the contract method 0xca2dfd0a.
//
// Solidity: function removeVerifier(address ) returns()
func (_Core *CoreSession) RemoveVerifier(arg0 common.Address) (*types.Transaction, error) {
	return _Core.Contract.RemoveVerifier(&_Core.TransactOpts, arg0)
}

// RemoveVerifier is a paid mutator transaction binding the contract method 0xca2dfd0a.
//
// Solidity: function removeVerifier(address ) returns()
func (_Core *CoreTransactorSession) RemoveVerifier(arg0 common.Address) (*types.Transaction, error) {
	return _Core.Contract.RemoveVerifier(&_Core.TransactOpts, arg0)
}

// UnFreeze is a paid mutator transaction binding the contract method 0x7cf12b90.
//
// Solidity: function unFreeze() returns()
func (_Core *CoreTransactor) UnFreeze(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "unFreeze")
}

// UnFreeze is a paid mutator transaction binding the contract method 0x7cf12b90.
//
// Solidity: function unFreeze() returns()
func (_Core *CoreSession) UnFreeze() (*types.Transaction, error) {
	return _Core.Contract.UnFreeze(&_Core.TransactOpts)
}

// UnFreeze is a paid mutator transaction binding the contract method 0x7cf12b90.
//
// Solidity: function unFreeze() returns()
func (_Core *CoreTransactorSession) UnFreeze() (*types.Transaction, error) {
	return _Core.Contract.UnFreeze(&_Core.TransactOpts)
}

// UnregisterOperator is a paid mutator transaction binding the contract method 0x96115bc2.
//
// Solidity: function unregisterOperator(address ) returns()
func (_Core *CoreTransactor) UnregisterOperator(opts *bind.TransactOpts, arg0 common.Address) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "unregisterOperator", arg0)
}

// UnregisterOperator is a paid mutator transaction binding the contract method 0x96115bc2.
//
// Solidity: function unregisterOperator(address ) returns()
func (_Core *CoreSession) UnregisterOperator(arg0 common.Address) (*types.Transaction, error) {
	return _Core.Contract.UnregisterOperator(&_Core.TransactOpts, arg0)
}

// UnregisterOperator is a paid mutator transaction binding the contract method 0x96115bc2.
//
// Solidity: function unregisterOperator(address ) returns()
func (_Core *CoreTransactorSession) UnregisterOperator(arg0 common.Address) (*types.Transaction, error) {
	return _Core.Contract.UnregisterOperator(&_Core.TransactOpts, arg0)
}

// UnregisterTokenAdmin is a paid mutator transaction binding the contract method 0xa6fa6e90.
//
// Solidity: function unregisterTokenAdmin(address ) returns()
func (_Core *CoreTransactor) UnregisterTokenAdmin(opts *bind.TransactOpts, arg0 common.Address) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "unregisterTokenAdmin", arg0)
}

// UnregisterTokenAdmin is a paid mutator transaction binding the contract method 0xa6fa6e90.
//
// Solidity: function unregisterTokenAdmin(address ) returns()
func (_Core *CoreSession) UnregisterTokenAdmin(arg0 common.Address) (*types.Transaction, error) {
	return _Core.Contract.UnregisterTokenAdmin(&_Core.TransactOpts, arg0)
}

// UnregisterTokenAdmin is a paid mutator transaction binding the contract method 0xa6fa6e90.
//
// Solidity: function unregisterTokenAdmin(address ) returns()
func (_Core *CoreTransactorSession) UnregisterTokenAdmin(arg0 common.Address) (*types.Transaction, error) {
	return _Core.Contract.UnregisterTokenAdmin(&_Core.TransactOpts, arg0)
}

// UnregisterUserAdmin is a paid mutator transaction binding the contract method 0xb04b6179.
//
// Solidity: function unregisterUserAdmin(address ) returns()
func (_Core *CoreTransactor) UnregisterUserAdmin(opts *bind.TransactOpts, arg0 common.Address) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "unregisterUserAdmin", arg0)
}

// UnregisterUserAdmin is a paid mutator transaction binding the contract method 0xb04b6179.
//
// Solidity: function unregisterUserAdmin(address ) returns()
func (_Core *CoreSession) UnregisterUserAdmin(arg0 common.Address) (*types.Transaction, error) {
	return _Core.Contract.UnregisterUserAdmin(&_Core.TransactOpts, arg0)
}

// UnregisterUserAdmin is a paid mutator transaction binding the contract method 0xb04b6179.
//
// Solidity: function unregisterUserAdmin(address ) returns()
func (_Core *CoreTransactorSession) UnregisterUserAdmin(arg0 common.Address) (*types.Transaction, error) {
	return _Core.Contract.UnregisterUserAdmin(&_Core.TransactOpts, arg0)
}

// UpdateState is a paid mutator transaction binding the contract method 0x538f9406.
//
// Solidity: function updateState(uint256[] publicInput, uint256[] applicationData) returns()
func (_Core *CoreTransactor) UpdateState(opts *bind.TransactOpts, publicInput []*big.Int, applicationData []*big.Int) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "updateState", publicInput, applicationData)
}

// UpdateState is a paid mutator transaction binding the contract method 0x538f9406.
//
// Solidity: function updateState(uint256[] publicInput, uint256[] applicationData) returns()
func (_Core *CoreSession) UpdateState(publicInput []*big.Int, applicationData []*big.Int) (*types.Transaction, error) {
	return _Core.Contract.UpdateState(&_Core.TransactOpts, publicInput, applicationData)
}

// UpdateState is a paid mutator transaction binding the contract method 0x538f9406.
//
// Solidity: function updateState(uint256[] publicInput, uint256[] applicationData) returns()
func (_Core *CoreTransactorSession) UpdateState(publicInput []*big.Int, applicationData []*big.Int) (*types.Transaction, error) {
	return _Core.Contract.UpdateState(&_Core.TransactOpts, publicInput, applicationData)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 ownerKey, uint256 assetType) returns()
func (_Core *CoreTransactor) Withdraw(opts *bind.TransactOpts, ownerKey *big.Int, assetType *big.Int) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "withdraw", ownerKey, assetType)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 ownerKey, uint256 assetType) returns()
func (_Core *CoreSession) Withdraw(ownerKey *big.Int, assetType *big.Int) (*types.Transaction, error) {
	return _Core.Contract.Withdraw(&_Core.TransactOpts, ownerKey, assetType)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 ownerKey, uint256 assetType) returns()
func (_Core *CoreTransactorSession) Withdraw(ownerKey *big.Int, assetType *big.Int) (*types.Transaction, error) {
	return _Core.Contract.Withdraw(&_Core.TransactOpts, ownerKey, assetType)
}

// WithdrawAndMint is a paid mutator transaction binding the contract method 0xd91443b7.
//
// Solidity: function withdrawAndMint(uint256 ownerKey, uint256 assetType, bytes mintingBlob) returns()
func (_Core *CoreTransactor) WithdrawAndMint(opts *bind.TransactOpts, ownerKey *big.Int, assetType *big.Int, mintingBlob []byte) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "withdrawAndMint", ownerKey, assetType, mintingBlob)
}

// WithdrawAndMint is a paid mutator transaction binding the contract method 0xd91443b7.
//
// Solidity: function withdrawAndMint(uint256 ownerKey, uint256 assetType, bytes mintingBlob) returns()
func (_Core *CoreSession) WithdrawAndMint(ownerKey *big.Int, assetType *big.Int, mintingBlob []byte) (*types.Transaction, error) {
	return _Core.Contract.WithdrawAndMint(&_Core.TransactOpts, ownerKey, assetType, mintingBlob)
}

// WithdrawAndMint is a paid mutator transaction binding the contract method 0xd91443b7.
//
// Solidity: function withdrawAndMint(uint256 ownerKey, uint256 assetType, bytes mintingBlob) returns()
func (_Core *CoreTransactorSession) WithdrawAndMint(ownerKey *big.Int, assetType *big.Int, mintingBlob []byte) (*types.Transaction, error) {
	return _Core.Contract.WithdrawAndMint(&_Core.TransactOpts, ownerKey, assetType, mintingBlob)
}

// WithdrawNft is a paid mutator transaction binding the contract method 0x019b417a.
//
// Solidity: function withdrawNft(uint256 ownerKey, uint256 assetType, uint256 tokenId) returns()
func (_Core *CoreTransactor) WithdrawNft(opts *bind.TransactOpts, ownerKey *big.Int, assetType *big.Int, tokenId *big.Int) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "withdrawNft", ownerKey, assetType, tokenId)
}

// WithdrawNft is a paid mutator transaction binding the contract method 0x019b417a.
//
// Solidity: function withdrawNft(uint256 ownerKey, uint256 assetType, uint256 tokenId) returns()
func (_Core *CoreSession) WithdrawNft(ownerKey *big.Int, assetType *big.Int, tokenId *big.Int) (*types.Transaction, error) {
	return _Core.Contract.WithdrawNft(&_Core.TransactOpts, ownerKey, assetType, tokenId)
}

// WithdrawNft is a paid mutator transaction binding the contract method 0x019b417a.
//
// Solidity: function withdrawNft(uint256 ownerKey, uint256 assetType, uint256 tokenId) returns()
func (_Core *CoreTransactorSession) WithdrawNft(ownerKey *big.Int, assetType *big.Int, tokenId *big.Int) (*types.Transaction, error) {
	return _Core.Contract.WithdrawNft(&_Core.TransactOpts, ownerKey, assetType, tokenId)
}

// WithdrawNftTo is a paid mutator transaction binding the contract method 0xebef0fd0.
//
// Solidity: function withdrawNftTo(uint256 , uint256 , uint256 , address ) returns()
func (_Core *CoreTransactor) WithdrawNftTo(opts *bind.TransactOpts, arg0 *big.Int, arg1 *big.Int, arg2 *big.Int, arg3 common.Address) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "withdrawNftTo", arg0, arg1, arg2, arg3)
}

// WithdrawNftTo is a paid mutator transaction binding the contract method 0xebef0fd0.
//
// Solidity: function withdrawNftTo(uint256 , uint256 , uint256 , address ) returns()
func (_Core *CoreSession) WithdrawNftTo(arg0 *big.Int, arg1 *big.Int, arg2 *big.Int, arg3 common.Address) (*types.Transaction, error) {
	return _Core.Contract.WithdrawNftTo(&_Core.TransactOpts, arg0, arg1, arg2, arg3)
}

// WithdrawNftTo is a paid mutator transaction binding the contract method 0xebef0fd0.
//
// Solidity: function withdrawNftTo(uint256 , uint256 , uint256 , address ) returns()
func (_Core *CoreTransactorSession) WithdrawNftTo(arg0 *big.Int, arg1 *big.Int, arg2 *big.Int, arg3 common.Address) (*types.Transaction, error) {
	return _Core.Contract.WithdrawNftTo(&_Core.TransactOpts, arg0, arg1, arg2, arg3)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0x14cd70e4.
//
// Solidity: function withdrawTo(uint256 , uint256 , address ) returns()
func (_Core *CoreTransactor) WithdrawTo(opts *bind.TransactOpts, arg0 *big.Int, arg1 *big.Int, arg2 common.Address) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "withdrawTo", arg0, arg1, arg2)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0x14cd70e4.
//
// Solidity: function withdrawTo(uint256 , uint256 , address ) returns()
func (_Core *CoreSession) WithdrawTo(arg0 *big.Int, arg1 *big.Int, arg2 common.Address) (*types.Transaction, error) {
	return _Core.Contract.WithdrawTo(&_Core.TransactOpts, arg0, arg1, arg2)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0x14cd70e4.
//
// Solidity: function withdrawTo(uint256 , uint256 , address ) returns()
func (_Core *CoreTransactorSession) WithdrawTo(arg0 *big.Int, arg1 *big.Int, arg2 common.Address) (*types.Transaction, error) {
	return _Core.Contract.WithdrawTo(&_Core.TransactOpts, arg0, arg1, arg2)
}

// CoreLogDepositIterator is returned from FilterLogDeposit and is used to iterate over the raw logs and unpacked data for LogDeposit events raised by the Core contract.
type CoreLogDepositIterator struct {
	Event *CoreLogDeposit // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CoreLogDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreLogDeposit)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CoreLogDeposit)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CoreLogDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreLogDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreLogDeposit represents a LogDeposit event raised by the Core contract.
type CoreLogDeposit struct {
	DepositorEthKey    common.Address
	StarkKey           *big.Int
	VaultId            *big.Int
	AssetType          *big.Int
	NonQuantizedAmount *big.Int
	QuantizedAmount    *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterLogDeposit is a free log retrieval operation binding the contract event 0x06724742ccc8c330a39a641ef02a0b419bd09248360680bb38159b0a8c2635d6.
//
// Solidity: event LogDeposit(address depositorEthKey, uint256 starkKey, uint256 vaultId, uint256 assetType, uint256 nonQuantizedAmount, uint256 quantizedAmount)
func (_Core *CoreFilterer) FilterLogDeposit(opts *bind.FilterOpts) (*CoreLogDepositIterator, error) {

	logs, sub, err := _Core.contract.FilterLogs(opts, "LogDeposit")
	if err != nil {
		return nil, err
	}
	return &CoreLogDepositIterator{contract: _Core.contract, event: "LogDeposit", logs: logs, sub: sub}, nil
}

// WatchLogDeposit is a free log subscription operation binding the contract event 0x06724742ccc8c330a39a641ef02a0b419bd09248360680bb38159b0a8c2635d6.
//
// Solidity: event LogDeposit(address depositorEthKey, uint256 starkKey, uint256 vaultId, uint256 assetType, uint256 nonQuantizedAmount, uint256 quantizedAmount)
func (_Core *CoreFilterer) WatchLogDeposit(opts *bind.WatchOpts, sink chan<- *CoreLogDeposit) (event.Subscription, error) {

	logs, sub, err := _Core.contract.WatchLogs(opts, "LogDeposit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreLogDeposit)
				if err := _Core.contract.UnpackLog(event, "LogDeposit", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogDeposit is a log parse operation binding the contract event 0x06724742ccc8c330a39a641ef02a0b419bd09248360680bb38159b0a8c2635d6.
//
// Solidity: event LogDeposit(address depositorEthKey, uint256 starkKey, uint256 vaultId, uint256 assetType, uint256 nonQuantizedAmount, uint256 quantizedAmount)
func (_Core *CoreFilterer) ParseLogDeposit(log types.Log) (*CoreLogDeposit, error) {
	event := new(CoreLogDeposit)
	if err := _Core.contract.UnpackLog(event, "LogDeposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreLogDepositCancelIterator is returned from FilterLogDepositCancel and is used to iterate over the raw logs and unpacked data for LogDepositCancel events raised by the Core contract.
type CoreLogDepositCancelIterator struct {
	Event *CoreLogDepositCancel // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CoreLogDepositCancelIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreLogDepositCancel)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CoreLogDepositCancel)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CoreLogDepositCancelIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreLogDepositCancelIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreLogDepositCancel represents a LogDepositCancel event raised by the Core contract.
type CoreLogDepositCancel struct {
	StarkKey *big.Int
	VaultId  *big.Int
	AssetId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLogDepositCancel is a free log retrieval operation binding the contract event 0x0bc1df35228095c37da66a6ffcc755ea79dfc437345685f618e05fafad6b445e.
//
// Solidity: event LogDepositCancel(uint256 starkKey, uint256 vaultId, uint256 assetId)
func (_Core *CoreFilterer) FilterLogDepositCancel(opts *bind.FilterOpts) (*CoreLogDepositCancelIterator, error) {

	logs, sub, err := _Core.contract.FilterLogs(opts, "LogDepositCancel")
	if err != nil {
		return nil, err
	}
	return &CoreLogDepositCancelIterator{contract: _Core.contract, event: "LogDepositCancel", logs: logs, sub: sub}, nil
}

// WatchLogDepositCancel is a free log subscription operation binding the contract event 0x0bc1df35228095c37da66a6ffcc755ea79dfc437345685f618e05fafad6b445e.
//
// Solidity: event LogDepositCancel(uint256 starkKey, uint256 vaultId, uint256 assetId)
func (_Core *CoreFilterer) WatchLogDepositCancel(opts *bind.WatchOpts, sink chan<- *CoreLogDepositCancel) (event.Subscription, error) {

	logs, sub, err := _Core.contract.WatchLogs(opts, "LogDepositCancel")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreLogDepositCancel)
				if err := _Core.contract.UnpackLog(event, "LogDepositCancel", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogDepositCancel is a log parse operation binding the contract event 0x0bc1df35228095c37da66a6ffcc755ea79dfc437345685f618e05fafad6b445e.
//
// Solidity: event LogDepositCancel(uint256 starkKey, uint256 vaultId, uint256 assetId)
func (_Core *CoreFilterer) ParseLogDepositCancel(log types.Log) (*CoreLogDepositCancel, error) {
	event := new(CoreLogDepositCancel)
	if err := _Core.contract.UnpackLog(event, "LogDepositCancel", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreLogDepositCancelReclaimedIterator is returned from FilterLogDepositCancelReclaimed and is used to iterate over the raw logs and unpacked data for LogDepositCancelReclaimed events raised by the Core contract.
type CoreLogDepositCancelReclaimedIterator struct {
	Event *CoreLogDepositCancelReclaimed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CoreLogDepositCancelReclaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreLogDepositCancelReclaimed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CoreLogDepositCancelReclaimed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CoreLogDepositCancelReclaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreLogDepositCancelReclaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreLogDepositCancelReclaimed represents a LogDepositCancelReclaimed event raised by the Core contract.
type CoreLogDepositCancelReclaimed struct {
	StarkKey           *big.Int
	VaultId            *big.Int
	AssetType          *big.Int
	NonQuantizedAmount *big.Int
	QuantizedAmount    *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterLogDepositCancelReclaimed is a free log retrieval operation binding the contract event 0xe3e46ecf1138180bf93cba62a0b7e661d976a8ab3d40243f7b082667d8f500af.
//
// Solidity: event LogDepositCancelReclaimed(uint256 starkKey, uint256 vaultId, uint256 assetType, uint256 nonQuantizedAmount, uint256 quantizedAmount)
func (_Core *CoreFilterer) FilterLogDepositCancelReclaimed(opts *bind.FilterOpts) (*CoreLogDepositCancelReclaimedIterator, error) {

	logs, sub, err := _Core.contract.FilterLogs(opts, "LogDepositCancelReclaimed")
	if err != nil {
		return nil, err
	}
	return &CoreLogDepositCancelReclaimedIterator{contract: _Core.contract, event: "LogDepositCancelReclaimed", logs: logs, sub: sub}, nil
}

// WatchLogDepositCancelReclaimed is a free log subscription operation binding the contract event 0xe3e46ecf1138180bf93cba62a0b7e661d976a8ab3d40243f7b082667d8f500af.
//
// Solidity: event LogDepositCancelReclaimed(uint256 starkKey, uint256 vaultId, uint256 assetType, uint256 nonQuantizedAmount, uint256 quantizedAmount)
func (_Core *CoreFilterer) WatchLogDepositCancelReclaimed(opts *bind.WatchOpts, sink chan<- *CoreLogDepositCancelReclaimed) (event.Subscription, error) {

	logs, sub, err := _Core.contract.WatchLogs(opts, "LogDepositCancelReclaimed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreLogDepositCancelReclaimed)
				if err := _Core.contract.UnpackLog(event, "LogDepositCancelReclaimed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogDepositCancelReclaimed is a log parse operation binding the contract event 0xe3e46ecf1138180bf93cba62a0b7e661d976a8ab3d40243f7b082667d8f500af.
//
// Solidity: event LogDepositCancelReclaimed(uint256 starkKey, uint256 vaultId, uint256 assetType, uint256 nonQuantizedAmount, uint256 quantizedAmount)
func (_Core *CoreFilterer) ParseLogDepositCancelReclaimed(log types.Log) (*CoreLogDepositCancelReclaimed, error) {
	event := new(CoreLogDepositCancelReclaimed)
	if err := _Core.contract.UnpackLog(event, "LogDepositCancelReclaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreLogDepositNftCancelReclaimedIterator is returned from FilterLogDepositNftCancelReclaimed and is used to iterate over the raw logs and unpacked data for LogDepositNftCancelReclaimed events raised by the Core contract.
type CoreLogDepositNftCancelReclaimedIterator struct {
	Event *CoreLogDepositNftCancelReclaimed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CoreLogDepositNftCancelReclaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreLogDepositNftCancelReclaimed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CoreLogDepositNftCancelReclaimed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CoreLogDepositNftCancelReclaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreLogDepositNftCancelReclaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreLogDepositNftCancelReclaimed represents a LogDepositNftCancelReclaimed event raised by the Core contract.
type CoreLogDepositNftCancelReclaimed struct {
	StarkKey  *big.Int
	VaultId   *big.Int
	AssetType *big.Int
	TokenId   *big.Int
	AssetId   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLogDepositNftCancelReclaimed is a free log retrieval operation binding the contract event 0xf00c0c1a754f6df7545d96a7e12aad552728b94ca6aa94f81e297bdbcf1dab9c.
//
// Solidity: event LogDepositNftCancelReclaimed(uint256 starkKey, uint256 vaultId, uint256 assetType, uint256 tokenId, uint256 assetId)
func (_Core *CoreFilterer) FilterLogDepositNftCancelReclaimed(opts *bind.FilterOpts) (*CoreLogDepositNftCancelReclaimedIterator, error) {

	logs, sub, err := _Core.contract.FilterLogs(opts, "LogDepositNftCancelReclaimed")
	if err != nil {
		return nil, err
	}
	return &CoreLogDepositNftCancelReclaimedIterator{contract: _Core.contract, event: "LogDepositNftCancelReclaimed", logs: logs, sub: sub}, nil
}

// WatchLogDepositNftCancelReclaimed is a free log subscription operation binding the contract event 0xf00c0c1a754f6df7545d96a7e12aad552728b94ca6aa94f81e297bdbcf1dab9c.
//
// Solidity: event LogDepositNftCancelReclaimed(uint256 starkKey, uint256 vaultId, uint256 assetType, uint256 tokenId, uint256 assetId)
func (_Core *CoreFilterer) WatchLogDepositNftCancelReclaimed(opts *bind.WatchOpts, sink chan<- *CoreLogDepositNftCancelReclaimed) (event.Subscription, error) {

	logs, sub, err := _Core.contract.WatchLogs(opts, "LogDepositNftCancelReclaimed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreLogDepositNftCancelReclaimed)
				if err := _Core.contract.UnpackLog(event, "LogDepositNftCancelReclaimed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogDepositNftCancelReclaimed is a log parse operation binding the contract event 0xf00c0c1a754f6df7545d96a7e12aad552728b94ca6aa94f81e297bdbcf1dab9c.
//
// Solidity: event LogDepositNftCancelReclaimed(uint256 starkKey, uint256 vaultId, uint256 assetType, uint256 tokenId, uint256 assetId)
func (_Core *CoreFilterer) ParseLogDepositNftCancelReclaimed(log types.Log) (*CoreLogDepositNftCancelReclaimed, error) {
	event := new(CoreLogDepositNftCancelReclaimed)
	if err := _Core.contract.UnpackLog(event, "LogDepositNftCancelReclaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreLogFullWithdrawalRequestIterator is returned from FilterLogFullWithdrawalRequest and is used to iterate over the raw logs and unpacked data for LogFullWithdrawalRequest events raised by the Core contract.
type CoreLogFullWithdrawalRequestIterator struct {
	Event *CoreLogFullWithdrawalRequest // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CoreLogFullWithdrawalRequestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreLogFullWithdrawalRequest)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CoreLogFullWithdrawalRequest)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CoreLogFullWithdrawalRequestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreLogFullWithdrawalRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreLogFullWithdrawalRequest represents a LogFullWithdrawalRequest event raised by the Core contract.
type CoreLogFullWithdrawalRequest struct {
	StarkKey *big.Int
	VaultId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLogFullWithdrawalRequest is a free log retrieval operation binding the contract event 0x08eb46dbb87dcfe92d4846e5766802051525fba08a9b48318f5e0fe41186d298.
//
// Solidity: event LogFullWithdrawalRequest(uint256 starkKey, uint256 vaultId)
func (_Core *CoreFilterer) FilterLogFullWithdrawalRequest(opts *bind.FilterOpts) (*CoreLogFullWithdrawalRequestIterator, error) {

	logs, sub, err := _Core.contract.FilterLogs(opts, "LogFullWithdrawalRequest")
	if err != nil {
		return nil, err
	}
	return &CoreLogFullWithdrawalRequestIterator{contract: _Core.contract, event: "LogFullWithdrawalRequest", logs: logs, sub: sub}, nil
}

// WatchLogFullWithdrawalRequest is a free log subscription operation binding the contract event 0x08eb46dbb87dcfe92d4846e5766802051525fba08a9b48318f5e0fe41186d298.
//
// Solidity: event LogFullWithdrawalRequest(uint256 starkKey, uint256 vaultId)
func (_Core *CoreFilterer) WatchLogFullWithdrawalRequest(opts *bind.WatchOpts, sink chan<- *CoreLogFullWithdrawalRequest) (event.Subscription, error) {

	logs, sub, err := _Core.contract.WatchLogs(opts, "LogFullWithdrawalRequest")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreLogFullWithdrawalRequest)
				if err := _Core.contract.UnpackLog(event, "LogFullWithdrawalRequest", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogFullWithdrawalRequest is a log parse operation binding the contract event 0x08eb46dbb87dcfe92d4846e5766802051525fba08a9b48318f5e0fe41186d298.
//
// Solidity: event LogFullWithdrawalRequest(uint256 starkKey, uint256 vaultId)
func (_Core *CoreFilterer) ParseLogFullWithdrawalRequest(log types.Log) (*CoreLogFullWithdrawalRequest, error) {
	event := new(CoreLogFullWithdrawalRequest)
	if err := _Core.contract.UnpackLog(event, "LogFullWithdrawalRequest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreLogMintWithdrawalPerformedIterator is returned from FilterLogMintWithdrawalPerformed and is used to iterate over the raw logs and unpacked data for LogMintWithdrawalPerformed events raised by the Core contract.
type CoreLogMintWithdrawalPerformedIterator struct {
	Event *CoreLogMintWithdrawalPerformed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CoreLogMintWithdrawalPerformedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreLogMintWithdrawalPerformed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CoreLogMintWithdrawalPerformed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CoreLogMintWithdrawalPerformedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreLogMintWithdrawalPerformedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreLogMintWithdrawalPerformed represents a LogMintWithdrawalPerformed event raised by the Core contract.
type CoreLogMintWithdrawalPerformed struct {
	OwnerKey           *big.Int
	AssetType          *big.Int
	NonQuantizedAmount *big.Int
	QuantizedAmount    *big.Int
	AssetId            *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterLogMintWithdrawalPerformed is a free log retrieval operation binding the contract event 0x7e6e15df814c1a309a57686de672b2bedd128eacde35c5370c36d6840d4e9a92.
//
// Solidity: event LogMintWithdrawalPerformed(uint256 ownerKey, uint256 assetType, uint256 nonQuantizedAmount, uint256 quantizedAmount, uint256 assetId)
func (_Core *CoreFilterer) FilterLogMintWithdrawalPerformed(opts *bind.FilterOpts) (*CoreLogMintWithdrawalPerformedIterator, error) {

	logs, sub, err := _Core.contract.FilterLogs(opts, "LogMintWithdrawalPerformed")
	if err != nil {
		return nil, err
	}
	return &CoreLogMintWithdrawalPerformedIterator{contract: _Core.contract, event: "LogMintWithdrawalPerformed", logs: logs, sub: sub}, nil
}

// WatchLogMintWithdrawalPerformed is a free log subscription operation binding the contract event 0x7e6e15df814c1a309a57686de672b2bedd128eacde35c5370c36d6840d4e9a92.
//
// Solidity: event LogMintWithdrawalPerformed(uint256 ownerKey, uint256 assetType, uint256 nonQuantizedAmount, uint256 quantizedAmount, uint256 assetId)
func (_Core *CoreFilterer) WatchLogMintWithdrawalPerformed(opts *bind.WatchOpts, sink chan<- *CoreLogMintWithdrawalPerformed) (event.Subscription, error) {

	logs, sub, err := _Core.contract.WatchLogs(opts, "LogMintWithdrawalPerformed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreLogMintWithdrawalPerformed)
				if err := _Core.contract.UnpackLog(event, "LogMintWithdrawalPerformed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogMintWithdrawalPerformed is a log parse operation binding the contract event 0x7e6e15df814c1a309a57686de672b2bedd128eacde35c5370c36d6840d4e9a92.
//
// Solidity: event LogMintWithdrawalPerformed(uint256 ownerKey, uint256 assetType, uint256 nonQuantizedAmount, uint256 quantizedAmount, uint256 assetId)
func (_Core *CoreFilterer) ParseLogMintWithdrawalPerformed(log types.Log) (*CoreLogMintWithdrawalPerformed, error) {
	event := new(CoreLogMintWithdrawalPerformed)
	if err := _Core.contract.UnpackLog(event, "LogMintWithdrawalPerformed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreLogMintableWithdrawalAllowedIterator is returned from FilterLogMintableWithdrawalAllowed and is used to iterate over the raw logs and unpacked data for LogMintableWithdrawalAllowed events raised by the Core contract.
type CoreLogMintableWithdrawalAllowedIterator struct {
	Event *CoreLogMintableWithdrawalAllowed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CoreLogMintableWithdrawalAllowedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreLogMintableWithdrawalAllowed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CoreLogMintableWithdrawalAllowed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CoreLogMintableWithdrawalAllowedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreLogMintableWithdrawalAllowedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreLogMintableWithdrawalAllowed represents a LogMintableWithdrawalAllowed event raised by the Core contract.
type CoreLogMintableWithdrawalAllowed struct {
	OwnerKey        *big.Int
	AssetId         *big.Int
	QuantizedAmount *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterLogMintableWithdrawalAllowed is a free log retrieval operation binding the contract event 0x2acce0cedb29dc4927e6c891b57ef5bc8858eea4bf52787ea94873aebd4aeca0.
//
// Solidity: event LogMintableWithdrawalAllowed(uint256 ownerKey, uint256 assetId, uint256 quantizedAmount)
func (_Core *CoreFilterer) FilterLogMintableWithdrawalAllowed(opts *bind.FilterOpts) (*CoreLogMintableWithdrawalAllowedIterator, error) {

	logs, sub, err := _Core.contract.FilterLogs(opts, "LogMintableWithdrawalAllowed")
	if err != nil {
		return nil, err
	}
	return &CoreLogMintableWithdrawalAllowedIterator{contract: _Core.contract, event: "LogMintableWithdrawalAllowed", logs: logs, sub: sub}, nil
}

// WatchLogMintableWithdrawalAllowed is a free log subscription operation binding the contract event 0x2acce0cedb29dc4927e6c891b57ef5bc8858eea4bf52787ea94873aebd4aeca0.
//
// Solidity: event LogMintableWithdrawalAllowed(uint256 ownerKey, uint256 assetId, uint256 quantizedAmount)
func (_Core *CoreFilterer) WatchLogMintableWithdrawalAllowed(opts *bind.WatchOpts, sink chan<- *CoreLogMintableWithdrawalAllowed) (event.Subscription, error) {

	logs, sub, err := _Core.contract.WatchLogs(opts, "LogMintableWithdrawalAllowed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreLogMintableWithdrawalAllowed)
				if err := _Core.contract.UnpackLog(event, "LogMintableWithdrawalAllowed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogMintableWithdrawalAllowed is a log parse operation binding the contract event 0x2acce0cedb29dc4927e6c891b57ef5bc8858eea4bf52787ea94873aebd4aeca0.
//
// Solidity: event LogMintableWithdrawalAllowed(uint256 ownerKey, uint256 assetId, uint256 quantizedAmount)
func (_Core *CoreFilterer) ParseLogMintableWithdrawalAllowed(log types.Log) (*CoreLogMintableWithdrawalAllowed, error) {
	event := new(CoreLogMintableWithdrawalAllowed)
	if err := _Core.contract.UnpackLog(event, "LogMintableWithdrawalAllowed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreLogNftDepositIterator is returned from FilterLogNftDeposit and is used to iterate over the raw logs and unpacked data for LogNftDeposit events raised by the Core contract.
type CoreLogNftDepositIterator struct {
	Event *CoreLogNftDeposit // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CoreLogNftDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreLogNftDeposit)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CoreLogNftDeposit)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CoreLogNftDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreLogNftDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreLogNftDeposit represents a LogNftDeposit event raised by the Core contract.
type CoreLogNftDeposit struct {
	DepositorEthKey common.Address
	StarkKey        *big.Int
	VaultId         *big.Int
	AssetType       *big.Int
	TokenId         *big.Int
	AssetId         *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterLogNftDeposit is a free log retrieval operation binding the contract event 0x0fcf2162832b2d6033d4d34d2f45a28d9cfee523f1899945bbdd32529cfda67b.
//
// Solidity: event LogNftDeposit(address depositorEthKey, uint256 starkKey, uint256 vaultId, uint256 assetType, uint256 tokenId, uint256 assetId)
func (_Core *CoreFilterer) FilterLogNftDeposit(opts *bind.FilterOpts) (*CoreLogNftDepositIterator, error) {

	logs, sub, err := _Core.contract.FilterLogs(opts, "LogNftDeposit")
	if err != nil {
		return nil, err
	}
	return &CoreLogNftDepositIterator{contract: _Core.contract, event: "LogNftDeposit", logs: logs, sub: sub}, nil
}

// WatchLogNftDeposit is a free log subscription operation binding the contract event 0x0fcf2162832b2d6033d4d34d2f45a28d9cfee523f1899945bbdd32529cfda67b.
//
// Solidity: event LogNftDeposit(address depositorEthKey, uint256 starkKey, uint256 vaultId, uint256 assetType, uint256 tokenId, uint256 assetId)
func (_Core *CoreFilterer) WatchLogNftDeposit(opts *bind.WatchOpts, sink chan<- *CoreLogNftDeposit) (event.Subscription, error) {

	logs, sub, err := _Core.contract.WatchLogs(opts, "LogNftDeposit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreLogNftDeposit)
				if err := _Core.contract.UnpackLog(event, "LogNftDeposit", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogNftDeposit is a log parse operation binding the contract event 0x0fcf2162832b2d6033d4d34d2f45a28d9cfee523f1899945bbdd32529cfda67b.
//
// Solidity: event LogNftDeposit(address depositorEthKey, uint256 starkKey, uint256 vaultId, uint256 assetType, uint256 tokenId, uint256 assetId)
func (_Core *CoreFilterer) ParseLogNftDeposit(log types.Log) (*CoreLogNftDeposit, error) {
	event := new(CoreLogNftDeposit)
	if err := _Core.contract.UnpackLog(event, "LogNftDeposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreLogNftWithdrawalAllowedIterator is returned from FilterLogNftWithdrawalAllowed and is used to iterate over the raw logs and unpacked data for LogNftWithdrawalAllowed events raised by the Core contract.
type CoreLogNftWithdrawalAllowedIterator struct {
	Event *CoreLogNftWithdrawalAllowed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CoreLogNftWithdrawalAllowedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreLogNftWithdrawalAllowed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CoreLogNftWithdrawalAllowed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CoreLogNftWithdrawalAllowedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreLogNftWithdrawalAllowedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreLogNftWithdrawalAllowed represents a LogNftWithdrawalAllowed event raised by the Core contract.
type CoreLogNftWithdrawalAllowed struct {
	OwnerKey *big.Int
	AssetId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLogNftWithdrawalAllowed is a free log retrieval operation binding the contract event 0xf07608f26256bce78d87220cfc0e7b1cc69b48e55104bfa591b2818161386627.
//
// Solidity: event LogNftWithdrawalAllowed(uint256 ownerKey, uint256 assetId)
func (_Core *CoreFilterer) FilterLogNftWithdrawalAllowed(opts *bind.FilterOpts) (*CoreLogNftWithdrawalAllowedIterator, error) {

	logs, sub, err := _Core.contract.FilterLogs(opts, "LogNftWithdrawalAllowed")
	if err != nil {
		return nil, err
	}
	return &CoreLogNftWithdrawalAllowedIterator{contract: _Core.contract, event: "LogNftWithdrawalAllowed", logs: logs, sub: sub}, nil
}

// WatchLogNftWithdrawalAllowed is a free log subscription operation binding the contract event 0xf07608f26256bce78d87220cfc0e7b1cc69b48e55104bfa591b2818161386627.
//
// Solidity: event LogNftWithdrawalAllowed(uint256 ownerKey, uint256 assetId)
func (_Core *CoreFilterer) WatchLogNftWithdrawalAllowed(opts *bind.WatchOpts, sink chan<- *CoreLogNftWithdrawalAllowed) (event.Subscription, error) {

	logs, sub, err := _Core.contract.WatchLogs(opts, "LogNftWithdrawalAllowed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreLogNftWithdrawalAllowed)
				if err := _Core.contract.UnpackLog(event, "LogNftWithdrawalAllowed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogNftWithdrawalAllowed is a log parse operation binding the contract event 0xf07608f26256bce78d87220cfc0e7b1cc69b48e55104bfa591b2818161386627.
//
// Solidity: event LogNftWithdrawalAllowed(uint256 ownerKey, uint256 assetId)
func (_Core *CoreFilterer) ParseLogNftWithdrawalAllowed(log types.Log) (*CoreLogNftWithdrawalAllowed, error) {
	event := new(CoreLogNftWithdrawalAllowed)
	if err := _Core.contract.UnpackLog(event, "LogNftWithdrawalAllowed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreLogNftWithdrawalPerformedIterator is returned from FilterLogNftWithdrawalPerformed and is used to iterate over the raw logs and unpacked data for LogNftWithdrawalPerformed events raised by the Core contract.
type CoreLogNftWithdrawalPerformedIterator struct {
	Event *CoreLogNftWithdrawalPerformed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CoreLogNftWithdrawalPerformedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreLogNftWithdrawalPerformed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CoreLogNftWithdrawalPerformed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CoreLogNftWithdrawalPerformedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreLogNftWithdrawalPerformedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreLogNftWithdrawalPerformed represents a LogNftWithdrawalPerformed event raised by the Core contract.
type CoreLogNftWithdrawalPerformed struct {
	OwnerKey  *big.Int
	AssetType *big.Int
	TokenId   *big.Int
	AssetId   *big.Int
	Recipient common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLogNftWithdrawalPerformed is a free log retrieval operation binding the contract event 0xa5cfa8e2199ec5b8ca319288bcab72734207d30569756ee594a74b4df7abbf41.
//
// Solidity: event LogNftWithdrawalPerformed(uint256 ownerKey, uint256 assetType, uint256 tokenId, uint256 assetId, address recipient)
func (_Core *CoreFilterer) FilterLogNftWithdrawalPerformed(opts *bind.FilterOpts) (*CoreLogNftWithdrawalPerformedIterator, error) {

	logs, sub, err := _Core.contract.FilterLogs(opts, "LogNftWithdrawalPerformed")
	if err != nil {
		return nil, err
	}
	return &CoreLogNftWithdrawalPerformedIterator{contract: _Core.contract, event: "LogNftWithdrawalPerformed", logs: logs, sub: sub}, nil
}

// WatchLogNftWithdrawalPerformed is a free log subscription operation binding the contract event 0xa5cfa8e2199ec5b8ca319288bcab72734207d30569756ee594a74b4df7abbf41.
//
// Solidity: event LogNftWithdrawalPerformed(uint256 ownerKey, uint256 assetType, uint256 tokenId, uint256 assetId, address recipient)
func (_Core *CoreFilterer) WatchLogNftWithdrawalPerformed(opts *bind.WatchOpts, sink chan<- *CoreLogNftWithdrawalPerformed) (event.Subscription, error) {

	logs, sub, err := _Core.contract.WatchLogs(opts, "LogNftWithdrawalPerformed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreLogNftWithdrawalPerformed)
				if err := _Core.contract.UnpackLog(event, "LogNftWithdrawalPerformed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogNftWithdrawalPerformed is a log parse operation binding the contract event 0xa5cfa8e2199ec5b8ca319288bcab72734207d30569756ee594a74b4df7abbf41.
//
// Solidity: event LogNftWithdrawalPerformed(uint256 ownerKey, uint256 assetType, uint256 tokenId, uint256 assetId, address recipient)
func (_Core *CoreFilterer) ParseLogNftWithdrawalPerformed(log types.Log) (*CoreLogNftWithdrawalPerformed, error) {
	event := new(CoreLogNftWithdrawalPerformed)
	if err := _Core.contract.UnpackLog(event, "LogNftWithdrawalPerformed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreLogRootUpdateIterator is returned from FilterLogRootUpdate and is used to iterate over the raw logs and unpacked data for LogRootUpdate events raised by the Core contract.
type CoreLogRootUpdateIterator struct {
	Event *CoreLogRootUpdate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CoreLogRootUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreLogRootUpdate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CoreLogRootUpdate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CoreLogRootUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreLogRootUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreLogRootUpdate represents a LogRootUpdate event raised by the Core contract.
type CoreLogRootUpdate struct {
	SequenceNumber *big.Int
	BatchId        *big.Int
	VaultRoot      *big.Int
	OrderRoot      *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterLogRootUpdate is a free log retrieval operation binding the contract event 0xd606ef105963a7e789d927c1d21df5111915b832996b92648138f59eb9763a20.
//
// Solidity: event LogRootUpdate(uint256 sequenceNumber, uint256 batchId, uint256 vaultRoot, uint256 orderRoot)
func (_Core *CoreFilterer) FilterLogRootUpdate(opts *bind.FilterOpts) (*CoreLogRootUpdateIterator, error) {

	logs, sub, err := _Core.contract.FilterLogs(opts, "LogRootUpdate")
	if err != nil {
		return nil, err
	}
	return &CoreLogRootUpdateIterator{contract: _Core.contract, event: "LogRootUpdate", logs: logs, sub: sub}, nil
}

// WatchLogRootUpdate is a free log subscription operation binding the contract event 0xd606ef105963a7e789d927c1d21df5111915b832996b92648138f59eb9763a20.
//
// Solidity: event LogRootUpdate(uint256 sequenceNumber, uint256 batchId, uint256 vaultRoot, uint256 orderRoot)
func (_Core *CoreFilterer) WatchLogRootUpdate(opts *bind.WatchOpts, sink chan<- *CoreLogRootUpdate) (event.Subscription, error) {

	logs, sub, err := _Core.contract.WatchLogs(opts, "LogRootUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreLogRootUpdate)
				if err := _Core.contract.UnpackLog(event, "LogRootUpdate", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogRootUpdate is a log parse operation binding the contract event 0xd606ef105963a7e789d927c1d21df5111915b832996b92648138f59eb9763a20.
//
// Solidity: event LogRootUpdate(uint256 sequenceNumber, uint256 batchId, uint256 vaultRoot, uint256 orderRoot)
func (_Core *CoreFilterer) ParseLogRootUpdate(log types.Log) (*CoreLogRootUpdate, error) {
	event := new(CoreLogRootUpdate)
	if err := _Core.contract.UnpackLog(event, "LogRootUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreLogStateTransitionFactIterator is returned from FilterLogStateTransitionFact and is used to iterate over the raw logs and unpacked data for LogStateTransitionFact events raised by the Core contract.
type CoreLogStateTransitionFactIterator struct {
	Event *CoreLogStateTransitionFact // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CoreLogStateTransitionFactIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreLogStateTransitionFact)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CoreLogStateTransitionFact)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CoreLogStateTransitionFactIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreLogStateTransitionFactIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreLogStateTransitionFact represents a LogStateTransitionFact event raised by the Core contract.
type CoreLogStateTransitionFact struct {
	StateTransitionFact [32]byte
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterLogStateTransitionFact is a free log retrieval operation binding the contract event 0x9866f8ddfe70bb512b2f2b28b49d4017c43f7ba775f1a20c61c13eea8cdac111.
//
// Solidity: event LogStateTransitionFact(bytes32 stateTransitionFact)
func (_Core *CoreFilterer) FilterLogStateTransitionFact(opts *bind.FilterOpts) (*CoreLogStateTransitionFactIterator, error) {

	logs, sub, err := _Core.contract.FilterLogs(opts, "LogStateTransitionFact")
	if err != nil {
		return nil, err
	}
	return &CoreLogStateTransitionFactIterator{contract: _Core.contract, event: "LogStateTransitionFact", logs: logs, sub: sub}, nil
}

// WatchLogStateTransitionFact is a free log subscription operation binding the contract event 0x9866f8ddfe70bb512b2f2b28b49d4017c43f7ba775f1a20c61c13eea8cdac111.
//
// Solidity: event LogStateTransitionFact(bytes32 stateTransitionFact)
func (_Core *CoreFilterer) WatchLogStateTransitionFact(opts *bind.WatchOpts, sink chan<- *CoreLogStateTransitionFact) (event.Subscription, error) {

	logs, sub, err := _Core.contract.WatchLogs(opts, "LogStateTransitionFact")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreLogStateTransitionFact)
				if err := _Core.contract.UnpackLog(event, "LogStateTransitionFact", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogStateTransitionFact is a log parse operation binding the contract event 0x9866f8ddfe70bb512b2f2b28b49d4017c43f7ba775f1a20c61c13eea8cdac111.
//
// Solidity: event LogStateTransitionFact(bytes32 stateTransitionFact)
func (_Core *CoreFilterer) ParseLogStateTransitionFact(log types.Log) (*CoreLogStateTransitionFact, error) {
	event := new(CoreLogStateTransitionFact)
	if err := _Core.contract.UnpackLog(event, "LogStateTransitionFact", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreLogVaultBalanceChangeAppliedIterator is returned from FilterLogVaultBalanceChangeApplied and is used to iterate over the raw logs and unpacked data for LogVaultBalanceChangeApplied events raised by the Core contract.
type CoreLogVaultBalanceChangeAppliedIterator struct {
	Event *CoreLogVaultBalanceChangeApplied // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CoreLogVaultBalanceChangeAppliedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreLogVaultBalanceChangeApplied)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CoreLogVaultBalanceChangeApplied)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CoreLogVaultBalanceChangeAppliedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreLogVaultBalanceChangeAppliedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreLogVaultBalanceChangeApplied represents a LogVaultBalanceChangeApplied event raised by the Core contract.
type CoreLogVaultBalanceChangeApplied struct {
	EthKey                common.Address
	AssetId               *big.Int
	VaultId               *big.Int
	QuantizedAmountChange *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterLogVaultBalanceChangeApplied is a free log retrieval operation binding the contract event 0x2b2b583bb5166d03b87a8e7c2785e61530ab776400fb69e1bcb13b33afef2b58.
//
// Solidity: event LogVaultBalanceChangeApplied(address ethKey, uint256 assetId, uint256 vaultId, int256 quantizedAmountChange)
func (_Core *CoreFilterer) FilterLogVaultBalanceChangeApplied(opts *bind.FilterOpts) (*CoreLogVaultBalanceChangeAppliedIterator, error) {

	logs, sub, err := _Core.contract.FilterLogs(opts, "LogVaultBalanceChangeApplied")
	if err != nil {
		return nil, err
	}
	return &CoreLogVaultBalanceChangeAppliedIterator{contract: _Core.contract, event: "LogVaultBalanceChangeApplied", logs: logs, sub: sub}, nil
}

// WatchLogVaultBalanceChangeApplied is a free log subscription operation binding the contract event 0x2b2b583bb5166d03b87a8e7c2785e61530ab776400fb69e1bcb13b33afef2b58.
//
// Solidity: event LogVaultBalanceChangeApplied(address ethKey, uint256 assetId, uint256 vaultId, int256 quantizedAmountChange)
func (_Core *CoreFilterer) WatchLogVaultBalanceChangeApplied(opts *bind.WatchOpts, sink chan<- *CoreLogVaultBalanceChangeApplied) (event.Subscription, error) {

	logs, sub, err := _Core.contract.WatchLogs(opts, "LogVaultBalanceChangeApplied")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreLogVaultBalanceChangeApplied)
				if err := _Core.contract.UnpackLog(event, "LogVaultBalanceChangeApplied", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogVaultBalanceChangeApplied is a log parse operation binding the contract event 0x2b2b583bb5166d03b87a8e7c2785e61530ab776400fb69e1bcb13b33afef2b58.
//
// Solidity: event LogVaultBalanceChangeApplied(address ethKey, uint256 assetId, uint256 vaultId, int256 quantizedAmountChange)
func (_Core *CoreFilterer) ParseLogVaultBalanceChangeApplied(log types.Log) (*CoreLogVaultBalanceChangeApplied, error) {
	event := new(CoreLogVaultBalanceChangeApplied)
	if err := _Core.contract.UnpackLog(event, "LogVaultBalanceChangeApplied", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreLogWithdrawalAllowedIterator is returned from FilterLogWithdrawalAllowed and is used to iterate over the raw logs and unpacked data for LogWithdrawalAllowed events raised by the Core contract.
type CoreLogWithdrawalAllowedIterator struct {
	Event *CoreLogWithdrawalAllowed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CoreLogWithdrawalAllowedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreLogWithdrawalAllowed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CoreLogWithdrawalAllowed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CoreLogWithdrawalAllowedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreLogWithdrawalAllowedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreLogWithdrawalAllowed represents a LogWithdrawalAllowed event raised by the Core contract.
type CoreLogWithdrawalAllowed struct {
	OwnerKey           *big.Int
	AssetType          *big.Int
	NonQuantizedAmount *big.Int
	QuantizedAmount    *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterLogWithdrawalAllowed is a free log retrieval operation binding the contract event 0x03c10a82c955f7bcd0c934147fb39cafca947a4294425b1751d884c8ac954287.
//
// Solidity: event LogWithdrawalAllowed(uint256 ownerKey, uint256 assetType, uint256 nonQuantizedAmount, uint256 quantizedAmount)
func (_Core *CoreFilterer) FilterLogWithdrawalAllowed(opts *bind.FilterOpts) (*CoreLogWithdrawalAllowedIterator, error) {

	logs, sub, err := _Core.contract.FilterLogs(opts, "LogWithdrawalAllowed")
	if err != nil {
		return nil, err
	}
	return &CoreLogWithdrawalAllowedIterator{contract: _Core.contract, event: "LogWithdrawalAllowed", logs: logs, sub: sub}, nil
}

// WatchLogWithdrawalAllowed is a free log subscription operation binding the contract event 0x03c10a82c955f7bcd0c934147fb39cafca947a4294425b1751d884c8ac954287.
//
// Solidity: event LogWithdrawalAllowed(uint256 ownerKey, uint256 assetType, uint256 nonQuantizedAmount, uint256 quantizedAmount)
func (_Core *CoreFilterer) WatchLogWithdrawalAllowed(opts *bind.WatchOpts, sink chan<- *CoreLogWithdrawalAllowed) (event.Subscription, error) {

	logs, sub, err := _Core.contract.WatchLogs(opts, "LogWithdrawalAllowed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreLogWithdrawalAllowed)
				if err := _Core.contract.UnpackLog(event, "LogWithdrawalAllowed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogWithdrawalAllowed is a log parse operation binding the contract event 0x03c10a82c955f7bcd0c934147fb39cafca947a4294425b1751d884c8ac954287.
//
// Solidity: event LogWithdrawalAllowed(uint256 ownerKey, uint256 assetType, uint256 nonQuantizedAmount, uint256 quantizedAmount)
func (_Core *CoreFilterer) ParseLogWithdrawalAllowed(log types.Log) (*CoreLogWithdrawalAllowed, error) {
	event := new(CoreLogWithdrawalAllowed)
	if err := _Core.contract.UnpackLog(event, "LogWithdrawalAllowed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreLogWithdrawalPerformedIterator is returned from FilterLogWithdrawalPerformed and is used to iterate over the raw logs and unpacked data for LogWithdrawalPerformed events raised by the Core contract.
type CoreLogWithdrawalPerformedIterator struct {
	Event *CoreLogWithdrawalPerformed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CoreLogWithdrawalPerformedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreLogWithdrawalPerformed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CoreLogWithdrawalPerformed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CoreLogWithdrawalPerformedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreLogWithdrawalPerformedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreLogWithdrawalPerformed represents a LogWithdrawalPerformed event raised by the Core contract.
type CoreLogWithdrawalPerformed struct {
	OwnerKey           *big.Int
	AssetType          *big.Int
	NonQuantizedAmount *big.Int
	QuantizedAmount    *big.Int
	Recipient          common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterLogWithdrawalPerformed is a free log retrieval operation binding the contract event 0xb7477a7b93b2addc5272bbd7ad0986ef1c0d0bd265f26c3dc4bbe42727c2ac0c.
//
// Solidity: event LogWithdrawalPerformed(uint256 ownerKey, uint256 assetType, uint256 nonQuantizedAmount, uint256 quantizedAmount, address recipient)
func (_Core *CoreFilterer) FilterLogWithdrawalPerformed(opts *bind.FilterOpts) (*CoreLogWithdrawalPerformedIterator, error) {

	logs, sub, err := _Core.contract.FilterLogs(opts, "LogWithdrawalPerformed")
	if err != nil {
		return nil, err
	}
	return &CoreLogWithdrawalPerformedIterator{contract: _Core.contract, event: "LogWithdrawalPerformed", logs: logs, sub: sub}, nil
}

// WatchLogWithdrawalPerformed is a free log subscription operation binding the contract event 0xb7477a7b93b2addc5272bbd7ad0986ef1c0d0bd265f26c3dc4bbe42727c2ac0c.
//
// Solidity: event LogWithdrawalPerformed(uint256 ownerKey, uint256 assetType, uint256 nonQuantizedAmount, uint256 quantizedAmount, address recipient)
func (_Core *CoreFilterer) WatchLogWithdrawalPerformed(opts *bind.WatchOpts, sink chan<- *CoreLogWithdrawalPerformed) (event.Subscription, error) {

	logs, sub, err := _Core.contract.WatchLogs(opts, "LogWithdrawalPerformed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreLogWithdrawalPerformed)
				if err := _Core.contract.UnpackLog(event, "LogWithdrawalPerformed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogWithdrawalPerformed is a log parse operation binding the contract event 0xb7477a7b93b2addc5272bbd7ad0986ef1c0d0bd265f26c3dc4bbe42727c2ac0c.
//
// Solidity: event LogWithdrawalPerformed(uint256 ownerKey, uint256 assetType, uint256 nonQuantizedAmount, uint256 quantizedAmount, address recipient)
func (_Core *CoreFilterer) ParseLogWithdrawalPerformed(log types.Log) (*CoreLogWithdrawalPerformed, error) {
	event := new(CoreLogWithdrawalPerformed)
	if err := _Core.contract.UnpackLog(event, "LogWithdrawalPerformed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
