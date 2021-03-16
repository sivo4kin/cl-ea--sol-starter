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

// BTCTokenABI is the input ABI used to generate the binding from.
const BTCTokenABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// BTCTokenBin is the compiled bytecode used for deploying new contracts.
var BTCTokenBin = "0x60806040523480156200001157600080fd5b5060405180604001604052806009815260200168212a21902a37b5b2b760b91b8152506040518060400160405280600381526020016242544360e81b8152508160039080519060200190620000689291906200024a565b5080516200007e9060049060208401906200024a565b50506005805460ff1916601217905550620000ac3368056bc75e2d631000006001600160e01b03620000b216565b620002ef565b620000c782826001600160e01b03620000cb16565b5050565b6001600160a01b03821662000127576040805162461bcd60e51b815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f206164647265737300604482015290519081900360640190fd5b6200013e600083836001600160e01b03620001e316565b6200015a81600254620001e860201b6200061d1790919060201c565b6002556001600160a01b038216600090815260208181526040909120546200018d9183906200061d620001e8821b17901c565b6001600160a01b0383166000818152602081815260408083209490945583518581529351929391927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9281900390910190a35050565b505050565b60008282018381101562000243576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b9392505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200028d57805160ff1916838001178555620002bd565b82800160010185558215620002bd579182015b82811115620002bd578251825591602001919060010190620002a0565b50620002cb929150620002cf565b5090565b620002ec91905b80821115620002cb5760008155600101620002d6565b90565b610d3180620002ff6000396000f3fe608060405234801561001057600080fd5b50600436106100cf5760003560e01c806340c10f191161008c5780639dc29fac116100665780639dc29fac14610287578063a457c2d7146102b3578063a9059cbb146102df578063dd62ed3e1461030b576100cf565b806340c10f191461022b57806370a082311461025957806395d89b411461027f576100cf565b806306fdde03146100d4578063095ea7b31461015157806318160ddd1461019157806323b872dd146101ab578063313ce567146101e157806339509351146101ff575b600080fd5b6100dc610339565b6040805160208082528351818301528351919283929083019185019080838360005b838110156101165781810151838201526020016100fe565b50505050905090810190601f1680156101435780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b61017d6004803603604081101561016757600080fd5b506001600160a01b0381351690602001356103cf565b604080519115158252519081900360200190f35b6101996103ec565b60408051918252519081900360200190f35b61017d600480360360608110156101c157600080fd5b506001600160a01b038135811691602081013590911690604001356103f2565b6101e961047f565b6040805160ff9092168252519081900360200190f35b61017d6004803603604081101561021557600080fd5b506001600160a01b038135169060200135610488565b6102576004803603604081101561024157600080fd5b506001600160a01b0381351690602001356104dc565b005b6101996004803603602081101561026f57600080fd5b50356001600160a01b03166104ea565b6100dc610505565b6102576004803603604081101561029d57600080fd5b506001600160a01b038135169060200135610566565b61017d600480360360408110156102c957600080fd5b506001600160a01b038135169060200135610570565b61017d600480360360408110156102f557600080fd5b506001600160a01b0381351690602001356105de565b6101996004803603604081101561032157600080fd5b506001600160a01b03813581169160200135166105f2565b60038054604080516020601f60026000196101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156103c55780601f1061039a576101008083540402835291602001916103c5565b820191906000526020600020905b8154815290600101906020018083116103a857829003601f168201915b5050505050905090565b60006103e36103dc61067e565b8484610682565b50600192915050565b60025490565b60006103ff84848461076e565b6104758461040b61067e565b61047085604051806060016040528060288152602001610c45602891396001600160a01b038a1660009081526001602052604081209061044961067e565b6001600160a01b03168152602081019190915260400160002054919063ffffffff6108d516565b610682565b5060019392505050565b60055460ff1690565b60006103e361049561067e565b8461047085600160006104a661067e565b6001600160a01b03908116825260208083019390935260409182016000908120918c16815292529020549063ffffffff61061d16565b6104e6828261096c565b5050565b6001600160a01b031660009081526020819052604090205490565b60048054604080516020601f60026000196101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156103c55780601f1061039a576101008083540402835291602001916103c5565b6104e68282610a68565b60006103e361057d61067e565b8461047085604051806060016040528060258152602001610cd760259139600160006105a761067e565b6001600160a01b03908116825260208083019390935260409182016000908120918d1681529252902054919063ffffffff6108d516565b60006103e36105eb61067e565b848461076e565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205490565b600082820183811015610677576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b9392505050565b3390565b6001600160a01b0383166106c75760405162461bcd60e51b8152600401808060200182810382526024815260200180610cb36024913960400191505060405180910390fd5b6001600160a01b03821661070c5760405162461bcd60e51b8152600401808060200182810382526022815260200180610bfd6022913960400191505060405180910390fd5b6001600160a01b03808416600081815260016020908152604080832094871680845294825291829020859055815185815291517f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259281900390910190a3505050565b6001600160a01b0383166107b35760405162461bcd60e51b8152600401808060200182810382526025815260200180610c8e6025913960400191505060405180910390fd5b6001600160a01b0382166107f85760405162461bcd60e51b8152600401808060200182810382526023815260200180610bb86023913960400191505060405180910390fd5b610803838383610b70565b61084681604051806060016040528060268152602001610c1f602691396001600160a01b038616600090815260208190526040902054919063ffffffff6108d516565b6001600160a01b03808516600090815260208190526040808220939093559084168152205461087b908263ffffffff61061d16565b6001600160a01b038084166000818152602081815260409182902094909455805185815290519193928716927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef92918290030190a3505050565b600081848411156109645760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b83811015610929578181015183820152602001610911565b50505050905090810190601f1680156109565780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b505050900390565b6001600160a01b0382166109c7576040805162461bcd60e51b815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f206164647265737300604482015290519081900360640190fd5b6109d360008383610b70565b6002546109e6908263ffffffff61061d16565b6002556001600160a01b038216600090815260208190526040902054610a12908263ffffffff61061d16565b6001600160a01b0383166000818152602081815260408083209490945583518581529351929391927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9281900390910190a35050565b6001600160a01b038216610aad5760405162461bcd60e51b8152600401808060200182810382526021815260200180610c6d6021913960400191505060405180910390fd5b610ab982600083610b70565b610afc81604051806060016040528060228152602001610bdb602291396001600160a01b038516600090815260208190526040902054919063ffffffff6108d516565b6001600160a01b038316600090815260208190526040902055600254610b28908263ffffffff610b7516565b6002556040805182815290516000916001600160a01b038516917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9181900360200190a35050565b505050565b600061067783836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f7700008152506108d556fe45524332303a207472616e7366657220746f20746865207a65726f206164647265737345524332303a206275726e20616d6f756e7420657863656564732062616c616e636545524332303a20617070726f766520746f20746865207a65726f206164647265737345524332303a207472616e7366657220616d6f756e7420657863656564732062616c616e636545524332303a207472616e7366657220616d6f756e74206578636565647320616c6c6f77616e636545524332303a206275726e2066726f6d20746865207a65726f206164647265737345524332303a207472616e736665722066726f6d20746865207a65726f206164647265737345524332303a20617070726f76652066726f6d20746865207a65726f206164647265737345524332303a2064656372656173656420616c6c6f77616e63652062656c6f77207a65726fa264697066735822122012f221827b3c5bc062071cbd1df29413b83eb4de4b4edaa1af8bb9d2e577e53c64736f6c63430006090033"

