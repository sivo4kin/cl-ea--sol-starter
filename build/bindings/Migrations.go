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

// MigrationsABI is the input ABI used to generate the binding from.
const MigrationsABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"constant\":true,\"inputs\":[],\"name\":\"last_completed_migration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"completed\",\"type\":\"uint256\"}],\"name\":\"setCompleted\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"new_address\",\"type\":\"address\"}],\"name\":\"upgrade\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// MigrationsBin is the compiled bytecode used for deploying new contracts.
var MigrationsBin = "0x608060405234801561001057600080fd5b50600080546001600160a01b031916331790556101b4806100326000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c80630900f01014610051578063445df0ac146100795780638da5cb5b14610093578063fdacd576146100b7575b600080fd5b6100776004803603602081101561006757600080fd5b50356001600160a01b03166100d4565b005b610081610151565b60408051918252519081900360200190f35b61009b610157565b604080516001600160a01b039092168252519081900360200190f35b610077600480360360208110156100cd57600080fd5b5035610166565b6000546001600160a01b031633141561014e576000819050806001600160a01b031663fdacd5766001546040518263ffffffff1660e01b815260040180828152602001915050600060405180830381600087803b15801561013457600080fd5b505af1158015610148573d6000803e3d6000fd5b50505050505b50565b60015481565b6000546001600160a01b031681565b6000546001600160a01b031633141561014e5760015556fea26469706673582212204630e432884a3938267c485125eae4b58589702b80746059bb7cf0d64fae945364736f6c63430006090033"

