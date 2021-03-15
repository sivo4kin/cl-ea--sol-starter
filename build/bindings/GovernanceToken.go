// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
	"strings"
	//"github.com/sivo4kin/digiu-cross-chain/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
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

// GovernanceTokenABI is the input ABI used to generate the binding from.
const GovernanceTokenABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// GovernanceTokenBin is the compiled bytecode used for deploying new contracts.
var GovernanceTokenBin = "0x608060405234801561001057600080fd5b50610228806100206000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c806306fdde031461006757806318160ddd146100e4578063313ce567146100fe57806370a082311461011c57806395d89b4114610142578063dd62ed3e1461014a575b600080fd5b61006f610178565b6040805160208082528351818301528351919283929083019185019080838360005b838110156100a9578181015183820152602001610091565b50505050905090810190601f1680156100d65780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6100ec610199565b60408051918252519081900360200190f35b61010661019f565b6040805160ff9092168252519081900360200190f35b6100ec6004803603602081101561013257600080fd5b50356001600160a01b03166101a4565b61006f6101b6565b6100ec6004803603604081101561016057600080fd5b506001600160a01b03813581169160200135166101d5565b60405180604001604052806005815260200164446967695560d81b81525081565b60005481565b601281565b60016020526000908152604090205481565b6040518060400160405280600381526020016244475560e81b81525081565b60026020908152600092835260408084209091529082529020548156fea2646970667358221220b6212a1321246139e04a13d1b62cf40b2d3c2050efc1b4ca09cabcf4ae7fe8a264736f6c63430006090033"

// DeployGovernanceToken deploys a new Ethereum contract, binding an instance of GovernanceToken to it.
func DeployGovernanceToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *GovernanceToken, error) {
	parsed, err := abi.JSON(strings.NewReader(GovernanceTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(GovernanceTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &GovernanceToken{GovernanceTokenCaller: GovernanceTokenCaller{contract: contract}, GovernanceTokenTransactor: GovernanceTokenTransactor{contract: contract}, GovernanceTokenFilterer: GovernanceTokenFilterer{contract: contract}}, nil
}

// DeployGovernanceTokenSync deploys a new Ethereum contract and waits for receipt, binding an instance of GovernanceTokenSession to it.
func DeployGovernanceTokenSync(session *bind.TransactSession, backend bind.ContractBackend) (*types.Transaction, *types.Receipt, *GovernanceTokenSession, error) {
	parsed, err := abi.JSON(strings.NewReader(GovernanceTokenABI))
	if err != nil {
		return nil, nil, nil, err
	}
	session.Lock()
	address, tx, _, err := bind.DeployContract(session.TransactOpts, parsed, common.FromHex(GovernanceTokenBin), backend)
	receipt, err := session.WaitTransaction(tx)
	if err != nil {
		session.Unlock()
		return nil, nil, nil, err
	}
	session.TransactOpts.Nonce.Add(session.TransactOpts.Nonce, big.NewInt(1))
	session.Unlock()
	contractSession, err := NewGovernanceTokenSession(address, backend, session)
	return tx, receipt, contractSession, err
}

// GovernanceToken is an auto generated Go binding around an Ethereum contract.
type GovernanceToken struct {
	GovernanceTokenCaller     // Read-only binding to the contract
	GovernanceTokenTransactor // Write-only binding to the contract
	GovernanceTokenFilterer   // Log filterer for contract events
}

// GovernanceTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type GovernanceTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GovernanceTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GovernanceTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GovernanceTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GovernanceTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GovernanceTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GovernanceTokenSession struct {
	Contract           *GovernanceToken // Generic contract binding to set the session for
	transactionSession *bind.TransactSession
	Address            common.Address
}

// GovernanceTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GovernanceTokenCallerSession struct {
	Contract *GovernanceTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// GovernanceTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GovernanceTokenTransactorSession struct {
	Contract     *GovernanceTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// GovernanceTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type GovernanceTokenRaw struct {
	Contract *GovernanceToken // Generic contract binding to access the raw methods on
}

// GovernanceTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GovernanceTokenCallerRaw struct {
	Contract *GovernanceTokenCaller // Generic read-only contract binding to access the raw methods on
}

// GovernanceTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GovernanceTokenTransactorRaw struct {
	Contract *GovernanceTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGovernanceToken creates a new instance of GovernanceToken, bound to a specific deployed contract.
func NewGovernanceToken(address common.Address, backend bind.ContractBackend) (*GovernanceToken, error) {
	contract, err := bindGovernanceToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GovernanceToken{GovernanceTokenCaller: GovernanceTokenCaller{contract: contract}, GovernanceTokenTransactor: GovernanceTokenTransactor{contract: contract}, GovernanceTokenFilterer: GovernanceTokenFilterer{contract: contract}}, nil
}

// NewGovernanceTokenCaller creates a new read-only instance of GovernanceToken, bound to a specific deployed contract.
func NewGovernanceTokenCaller(address common.Address, caller bind.ContractCaller) (*GovernanceTokenCaller, error) {
	contract, err := bindGovernanceToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GovernanceTokenCaller{contract: contract}, nil
}

// NewGovernanceTokenTransactor creates a new write-only instance of GovernanceToken, bound to a specific deployed contract.
func NewGovernanceTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*GovernanceTokenTransactor, error) {
	contract, err := bindGovernanceToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GovernanceTokenTransactor{contract: contract}, nil
}

