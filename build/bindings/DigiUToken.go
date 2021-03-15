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

// DigiUTokenABI is the input ABI used to generate the binding from.
const DigiUTokenABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"fromDelegate\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"toDelegate\",\"type\":\"address\"}],\"name\":\"DelegateChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"previousBalance\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newBalance\",\"type\":\"uint256\"}],\"name\":\"DelegateVotesChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DELEGATION_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DOMAIN_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"checkpoints\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"fromBlock\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"votes\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"numCheckpoints\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"}],\"name\":\"delegates\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegatee\",\"type\":\"address\"}],\"name\":\"delegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegatee\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"delegateBySig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getCurrentVotes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"getPriorVotes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// DigiUTokenBin is the compiled bytecode used for deploying new contracts.
var DigiUTokenBin = "0x60806040523480156200001157600080fd5b50604080518082018252600a8152692234b3b4aaaa37b5b2b760b11b602080830191825283518085019094526005845264446967695560d81b9084015281519192916200006191600391620000fd565b50805162000077906004906020840190620000fd565b50506005805460ff191660121790555060006200009c6001600160e01b03620000f816565b60058054610100600160a81b0319166101006001600160a01b03841690810291909117909155604051919250906000907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a3506200019f565b335b90565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200014057805160ff191683800117855562000170565b8280016001018555821562000170579182015b828111156200017057825182559160200191906001019062000153565b506200017e92915062000182565b5090565b620000fa91905b808211156200017e576000815560010162000189565b611b9580620001af6000396000f3fe608060405234801561001057600080fd5b50600436106101735760003560e01c8063715018a6116100de578063a9059cbb11610097578063dd62ed3e11610071578063dd62ed3e14610501578063e7a324dc1461052f578063f1127ed814610537578063f2fde38b1461058957610173565b8063a9059cbb14610468578063b4b5ea5714610494578063c3cda520146104ba57610173565b8063715018a6146103d2578063782d6fe1146103da5780637ecebe00146104065780638da5cb5b1461042c57806395d89b4114610434578063a457c2d71461043c57610173565b8063395093511161013057806339509351146102ab57806340c10f19146102d7578063587cde1e146103055780635c19a95c146103475780636fcfff451461036d57806370a08231146103ac57610173565b806306fdde0314610178578063095ea7b3146101f557806318160ddd1461023557806320606b701461024f57806323b872dd14610257578063313ce5671461028d575b600080fd5b6101806105af565b6040805160208082528351818301528351919283929083019185019080838360005b838110156101ba5781810151838201526020016101a2565b50505050905090810190601f1680156101e75780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6102216004803603604081101561020b57600080fd5b506001600160a01b038135169060200135610645565b604080519115158252519081900360200190f35b61023d610663565b60408051918252519081900360200190f35b61023d610669565b6102216004803603606081101561026d57600080fd5b506001600160a01b03813581169160208101359091169060400135610684565b610295610711565b6040805160ff9092168252519081900360200190f35b610221600480360360408110156102c157600080fd5b506001600160a01b03813516906020013561071a565b610303600480360360408110156102ed57600080fd5b506001600160a01b03813516906020013561076e565b005b61032b6004803603602081101561031b57600080fd5b50356001600160a01b0316610810565b604080516001600160a01b039092168252519081900360200190f35b6103036004803603602081101561035d57600080fd5b50356001600160a01b031661082e565b6103936004803603602081101561038357600080fd5b50356001600160a01b031661083b565b6040805163ffffffff9092168252519081900360200190f35b61023d600480360360208110156103c257600080fd5b50356001600160a01b0316610853565b61030361086e565b61023d600480360360408110156103f057600080fd5b506001600160a01b03813516906020013561092d565b61023d6004803603602081101561041c57600080fd5b50356001600160a01b0316610b35565b61032b610b47565b610180610b5b565b6102216004803603604081101561045257600080fd5b506001600160a01b038135169060200135610bbc565b6102216004803603604081101561047e57600080fd5b506001600160a01b038135169060200135610c2a565b61023d600480360360208110156104aa57600080fd5b50356001600160a01b0316610c3e565b610303600480360360c08110156104d057600080fd5b506001600160a01b038135169060208101359060408101359060ff6060820135169060808101359060a00135610ca2565b61023d6004803603604081101561051757600080fd5b506001600160a01b0381358116916020013516610f18565b61023d610f43565b6105696004803603604081101561054d57600080fd5b5080356001600160a01b0316906020013563ffffffff16610f5e565b6040805163ffffffff909316835260208301919091528051918290030190f35b6103036004803603602081101561059f57600080fd5b50356001600160a01b0316610f8b565b60038054604080516020601f600260001961010060018816150201909516949094049384018190048102820181019092528281526060939092909183018282801561063b5780601f106106105761010080835404028352916020019161063b565b820191906000526020600020905b81548152906001019060200180831161061e57829003601f168201915b5050505050905090565b60006106596106526110a6565b84846110aa565b5060015b92915050565b60025490565b6040518060436119db82396043019050604051809103902081565b6000610691848484611196565b6107078461069d6110a6565b61070285604051806060016040528060288152602001611a46602891396001600160a01b038a166000908152600160205260408120906106db6110a6565b6001600160a01b03168152602081019190915260400160002054919063ffffffff6112fd16565b6110aa565b5060019392505050565b60055460ff1690565b60006106596107276110a6565b8461070285600160006107386110a6565b6001600160a01b03908116825260208083019390935260409182016000908120918c16815292529020549063ffffffff61139416565b6107766110a6565b60055461010090046001600160a01b039081169116146107dd576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b6107e782826113ee565b6001600160a01b0380831660009081526006602052604081205461080c9216836114ea565b5050565b6001600160a01b039081166000908152600660205260409020541690565b6108383382611638565b50565b60086020526000908152604090205463ffffffff1681565b6001600160a01b031660009081526020819052604090205490565b6108766110a6565b60055461010090046001600160a01b039081169116146108dd576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b60055460405160009161010090046001600160a01b0316907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a360058054610100600160a81b0319169055565b600043821061096d5760405162461bcd60e51b8152600401808060200182810382526028815260200180611a1e6028913960400191505060405180910390fd5b6001600160a01b03831660009081526008602052604090205463ffffffff168061099b57600091505061065d565b6001600160a01b038416600090815260076020908152604080832063ffffffff600019860181168552925290912054168310610a0a576001600160a01b03841660009081526007602090815260408083206000199490940163ffffffff1683529290522060010154905061065d565b6001600160a01b038416600090815260076020908152604080832083805290915290205463ffffffff16831015610a4557600091505061065d565b600060001982015b8163ffffffff168163ffffffff161115610afe57600282820363ffffffff16048103610a776118d6565b506001600160a01b038716600090815260076020908152604080832063ffffffff808616855290835292819020815180830190925280549093168082526001909301549181019190915290871415610ad95760200151945061065d9350505050565b805163ffffffff16871115610af057819350610af7565b6001820392505b5050610a4d565b506001600160a01b038516600090815260076020908152604080832063ffffffff9094168352929052206001015491505092915050565b60096020526000908152604090205481565b60055461010090046001600160a01b031690565b60048054604080516020601f600260001961010060018816150201909516949094049384018190048102820181019092528281526060939092909183018282801561063b5780601f106106105761010080835404028352916020019161063b565b6000610659610bc96110a6565b8461070285604051806060016040528060258152602001611b3b6025913960016000610bf36110a6565b6001600160a01b03908116825260208083019390935260409182016000908120918d1681529252902054919063ffffffff6112fd16565b6000610659610c376110a6565b8484611196565b6001600160a01b03811660009081526008602052604081205463ffffffff1680610c69576000610c9b565b6001600160a01b038316600090815260076020908152604080832063ffffffff60001986011684529091529020600101545b9392505050565b600060405180806119db6043913960430190506040518091039020610cc56105af565b80519060200120610cd46116cd565b3060405160200180858152602001848152602001838152602001826001600160a01b03166001600160a01b0316815260200194505050505060405160208183030381529060405280519060200120905060006040518080611b01603a91396040805191829003603a0182206020808401919091526001600160a01b038c1683830152606083018b905260808084018b90528251808503909101815260a08401835280519082012061190160f01b60c085015260c2840187905260e2808501829052835180860390910181526101028501808552815191840191909120600091829052610122860180865281905260ff8c1661014287015261016286018b905261018286018a9052935191965092945091926001926101a28083019392601f198301929081900390910190855afa158015610e12573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b038116610e645760405162461bcd60e51b8152600401808060200182810382526027815260200180611a916027913960400191505060405180910390fd5b6001600160a01b03811660009081526009602052604090208054600181019091558914610ec25760405162461bcd60e51b8152600401808060200182810382526023815260200180611a6e6023913960400191505060405180910390fd5b87421115610f015760405162461bcd60e51b81526004018080602001828103825260278152602001806119116027913960400191505060405180910390fd5b610f0b818b611638565b505050505b505050505050565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205490565b60405180603a611b018239603a019050604051809103902081565b60076020908152600092835260408084209091529082529020805460019091015463ffffffff9091169082565b610f936110a6565b60055461010090046001600160a01b03908116911614610ffa576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b6001600160a01b03811661103f5760405162461bcd60e51b81526004018080602001828103825260268152602001806119386026913960400191505060405180910390fd5b6005546040516001600160a01b0380841692610100900416907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a3600580546001600160a01b0390921661010002610100600160a81b0319909216919091179055565b3390565b6001600160a01b0383166110ef5760405162461bcd60e51b8152600401808060200182810382526024815260200180611add6024913960400191505060405180910390fd5b6001600160a01b0382166111345760405162461bcd60e51b815260040180806020018281038252602281526020018061195e6022913960400191505060405180910390fd5b6001600160a01b03808416600081815260016020908152604080832094871680845294825291829020859055815185815291517f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259281900390910190a3505050565b6001600160a01b0383166111db5760405162461bcd60e51b8152600401808060200182810382526025815260200180611ab86025913960400191505060405180910390fd5b6001600160a01b0382166112205760405162461bcd60e51b81526004018080602001828103825260238152602001806118ee6023913960400191505060405180910390fd5b61122b838383611633565b61126e81604051806060016040528060268152602001611980602691396001600160a01b038616600090815260208190526040902054919063ffffffff6112fd16565b6001600160a01b0380851660009081526020819052604080822093909355908416815220546112a3908263ffffffff61139416565b6001600160a01b038084166000818152602081815260409182902094909455805185815290519193928716927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef92918290030190a3505050565b6000818484111561138c5760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b83811015611351578181015183820152602001611339565b50505050905090810190601f16801561137e5780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b505050900390565b600082820183811015610c9b576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b6001600160a01b038216611449576040805162461bcd60e51b815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f206164647265737300604482015290519081900360640190fd5b61145560008383611633565b600254611468908263ffffffff61139416565b6002556001600160a01b038216600090815260208190526040902054611494908263ffffffff61139416565b6001600160a01b0383166000818152602081815260408083209490945583518581529351929391927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9281900390910190a35050565b816001600160a01b0316836001600160a01b03161415801561150c5750600081115b15611633576001600160a01b038316156115a4576001600160a01b03831660009081526008602052604081205463ffffffff16908161154c57600061157e565b6001600160a01b038516600090815260076020908152604080832063ffffffff60001987011684529091529020600101545b90506000611592828563ffffffff6116d116565b90506115a086848484611713565b5050505b6001600160a01b03821615611633576001600160a01b03821660009081526008602052604081205463ffffffff1690816115df576000611611565b6001600160a01b038416600090815260076020908152604080832063ffffffff60001987011684529091529020600101545b90506000611625828563ffffffff61139416565b9050610f1085848484611713565b505050565b6001600160a01b038083166000908152600660205260408120549091169061165f84610853565b6001600160a01b0385811660008181526006602052604080822080546001600160a01b031916898616908117909155905194955093928616927f3134e8a2e6d97e929a7e54011ea5485d7d196dd5f0ba4d4ef95803e8e3fc257f9190a46116c78284836114ea565b50505050565b4690565b6000610c9b83836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f7700008152506112fd565b6000611737436040518060600160405280603581526020016119a660359139611878565b905060008463ffffffff1611801561178057506001600160a01b038516600090815260076020908152604080832063ffffffff6000198901811685529252909120548282169116145b156117bd576001600160a01b038516600090815260076020908152604080832063ffffffff6000198901168452909152902060010182905561182e565b60408051808201825263ffffffff808416825260208083018681526001600160a01b038a166000818152600784528681208b8616825284528681209551865490861663ffffffff19918216178755925160019687015590815260089092529390208054928801909116919092161790555b604080518481526020810184905281516001600160a01b038816927fdec2bacdd2f05b59de34da9b523dff8be42e5e38e818c82fdb0bae774387a724928290030190a25050505050565b60008164010000000084106118ce5760405162461bcd60e51b8152602060048201818152835160248401528351909283926044909101919085019080838360008315611351578181015183820152602001611339565b509192915050565b60408051808201909152600080825260208201529056fe45524332303a207472616e7366657220746f20746865207a65726f206164647265737353555348493a3a64656c656761746542795369673a207369676e617475726520657870697265644f776e61626c653a206e6577206f776e657220697320746865207a65726f206164647265737345524332303a20617070726f766520746f20746865207a65726f206164647265737345524332303a207472616e7366657220616d6f756e7420657863656564732062616c616e636553555348493a3a5f7772697465436865636b706f696e743a20626c6f636b206e756d62657220657863656564732033322062697473454950373132446f6d61696e28737472696e67206e616d652c75696e7432353620636861696e49642c6164647265737320766572696679696e67436f6e74726163742953555348493a3a6765745072696f72566f7465733a206e6f74207965742064657465726d696e656445524332303a207472616e7366657220616d6f756e74206578636565647320616c6c6f77616e636553555348493a3a64656c656761746542795369673a20696e76616c6964206e6f6e636553555348493a3a64656c656761746542795369673a20696e76616c6964207369676e617475726545524332303a207472616e736665722066726f6d20746865207a65726f206164647265737345524332303a20617070726f76652066726f6d20746865207a65726f206164647265737344656c65676174696f6e28616464726573732064656c6567617465652c75696e74323536206e6f6e63652c75696e74323536206578706972792945524332303a2064656372656173656420616c6c6f77616e63652062656c6f77207a65726fa2646970667358221220fa33adcd101f3dc7f0fff6c50c5b0f4d96fc57ab13202fdba64124d80aba065764736f6c63430006090033"

