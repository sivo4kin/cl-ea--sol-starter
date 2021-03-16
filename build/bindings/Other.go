// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// OtherABI is the input ABI used to generate the binding from.
const OtherABI = "[]"

// OtherBin is the compiled bytecode used for deploying new contracts.
var OtherBin = "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220d881428fab066ffedf049cfe3a9d1b053a7d0a64be89b576bbc450304272b47d64736f6c63430006090033"

// DeployOther deploys a new Ethereum contract, binding an instance of Other to it.
func DeployOther(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Other, error) {
	parsed, err := abi.JSON(strings.NewReader(OtherABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OtherBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Other{OtherCaller: OtherCaller{contract: contract}, OtherTransactor: OtherTransactor{contract: contract}, OtherFilterer: OtherFilterer{contract: contract}}, nil
}

// Other is an auto generated Go binding around an Ethereum contract.
type Other struct {
	OtherCaller     // Read-only binding to the contract
	OtherTransactor // Write-only binding to the contract
	OtherFilterer   // Log filterer for contract events
}

// OtherCaller is an auto generated read-only Go binding around an Ethereum contract.
type OtherCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OtherTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OtherTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OtherFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OtherFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OtherSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OtherSession struct {
	Contract     *Other            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OtherCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OtherCallerSession struct {
	Contract *OtherCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// OtherTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OtherTransactorSession struct {
	Contract     *OtherTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OtherRaw is an auto generated low-level Go binding around an Ethereum contract.
type OtherRaw struct {
	Contract *Other // Generic contract binding to access the raw methods on
}

// OtherCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OtherCallerRaw struct {
	Contract *OtherCaller // Generic read-only contract binding to access the raw methods on
}

// OtherTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OtherTransactorRaw struct {
	Contract *OtherTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOther creates a new instance of Other, bound to a specific deployed contract.
func NewOther(address common.Address, backend bind.ContractBackend) (*Other, error) {
	contract, err := bindOther(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Other{OtherCaller: OtherCaller{contract: contract}, OtherTransactor: OtherTransactor{contract: contract}, OtherFilterer: OtherFilterer{contract: contract}}, nil
}

// NewOtherCaller creates a new read-only instance of Other, bound to a specific deployed contract.
func NewOtherCaller(address common.Address, caller bind.ContractCaller) (*OtherCaller, error) {
	contract, err := bindOther(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OtherCaller{contract: contract}, nil
}

// NewOtherTransactor creates a new write-only instance of Other, bound to a specific deployed contract.
func NewOtherTransactor(address common.Address, transactor bind.ContractTransactor) (*OtherTransactor, error) {
	contract, err := bindOther(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OtherTransactor{contract: contract}, nil
}

// NewOtherFilterer creates a new log filterer instance of Other, bound to a specific deployed contract.
func NewOtherFilterer(address common.Address, filterer bind.ContractFilterer) (*OtherFilterer, error) {
	contract, err := bindOther(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OtherFilterer{contract: contract}, nil
}

// bindOther binds a generic wrapper to an already deployed contract.
func bindOther(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OtherABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Other *OtherRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Other.Contract.OtherCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Other *OtherRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Other.Contract.OtherTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Other *OtherRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Other.Contract.OtherTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Other *OtherCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Other.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Other *OtherTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Other.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Other *OtherTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Other.Contract.contract.Transact(opts, method, params...)
}