// NewGovernanceTokenFilterer creates a new log filterer instance of GovernanceToken, bound to a specific deployed contract.
func NewGovernanceTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*GovernanceTokenFilterer, error) {
	contract, err := bindGovernanceToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GovernanceTokenFilterer{contract: contract}, nil
}

func NewGovernanceTokenSession(address common.Address, backend bind.ContractBackend, transactionSession *bind.TransactSession) (*GovernanceTokenSession, error) {
	GovernanceTokenInstance, err := NewGovernanceToken(address, backend)
	if err != nil {
		return nil, err
	}
	return &GovernanceTokenSession{
		Contract:           GovernanceTokenInstance,
		transactionSession: transactionSession,
		Address:            address,
	}, nil
}

// bindGovernanceToken binds a generic wrapper to an already deployed contract.
func bindGovernanceToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GovernanceTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GovernanceToken *GovernanceTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _GovernanceToken.Contract.GovernanceTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GovernanceToken *GovernanceTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GovernanceToken.Contract.GovernanceTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GovernanceToken *GovernanceTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GovernanceToken.Contract.GovernanceTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GovernanceToken *GovernanceTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _GovernanceToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GovernanceToken *GovernanceTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GovernanceToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GovernanceToken *GovernanceTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GovernanceToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_GovernanceToken *GovernanceTokenCaller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _GovernanceToken.contract.Call(opts, out, "allowance", arg0, arg1)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_GovernanceToken *GovernanceTokenSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _GovernanceToken.Contract.Allowance(_GovernanceToken.transactionSession.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_GovernanceToken *GovernanceTokenCallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _GovernanceToken.Contract.Allowance(&_GovernanceToken.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_GovernanceToken *GovernanceTokenCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _GovernanceToken.contract.Call(opts, out, "balanceOf", arg0)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_GovernanceToken *GovernanceTokenSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _GovernanceToken.Contract.BalanceOf(_GovernanceToken.transactionSession.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_GovernanceToken *GovernanceTokenCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _GovernanceToken.Contract.BalanceOf(&_GovernanceToken.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_GovernanceToken *GovernanceTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _GovernanceToken.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_GovernanceToken *GovernanceTokenSession) Decimals() (uint8, error) {
	return _GovernanceToken.Contract.Decimals(_GovernanceToken.transactionSession.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_GovernanceToken *GovernanceTokenCallerSession) Decimals() (uint8, error) {
	return _GovernanceToken.Contract.Decimals(&_GovernanceToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_GovernanceToken *GovernanceTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _GovernanceToken.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_GovernanceToken *GovernanceTokenSession) Name() (string, error) {
	return _GovernanceToken.Contract.Name(_GovernanceToken.transactionSession.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_GovernanceToken *GovernanceTokenCallerSession) Name() (string, error) {
	return _GovernanceToken.Contract.Name(&_GovernanceToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_GovernanceToken *GovernanceTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _GovernanceToken.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_GovernanceToken *GovernanceTokenSession) Symbol() (string, error) {
	return _GovernanceToken.Contract.Symbol(_GovernanceToken.transactionSession.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_GovernanceToken *GovernanceTokenCallerSession) Symbol() (string, error) {
	return _GovernanceToken.Contract.Symbol(&_GovernanceToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_GovernanceToken *GovernanceTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _GovernanceToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_GovernanceToken *GovernanceTokenSession) TotalSupply() (*big.Int, error) {
	return _GovernanceToken.Contract.TotalSupply(_GovernanceToken.transactionSession.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_GovernanceToken *GovernanceTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _GovernanceToken.Contract.TotalSupply(&_GovernanceToken.CallOpts)
}

// GovernanceTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the GovernanceToken contract.
type GovernanceTokenApprovalIterator struct {
	Event *GovernanceTokenApproval // Event containing the contract specifics and raw log

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
func (it *GovernanceTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceTokenApproval)
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
		it.Event = new(GovernanceTokenApproval)
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
func (it *GovernanceTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceTokenApproval represents a Approval event raised by the GovernanceToken contract.
type GovernanceTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_GovernanceToken *GovernanceTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*GovernanceTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _GovernanceToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &GovernanceTokenApprovalIterator{contract: _GovernanceToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_GovernanceToken *GovernanceTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *GovernanceTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _GovernanceToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceTokenApproval)
				if err := _GovernanceToken.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_GovernanceToken *GovernanceTokenFilterer) ParseApproval(log types.Log) (*GovernanceTokenApproval, error) {
	event := new(GovernanceTokenApproval)
	if err := _GovernanceToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GovernanceTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the GovernanceToken contract.
type GovernanceTokenTransferIterator struct {
	Event *GovernanceTokenTransfer // Event containing the contract specifics and raw log

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
func (it *GovernanceTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceTokenTransfer)
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
		it.Event = new(GovernanceTokenTransfer)
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
func (it *GovernanceTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceTokenTransfer represents a Transfer event raised by the GovernanceToken contract.
type GovernanceTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_GovernanceToken *GovernanceTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*GovernanceTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _GovernanceToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &GovernanceTokenTransferIterator{contract: _GovernanceToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_GovernanceToken *GovernanceTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *GovernanceTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _GovernanceToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceTokenTransfer)
				if err := _GovernanceToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_GovernanceToken *GovernanceTokenFilterer) ParseTransfer(log types.Log) (*GovernanceTokenTransfer, error) {
	event := new(GovernanceTokenTransfer)
	if err := _GovernanceToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}
