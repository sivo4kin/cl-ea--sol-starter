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

// DaiTokenABI is the input ABI used to generate the binding from.
const DaiTokenABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// DaiTokenBin is the compiled bytecode used for deploying new contracts.
var DaiTokenBin = "0x60806040523480156200001157600080fd5b50604051806040016040528060098152602001682230b4902a37b5b2b760b91b8152506040518060400160405280600381526020016244414960e81b8152508160039080519060200190620000689291906200024a565b5080516200007e9060049060208401906200024a565b50506005805460ff1916601217905550620000ac3368056bc75e2d631000006001600160e01b03620000b216565b620002ef565b620000c782826001600160e01b03620000cb16565b5050565b6001600160a01b03821662000127576040805162461bcd60e51b815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f206164647265737300604482015290519081900360640190fd5b6200013e600083836001600160e01b03620001e316565b6200015a81600254620001e860201b6200061d1790919060201c565b6002556001600160a01b038216600090815260208181526040909120546200018d9183906200061d620001e8821b17901c565b6001600160a01b0383166000818152602081815260408083209490945583518581529351929391927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9281900390910190a35050565b505050565b60008282018381101562000243576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b9392505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200028d57805160ff1916838001178555620002bd565b82800160010185558215620002bd579182015b82811115620002bd578251825591602001919060010190620002a0565b50620002cb929150620002cf565b5090565b620002ec91905b80821115620002cb5760008155600101620002d6565b90565b610d3180620002ff6000396000f3fe608060405234801561001057600080fd5b50600436106100cf5760003560e01c806340c10f191161008c5780639dc29fac116100665780639dc29fac14610287578063a457c2d7146102b3578063a9059cbb146102df578063dd62ed3e1461030b576100cf565b806340c10f191461022b57806370a082311461025957806395d89b411461027f576100cf565b806306fdde03146100d4578063095ea7b31461015157806318160ddd1461019157806323b872dd146101ab578063313ce567146101e157806339509351146101ff575b600080fd5b6100dc610339565b6040805160208082528351818301528351919283929083019185019080838360005b838110156101165781810151838201526020016100fe565b50505050905090810190601f1680156101435780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b61017d6004803603604081101561016757600080fd5b506001600160a01b0381351690602001356103cf565b604080519115158252519081900360200190f35b6101996103ec565b60408051918252519081900360200190f35b61017d600480360360608110156101c157600080fd5b506001600160a01b038135811691602081013590911690604001356103f2565b6101e961047f565b6040805160ff9092168252519081900360200190f35b61017d6004803603604081101561021557600080fd5b506001600160a01b038135169060200135610488565b6102576004803603604081101561024157600080fd5b506001600160a01b0381351690602001356104dc565b005b6101996004803603602081101561026f57600080fd5b50356001600160a01b03166104ea565b6100dc610505565b6102576004803603604081101561029d57600080fd5b506001600160a01b038135169060200135610566565b61017d600480360360408110156102c957600080fd5b506001600160a01b038135169060200135610570565b61017d600480360360408110156102f557600080fd5b506001600160a01b0381351690602001356105de565b6101996004803603604081101561032157600080fd5b506001600160a01b03813581169160200135166105f2565b60038054604080516020601f60026000196101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156103c55780601f1061039a576101008083540402835291602001916103c5565b820191906000526020600020905b8154815290600101906020018083116103a857829003601f168201915b5050505050905090565b60006103e36103dc61067e565b8484610682565b50600192915050565b60025490565b60006103ff84848461076e565b6104758461040b61067e565b61047085604051806060016040528060288152602001610c45602891396001600160a01b038a1660009081526001602052604081209061044961067e565b6001600160a01b03168152602081019190915260400160002054919063ffffffff6108d516565b610682565b5060019392505050565b60055460ff1690565b60006103e361049561067e565b8461047085600160006104a661067e565b6001600160a01b03908116825260208083019390935260409182016000908120918c16815292529020549063ffffffff61061d16565b6104e6828261096c565b5050565b6001600160a01b031660009081526020819052604090205490565b60048054604080516020601f60026000196101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156103c55780601f1061039a576101008083540402835291602001916103c5565b6104e68282610a68565b60006103e361057d61067e565b8461047085604051806060016040528060258152602001610cd760259139600160006105a761067e565b6001600160a01b03908116825260208083019390935260409182016000908120918d1681529252902054919063ffffffff6108d516565b60006103e36105eb61067e565b848461076e565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205490565b600082820183811015610677576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b9392505050565b3390565b6001600160a01b0383166106c75760405162461bcd60e51b8152600401808060200182810382526024815260200180610cb36024913960400191505060405180910390fd5b6001600160a01b03821661070c5760405162461bcd60e51b8152600401808060200182810382526022815260200180610bfd6022913960400191505060405180910390fd5b6001600160a01b03808416600081815260016020908152604080832094871680845294825291829020859055815185815291517f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259281900390910190a3505050565b6001600160a01b0383166107b35760405162461bcd60e51b8152600401808060200182810382526025815260200180610c8e6025913960400191505060405180910390fd5b6001600160a01b0382166107f85760405162461bcd60e51b8152600401808060200182810382526023815260200180610bb86023913960400191505060405180910390fd5b610803838383610b70565b61084681604051806060016040528060268152602001610c1f602691396001600160a01b038616600090815260208190526040902054919063ffffffff6108d516565b6001600160a01b03808516600090815260208190526040808220939093559084168152205461087b908263ffffffff61061d16565b6001600160a01b038084166000818152602081815260409182902094909455805185815290519193928716927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef92918290030190a3505050565b600081848411156109645760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b83811015610929578181015183820152602001610911565b50505050905090810190601f1680156109565780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b505050900390565b6001600160a01b0382166109c7576040805162461bcd60e51b815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f206164647265737300604482015290519081900360640190fd5b6109d360008383610b70565b6002546109e6908263ffffffff61061d16565b6002556001600160a01b038216600090815260208190526040902054610a12908263ffffffff61061d16565b6001600160a01b0383166000818152602081815260408083209490945583518581529351929391927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9281900390910190a35050565b6001600160a01b038216610aad5760405162461bcd60e51b8152600401808060200182810382526021815260200180610c6d6021913960400191505060405180910390fd5b610ab982600083610b70565b610afc81604051806060016040528060228152602001610bdb602291396001600160a01b038516600090815260208190526040902054919063ffffffff6108d516565b6001600160a01b038316600090815260208190526040902055600254610b28908263ffffffff610b7516565b6002556040805182815290516000916001600160a01b038516917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9181900360200190a35050565b505050565b600061067783836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f7700008152506108d556fe45524332303a207472616e7366657220746f20746865207a65726f206164647265737345524332303a206275726e20616d6f756e7420657863656564732062616c616e636545524332303a20617070726f766520746f20746865207a65726f206164647265737345524332303a207472616e7366657220616d6f756e7420657863656564732062616c616e636545524332303a207472616e7366657220616d6f756e74206578636565647320616c6c6f77616e636545524332303a206275726e2066726f6d20746865207a65726f206164647265737345524332303a207472616e736665722066726f6d20746865207a65726f206164647265737345524332303a20617070726f76652066726f6d20746865207a65726f206164647265737345524332303a2064656372656173656420616c6c6f77616e63652062656c6f77207a65726fa264697066735822122020f30d02e00672c148f17e8eefb64cb32f55c63104f396a2dab3373d2d23dfba64736f6c63430006090033"