// DeployMigrations deploys a new Ethereum contract, binding an instance of Migrations to it.
func DeployMigrations(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Migrations, error) {
	parsed, err := abi.JSON(strings.NewReader(MigrationsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MigrationsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Migrations{MigrationsCaller: MigrationsCaller{contract: contract}, MigrationsTransactor: MigrationsTransactor{contract: contract}, MigrationsFilterer: MigrationsFilterer{contract: contract}}, nil
}

// DeployMigrationsSync deploys a new Ethereum contract and waits for receipt, binding an instance of MigrationsSession to it.
func DeployMigrationsSync(session *bind.TransactSession, backend bind.ContractBackend) (*types.Transaction, *types.Receipt, *MigrationsSession, error) {
	parsed, err := abi.JSON(strings.NewReader(MigrationsABI))
	if err != nil {
		return nil, nil, nil, err
	}
	session.Lock()
	address, tx, _, err := bind.DeployContract(session.TransactOpts, parsed, common.FromHex(MigrationsBin), backend)
	receipt, err := session.WaitTransaction(tx)
	if err != nil {
		session.Unlock()
		return nil, nil, nil, err
	}
	session.TransactOpts.Nonce.Add(session.TransactOpts.Nonce, big.NewInt(1))
	session.Unlock()
	contractSession, err := NewMigrationsSession(address, backend, session)
	return tx, receipt, contractSession, err
}

// Migrations is an auto generated Go binding around an Ethereum contract.
type Migrations struct {
	MigrationsCaller     // Read-only binding to the contract
	MigrationsTransactor // Write-only binding to the contract
	MigrationsFilterer   // Log filterer for contract events
}

// MigrationsCaller is an auto generated read-only Go binding around an Ethereum contract.
type MigrationsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MigrationsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MigrationsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MigrationsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MigrationsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MigrationsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MigrationsSession struct {
	Contract           *Migrations // Generic contract binding to set the session for
	transactionSession *bind.TransactSession
	Address            common.Address
}

// MigrationsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MigrationsCallerSession struct {
	Contract *MigrationsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// MigrationsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MigrationsTransactorSession struct {
	Contract     *MigrationsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// MigrationsRaw is an auto generated low-level Go binding around an Ethereum contract.
type MigrationsRaw struct {
	Contract *Migrations // Generic contract binding to access the raw methods on
}

// MigrationsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MigrationsCallerRaw struct {
	Contract *MigrationsCaller // Generic read-only contract binding to access the raw methods on
}

// MigrationsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MigrationsTransactorRaw struct {
	Contract *MigrationsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMigrations creates a new instance of Migrations, bound to a specific deployed contract.
func NewMigrations(address common.Address, backend bind.ContractBackend) (*Migrations, error) {
	contract, err := bindMigrations(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Migrations{MigrationsCaller: MigrationsCaller{contract: contract}, MigrationsTransactor: MigrationsTransactor{contract: contract}, MigrationsFilterer: MigrationsFilterer{contract: contract}}, nil
}

// NewMigrationsCaller creates a new read-only instance of Migrations, bound to a specific deployed contract.
func NewMigrationsCaller(address common.Address, caller bind.ContractCaller) (*MigrationsCaller, error) {
	contract, err := bindMigrations(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MigrationsCaller{contract: contract}, nil
}

// NewMigrationsTransactor creates a new write-only instance of Migrations, bound to a specific deployed contract.
func NewMigrationsTransactor(address common.Address, transactor bind.ContractTransactor) (*MigrationsTransactor, error) {
	contract, err := bindMigrations(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MigrationsTransactor{contract: contract}, nil
}

// NewMigrationsFilterer creates a new log filterer instance of Migrations, bound to a specific deployed contract.
func NewMigrationsFilterer(address common.Address, filterer bind.ContractFilterer) (*MigrationsFilterer, error) {
	contract, err := bindMigrations(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MigrationsFilterer{contract: contract}, nil
}

func NewMigrationsSession(address common.Address, backend bind.ContractBackend, transactionSession *bind.TransactSession) (*MigrationsSession, error) {
	MigrationsInstance, err := NewMigrations(address, backend)
	if err != nil {
		return nil, err
	}
	return &MigrationsSession{
		Contract:           MigrationsInstance,
		transactionSession: transactionSession,
		Address:            address,
	}, nil
}

// bindMigrations binds a generic wrapper to an already deployed contract.
func bindMigrations(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MigrationsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Migrations *MigrationsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Migrations.Contract.MigrationsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Migrations *MigrationsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Migrations.Contract.MigrationsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Migrations *MigrationsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Migrations.Contract.MigrationsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Migrations *MigrationsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Migrations.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Migrations *MigrationsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Migrations.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Migrations *MigrationsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Migrations.Contract.contract.Transact(opts, method, params...)
}

// LastCompletedMigration is a free data retrieval call binding the contract method 0x445df0ac.
//
// Solidity: function last_completed_migration() view returns(uint256)
func (_Migrations *MigrationsCaller) LastCompletedMigration(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Migrations.contract.Call(opts, out, "last_completed_migration")
	return *ret0, err
}

// LastCompletedMigration is a free data retrieval call binding the contract method 0x445df0ac.
//
// Solidity: function last_completed_migration() view returns(uint256)
func (_Migrations *MigrationsSession) LastCompletedMigration() (*big.Int, error) {
	return _Migrations.Contract.LastCompletedMigration(_Migrations.transactionSession.CallOpts)
}

// LastCompletedMigration is a free data retrieval call binding the contract method 0x445df0ac.
//
// Solidity: function last_completed_migration() view returns(uint256)
func (_Migrations *MigrationsCallerSession) LastCompletedMigration() (*big.Int, error) {
	return _Migrations.Contract.LastCompletedMigration(&_Migrations.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Migrations *MigrationsCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Migrations.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Migrations *MigrationsSession) Owner() (common.Address, error) {
	return _Migrations.Contract.Owner(_Migrations.transactionSession.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Migrations *MigrationsCallerSession) Owner() (common.Address, error) {
	return _Migrations.Contract.Owner(&_Migrations.CallOpts)
}

// SetCompleted is a paid mutator transaction binding the contract method 0xfdacd576.
//
// Solidity: function setCompleted(uint256 completed) returns()
func (_Migrations *MigrationsTransactor) SetCompleted(opts *bind.TransactOpts, completed *big.Int) (*types.Transaction, error) {
	return _Migrations.contract.Transact(opts, "setCompleted", completed)
}

func (_Migrations *MigrationsTransactor) SetCompletedRawTx(opts *bind.TransactOpts, completed *big.Int) (*types.Transaction, error) {
	return _Migrations.contract.RawTx(opts, "setCompleted", completed)
}

// SetCompleted is a paid mutator transaction binding the contract method 0xfdacd576.
// Will wait for tx receipt
//
// Solidity: function setCompleted(uint256 completed) returns()
func (_Migrations *MigrationsSession) SetCompleted(completed *big.Int) (*types.Transaction, *types.Receipt, error) {
	_Migrations.transactionSession.Lock()
	tx, err := _Migrations.Contract.SetCompleted(_Migrations.transactionSession.TransactOpts, completed)
	if err != nil {
		_Migrations.transactionSession.Unlock()
		return nil, nil, err
	}
	_Migrations.transactionSession.TransactOpts.Nonce.Add(_Migrations.transactionSession.TransactOpts.Nonce, big.NewInt(1))
	_Migrations.transactionSession.Unlock()
	receipt, err := _Migrations.transactionSession.WaitTransaction(tx)
	return tx, receipt, err
}

// SetCompleted returns raw transaction bound to the contract method 0xfdacd576.
//
// Solidity: function setCompleted(uint256 completed) returns()
func (_Migrations *MigrationsSession) SetCompletedRawTx(completed *big.Int) (*types.Transaction, error) {
	tx, err := _Migrations.Contract.SetCompletedRawTx(_Migrations.transactionSession.TransactOpts, completed)
	return tx, err
}

// SetCompleted is a paid mutator transaction binding the contract method 0xfdacd576.
// Will not wait for tx, but put it to ch
//
// Solidity: function setCompleted(uint256 completed) returns()
func (_Migrations *MigrationsSession) SetCompletedAsync(receiptCh chan *types.ReceiptResult, completed *big.Int) (*types.Transaction, error) {
	_Migrations.transactionSession.Lock()
	tx, err := _Migrations.Contract.SetCompleted(_Migrations.transactionSession.TransactOpts, completed)
	if err != nil {
		_Migrations.transactionSession.Unlock()
		return nil, err
	}
	_Migrations.transactionSession.TransactOpts.Nonce.Add(_Migrations.transactionSession.TransactOpts.Nonce, big.NewInt(1))
	_Migrations.transactionSession.Unlock()
	go func() {
		receipt, err := _Migrations.transactionSession.WaitTransaction(tx)
		receiptCh <- &types.ReceiptResult{
			Receipt: *receipt,
			Err:     err,
		}
	}()
	return tx, err
}

// SetCompleted is a paid mutator transaction binding the contract method 0xfdacd576.
//
// Solidity: function setCompleted(uint256 completed) returns()
func (_Migrations *MigrationsTransactorSession) SetCompleted(wait bool, completed *big.Int) (*types.Transaction, error) {
	return _Migrations.Contract.SetCompleted(&_Migrations.TransactOpts, completed)
}

// Upgrade is a paid mutator transaction binding the contract method 0x0900f010.
//
// Solidity: function upgrade(address new_address) returns()
func (_Migrations *MigrationsTransactor) Upgrade(opts *bind.TransactOpts, new_address common.Address) (*types.Transaction, error) {
	return _Migrations.contract.Transact(opts, "upgrade", new_address)
}

func (_Migrations *MigrationsTransactor) UpgradeRawTx(opts *bind.TransactOpts, new_address common.Address) (*types.Transaction, error) {
	return _Migrations.contract.RawTx(opts, "upgrade", new_address)
}

// Upgrade is a paid mutator transaction binding the contract method 0x0900f010.
// Will wait for tx receipt
//
// Solidity: function upgrade(address new_address) returns()
func (_Migrations *MigrationsSession) Upgrade(new_address common.Address) (*types.Transaction, *types.Receipt, error) {
	_Migrations.transactionSession.Lock()
	tx, err := _Migrations.Contract.Upgrade(_Migrations.transactionSession.TransactOpts, new_address)
	if err != nil {
		_Migrations.transactionSession.Unlock()
		return nil, nil, err
	}
	_Migrations.transactionSession.TransactOpts.Nonce.Add(_Migrations.transactionSession.TransactOpts.Nonce, big.NewInt(1))
	_Migrations.transactionSession.Unlock()
	receipt, err := _Migrations.transactionSession.WaitTransaction(tx)
	return tx, receipt, err
}

// Upgrade returns raw transaction bound to the contract method 0x0900f010.
//
// Solidity: function upgrade(address new_address) returns()
func (_Migrations *MigrationsSession) UpgradeRawTx(new_address common.Address) (*types.Transaction, error) {
	tx, err := _Migrations.Contract.UpgradeRawTx(_Migrations.transactionSession.TransactOpts, new_address)
	return tx, err
}

// Upgrade is a paid mutator transaction binding the contract method 0x0900f010.
// Will not wait for tx, but put it to ch
//
// Solidity: function upgrade(address new_address) returns()
func (_Migrations *MigrationsSession) UpgradeAsync(receiptCh chan *types.ReceiptResult, new_address common.Address) (*types.Transaction, error) {
	_Migrations.transactionSession.Lock()
	tx, err := _Migrations.Contract.Upgrade(_Migrations.transactionSession.TransactOpts, new_address)
	if err != nil {
		_Migrations.transactionSession.Unlock()
		return nil, err
	}
	_Migrations.transactionSession.TransactOpts.Nonce.Add(_Migrations.transactionSession.TransactOpts.Nonce, big.NewInt(1))
	_Migrations.transactionSession.Unlock()
	go func() {
		receipt, err := _Migrations.transactionSession.WaitTransaction(tx)
		receiptCh <- &types.ReceiptResult{
			Receipt: *receipt,
			Err:     err,
		}
	}()
	return tx, err
}

// Upgrade is a paid mutator transaction binding the contract method 0x0900f010.
//
// Solidity: function upgrade(address new_address) returns()
func (_Migrations *MigrationsTransactorSession) Upgrade(wait bool, new_address common.Address) (*types.Transaction, error) {
	return _Migrations.Contract.Upgrade(&_Migrations.TransactOpts, new_address)
}