// DeployBTCToken deploys a new Ethereum contract, binding an instance of BTCToken to it.
func DeployBTCToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BTCToken, error) {
	parsed, err := abi.JSON(strings.NewReader(BTCTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BTCTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BTCToken{BTCTokenCaller: BTCTokenCaller{contract: contract}, BTCTokenTransactor: BTCTokenTransactor{contract: contract}, BTCTokenFilterer: BTCTokenFilterer{contract: contract}}, nil
}

// BTCToken is an auto generated Go binding around an Ethereum contract.
type BTCToken struct {
	BTCTokenCaller     // Read-only binding to the contract
	BTCTokenTransactor // Write-only binding to the contract
	BTCTokenFilterer   // Log filterer for contract events
}

// BTCTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type BTCTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BTCTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BTCTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BTCTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BTCTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BTCTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BTCTokenSession struct {
	Contract     *BTCToken         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BTCTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BTCTokenCallerSession struct {
	Contract *BTCTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// BTCTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BTCTokenTransactorSession struct {
	Contract     *BTCTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// BTCTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type BTCTokenRaw struct {
	Contract *BTCToken // Generic contract binding to access the raw methods on
}

// BTCTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BTCTokenCallerRaw struct {
	Contract *BTCTokenCaller // Generic read-only contract binding to access the raw methods on
}

// BTCTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BTCTokenTransactorRaw struct {
	Contract *BTCTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBTCToken creates a new instance of BTCToken, bound to a specific deployed contract.
func NewBTCToken(address common.Address, backend bind.ContractBackend) (*BTCToken, error) {
	contract, err := bindBTCToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BTCToken{BTCTokenCaller: BTCTokenCaller{contract: contract}, BTCTokenTransactor: BTCTokenTransactor{contract: contract}, BTCTokenFilterer: BTCTokenFilterer{contract: contract}}, nil
}

// NewBTCTokenCaller creates a new read-only instance of BTCToken, bound to a specific deployed contract.
func NewBTCTokenCaller(address common.Address, caller bind.ContractCaller) (*BTCTokenCaller, error) {
	contract, err := bindBTCToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BTCTokenCaller{contract: contract}, nil
}

// NewBTCTokenTransactor creates a new write-only instance of BTCToken, bound to a specific deployed contract.
func NewBTCTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*BTCTokenTransactor, error) {
	contract, err := bindBTCToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BTCTokenTransactor{contract: contract}, nil
}