// DeployDaiToken deploys a new Ethereum contract, binding an instance of DaiToken to it.
func DeployDaiToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DaiToken, error) {
	parsed, err := abi.JSON(strings.NewReader(DaiTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DaiTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DaiToken{DaiTokenCaller: DaiTokenCaller{contract: contract}, DaiTokenTransactor: DaiTokenTransactor{contract: contract}, DaiTokenFilterer: DaiTokenFilterer{contract: contract}}, nil
}

// DeployDaiTokenSync deploys a new Ethereum contract and waits for receipt, binding an instance of DaiTokenSession to it.
func DeployDaiTokenSync(session *bind.TransactSession, backend bind.ContractBackend) (*types.Transaction, *types.Receipt, *DaiTokenSession, error) {
	parsed, err := abi.JSON(strings.NewReader(DaiTokenABI))
	if err != nil {
		return nil, nil, nil, err
	}
	session.Lock()
	address, tx, _, err := bind.DeployContract(session.TransactOpts, parsed, common.FromHex(DaiTokenBin), backend)
	receipt, err := session.WaitTransaction(tx)
	if err != nil {
		session.Unlock()
		return nil, nil, nil, err
	}
	session.TransactOpts.Nonce.Add(session.TransactOpts.Nonce, big.NewInt(1))
	session.Unlock()
	contractSession, err := NewDaiTokenSession(address, backend, session)
	return tx, receipt, contractSession, err
}

// DaiToken is an auto generated Go binding around an Ethereum contract.
type DaiToken struct {
	DaiTokenCaller     // Read-only binding to the contract
	DaiTokenTransactor // Write-only binding to the contract
	DaiTokenFilterer   // Log filterer for contract events
}

// DaiTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type DaiTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DaiTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DaiTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DaiTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DaiTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DaiTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DaiTokenSession struct {
	Contract           *DaiToken // Generic contract binding to set the session for
	transactionSession *bind.TransactSession
	Address            common.Address
}

// DaiTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DaiTokenCallerSession struct {
	Contract *DaiTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// DaiTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DaiTokenTransactorSession struct {
	Contract     *DaiTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// DaiTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type DaiTokenRaw struct {
	Contract *DaiToken // Generic contract binding to access the raw methods on
}

// DaiTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DaiTokenCallerRaw struct {
	Contract *DaiTokenCaller // Generic read-only contract binding to access the raw methods on
}

// DaiTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DaiTokenTransactorRaw struct {
	Contract *DaiTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDaiToken creates a new instance of DaiToken, bound to a specific deployed contract.
func NewDaiToken(address common.Address, backend bind.ContractBackend) (*DaiToken, error) {
	contract, err := bindDaiToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DaiToken{DaiTokenCaller: DaiTokenCaller{contract: contract}, DaiTokenTransactor: DaiTokenTransactor{contract: contract}, DaiTokenFilterer: DaiTokenFilterer{contract: contract}}, nil
}

// NewDaiTokenCaller creates a new read-only instance of DaiToken, bound to a specific deployed contract.
func NewDaiTokenCaller(address common.Address, caller bind.ContractCaller) (*DaiTokenCaller, error) {
	contract, err := bindDaiToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DaiTokenCaller{contract: contract}, nil
}

// NewDaiTokenTransactor creates a new write-only instance of DaiToken, bound to a specific deployed contract.
func NewDaiTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*DaiTokenTransactor, error) {
	contract, err := bindDaiToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DaiTokenTransactor{contract: contract}, nil
}

// NewDaiTokenFilterer creates a new log filterer instance of DaiToken, bound to a specific deployed contract.
func NewDaiTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*DaiTokenFilterer, error) {
	contract, err := bindDaiToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DaiTokenFilterer{contract: contract}, nil
}

func NewDaiTokenSession(address common.Address, backend bind.ContractBackend, transactionSession *bind.TransactSession) (*DaiTokenSession, error) {
	DaiTokenInstance, err := NewDaiToken(address, backend)
	if err != nil {
		return nil, err
	}
	return &DaiTokenSession{
		Contract:           DaiTokenInstance,
		transactionSession: transactionSession,
		Address:            address,
	}, nil
}

