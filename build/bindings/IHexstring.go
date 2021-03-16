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

// IHexstringABI is the input ABI used to generate the binding from.
const IHexstringABI = "[{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"bytesToHexString\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// IHexstring is an auto generated Go binding around an Ethereum contract.
type IHexstring struct {
	IHexstringCaller     // Read-only binding to the contract
	IHexstringTransactor // Write-only binding to the contract
	IHexstringFilterer   // Log filterer for contract events
}

// IHexstringCaller is an auto generated read-only Go binding around an Ethereum contract.
type IHexstringCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IHexstringTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IHexstringTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IHexstringFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IHexstringFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IHexstringSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IHexstringSession struct {
	Contract     *IHexstring       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IHexstringCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IHexstringCallerSession struct {
	Contract *IHexstringCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// IHexstringTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IHexstringTransactorSession struct {
	Contract     *IHexstringTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// IHexstringRaw is an auto generated low-level Go binding around an Ethereum contract.
type IHexstringRaw struct {
	Contract *IHexstring // Generic contract binding to access the raw methods on
}

// IHexstringCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IHexstringCallerRaw struct {
	Contract *IHexstringCaller // Generic read-only contract binding to access the raw methods on
}

// IHexstringTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IHexstringTransactorRaw struct {
	Contract *IHexstringTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIHexstring creates a new instance of IHexstring, bound to a specific deployed contract.
func NewIHexstring(address common.Address, backend bind.ContractBackend) (*IHexstring, error) {
	contract, err := bindIHexstring(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IHexstring{IHexstringCaller: IHexstringCaller{contract: contract}, IHexstringTransactor: IHexstringTransactor{contract: contract}, IHexstringFilterer: IHexstringFilterer{contract: contract}}, nil
}

// NewIHexstringCaller creates a new read-only instance of IHexstring, bound to a specific deployed contract.
func NewIHexstringCaller(address common.Address, caller bind.ContractCaller) (*IHexstringCaller, error) {
	contract, err := bindIHexstring(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IHexstringCaller{contract: contract}, nil
}

// NewIHexstringTransactor creates a new write-only instance of IHexstring, bound to a specific deployed contract.
func NewIHexstringTransactor(address common.Address, transactor bind.ContractTransactor) (*IHexstringTransactor, error) {
	contract, err := bindIHexstring(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IHexstringTransactor{contract: contract}, nil
}

// NewIHexstringFilterer creates a new log filterer instance of IHexstring, bound to a specific deployed contract.
func NewIHexstringFilterer(address common.Address, filterer bind.ContractFilterer) (*IHexstringFilterer, error) {
	contract, err := bindIHexstring(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IHexstringFilterer{contract: contract}, nil
}

// bindIHexstring binds a generic wrapper to an already deployed contract.
func bindIHexstring(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IHexstringABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IHexstring *IHexstringRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IHexstring.Contract.IHexstringCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IHexstring *IHexstringRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IHexstring.Contract.IHexstringTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IHexstring *IHexstringRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IHexstring.Contract.IHexstringTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IHexstring *IHexstringCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IHexstring.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IHexstring *IHexstringTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IHexstring.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IHexstring *IHexstringTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IHexstring.Contract.contract.Transact(opts, method, params...)
}

// BytesToHexString is a free data retrieval call binding the contract method 0x7d520c22.
//
// Solidity: function bytesToHexString(bytes value) view returns(string)
func (_IHexstring *IHexstringCaller) BytesToHexString(opts *bind.CallOpts, value []byte) (string, error) {
	var out []interface{}
	err := _IHexstring.contract.Call(opts, &out, "bytesToHexString", value)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// BytesToHexString is a free data retrieval call binding the contract method 0x7d520c22.
//
// Solidity: function bytesToHexString(bytes value) view returns(string)
func (_IHexstring *IHexstringSession) BytesToHexString(value []byte) (string, error) {
	return _IHexstring.Contract.BytesToHexString(&_IHexstring.CallOpts, value)
}

// BytesToHexString is a free data retrieval call binding the contract method 0x7d520c22.
//
// Solidity: function bytesToHexString(bytes value) view returns(string)
func (_IHexstring *IHexstringCallerSession) BytesToHexString(value []byte) (string, error) {
	return _IHexstring.Contract.BytesToHexString(&_IHexstring.CallOpts, value)
}