// DeployDigiUToken deploys a new Ethereum contract, binding an instance of DigiUToken to it.
func DeployDigiUToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DigiUToken, error) {
	parsed, err := abi.JSON(strings.NewReader(DigiUTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DigiUTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DigiUToken{DigiUTokenCaller: DigiUTokenCaller{contract: contract}, DigiUTokenTransactor: DigiUTokenTransactor{contract: contract}, DigiUTokenFilterer: DigiUTokenFilterer{contract: contract}}, nil
}

// DeployDigiUTokenSync deploys a new Ethereum contract and waits for receipt, binding an instance of DigiUTokenSession to it.
func DeployDigiUTokenSync(session *bind.TransactSession, backend bind.ContractBackend) (*types.Transaction, *types.Receipt, *DigiUTokenSession, error) {
	parsed, err := abi.JSON(strings.NewReader(DigiUTokenABI))
	if err != nil {
		return nil, nil, nil, err
	}
	session.Lock()
	address, tx, _, err := bind.DeployContract(session.TransactOpts, parsed, common.FromHex(DigiUTokenBin), backend)
	receipt, err := session.WaitTransaction(tx)
	if err != nil {
		session.Unlock()
		return nil, nil, nil, err
	}
	session.TransactOpts.Nonce.Add(session.TransactOpts.Nonce, big.NewInt(1))
	session.Unlock()
	contractSession, err := NewDigiUTokenSession(address, backend, session)
	return tx, receipt, contractSession, err
}

// DigiUToken is an auto generated Go binding around an Ethereum contract.
type DigiUToken struct {
	DigiUTokenCaller     // Read-only binding to the contract
	DigiUTokenTransactor // Write-only binding to the contract
	DigiUTokenFilterer   // Log filterer for contract events
}

// DigiUTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type DigiUTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DigiUTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DigiUTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DigiUTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DigiUTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DigiUTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DigiUTokenSession struct {
	Contract           *DigiUToken // Generic contract binding to set the session for
	transactionSession *bind.TransactSession
	Address            common.Address
}

// DigiUTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DigiUTokenCallerSession struct {
	Contract *DigiUTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// DigiUTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DigiUTokenTransactorSession struct {
	Contract     *DigiUTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// DigiUTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type DigiUTokenRaw struct {
	Contract *DigiUToken // Generic contract binding to access the raw methods on
}

// DigiUTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DigiUTokenCallerRaw struct {
	Contract *DigiUTokenCaller // Generic read-only contract binding to access the raw methods on
}

// DigiUTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DigiUTokenTransactorRaw struct {
	Contract *DigiUTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDigiUToken creates a new instance of DigiUToken, bound to a specific deployed contract.
func NewDigiUToken(address common.Address, backend bind.ContractBackend) (*DigiUToken, error) {
	contract, err := bindDigiUToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DigiUToken{DigiUTokenCaller: DigiUTokenCaller{contract: contract}, DigiUTokenTransactor: DigiUTokenTransactor{contract: contract}, DigiUTokenFilterer: DigiUTokenFilterer{contract: contract}}, nil
}

// NewDigiUTokenCaller creates a new read-only instance of DigiUToken, bound to a specific deployed contract.
func NewDigiUTokenCaller(address common.Address, caller bind.ContractCaller) (*DigiUTokenCaller, error) {
	contract, err := bindDigiUToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DigiUTokenCaller{contract: contract}, nil
}

// NewDigiUTokenTransactor creates a new write-only instance of DigiUToken, bound to a specific deployed contract.
func NewDigiUTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*DigiUTokenTransactor, error) {
	contract, err := bindDigiUToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DigiUTokenTransactor{contract: contract}, nil
}

// NewDigiUTokenFilterer creates a new log filterer instance of DigiUToken, bound to a specific deployed contract.
func NewDigiUTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*DigiUTokenFilterer, error) {
	contract, err := bindDigiUToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DigiUTokenFilterer{contract: contract}, nil
}

func NewDigiUTokenSession(address common.Address, backend bind.ContractBackend, transactionSession *bind.TransactSession) (*DigiUTokenSession, error) {
	DigiUTokenInstance, err := NewDigiUToken(address, backend)
	if err != nil {
		return nil, err
	}
	return &DigiUTokenSession{
		Contract:           DigiUTokenInstance,
		transactionSession: transactionSession,
		Address:            address,
	}, nil
}

// bindDigiUToken binds a generic wrapper to an already deployed contract.
func bindDigiUToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DigiUTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DigiUToken *DigiUTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DigiUToken.Contract.DigiUTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DigiUToken *DigiUTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DigiUToken.Contract.DigiUTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DigiUToken *DigiUTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DigiUToken.Contract.DigiUTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DigiUToken *DigiUTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DigiUToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DigiUToken *DigiUTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DigiUToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DigiUToken *DigiUTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DigiUToken.Contract.contract.Transact(opts, method, params...)
}

// DELEGATIONTYPEHASH is a free data retrieval call binding the contract method 0xe7a324dc.
//
// Solidity: function DELEGATION_TYPEHASH() view returns(bytes32)
func (_DigiUToken *DigiUTokenCaller) DELEGATIONTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _DigiUToken.contract.Call(opts, out, "DELEGATION_TYPEHASH")
	return *ret0, err
}

// DELEGATIONTYPEHASH is a free data retrieval call binding the contract method 0xe7a324dc.
//
// Solidity: function DELEGATION_TYPEHASH() view returns(bytes32)
func (_DigiUToken *DigiUTokenSession) DELEGATIONTYPEHASH() ([32]byte, error) {
	return _DigiUToken.Contract.DELEGATIONTYPEHASH(_DigiUToken.transactionSession.CallOpts)
}

// DELEGATIONTYPEHASH is a free data retrieval call binding the contract method 0xe7a324dc.
//
// Solidity: function DELEGATION_TYPEHASH() view returns(bytes32)
func (_DigiUToken *DigiUTokenCallerSession) DELEGATIONTYPEHASH() ([32]byte, error) {
	return _DigiUToken.Contract.DELEGATIONTYPEHASH(&_DigiUToken.CallOpts)
}

// DOMAINTYPEHASH is a free data retrieval call binding the contract method 0x20606b70.
//
// Solidity: function DOMAIN_TYPEHASH() view returns(bytes32)
func (_DigiUToken *DigiUTokenCaller) DOMAINTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _DigiUToken.contract.Call(opts, out, "DOMAIN_TYPEHASH")
	return *ret0, err
}

// DOMAINTYPEHASH is a free data retrieval call binding the contract method 0x20606b70.
//
// Solidity: function DOMAIN_TYPEHASH() view returns(bytes32)
func (_DigiUToken *DigiUTokenSession) DOMAINTYPEHASH() ([32]byte, error) {
	return _DigiUToken.Contract.DOMAINTYPEHASH(_DigiUToken.transactionSession.CallOpts)
}