// NewBTCTokenFilterer creates a new log filterer instance of BTCToken, bound to a specific deployed contract.
func NewBTCTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*BTCTokenFilterer, error) {
	contract, err := bindBTCToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BTCTokenFilterer{contract: contract}, nil
}

// bindBTCToken binds a generic wrapper to an already deployed contract.
func bindBTCToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BTCTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BTCToken *BTCTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BTCToken.Contract.BTCTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BTCToken *BTCTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BTCToken.Contract.BTCTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BTCToken *BTCTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BTCToken.Contract.BTCTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BTCToken *BTCTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BTCToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BTCToken *BTCTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BTCToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BTCToken *BTCTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BTCToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_BTCToken *BTCTokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BTCToken.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_BTCToken *BTCTokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _BTCToken.Contract.Allowance(&_BTCToken.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_BTCToken *BTCTokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _BTCToken.Contract.Allowance(&_BTCToken.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_BTCToken *BTCTokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BTCToken.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_BTCToken *BTCTokenSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _BTCToken.Contract.BalanceOf(&_BTCToken.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_BTCToken *BTCTokenCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _BTCToken.Contract.BalanceOf(&_BTCToken.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_BTCToken *BTCTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _BTCToken.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_BTCToken *BTCTokenSession) Decimals() (uint8, error) {
	return _BTCToken.Contract.Decimals(&_BTCToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_BTCToken *BTCTokenCallerSession) Decimals() (uint8, error) {
	return _BTCToken.Contract.Decimals(&_BTCToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BTCToken *BTCTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BTCToken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BTCToken *BTCTokenSession) Name() (string, error) {
	return _BTCToken.Contract.Name(&_BTCToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BTCToken *BTCTokenCallerSession) Name() (string, error) {
	return _BTCToken.Contract.Name(&_BTCToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BTCToken *BTCTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BTCToken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BTCToken *BTCTokenSession) Symbol() (string, error) {
	return _BTCToken.Contract.Symbol(&_BTCToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BTCToken *BTCTokenCallerSession) Symbol() (string, error) {
	return _BTCToken.Contract.Symbol(&_BTCToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BTCToken *BTCTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BTCToken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BTCToken *BTCTokenSession) TotalSupply() (*big.Int, error) {
	return _BTCToken.Contract.TotalSupply(&_BTCToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BTCToken *BTCTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _BTCToken.Contract.TotalSupply(&_BTCToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_BTCToken *BTCTokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BTCToken.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_BTCToken *BTCTokenSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BTCToken.Contract.Approve(&_BTCToken.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_BTCToken *BTCTokenTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BTCToken.Contract.Approve(&_BTCToken.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address _account, uint256 _amount) returns()
func (_BTCToken *BTCTokenTransactor) Burn(opts *bind.TransactOpts, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _BTCToken.contract.Transact(opts, "burn", _account, _amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address _account, uint256 _amount) returns()
func (_BTCToken *BTCTokenSession) Burn(_account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _BTCToken.Contract.Burn(&_BTCToken.TransactOpts, _account, _amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address _account, uint256 _amount) returns()
func (_BTCToken *BTCTokenTransactorSession) Burn(_account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _BTCToken.Contract.Burn(&_BTCToken.TransactOpts, _account, _amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_BTCToken *BTCTokenTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _BTCToken.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_BTCToken *BTCTokenSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _BTCToken.Contract.DecreaseAllowance(&_BTCToken.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_BTCToken *BTCTokenTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _BTCToken.Contract.DecreaseAllowance(&_BTCToken.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_BTCToken *BTCTokenTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _BTCToken.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_BTCToken *BTCTokenSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _BTCToken.Contract.IncreaseAllowance(&_BTCToken.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_BTCToken *BTCTokenTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _BTCToken.Contract.IncreaseAllowance(&_BTCToken.TransactOpts, spender, addedValue)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _account, uint256 _amount) returns()
func (_BTCToken *BTCTokenTransactor) Mint(opts *bind.TransactOpts, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _BTCToken.contract.Transact(opts, "mint", _account, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _account, uint256 _amount) returns()
func (_BTCToken *BTCTokenSession) Mint(_account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _BTCToken.Contract.Mint(&_BTCToken.TransactOpts, _account, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _account, uint256 _amount) returns()
func (_BTCToken *BTCTokenTransactorSession) Mint(_account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _BTCToken.Contract.Mint(&_BTCToken.TransactOpts, _account, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_BTCToken *BTCTokenTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BTCToken.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_BTCToken *BTCTokenSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BTCToken.Contract.Transfer(&_BTCToken.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_BTCToken *BTCTokenTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BTCToken.Contract.Transfer(&_BTCToken.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_BTCToken *BTCTokenTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BTCToken.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_BTCToken *BTCTokenSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BTCToken.Contract.TransferFrom(&_BTCToken.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_BTCToken *BTCTokenTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BTCToken.Contract.TransferFrom(&_BTCToken.TransactOpts, sender, recipient, amount)
}

// BTCTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the BTCToken contract.
type BTCTokenApprovalIterator struct {
	Event *BTCTokenApproval // Event containing the contract specifics and raw log

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
func (it *BTCTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BTCTokenApproval)
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
		it.Event = new(BTCTokenApproval)
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
func (it *BTCTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BTCTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BTCTokenApproval represents a Approval event raised by the BTCToken contract.
type BTCTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_BTCToken *BTCTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*BTCTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _BTCToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &BTCTokenApprovalIterator{contract: _BTCToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_BTCToken *BTCTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *BTCTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _BTCToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BTCTokenApproval)
				if err := _BTCToken.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_BTCToken *BTCTokenFilterer) ParseApproval(log types.Log) (*BTCTokenApproval, error) {
	event := new(BTCTokenApproval)
	if err := _BTCToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BTCTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the BTCToken contract.
type BTCTokenTransferIterator struct {
	Event *BTCTokenTransfer // Event containing the contract specifics and raw log

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
func (it *BTCTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BTCTokenTransfer)
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
		it.Event = new(BTCTokenTransfer)
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
func (it *BTCTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BTCTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BTCTokenTransfer represents a Transfer event raised by the BTCToken contract.
type BTCTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_BTCToken *BTCTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BTCTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BTCToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BTCTokenTransferIterator{contract: _BTCToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_BTCToken *BTCTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *BTCTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BTCToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BTCTokenTransfer)
				if err := _BTCToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_BTCToken *BTCTokenFilterer) ParseTransfer(log types.Log) (*BTCTokenTransfer, error) {
	event := new(BTCTokenTransfer)
	if err := _BTCToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