// bindDaiToken binds a generic wrapper to an already deployed contract.
func bindDaiToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DaiTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DaiToken *DaiTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DaiToken.Contract.DaiTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DaiToken *DaiTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DaiToken.Contract.DaiTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DaiToken *DaiTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DaiToken.Contract.DaiTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DaiToken *DaiTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DaiToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DaiToken *DaiTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DaiToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DaiToken *DaiTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DaiToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_DaiToken *DaiTokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DaiToken.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_DaiToken *DaiTokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _DaiToken.Contract.Allowance(_DaiToken.transactionSession.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_DaiToken *DaiTokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _DaiToken.Contract.Allowance(&_DaiToken.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_DaiToken *DaiTokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DaiToken.contract.Call(opts, out, "balanceOf", account)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_DaiToken *DaiTokenSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _DaiToken.Contract.BalanceOf(_DaiToken.transactionSession.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_DaiToken *DaiTokenCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _DaiToken.Contract.BalanceOf(&_DaiToken.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_DaiToken *DaiTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _DaiToken.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_DaiToken *DaiTokenSession) Decimals() (uint8, error) {
	return _DaiToken.Contract.Decimals(_DaiToken.transactionSession.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_DaiToken *DaiTokenCallerSession) Decimals() (uint8, error) {
	return _DaiToken.Contract.Decimals(&_DaiToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_DaiToken *DaiTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _DaiToken.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_DaiToken *DaiTokenSession) Name() (string, error) {
	return _DaiToken.Contract.Name(_DaiToken.transactionSession.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_DaiToken *DaiTokenCallerSession) Name() (string, error) {
	return _DaiToken.Contract.Name(&_DaiToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_DaiToken *DaiTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _DaiToken.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_DaiToken *DaiTokenSession) Symbol() (string, error) {
	return _DaiToken.Contract.Symbol(_DaiToken.transactionSession.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_DaiToken *DaiTokenCallerSession) Symbol() (string, error) {
	return _DaiToken.Contract.Symbol(&_DaiToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_DaiToken *DaiTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DaiToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_DaiToken *DaiTokenSession) TotalSupply() (*big.Int, error) {
	return _DaiToken.Contract.TotalSupply(_DaiToken.transactionSession.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_DaiToken *DaiTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _DaiToken.Contract.TotalSupply(&_DaiToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_DaiToken *DaiTokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DaiToken.contract.Transact(opts, "approve", spender, amount)
}

func (_DaiToken *DaiTokenTransactor) ApproveRawTx(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DaiToken.contract.RawTx(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
// Will wait for tx receipt
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_DaiToken *DaiTokenSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, *types.Receipt, error) {
	_DaiToken.transactionSession.Lock()
	tx, err := _DaiToken.Contract.Approve(_DaiToken.transactionSession.TransactOpts, spender, amount)
	if err != nil {
		_DaiToken.transactionSession.Unlock()
		return nil, nil, err
	}
	_DaiToken.transactionSession.TransactOpts.Nonce.Add(_DaiToken.transactionSession.TransactOpts.Nonce, big.NewInt(1))
	_DaiToken.transactionSession.Unlock()
	receipt, err := _DaiToken.transactionSession.WaitTransaction(tx)
	return tx, receipt, err
}

// Approve returns raw transaction bound to the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_DaiToken *DaiTokenSession) ApproveRawTx(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	tx, err := _DaiToken.Contract.ApproveRawTx(_DaiToken.transactionSession.TransactOpts, spender, amount)
	return tx, err
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
// Will not wait for tx, but put it to ch
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_DaiToken *DaiTokenSession) ApproveAsync(receiptCh chan *types.ReceiptResult, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	_DaiToken.transactionSession.Lock()
	tx, err := _DaiToken.Contract.Approve(_DaiToken.transactionSession.TransactOpts, spender, amount)
	if err != nil {
		_DaiToken.transactionSession.Unlock()
		return nil, err
	}
	_DaiToken.transactionSession.TransactOpts.Nonce.Add(_DaiToken.transactionSession.TransactOpts.Nonce, big.NewInt(1))
	_DaiToken.transactionSession.Unlock()
	go func() {
		receipt, err := _DaiToken.transactionSession.WaitTransaction(tx)
		receiptCh <- &types.ReceiptResult{
			Receipt: *receipt,
			Err:     err,
		}
	}()
	return tx, err
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_DaiToken *DaiTokenTransactorSession) Approve(wait bool, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DaiToken.Contract.Approve(&_DaiToken.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address _account, uint256 _amount) returns()
func (_DaiToken *DaiTokenTransactor) Burn(opts *bind.TransactOpts, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _DaiToken.contract.Transact(opts, "burn", _account, _amount)
}

func (_DaiToken *DaiTokenTransactor) BurnRawTx(opts *bind.TransactOpts, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _DaiToken.contract.RawTx(opts, "burn", _account, _amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
// Will wait for tx receipt
//
// Solidity: function burn(address _account, uint256 _amount) returns()
func (_DaiToken *DaiTokenSession) Burn(_account common.Address, _amount *big.Int) (*types.Transaction, *types.Receipt, error) {
	_DaiToken.transactionSession.Lock()
	tx, err := _DaiToken.Contract.Burn(_DaiToken.transactionSession.TransactOpts, _account, _amount)
	if err != nil {
		_DaiToken.transactionSession.Unlock()
		return nil, nil, err
	}
	_DaiToken.transactionSession.TransactOpts.Nonce.Add(_DaiToken.transactionSession.TransactOpts.Nonce, big.NewInt(1))
	_DaiToken.transactionSession.Unlock()
	receipt, err := _DaiToken.transactionSession.WaitTransaction(tx)
	return tx, receipt, err
}

// Burn returns raw transaction bound to the contract method 0x9dc29fac.
//
// Solidity: function burn(address _account, uint256 _amount) returns()
func (_DaiToken *DaiTokenSession) BurnRawTx(_account common.Address, _amount *big.Int) (*types.Transaction, error) {
	tx, err := _DaiToken.Contract.BurnRawTx(_DaiToken.transactionSession.TransactOpts, _account, _amount)
	return tx, err
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
// Will not wait for tx, but put it to ch
//
// Solidity: function burn(address _account, uint256 _amount) returns()
func (_DaiToken *DaiTokenSession) BurnAsync(receiptCh chan *types.ReceiptResult, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	_DaiToken.transactionSession.Lock()
	tx, err := _DaiToken.Contract.Burn(_DaiToken.transactionSession.TransactOpts, _account, _amount)
	if err != nil {
		_DaiToken.transactionSession.Unlock()
		return nil, err
	}
	_DaiToken.transactionSession.TransactOpts.Nonce.Add(_DaiToken.transactionSession.TransactOpts.Nonce, big.NewInt(1))
	_DaiToken.transactionSession.Unlock()
	go func() {
		receipt, err := _DaiToken.transactionSession.WaitTransaction(tx)
		receiptCh <- &types.ReceiptResult{
			Receipt: *receipt,
			Err:     err,
		}
	}()
	return tx, err
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address _account, uint256 _amount) returns()
func (_DaiToken *DaiTokenTransactorSession) Burn(wait bool, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _DaiToken.Contract.Burn(&_DaiToken.TransactOpts, _account, _amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_DaiToken *DaiTokenTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _DaiToken.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

func (_DaiToken *DaiTokenTransactor) DecreaseAllowanceRawTx(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _DaiToken.contract.RawTx(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
// Will wait for tx receipt
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_DaiToken *DaiTokenSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, *types.Receipt, error) {
	_DaiToken.transactionSession.Lock()
	tx, err := _DaiToken.Contract.DecreaseAllowance(_DaiToken.transactionSession.TransactOpts, spender, subtractedValue)
	if err != nil {
		_DaiToken.transactionSession.Unlock()
		return nil, nil, err
	}
	_DaiToken.transactionSession.TransactOpts.Nonce.Add(_DaiToken.transactionSession.TransactOpts.Nonce, big.NewInt(1))
	_DaiToken.transactionSession.Unlock()
	receipt, err := _DaiToken.transactionSession.WaitTransaction(tx)
	return tx, receipt, err
}

// DecreaseAllowance returns raw transaction bound to the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_DaiToken *DaiTokenSession) DecreaseAllowanceRawTx(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	tx, err := _DaiToken.Contract.DecreaseAllowanceRawTx(_DaiToken.transactionSession.TransactOpts, spender, subtractedValue)
	return tx, err
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
// Will not wait for tx, but put it to ch
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_DaiToken *DaiTokenSession) DecreaseAllowanceAsync(receiptCh chan *types.ReceiptResult, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	_DaiToken.transactionSession.Lock()
	tx, err := _DaiToken.Contract.DecreaseAllowance(_DaiToken.transactionSession.TransactOpts, spender, subtractedValue)
	if err != nil {
		_DaiToken.transactionSession.Unlock()
		return nil, err
	}
	_DaiToken.transactionSession.TransactOpts.Nonce.Add(_DaiToken.transactionSession.TransactOpts.Nonce, big.NewInt(1))
	_DaiToken.transactionSession.Unlock()
	go func() {
		receipt, err := _DaiToken.transactionSession.WaitTransaction(tx)
		receiptCh <- &types.ReceiptResult{
			Receipt: *receipt,
			Err:     err,
		}
	}()
	return tx, err
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_DaiToken *DaiTokenTransactorSession) DecreaseAllowance(wait bool, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _DaiToken.Contract.DecreaseAllowance(&_DaiToken.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_DaiToken *DaiTokenTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _DaiToken.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

func (_DaiToken *DaiTokenTransactor) IncreaseAllowanceRawTx(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _DaiToken.contract.RawTx(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
// Will wait for tx receipt
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_DaiToken *DaiTokenSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, *types.Receipt, error) {
	_DaiToken.transactionSession.Lock()
	tx, err := _DaiToken.Contract.IncreaseAllowance(_DaiToken.transactionSession.TransactOpts, spender, addedValue)
	if err != nil {
		_DaiToken.transactionSession.Unlock()
		return nil, nil, err
	}
	_DaiToken.transactionSession.TransactOpts.Nonce.Add(_DaiToken.transactionSession.TransactOpts.Nonce, big.NewInt(1))
	_DaiToken.transactionSession.Unlock()
	receipt, err := _DaiToken.transactionSession.WaitTransaction(tx)
	return tx, receipt, err
}

// IncreaseAllowance returns raw transaction bound to the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_DaiToken *DaiTokenSession) IncreaseAllowanceRawTx(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	tx, err := _DaiToken.Contract.IncreaseAllowanceRawTx(_DaiToken.transactionSession.TransactOpts, spender, addedValue)
	return tx, err
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
// Will not wait for tx, but put it to ch
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_DaiToken *DaiTokenSession) IncreaseAllowanceAsync(receiptCh chan *types.ReceiptResult, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	_DaiToken.transactionSession.Lock()
	tx, err := _DaiToken.Contract.IncreaseAllowance(_DaiToken.transactionSession.TransactOpts, spender, addedValue)
	if err != nil {
		_DaiToken.transactionSession.Unlock()
		return nil, err
	}
	_DaiToken.transactionSession.TransactOpts.Nonce.Add(_DaiToken.transactionSession.TransactOpts.Nonce, big.NewInt(1))
	_DaiToken.transactionSession.Unlock()
	go func() {
		receipt, err := _DaiToken.transactionSession.WaitTransaction(tx)
		receiptCh <- &types.ReceiptResult{
			Receipt: *receipt,
			Err:     err,
		}
	}()
	return tx, err
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_DaiToken *DaiTokenTransactorSession) IncreaseAllowance(wait bool, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _DaiToken.Contract.IncreaseAllowance(&_DaiToken.TransactOpts, spender, addedValue)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _account, uint256 _amount) returns()
func (_DaiToken *DaiTokenTransactor) Mint(opts *bind.TransactOpts, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _DaiToken.contract.Transact(opts, "mint", _account, _amount)
}

func (_DaiToken *DaiTokenTransactor) MintRawTx(opts *bind.TransactOpts, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _DaiToken.contract.RawTx(opts, "mint", _account, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
// Will wait for tx receipt
//
// Solidity: function mint(address _account, uint256 _amount) returns()
func (_DaiToken *DaiTokenSession) Mint(_account common.Address, _amount *big.Int) (*types.Transaction, *types.Receipt, error) {
	_DaiToken.transactionSession.Lock()
	tx, err := _DaiToken.Contract.Mint(_DaiToken.transactionSession.TransactOpts, _account, _amount)
	if err != nil {
		_DaiToken.transactionSession.Unlock()
		return nil, nil, err
	}
	_DaiToken.transactionSession.TransactOpts.Nonce.Add(_DaiToken.transactionSession.TransactOpts.Nonce, big.NewInt(1))
	_DaiToken.transactionSession.Unlock()
	receipt, err := _DaiToken.transactionSession.WaitTransaction(tx)
	return tx, receipt, err
}

// Mint returns raw transaction bound to the contract method 0x40c10f19.
//
// Solidity: function mint(address _account, uint256 _amount) returns()
func (_DaiToken *DaiTokenSession) MintRawTx(_account common.Address, _amount *big.Int) (*types.Transaction, error) {
	tx, err := _DaiToken.Contract.MintRawTx(_DaiToken.transactionSession.TransactOpts, _account, _amount)
	return tx, err
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
// Will not wait for tx, but put it to ch
//
// Solidity: function mint(address _account, uint256 _amount) returns()
func (_DaiToken *DaiTokenSession) MintAsync(receiptCh chan *types.ReceiptResult, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	_DaiToken.transactionSession.Lock()
	tx, err := _DaiToken.Contract.Mint(_DaiToken.transactionSession.TransactOpts, _account, _amount)
	if err != nil {
		_DaiToken.transactionSession.Unlock()
		return nil, err
	}
	_DaiToken.transactionSession.TransactOpts.Nonce.Add(_DaiToken.transactionSession.TransactOpts.Nonce, big.NewInt(1))
	_DaiToken.transactionSession.Unlock()
	go func() {
		receipt, err := _DaiToken.transactionSession.WaitTransaction(tx)
		receiptCh <- &types.ReceiptResult{
			Receipt: *receipt,
			Err:     err,
		}
	}()
	return tx, err
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _account, uint256 _amount) returns()
func (_DaiToken *DaiTokenTransactorSession) Mint(wait bool, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _DaiToken.Contract.Mint(&_DaiToken.TransactOpts, _account, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_DaiToken *DaiTokenTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DaiToken.contract.Transact(opts, "transfer", recipient, amount)
}

func (_DaiToken *DaiTokenTransactor) TransferRawTx(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DaiToken.contract.RawTx(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
// Will wait for tx receipt
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_DaiToken *DaiTokenSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, *types.Receipt, error) {
	_DaiToken.transactionSession.Lock()
	tx, err := _DaiToken.Contract.Transfer(_DaiToken.transactionSession.TransactOpts, recipient, amount)
	if err != nil {
		_DaiToken.transactionSession.Unlock()
		return nil, nil, err
	}
	_DaiToken.transactionSession.TransactOpts.Nonce.Add(_DaiToken.transactionSession.TransactOpts.Nonce, big.NewInt(1))
	_DaiToken.transactionSession.Unlock()
	receipt, err := _DaiToken.transactionSession.WaitTransaction(tx)
	return tx, receipt, err
}

// Transfer returns raw transaction bound to the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_DaiToken *DaiTokenSession) TransferRawTx(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	tx, err := _DaiToken.Contract.TransferRawTx(_DaiToken.transactionSession.TransactOpts, recipient, amount)
	return tx, err
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
// Will not wait for tx, but put it to ch
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_DaiToken *DaiTokenSession) TransferAsync(receiptCh chan *types.ReceiptResult, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	_DaiToken.transactionSession.Lock()
	tx, err := _DaiToken.Contract.Transfer(_DaiToken.transactionSession.TransactOpts, recipient, amount)
	if err != nil {
		_DaiToken.transactionSession.Unlock()
		return nil, err
	}
	_DaiToken.transactionSession.TransactOpts.Nonce.Add(_DaiToken.transactionSession.TransactOpts.Nonce, big.NewInt(1))
	_DaiToken.transactionSession.Unlock()
	go func() {
		receipt, err := _DaiToken.transactionSession.WaitTransaction(tx)
		receiptCh <- &types.ReceiptResult{
			Receipt: *receipt,
			Err:     err,
		}
	}()
	return tx, err
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_DaiToken *DaiTokenTransactorSession) Transfer(wait bool, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DaiToken.Contract.Transfer(&_DaiToken.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_DaiToken *DaiTokenTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DaiToken.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

func (_DaiToken *DaiTokenTransactor) TransferFromRawTx(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DaiToken.contract.RawTx(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
// Will wait for tx receipt
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_DaiToken *DaiTokenSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, *types.Receipt, error) {
	_DaiToken.transactionSession.Lock()
	tx, err := _DaiToken.Contract.TransferFrom(_DaiToken.transactionSession.TransactOpts, sender, recipient, amount)
	if err != nil {
		_DaiToken.transactionSession.Unlock()
		return nil, nil, err
	}
	_DaiToken.transactionSession.TransactOpts.Nonce.Add(_DaiToken.transactionSession.TransactOpts.Nonce, big.NewInt(1))
	_DaiToken.transactionSession.Unlock()
	receipt, err := _DaiToken.transactionSession.WaitTransaction(tx)
	return tx, receipt, err
}

// TransferFrom returns raw transaction bound to the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_DaiToken *DaiTokenSession) TransferFromRawTx(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	tx, err := _DaiToken.Contract.TransferFromRawTx(_DaiToken.transactionSession.TransactOpts, sender, recipient, amount)
	return tx, err
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
// Will not wait for tx, but put it to ch
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_DaiToken *DaiTokenSession) TransferFromAsync(receiptCh chan *types.ReceiptResult, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	_DaiToken.transactionSession.Lock()
	tx, err := _DaiToken.Contract.TransferFrom(_DaiToken.transactionSession.TransactOpts, sender, recipient, amount)
	if err != nil {
		_DaiToken.transactionSession.Unlock()
		return nil, err
	}
	_DaiToken.transactionSession.TransactOpts.Nonce.Add(_DaiToken.transactionSession.TransactOpts.Nonce, big.NewInt(1))
	_DaiToken.transactionSession.Unlock()
	go func() {
		receipt, err := _DaiToken.transactionSession.WaitTransaction(tx)
		receiptCh <- &types.ReceiptResult{
			Receipt: *receipt,
			Err:     err,
		}
	}()
	return tx, err
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_DaiToken *DaiTokenTransactorSession) TransferFrom(wait bool, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DaiToken.Contract.TransferFrom(&_DaiToken.TransactOpts, sender, recipient, amount)
}

// DaiTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the DaiToken contract.
type DaiTokenApprovalIterator struct {
	Event *DaiTokenApproval // Event containing the contract specifics and raw log

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
func (it *DaiTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DaiTokenApproval)
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
		it.Event = new(DaiTokenApproval)
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
func (it *DaiTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DaiTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DaiTokenApproval represents a Approval event raised by the DaiToken contract.
type DaiTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_DaiToken *DaiTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*DaiTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _DaiToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &DaiTokenApprovalIterator{contract: _DaiToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_DaiToken *DaiTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *DaiTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _DaiToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DaiTokenApproval)
				if err := _DaiToken.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_DaiToken *DaiTokenFilterer) ParseApproval(log types.Log) (*DaiTokenApproval, error) {
	event := new(DaiTokenApproval)
	if err := _DaiToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DaiTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the DaiToken contract.
type DaiTokenTransferIterator struct {
	Event *DaiTokenTransfer // Event containing the contract specifics and raw log

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
func (it *DaiTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DaiTokenTransfer)
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
		it.Event = new(DaiTokenTransfer)
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
func (it *DaiTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DaiTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DaiTokenTransfer represents a Transfer event raised by the DaiToken contract.
type DaiTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_DaiToken *DaiTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*DaiTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _DaiToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &DaiTokenTransferIterator{contract: _DaiToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_DaiToken *DaiTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *DaiTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _DaiToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DaiTokenTransfer)
				if err := _DaiToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_DaiToken *DaiTokenFilterer) ParseTransfer(log types.Log) (*DaiTokenTransfer, error) {
	event := new(DaiTokenTransfer)
	if err := _DaiToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}