// DOMAINTYPEHASH is a free data retrieval call binding the contract method 0x20606b70.
//
// Solidity: function DOMAIN_TYPEHASH() view returns(bytes32)
func (_DigiUToken *DigiUTokenCallerSession) DOMAINTYPEHASH() ([32]byte, error) {
	return _DigiUToken.Contract.DOMAINTYPEHASH(&_DigiUToken.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_DigiUToken *DigiUTokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DigiUToken.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_DigiUToken *DigiUTokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _DigiUToken.Contract.Allowance(_DigiUToken.transactionSession.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_DigiUToken *DigiUTokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _DigiUToken.Contract.Allowance(&_DigiUToken.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_DigiUToken *DigiUTokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DigiUToken.contract.Call(opts, out, "balanceOf", account)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_DigiUToken *DigiUTokenSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _DigiUToken.Contract.BalanceOf(_DigiUToken.transactionSession.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_DigiUToken *DigiUTokenCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _DigiUToken.Contract.BalanceOf(&_DigiUToken.CallOpts, account)
}

// Checkpoints is a free data retrieval call binding the contract method 0xf1127ed8.
//
// Solidity: function checkpoints(address , uint32 ) view returns(uint32 fromBlock, uint256 votes)
func (_DigiUToken *DigiUTokenCaller) Checkpoints(opts *bind.CallOpts, arg0 common.Address, arg1 uint32) (struct {
	FromBlock uint32
	Votes     *big.Int
}, error) {
	ret := new(struct {
		FromBlock uint32
		Votes     *big.Int
	})
	out := ret
	err := _DigiUToken.contract.Call(opts, out, "checkpoints", arg0, arg1)
	return *ret, err
}

// Checkpoints is a free data retrieval call binding the contract method 0xf1127ed8.
//
// Solidity: function checkpoints(address , uint32 ) view returns(uint32 fromBlock, uint256 votes)
func (_DigiUToken *DigiUTokenSession) Checkpoints(arg0 common.Address, arg1 uint32) (struct {
	FromBlock uint32
	Votes     *big.Int
}, error) {
	return _DigiUToken.Contract.Checkpoints(_DigiUToken.transactionSession.CallOpts, arg0, arg1)
}

// Checkpoints is a free data retrieval call binding the contract method 0xf1127ed8.
//
// Solidity: function checkpoints(address , uint32 ) view returns(uint32 fromBlock, uint256 votes)
func (_DigiUToken *DigiUTokenCallerSession) Checkpoints(arg0 common.Address, arg1 uint32) (struct {
	FromBlock uint32
	Votes     *big.Int
}, error) {
	return _DigiUToken.Contract.Checkpoints(&_DigiUToken.CallOpts, arg0, arg1)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_DigiUToken *DigiUTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _DigiUToken.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_DigiUToken *DigiUTokenSession) Decimals() (uint8, error) {
	return _DigiUToken.Contract.Decimals(_DigiUToken.transactionSession.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_DigiUToken *DigiUTokenCallerSession) Decimals() (uint8, error) {
	return _DigiUToken.Contract.Decimals(&_DigiUToken.CallOpts)
}

// Delegates is a free data retrieval call binding the contract method 0x587cde1e.
//
// Solidity: function delegates(address delegator) view returns(address)
func (_DigiUToken *DigiUTokenCaller) Delegates(opts *bind.CallOpts, delegator common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DigiUToken.contract.Call(opts, out, "delegates", delegator)
	return *ret0, err
}

// Delegates is a free data retrieval call binding the contract method 0x587cde1e.
//
// Solidity: function delegates(address delegator) view returns(address)
func (_DigiUToken *DigiUTokenSession) Delegates(delegator common.Address) (common.Address, error) {
	return _DigiUToken.Contract.Delegates(_DigiUToken.transactionSession.CallOpts, delegator)
}

// Delegates is a free data retrieval call binding the contract method 0x587cde1e.
//
// Solidity: function delegates(address delegator) view returns(address)
func (_DigiUToken *DigiUTokenCallerSession) Delegates(delegator common.Address) (common.Address, error) {
	return _DigiUToken.Contract.Delegates(&_DigiUToken.CallOpts, delegator)
}

// GetCurrentVotes is a free data retrieval call binding the contract method 0xb4b5ea57.
//
// Solidity: function getCurrentVotes(address account) view returns(uint256)
func (_DigiUToken *DigiUTokenCaller) GetCurrentVotes(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DigiUToken.contract.Call(opts, out, "getCurrentVotes", account)
	return *ret0, err
}

// GetCurrentVotes is a free data retrieval call binding the contract method 0xb4b5ea57.
//
// Solidity: function getCurrentVotes(address account) view returns(uint256)
func (_DigiUToken *DigiUTokenSession) GetCurrentVotes(account common.Address) (*big.Int, error) {
	return _DigiUToken.Contract.GetCurrentVotes(_DigiUToken.transactionSession.CallOpts, account)
}

// GetCurrentVotes is a free data retrieval call binding the contract method 0xb4b5ea57.
//
// Solidity: function getCurrentVotes(address account) view returns(uint256)
func (_DigiUToken *DigiUTokenCallerSession) GetCurrentVotes(account common.Address) (*big.Int, error) {
	return _DigiUToken.Contract.GetCurrentVotes(&_DigiUToken.CallOpts, account)
}

// GetPriorVotes is a free data retrieval call binding the contract method 0x782d6fe1.
//
// Solidity: function getPriorVotes(address account, uint256 blockNumber) view returns(uint256)
func (_DigiUToken *DigiUTokenCaller) GetPriorVotes(opts *bind.CallOpts, account common.Address, blockNumber *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DigiUToken.contract.Call(opts, out, "getPriorVotes", account, blockNumber)
	return *ret0, err
}

// GetPriorVotes is a free data retrieval call binding the contract method 0x782d6fe1.
//
// Solidity: function getPriorVotes(address account, uint256 blockNumber) view returns(uint256)
func (_DigiUToken *DigiUTokenSession) GetPriorVotes(account common.Address, blockNumber *big.Int) (*big.Int, error) {
	return _DigiUToken.Contract.GetPriorVotes(_DigiUToken.transactionSession.CallOpts, account, blockNumber)
}

// GetPriorVotes is a free data retrieval call binding the contract method 0x782d6fe1.
//
// Solidity: function getPriorVotes(address account, uint256 blockNumber) view returns(uint256)
func (_DigiUToken *DigiUTokenCallerSession) GetPriorVotes(account common.Address, blockNumber *big.Int) (*big.Int, error) {
	return _DigiUToken.Contract.GetPriorVotes(&_DigiUToken.CallOpts, account, blockNumber)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_DigiUToken *DigiUTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _DigiUToken.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_DigiUToken *DigiUTokenSession) Name() (string, error) {
	return _DigiUToken.Contract.Name(_DigiUToken.transactionSession.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_DigiUToken *DigiUTokenCallerSession) Name() (string, error) {
	return _DigiUToken.Contract.Name(&_DigiUToken.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_DigiUToken *DigiUTokenCaller) Nonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DigiUToken.contract.Call(opts, out, "nonces", arg0)
	return *ret0, err
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_DigiUToken *DigiUTokenSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _DigiUToken.Contract.Nonces(_DigiUToken.transactionSession.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_DigiUToken *DigiUTokenCallerSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _DigiUToken.Contract.Nonces(&_DigiUToken.CallOpts, arg0)
}

// NumCheckpoints is a free data retrieval call binding the contract method 0x6fcfff45.
//
// Solidity: function numCheckpoints(address ) view returns(uint32)
func (_DigiUToken *DigiUTokenCaller) NumCheckpoints(opts *bind.CallOpts, arg0 common.Address) (uint32, error) {
	var (
		ret0 = new(uint32)
	)
	out := ret0
	err := _DigiUToken.contract.Call(opts, out, "numCheckpoints", arg0)
	return *ret0, err
}

// NumCheckpoints is a free data retrieval call binding the contract method 0x6fcfff45.
//
// Solidity: function numCheckpoints(address ) view returns(uint32)
func (_DigiUToken *DigiUTokenSession) NumCheckpoints(arg0 common.Address) (uint32, error) {
	return _DigiUToken.Contract.NumCheckpoints(_DigiUToken.transactionSession.CallOpts, arg0)
}

// NumCheckpoints is a free data retrieval call binding the contract method 0x6fcfff45.
//
// Solidity: function numCheckpoints(address ) view returns(uint32)
func (_DigiUToken *DigiUTokenCallerSession) NumCheckpoints(arg0 common.Address) (uint32, error) {
	return _DigiUToken.Contract.NumCheckpoints(&_DigiUToken.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DigiUToken *DigiUTokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DigiUToken.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DigiUToken *DigiUTokenSession) Owner() (common.Address, error) {
	return _DigiUToken.Contract.Owner(_DigiUToken.transactionSession.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DigiUToken *DigiUTokenCallerSession) Owner() (common.Address, error) {
	return _DigiUToken.Contract.Owner(&_DigiUToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_DigiUToken *DigiUTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _DigiUToken.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_DigiUToken *DigiUTokenSession) Symbol() (string, error) {
	return _DigiUToken.Contract.Symbol(_DigiUToken.transactionSession.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_DigiUToken *DigiUTokenCallerSession) Symbol() (string, error) {
	return _DigiUToken.Contract.Symbol(&_DigiUToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_DigiUToken *DigiUTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DigiUToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_DigiUToken *DigiUTokenSession) TotalSupply() (*big.Int, error) {
	return _DigiUToken.Contract.TotalSupply(_DigiUToken.transactionSession.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_DigiUToken *DigiUTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _DigiUToken.Contract.TotalSupply(&_DigiUToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_DigiUToken *DigiUTokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DigiUToken.contract.Transact(opts, "approve", spender, amount)
}

func (_DigiUToken *DigiUTokenTransactor) ApproveRawTx(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DigiUToken.contract.RawTx(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
// Will wait for tx receipt
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_DigiUToken *DigiUTokenSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, *types.Receipt, error) {
	_DigiUToken.transactionSession.Lock()
	tx, err := _DigiUToken.Contract.Approve(_DigiUToken.transactionSession.TransactOpts, spender, amount)
	if err != nil {
		_DigiUToken.transactionSession.Unlock()
		return nil, nil, err
	}
	_DigiUToken.transactionSession.TransactOpts.Nonce.Add(_DigiUToken.transactionSession.TransactOpts.Nonce, big.NewInt(1))
	_DigiUToken.transactionSession.Unlock()
	receipt, err := _DigiUToken.transactionSession.WaitTransaction(tx)
	return tx, receipt, err
}

// Approve returns raw transaction bound to the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_DigiUToken *DigiUTokenSession) ApproveRawTx(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	tx, err := _DigiUToken.Contract.ApproveRawTx(_DigiUToken.transactionSession.TransactOpts, spender, amount)
	return tx, err
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
// Will not wait for tx, but put it to ch
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_DigiUToken *DigiUTokenSession) ApproveAsync(receiptCh chan *types.ReceiptResult, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	_DigiUToken.transactionSession.Lock()
	tx, err := _DigiUToken.Contract.Approve(_DigiUToken.transactionSession.TransactOpts, spender, amount)
	if err != nil {
		_DigiUToken.transactionSession.Unlock()
		return nil, err
	}
	_DigiUToken.transactionSession.TransactOpts.Nonce.Add(_DigiUToken.transactionSession.TransactOpts.Nonce, big.NewInt(1))
	_DigiUToken.transactionSession.Unlock()
	go func() {
		receipt, err := _DigiUToken.transactionSession.WaitTransaction(tx)
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
func (_DigiUToken *DigiUTokenTransactorSession) Approve(wait bool, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DigiUToken.Contract.Approve(&_DigiUToken.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_DigiUToken *DigiUTokenTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _DigiUToken.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

func (_DigiUToken *DigiUTokenTransactor) DecreaseAllowanceRawTx(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _DigiUToken.contract.RawTx(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
// Will wait for tx receipt
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_DigiUToken *DigiUTokenSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, *types.Receipt, error) {
	_DigiUToken.transactionSession.Lock()
	tx, err := _DigiUToken.Contract.DecreaseAllowance(_DigiUToken.transactionSession.TransactOpts, spender, subtractedValue)
	if err != nil {
		_DigiUToken.transactionSession.Unlock()
		return nil, nil, err
	}
	_DigiUToken.transactionSession.TransactOpts.Nonce.Add(_DigiUToken.transactionSession.TransactOpts.Nonce, big.NewInt(1))
	_DigiUToken.transactionSession.Unlock()
	receipt, err := _DigiUToken.transactionSession.WaitTransaction(tx)
	return tx, receipt, err
}

// DecreaseAllowance returns raw transaction bound to the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_DigiUToken *DigiUTokenSession) DecreaseAllowanceRawTx(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	tx, err := _DigiUToken.Contract.DecreaseAllowanceRawTx(_DigiUToken.transactionSession.TransactOpts, spender, subtractedValue)
	return tx, err
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
// Will not wait for tx, but put it to ch
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_DigiUToken *DigiUTokenSession) DecreaseAllowanceAsync(receiptCh chan *types.ReceiptResult, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	_DigiUToken.transactionSession.Lock()
	tx, err := _DigiUToken.Contract.DecreaseAllowance(_DigiUToken.transactionSession.TransactOpts, spender, subtractedValue)
	if err != nil {
		_DigiUToken.transactionSession.Unlock()
		return nil, err
	}
	_DigiUToken.transactionSession.TransactOpts.Nonce.Add(_DigiUToken.transactionSession.TransactOpts.Nonce, big.NewInt(1))
	_DigiUToken.transactionSession.Unlock()
	go func() {
		receipt, err := _DigiUToken.transactionSession.WaitTransaction(tx)
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
func (_DigiUToken *DigiUTokenTransactorSession) DecreaseAllowance(wait bool, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _DigiUToken.Contract.DecreaseAllowance(&_DigiUToken.TransactOpts, spender, subtractedValue)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address delegatee) returns()
func (_DigiUToken *DigiUTokenTransactor) Delegate(opts *bind.TransactOpts, delegatee common.Address) (*types.Transaction, error) {
	return _DigiUToken.contract.Transact(opts, "delegate", delegatee)
}

func (_DigiUToken *DigiUTokenTransactor) DelegateRawTx(opts *bind.TransactOpts, delegatee common.Address) (*types.Transaction, error) {
	return _DigiUToken.contract.RawTx(opts, "delegate", delegatee)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
// Will wait for tx receipt
//
// Solidity: function delegate(address delegatee) returns()
func (_DigiUToken *DigiUTokenSession) Delegate(delegatee common.Address) (*types.Transaction, *types.Receipt, error) {
	_DigiUToken.transactionSession.Lock()
	tx, err := _DigiUToken.Contract.Delegate(_DigiUToken.transactionSession.TransactOpts, delegatee)
	if err != nil {
		_DigiUToken.transactionSession.Unlock()
		return nil, nil, err
	}
	_DigiUToken.transactionSession.TransactOpts.Nonce.Add(_DigiUToken.transactionSession.TransactOpts.Nonce, big.NewInt(1))
	_DigiUToken.transactionSession.Unlock()
	receipt, err := _DigiUToken.transactionSession.WaitTransaction(tx)
	return tx, receipt, err
}

// Delegate returns raw transaction bound to the contract method 0x5c19a95c.
//
// Solidity: function delegate(address delegatee) returns()
func (_DigiUToken *DigiUTokenSession) DelegateRawTx(delegatee common.Address) (*types.Transaction, error) {
	tx, err := _DigiUToken.Contract.DelegateRawTx(_DigiUToken.transactionSession.TransactOpts, delegatee)
	return tx, err
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
// Will not wait for tx, but put it to ch
//
// Solidity: function delegate(address delegatee) returns()
func (_DigiUToken *DigiUTokenSession) DelegateAsync(receiptCh chan *types.ReceiptResult, delegatee common.Address) (*types.Transaction, error) {
	_DigiUToken.transactionSession.Lock()
	tx, err := _DigiUToken.Contract.Delegate(_DigiUToken.transactionSession.TransactOpts, delegatee)
	if err != nil {
		_DigiUToken.transactionSession.Unlock()
		return nil, err
	}
	_DigiUToken.transactionSession.TransactOpts.Nonce.Add(_DigiUToken.transactionSession.TransactOpts.Nonce, big.NewInt(1))
	_DigiUToken.transactionSession.Unlock()
	go func() {
		receipt, err := _DigiUToken.transactionSession.WaitTransaction(tx)
		receiptCh <- &types.ReceiptResult{
			Receipt: *receipt,
			Err:     err,
		}
	}()
	return tx, err
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address delegatee) returns()
func (_DigiUToken *DigiUTokenTransactorSession) Delegate(wait bool, delegatee common.Address) (*types.Transaction, error) {
	return _DigiUToken.Contract.Delegate(&_DigiUToken.TransactOpts, delegatee)
}

// DelegateBySig is a paid mutator transaction binding the contract method 0xc3cda520.
//
// Solidity: function delegateBySig(address delegatee, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) returns()
func (_DigiUToken *DigiUTokenTransactor) DelegateBySig(opts *bind.TransactOpts, delegatee common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _DigiUToken.contract.Transact(opts, "delegateBySig", delegatee, nonce, expiry, v, r, s)
}

func (_DigiUToken *DigiUTokenTransactor) DelegateBySigRawTx(opts *bind.TransactOpts, delegatee common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _DigiUToken.contract.RawTx(opts, "delegateBySig", delegatee, nonce, expiry, v, r, s)
}

// DelegateBySig is a paid mutator transaction binding the contract method 0xc3cda520.
// Will wait for tx receipt
//
// Solidity: function delegateBySig(address delegatee, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) returns()
func (_DigiUToken *DigiUTokenSession) DelegateBySig(delegatee common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, *types.Receipt, error) {
	_DigiUToken.transactionSession.Lock()
	tx, err := _DigiUToken.Contract.DelegateBySig(_DigiUToken.transactionSession.TransactOpts, delegatee, nonce, expiry, v, r, s)
	if err != nil {
		_DigiUToken.transactionSession.Unlock()
		return nil, nil, err
	}
	_DigiUToken.transactionSession.TransactOpts.Nonce.Add(_DigiUToken.transactionSession.TransactOpts.Nonce, big.NewInt(1))
	_DigiUToken.transactionSession.Unlock()
	receipt, err := _DigiUToken.transactionSession.WaitTransaction(tx)
	return tx, receipt, err
}

// DelegateBySig returns raw transaction bound to the contract method 0xc3cda520.
//
// Solidity: function delegateBySig(address delegatee, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) returns()
func (_DigiUToken *DigiUTokenSession) DelegateBySigRawTx(delegatee common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	tx, err := _DigiUToken.Contract.DelegateBySigRawTx(_DigiUToken.transactionSession.TransactOpts, delegatee, nonce, expiry, v, r, s)
	return tx, err
}

// DelegateBySig is a paid mutator transaction binding the contract method 0xc3cda520.
// Will not wait for tx, but put it to ch
//
// Solidity: function delegateBySig(address delegatee, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) returns()
func (_DigiUToken *DigiUTokenSession) DelegateBySigAsync(receiptCh chan *types.ReceiptResult, delegatee common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	_DigiUToken.transactionSession.Lock()
	tx, err := _DigiUToken.Contract.DelegateBySig(_DigiUToken.transactionSession.TransactOpts, delegatee, nonce, expiry, v, r, s)
	if err != nil {
		_DigiUToken.transactionSession.Unlock()
		return nil, err
	}
	_DigiUToken.transactionSession.TransactOpts.Nonce.Add(_DigiUToken.transactionSession.TransactOpts.Nonce, big.NewInt(1))
	_DigiUToken.transactionSession.Unlock()
	go func() {
		receipt, err := _DigiUToken.transactionSession.WaitTransaction(tx)
		receiptCh <- &types.ReceiptResult{
			Receipt: *receipt,
			Err:     err,
		}
	}()
	return tx, err
}

// DelegateBySig is a paid mutator transaction binding the contract method 0xc3cda520.
//
// Solidity: function delegateBySig(address delegatee, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) returns()
func (_DigiUToken *DigiUTokenTransactorSession) DelegateBySig(wait bool, delegatee common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _DigiUToken.Contract.DelegateBySig(&_DigiUToken.TransactOpts, delegatee, nonce, expiry, v, r, s)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_DigiUToken *DigiUTokenTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _DigiUToken.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

func (_DigiUToken *DigiUTokenTransactor) IncreaseAllowanceRawTx(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _DigiUToken.contract.RawTx(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
// Will wait for tx receipt
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_DigiUToken *DigiUTokenSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, *types.Receipt, error) {
	_DigiUToken.transactionSession.Lock()
	tx, err := _DigiUToken.Contract.IncreaseAllowance(_DigiUToken.transactionSession.TransactOpts, spender, addedValue)
	if err != nil {
		_DigiUToken.transactionSession.Unlock()
		return nil, nil, err
	}
	_DigiUToken.transactionSession.TransactOpts.Nonce.Add(_DigiUToken.transactionSession.TransactOpts.Nonce, big.NewInt(1))
	_DigiUToken.transactionSession.Unlock()
	receipt, err := _DigiUToken.transactionSession.WaitTransaction(tx)
	return tx, receipt, err
}

// IncreaseAllowance returns raw transaction bound to the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_DigiUToken *DigiUTokenSession) IncreaseAllowanceRawTx(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	tx, err := _DigiUToken.Contract.IncreaseAllowanceRawTx(_DigiUToken.transactionSession.TransactOpts, spender, addedValue)
	return tx, err
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
// Will not wait for tx, but put it to ch
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_DigiUToken *DigiUTokenSession) IncreaseAllowanceAsync(receiptCh chan *types.ReceiptResult, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	_DigiUToken.transactionSession.Lock()
	tx, err := _DigiUToken.Contract.IncreaseAllowance(_DigiUToken.transactionSession.TransactOpts, spender, addedValue)
	if err != nil {
		_DigiUToken.transactionSession.Unlock()
		return nil, err
	}
	_DigiUToken.transactionSession.TransactOpts.Nonce.Add(_DigiUToken.transactionSession.TransactOpts.Nonce, big.NewInt(1))
	_DigiUToken.transactionSession.Unlock()
	go func() {
		receipt, err := _DigiUToken.transactionSession.WaitTransaction(tx)
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
func (_DigiUToken *DigiUTokenTransactorSession) IncreaseAllowance(wait bool, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _DigiUToken.Contract.IncreaseAllowance(&_DigiUToken.TransactOpts, spender, addedValue)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _amount) returns()
func (_DigiUToken *DigiUTokenTransactor) Mint(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _DigiUToken.contract.Transact(opts, "mint", _to, _amount)
}

func (_DigiUToken *DigiUTokenTransactor) MintRawTx(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _DigiUToken.contract.RawTx(opts, "mint", _to, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
// Will wait for tx receipt
//
// Solidity: function mint(address _to, uint256 _amount) returns()
func (_DigiUToken *DigiUTokenSession) Mint(_to common.Address, _amount *big.Int) (*types.Transaction, *types.Receipt, error) {
	_DigiUToken.transactionSession.Lock()
	tx, err := _DigiUToken.Contract.Mint(_DigiUToken.transactionSession.TransactOpts, _to, _amount)
	if err != nil {
		_DigiUToken.transactionSession.Unlock()
		return nil, nil, err
	}
	_DigiUToken.transactionSession.TransactOpts.Nonce.Add(_DigiUToken.transactionSession.TransactOpts.Nonce, big.NewInt(1))
	_DigiUToken.transactionSession.Unlock()
	receipt, err := _DigiUToken.transactionSession.WaitTransaction(tx)
	return tx, receipt, err
}

// Mint returns raw transaction bound to the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _amount) returns()
func (_DigiUToken *DigiUTokenSession) MintRawTx(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	tx, err := _DigiUToken.Contract.MintRawTx(_DigiUToken.transactionSession.TransactOpts, _to, _amount)
	return tx, err
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
// Will not wait for tx, but put it to ch
//
// Solidity: function mint(address _to, uint256 _amount) returns()
func (_DigiUToken *DigiUTokenSession) MintAsync(receiptCh chan *types.ReceiptResult, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	_DigiUToken.transactionSession.Lock()
	tx, err := _DigiUToken.Contract.Mint(_DigiUToken.transactionSession.TransactOpts, _to, _amount)
	if err != nil {
		_DigiUToken.transactionSession.Unlock()
		return nil, err
	}
	_DigiUToken.transactionSession.TransactOpts.Nonce.Add(_DigiUToken.transactionSession.TransactOpts.Nonce, big.NewInt(1))
	_DigiUToken.transactionSession.Unlock()
	go func() {
		receipt, err := _DigiUToken.transactionSession.WaitTransaction(tx)
		receiptCh <- &types.ReceiptResult{
			Receipt: *receipt,
			Err:     err,
		}
	}()
	return tx, err
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _amount) returns()
func (_DigiUToken *DigiUTokenTransactorSession) Mint(wait bool, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _DigiUToken.Contract.Mint(&_DigiUToken.TransactOpts, _to, _amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DigiUToken *DigiUTokenTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DigiUToken.contract.Transact(opts, "renounceOwnership")
}

func (_DigiUToken *DigiUTokenTransactor) RenounceOwnershipRawTx(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DigiUToken.contract.RawTx(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
// Will wait for tx receipt
//
// Solidity: function renounceOwnership() returns()
func (_DigiUToken *DigiUTokenSession) RenounceOwnership() (*types.Transaction, *types.Receipt, error) {
	_DigiUToken.transactionSession.Lock()
	tx, err := _DigiUToken.Contract.RenounceOwnership(_DigiUToken.transactionSession.TransactOpts)
	if err != nil {
		_DigiUToken.transactionSession.Unlock()
		return nil, nil, err
	}
	_DigiUToken.transactionSession.TransactOpts.Nonce.Add(_DigiUToken.transactionSession.TransactOpts.Nonce, big.NewInt(1))
	_DigiUToken.transactionSession.Unlock()
	receipt, err := _DigiUToken.transactionSession.WaitTransaction(tx)
	return tx, receipt, err
}

// RenounceOwnership returns raw transaction bound to the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DigiUToken *DigiUTokenSession) RenounceOwnershipRawTx() (*types.Transaction, error) {
	tx, err := _DigiUToken.Contract.RenounceOwnershipRawTx(_DigiUToken.transactionSession.TransactOpts)
	return tx, err
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
// Will not wait for tx, but put it to ch
//
// Solidity: function renounceOwnership() returns()
func (_DigiUToken *DigiUTokenSession) RenounceOwnershipAsync(receiptCh chan *types.ReceiptResult) (*types.Transaction, error) {
	_DigiUToken.transactionSession.Lock()
	tx, err := _DigiUToken.Contract.RenounceOwnership(_DigiUToken.transactionSession.TransactOpts)
	if err != nil {
		_DigiUToken.transactionSession.Unlock()
		return nil, err
	}
	_DigiUToken.transactionSession.TransactOpts.Nonce.Add(_DigiUToken.transactionSession.TransactOpts.Nonce, big.NewInt(1))
	_DigiUToken.transactionSession.Unlock()
	go func() {
		receipt, err := _DigiUToken.transactionSession.WaitTransaction(tx)
		receiptCh <- &types.ReceiptResult{
			Receipt: *receipt,
			Err:     err,
		}
	}()
	return tx, err
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DigiUToken *DigiUTokenTransactorSession) RenounceOwnership(wait bool) (*types.Transaction, error) {
	return _DigiUToken.Contract.RenounceOwnership(&_DigiUToken.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_DigiUToken *DigiUTokenTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DigiUToken.contract.Transact(opts, "transfer", recipient, amount)
}

func (_DigiUToken *DigiUTokenTransactor) TransferRawTx(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DigiUToken.contract.RawTx(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
// Will wait for tx receipt
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_DigiUToken *DigiUTokenSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, *types.Receipt, error) {
	_DigiUToken.transactionSession.Lock()
	tx, err := _DigiUToken.Contract.Transfer(_DigiUToken.transactionSession.TransactOpts, recipient, amount)
	if err != nil {
		_DigiUToken.transactionSession.Unlock()
		return nil, nil, err
	}
	_DigiUToken.transactionSession.TransactOpts.Nonce.Add(_DigiUToken.transactionSession.TransactOpts.Nonce, big.NewInt(1))
	_DigiUToken.transactionSession.Unlock()
	receipt, err := _DigiUToken.transactionSession.WaitTransaction(tx)
	return tx, receipt, err
}

// Transfer returns raw transaction bound to the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_DigiUToken *DigiUTokenSession) TransferRawTx(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	tx, err := _DigiUToken.Contract.TransferRawTx(_DigiUToken.transactionSession.TransactOpts, recipient, amount)
	return tx, err
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
// Will not wait for tx, but put it to ch
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_DigiUToken *DigiUTokenSession) TransferAsync(receiptCh chan *types.ReceiptResult, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	_DigiUToken.transactionSession.Lock()
	tx, err := _DigiUToken.Contract.Transfer(_DigiUToken.transactionSession.TransactOpts, recipient, amount)
	if err != nil {
		_DigiUToken.transactionSession.Unlock()
		return nil, err
	}
	_DigiUToken.transactionSession.TransactOpts.Nonce.Add(_DigiUToken.transactionSession.TransactOpts.Nonce, big.NewInt(1))
	_DigiUToken.transactionSession.Unlock()
	go func() {
		receipt, err := _DigiUToken.transactionSession.WaitTransaction(tx)
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
func (_DigiUToken *DigiUTokenTransactorSession) Transfer(wait bool, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DigiUToken.Contract.Transfer(&_DigiUToken.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_DigiUToken *DigiUTokenTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DigiUToken.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

func (_DigiUToken *DigiUTokenTransactor) TransferFromRawTx(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DigiUToken.contract.RawTx(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
// Will wait for tx receipt
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_DigiUToken *DigiUTokenSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, *types.Receipt, error) {
	_DigiUToken.transactionSession.Lock()
	tx, err := _DigiUToken.Contract.TransferFrom(_DigiUToken.transactionSession.TransactOpts, sender, recipient, amount)
	if err != nil {
		_DigiUToken.transactionSession.Unlock()
		return nil, nil, err
	}
	_DigiUToken.transactionSession.TransactOpts.Nonce.Add(_DigiUToken.transactionSession.TransactOpts.Nonce, big.NewInt(1))
	_DigiUToken.transactionSession.Unlock()
	receipt, err := _DigiUToken.transactionSession.WaitTransaction(tx)
	return tx, receipt, err
}

// TransferFrom returns raw transaction bound to the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_DigiUToken *DigiUTokenSession) TransferFromRawTx(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	tx, err := _DigiUToken.Contract.TransferFromRawTx(_DigiUToken.transactionSession.TransactOpts, sender, recipient, amount)
	return tx, err
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
// Will not wait for tx, but put it to ch
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_DigiUToken *DigiUTokenSession) TransferFromAsync(receiptCh chan *types.ReceiptResult, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	_DigiUToken.transactionSession.Lock()
	tx, err := _DigiUToken.Contract.TransferFrom(_DigiUToken.transactionSession.TransactOpts, sender, recipient, amount)
	if err != nil {
		_DigiUToken.transactionSession.Unlock()
		return nil, err
	}
	_DigiUToken.transactionSession.TransactOpts.Nonce.Add(_DigiUToken.transactionSession.TransactOpts.Nonce, big.NewInt(1))
	_DigiUToken.transactionSession.Unlock()
	go func() {
		receipt, err := _DigiUToken.transactionSession.WaitTransaction(tx)
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
func (_DigiUToken *DigiUTokenTransactorSession) TransferFrom(wait bool, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DigiUToken.Contract.TransferFrom(&_DigiUToken.TransactOpts, sender, recipient, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DigiUToken *DigiUTokenTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _DigiUToken.contract.Transact(opts, "transferOwnership", newOwner)
}

func (_DigiUToken *DigiUTokenTransactor) TransferOwnershipRawTx(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _DigiUToken.contract.RawTx(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
// Will wait for tx receipt
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DigiUToken *DigiUTokenSession) TransferOwnership(newOwner common.Address) (*types.Transaction, *types.Receipt, error) {
	_DigiUToken.transactionSession.Lock()
	tx, err := _DigiUToken.Contract.TransferOwnership(_DigiUToken.transactionSession.TransactOpts, newOwner)
	if err != nil {
		_DigiUToken.transactionSession.Unlock()
		return nil, nil, err
	}
	_DigiUToken.transactionSession.TransactOpts.Nonce.Add(_DigiUToken.transactionSession.TransactOpts.Nonce, big.NewInt(1))
	_DigiUToken.transactionSession.Unlock()
	receipt, err := _DigiUToken.transactionSession.WaitTransaction(tx)
	return tx, receipt, err
}

// TransferOwnership returns raw transaction bound to the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DigiUToken *DigiUTokenSession) TransferOwnershipRawTx(newOwner common.Address) (*types.Transaction, error) {
	tx, err := _DigiUToken.Contract.TransferOwnershipRawTx(_DigiUToken.transactionSession.TransactOpts, newOwner)
	return tx, err
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
// Will not wait for tx, but put it to ch
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DigiUToken *DigiUTokenSession) TransferOwnershipAsync(receiptCh chan *types.ReceiptResult, newOwner common.Address) (*types.Transaction, error) {
	_DigiUToken.transactionSession.Lock()
	tx, err := _DigiUToken.Contract.TransferOwnership(_DigiUToken.transactionSession.TransactOpts, newOwner)
	if err != nil {
		_DigiUToken.transactionSession.Unlock()
		return nil, err
	}
	_DigiUToken.transactionSession.TransactOpts.Nonce.Add(_DigiUToken.transactionSession.TransactOpts.Nonce, big.NewInt(1))
	_DigiUToken.transactionSession.Unlock()
	go func() {
		receipt, err := _DigiUToken.transactionSession.WaitTransaction(tx)
		receiptCh <- &types.ReceiptResult{
			Receipt: *receipt,
			Err:     err,
		}
	}()
	return tx, err
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DigiUToken *DigiUTokenTransactorSession) TransferOwnership(wait bool, newOwner common.Address) (*types.Transaction, error) {
	return _DigiUToken.Contract.TransferOwnership(&_DigiUToken.TransactOpts, newOwner)
}

// DigiUTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the DigiUToken contract.
type DigiUTokenApprovalIterator struct {
	Event *DigiUTokenApproval // Event containing the contract specifics and raw log

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
func (it *DigiUTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DigiUTokenApproval)
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
		it.Event = new(DigiUTokenApproval)
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
func (it *DigiUTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DigiUTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DigiUTokenApproval represents a Approval event raised by the DigiUToken contract.
type DigiUTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_DigiUToken *DigiUTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*DigiUTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _DigiUToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &DigiUTokenApprovalIterator{contract: _DigiUToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_DigiUToken *DigiUTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *DigiUTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _DigiUToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DigiUTokenApproval)
				if err := _DigiUToken.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_DigiUToken *DigiUTokenFilterer) ParseApproval(log types.Log) (*DigiUTokenApproval, error) {
	event := new(DigiUTokenApproval)
	if err := _DigiUToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DigiUTokenDelegateChangedIterator is returned from FilterDelegateChanged and is used to iterate over the raw logs and unpacked data for DelegateChanged events raised by the DigiUToken contract.
type DigiUTokenDelegateChangedIterator struct {
	Event *DigiUTokenDelegateChanged // Event containing the contract specifics and raw log

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
func (it *DigiUTokenDelegateChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DigiUTokenDelegateChanged)
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
		it.Event = new(DigiUTokenDelegateChanged)
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
func (it *DigiUTokenDelegateChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DigiUTokenDelegateChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DigiUTokenDelegateChanged represents a DelegateChanged event raised by the DigiUToken contract.
type DigiUTokenDelegateChanged struct {
	Delegator    common.Address
	FromDelegate common.Address
	ToDelegate   common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDelegateChanged is a free log retrieval operation binding the contract event 0x3134e8a2e6d97e929a7e54011ea5485d7d196dd5f0ba4d4ef95803e8e3fc257f.
//
// Solidity: event DelegateChanged(address indexed delegator, address indexed fromDelegate, address indexed toDelegate)
func (_DigiUToken *DigiUTokenFilterer) FilterDelegateChanged(opts *bind.FilterOpts, delegator []common.Address, fromDelegate []common.Address, toDelegate []common.Address) (*DigiUTokenDelegateChangedIterator, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var fromDelegateRule []interface{}
	for _, fromDelegateItem := range fromDelegate {
		fromDelegateRule = append(fromDelegateRule, fromDelegateItem)
	}
	var toDelegateRule []interface{}
	for _, toDelegateItem := range toDelegate {
		toDelegateRule = append(toDelegateRule, toDelegateItem)
	}

	logs, sub, err := _DigiUToken.contract.FilterLogs(opts, "DelegateChanged", delegatorRule, fromDelegateRule, toDelegateRule)
	if err != nil {
		return nil, err
	}
	return &DigiUTokenDelegateChangedIterator{contract: _DigiUToken.contract, event: "DelegateChanged", logs: logs, sub: sub}, nil
}

// WatchDelegateChanged is a free log subscription operation binding the contract event 0x3134e8a2e6d97e929a7e54011ea5485d7d196dd5f0ba4d4ef95803e8e3fc257f.
//
// Solidity: event DelegateChanged(address indexed delegator, address indexed fromDelegate, address indexed toDelegate)
func (_DigiUToken *DigiUTokenFilterer) WatchDelegateChanged(opts *bind.WatchOpts, sink chan<- *DigiUTokenDelegateChanged, delegator []common.Address, fromDelegate []common.Address, toDelegate []common.Address) (event.Subscription, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var fromDelegateRule []interface{}
	for _, fromDelegateItem := range fromDelegate {
		fromDelegateRule = append(fromDelegateRule, fromDelegateItem)
	}
	var toDelegateRule []interface{}
	for _, toDelegateItem := range toDelegate {
		toDelegateRule = append(toDelegateRule, toDelegateItem)
	}

	logs, sub, err := _DigiUToken.contract.WatchLogs(opts, "DelegateChanged", delegatorRule, fromDelegateRule, toDelegateRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DigiUTokenDelegateChanged)
				if err := _DigiUToken.contract.UnpackLog(event, "DelegateChanged", log); err != nil {
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

// ParseDelegateChanged is a log parse operation binding the contract event 0x3134e8a2e6d97e929a7e54011ea5485d7d196dd5f0ba4d4ef95803e8e3fc257f.
//
// Solidity: event DelegateChanged(address indexed delegator, address indexed fromDelegate, address indexed toDelegate)
func (_DigiUToken *DigiUTokenFilterer) ParseDelegateChanged(log types.Log) (*DigiUTokenDelegateChanged, error) {
	event := new(DigiUTokenDelegateChanged)
	if err := _DigiUToken.contract.UnpackLog(event, "DelegateChanged", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DigiUTokenDelegateVotesChangedIterator is returned from FilterDelegateVotesChanged and is used to iterate over the raw logs and unpacked data for DelegateVotesChanged events raised by the DigiUToken contract.
type DigiUTokenDelegateVotesChangedIterator struct {
	Event *DigiUTokenDelegateVotesChanged // Event containing the contract specifics and raw log

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
func (it *DigiUTokenDelegateVotesChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DigiUTokenDelegateVotesChanged)
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
		it.Event = new(DigiUTokenDelegateVotesChanged)
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
func (it *DigiUTokenDelegateVotesChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DigiUTokenDelegateVotesChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DigiUTokenDelegateVotesChanged represents a DelegateVotesChanged event raised by the DigiUToken contract.
type DigiUTokenDelegateVotesChanged struct {
	Delegate        common.Address
	PreviousBalance *big.Int
	NewBalance      *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterDelegateVotesChanged is a free log retrieval operation binding the contract event 0xdec2bacdd2f05b59de34da9b523dff8be42e5e38e818c82fdb0bae774387a724.
//
// Solidity: event DelegateVotesChanged(address indexed delegate, uint256 previousBalance, uint256 newBalance)
func (_DigiUToken *DigiUTokenFilterer) FilterDelegateVotesChanged(opts *bind.FilterOpts, delegate []common.Address) (*DigiUTokenDelegateVotesChangedIterator, error) {

	var delegateRule []interface{}
	for _, delegateItem := range delegate {
		delegateRule = append(delegateRule, delegateItem)
	}

	logs, sub, err := _DigiUToken.contract.FilterLogs(opts, "DelegateVotesChanged", delegateRule)
	if err != nil {
		return nil, err
	}
	return &DigiUTokenDelegateVotesChangedIterator{contract: _DigiUToken.contract, event: "DelegateVotesChanged", logs: logs, sub: sub}, nil
}

// WatchDelegateVotesChanged is a free log subscription operation binding the contract event 0xdec2bacdd2f05b59de34da9b523dff8be42e5e38e818c82fdb0bae774387a724.
//
// Solidity: event DelegateVotesChanged(address indexed delegate, uint256 previousBalance, uint256 newBalance)
func (_DigiUToken *DigiUTokenFilterer) WatchDelegateVotesChanged(opts *bind.WatchOpts, sink chan<- *DigiUTokenDelegateVotesChanged, delegate []common.Address) (event.Subscription, error) {

	var delegateRule []interface{}
	for _, delegateItem := range delegate {
		delegateRule = append(delegateRule, delegateItem)
	}

	logs, sub, err := _DigiUToken.contract.WatchLogs(opts, "DelegateVotesChanged", delegateRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DigiUTokenDelegateVotesChanged)
				if err := _DigiUToken.contract.UnpackLog(event, "DelegateVotesChanged", log); err != nil {
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

// ParseDelegateVotesChanged is a log parse operation binding the contract event 0xdec2bacdd2f05b59de34da9b523dff8be42e5e38e818c82fdb0bae774387a724.
//
// Solidity: event DelegateVotesChanged(address indexed delegate, uint256 previousBalance, uint256 newBalance)
func (_DigiUToken *DigiUTokenFilterer) ParseDelegateVotesChanged(log types.Log) (*DigiUTokenDelegateVotesChanged, error) {
	event := new(DigiUTokenDelegateVotesChanged)
	if err := _DigiUToken.contract.UnpackLog(event, "DelegateVotesChanged", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DigiUTokenOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the DigiUToken contract.
type DigiUTokenOwnershipTransferredIterator struct {
	Event *DigiUTokenOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *DigiUTokenOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DigiUTokenOwnershipTransferred)
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
		it.Event = new(DigiUTokenOwnershipTransferred)
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
func (it *DigiUTokenOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DigiUTokenOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DigiUTokenOwnershipTransferred represents a OwnershipTransferred event raised by the DigiUToken contract.
type DigiUTokenOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DigiUToken *DigiUTokenFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*DigiUTokenOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DigiUToken.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DigiUTokenOwnershipTransferredIterator{contract: _DigiUToken.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DigiUToken *DigiUTokenFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DigiUTokenOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DigiUToken.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DigiUTokenOwnershipTransferred)
				if err := _DigiUToken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DigiUToken *DigiUTokenFilterer) ParseOwnershipTransferred(log types.Log) (*DigiUTokenOwnershipTransferred, error) {
	event := new(DigiUTokenOwnershipTransferred)
	if err := _DigiUToken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DigiUTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the DigiUToken contract.
type DigiUTokenTransferIterator struct {
	Event *DigiUTokenTransfer // Event containing the contract specifics and raw log

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
func (it *DigiUTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DigiUTokenTransfer)
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
		it.Event = new(DigiUTokenTransfer)
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
func (it *DigiUTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DigiUTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DigiUTokenTransfer represents a Transfer event raised by the DigiUToken contract.
type DigiUTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_DigiUToken *DigiUTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*DigiUTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _DigiUToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &DigiUTokenTransferIterator{contract: _DigiUToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_DigiUToken *DigiUTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *DigiUTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _DigiUToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DigiUTokenTransfer)
				if err := _DigiUToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_DigiUToken *DigiUTokenFilterer) ParseTransfer(log types.Log) (*DigiUTokenTransfer, error) {
	event := new(DigiUTokenTransfer)
	if err := _DigiUToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}
