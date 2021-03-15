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

// MyContractABI is the input ABI used to generate the binding from.
const MyContractABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_link\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_oracle\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"ChainlinkCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"ChainlinkFulfilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"ChainlinkRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"specId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"payment\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"callbackAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"callbackFunctionId\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cancelExpiration\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"dataVersion\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"OracleRequest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"EXPIRY_TIME\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"oracle\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"jobId\",\"type\":\"bytes32\"}],\"name\":\"setPermissionJobId\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"setControl\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getPermissionJobId\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"val\",\"type\":\"address\"}],\"name\":\"setOracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getChainlinkToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"rqt\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_selector\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"receiveSide\",\"type\":\"string\"}],\"name\":\"transmitRequest\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"reqId\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"b\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"tx\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"receiveSide\",\"type\":\"address\"}],\"name\":\"receiveRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"correlationId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"tx\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"reqId\",\"type\":\"bytes32\"}],\"name\":\"receiveResponse\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_requestId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_data\",\"type\":\"bytes32\"}],\"name\":\"callback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawLink\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_requestId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_payment\",\"type\":\"uint256\"},{\"internalType\":\"bytes4\",\"name\":\"_callbackFunctionId\",\"type\":\"bytes4\"},{\"internalType\":\"uint256\",\"name\":\"_expiration\",\"type\":\"uint256\"}],\"name\":\"cancelRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// MyContractBin is the compiled bytecode used for deploying new contracts.
var MyContractBin = "0x608060405260016004553480156200001657600080fd5b506040516200229238038062002292833981810160405260408110156200003c57600080fd5b50805160209091015160006200005a6001600160e01b036200010416565b600680546001600160a01b0319166001600160a01b038316908117909155604051919250906000907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a350600780546001600160a01b0319166001600160a01b03838116919091179091558216620000e857620000e26001600160e01b036200010816565b620000fc565b620000fc826001600160e01b036200019916565b5050620001bb565b3390565b6200019773c89bd4e1632d3a43cb03aaad5262cbe4038bc5716001600160a01b03166338cc48316040518163ffffffff1660e01b815260040160206040518083038186803b1580156200015a57600080fd5b505afa1580156200016f573d6000803e3d6000fd5b505050506040513d60208110156200018657600080fd5b50516001600160e01b036200019916565b565b600280546001600160a01b0319166001600160a01b0392909216919091179055565b6120c780620001cb6000396000f3fe608060405234801561001057600080fd5b50600436106101005760003560e01c80637adbf973116100975780638f94bb87116100665780638f94bb871461044f578063acdc10981461060b578063ec65d0f814610663578063f2fde38b1461069c57610100565b80637adbf973146104115780637dc0d1d0146104375780638da5cb5b1461043f5780638dc654a21461044757610100565b80635ac8260c116100d35780635ac8260c146103a357806365f4916c146103c9578063715018a6146103e657806378626293146103ee57610100565b80631503aef914610105578063165d35e1146102c557806347bc6e9a146102e95780634b6022821461039b575b600080fd5b6102b36004803603606081101561011b57600080fd5b810190602081018135600160201b81111561013557600080fd5b82018360208201111561014757600080fd5b803590602001918460018302840111600160201b8311171561016857600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b8111156101ba57600080fd5b8201836020820111156101cc57600080fd5b803590602001918460018302840111600160201b831117156101ed57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b81111561023f57600080fd5b82018360208201111561025157600080fd5b803590602001918460018302840111600160201b8311171561027257600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506106c2945050505050565b60408051918252519081900360200190f35b6102cd61080b565b604080516001600160a01b039092168252519081900360200190f35b610399600480360360808110156102ff57600080fd5b81359190810190604081016020820135600160201b81111561032057600080fd5b82018360208201111561033257600080fd5b803590602001918460018302840111600160201b8311171561035357600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550508235935050506020013561081a565b005b6102b36109bc565b610399600480360360208110156103b957600080fd5b50356001600160a01b03166109c2565b610399600480360360208110156103df57600080fd5b5035610a3e565b610399610acb565b6103996004803603604081101561040457600080fd5b5080359060200135610b6d565b6103996004803603602081101561042757600080fd5b50356001600160a01b0316610c0c565b6102cd610c86565b6102cd610c95565b610399610ca4565b610399600480360360a081101561046557600080fd5b810190602081018135600160201b81111561047f57600080fd5b82018360208201111561049157600080fd5b803590602001918460018302840111600160201b831117156104b257600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b81111561050457600080fd5b82018360208201111561051657600080fd5b803590602001918460018302840111600160201b8311171561053757600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b81111561058957600080fd5b82018360208201111561059b57600080fd5b803590602001918460018302840111600160201b831117156105bc57600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955050823593505050602001356001600160a01b0316610e49565b610613611150565b60408051602080825283518183015283519192839290830191858101910280838360005b8381101561064f578181015183820152602001610637565b505050509050019250505060405180910390f35b6103996004803603608081101561067957600080fd5b508035906020810135906001600160e01b031960408201351690606001356111a8565b610399600480360360208110156106b257600080fd5b50356001600160a01b0316611212565b60006106cc611f4c565b6106f660086000815481106106dd57fe5b906000526020600020015430637862629360e01b61130b565b60408051808201909152600881526739b2b632b1ba37b960c11b602082015290915061072a9082908663ffffffff61133616565b60408051808201909152600c81526b726571756573745f7479706560a01b602082015261075f9082908763ffffffff61133616565b60408051808201909152600c81526b726563656976655f7369646560a01b60208201526107949082908563ffffffff61133616565b6007546107b3906001600160a01b031682670de0b6b3a7640000611365565b600081815260096020526040902080546001600160a01b031916331790556080820151519092506108039030908490670de0b6b3a7640000908390637862629360e01b9061012c9060019061153c565b509392505050565b6000610815611673565b905090565b600061085b85848460405160200180848152602001838152602001828152602001935050505060405160208183030381529060405280519060200120611682565b9050600061086982866116d3565b6001600160a01b0381166000908152600a602052604090205490915060ff1615156001146108cf576040805162461bcd60e51b815260206004820152600e60248201526d14d150d55492551648115591539560921b604482015290519081900360640190fd5b6000868152600960205260409020546001600160a01b03166109225760405162461bcd60e51b8152600401808060200182810382526021815260200180611f9c6021913960400191505060405180910390fd5b6000838152600b60205260409020805460010190819055600214156109b457600086815260096020526040808220548151630c262e5560e21b8152600481018a90526024810188905291516001600160a01b0390911692633098b954926044808201939182900301818387803b15801561099b57600080fd5b505af11580156109af573d6000803e3d6000fd5b505050505b505050505050565b61012c81565b6109ca6118b8565b6006546001600160a01b03908116911614610a1a576040805162461bcd60e51b8152602060048201819052602482015260008051602061204a833981519152604482015290519081900360640190fd5b6001600160a01b03166000908152600a60205260409020805460ff19166001179055565b610a466118b8565b6006546001600160a01b03908116911614610a96576040805162461bcd60e51b8152602060048201819052602482015260008051602061204a833981519152604482015290519081900360640190fd5b600880546001810182556000919091527ff3f7a9fe364faab93b216da50a3214154f22a0a2b415b23a84c8169e8b636ee30155565b610ad36118b8565b6006546001600160a01b03908116911614610b23576040805162461bcd60e51b8152602060048201819052602482015260008051602061204a833981519152604482015290519081900360640190fd5b6006546040516000916001600160a01b0316907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3600680546001600160a01b0319169055565b60008281526005602052604090205482906001600160a01b03163314610bc45760405162461bcd60e51b815260040180806020018281038252602881526020018061206a6028913960400191505060405180910390fd5b60008181526005602052604080822080546001600160a01b03191690555182917f7cc135e0cebb02c3480ae5d74d377283180a2601f8f644edf7987b009316c63a91a2505050565b610c146118b8565b6006546001600160a01b03908116911614610c64576040805162461bcd60e51b8152602060048201819052602482015260008051602061204a833981519152604482015290519081900360640190fd5b600780546001600160a01b0319166001600160a01b0392909216919091179055565b6007546001600160a01b031681565b6006546001600160a01b031690565b610cac6118b8565b6006546001600160a01b03908116911614610cfc576040805162461bcd60e51b8152602060048201819052602482015260008051602061204a833981519152604482015290519081900360640190fd5b6000610d06611673565b604080516370a0823160e01b815230600482015290519192506001600160a01b0383169163a9059cbb91339184916370a08231916024808301926020929190829003018186803b158015610d5957600080fd5b505afa158015610d6d573d6000803e3d6000fd5b505050506040513d6020811015610d8357600080fd5b5051604080516001600160e01b031960e086901b1681526001600160a01b03909316600484015260248301919091525160448083019260209291908290030181600087803b158015610dd457600080fd5b505af1158015610de8573d6000803e3d6000fd5b505050506040513d6020811015610dfe57600080fd5b5051610e46576040805162461bcd60e51b81526020600482015260126024820152712ab730b13632903a37903a3930b739b332b960711b604482015290519081900360640190fd5b50565b6000856040516020018082805190602001908083835b60208310610e7e5780518252601f199092019160209182019101610e5f565b6001836020036101000a0380198251168184511680821785525050505050509050019150506040516020818303038152906040528051906020012090506000610f99878686866040516020018085805190602001908083835b60208310610ef65780518252601f199092019160209182019101610ed7565b51815160209384036101000a600019018019909216911617905287519190930192870191508083835b60208310610f3e5780518252601f199092019160209182019101610f1f565b51815160209384036101000a600019018019909216911617905292019485525060609290921b6bffffffffffffffffffffffff1916838301525060408051601481850301815260349093019052815191012091506116829050565b90506000610fa782886116d3565b6001600160a01b0381166000908152600a602052604090205490915060ff16151560011461100d576040805162461bcd60e51b815260206004820152600e60248201526d14d150d55492551648115591539560921b604482015290519081900360640190fd5b6000838152600b60205260409020805460010190819055600214156111465760006060856001600160a01b0316886040518082805190602001908083835b6020831061106a5780518252601f19909201916020918201910161104b565b6001836020036101000a0380198251168184511680821785525050505050509050019150506000604051808303816000865af19150503d80600081146110cc576040519150601f19603f3d011682016040523d82523d6000602084013e6110d1565b606091505b50915091508180156110ff5750805115806110ff57508080602001905160208110156110fc57600080fd5b50515b611139576040805162461bcd60e51b815260206004820152600660248201526511905253115160d21b604482015290519081900360640190fd5b6111428a6118bc565b5050505b5050505050505050565b6060600880548060200260200160405190810160405280929190818152602001828054801561119e57602002820191906000526020600020905b81548152602001906001019080831161118a575b5050505050905090565b6111b06118b8565b6006546001600160a01b03908116911614611200576040805162461bcd60e51b8152602060048201819052602482015260008051602061204a833981519152604482015290519081900360640190fd5b61120c848484846119bf565b50505050565b61121a6118b8565b6006546001600160a01b0390811691161461126a576040805162461bcd60e51b8152602060048201819052602482015260008051602061204a833981519152604482015290519081900360640190fd5b6001600160a01b0381166112af5760405162461bcd60e51b8152600401808060200182810382526026815260200180611fbd6026913960400191505060405180910390fd5b6006546040516001600160a01b038084169216907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a3600680546001600160a01b0319166001600160a01b0392909216919091179055565b611313611f4c565b61131b611f4c565b61132d8186868663ffffffff611a9716565b95945050505050565b608083015161134b908363ffffffff611ad416565b6080830151611360908263ffffffff611ad416565b505050565b6004546040805130606090811b60208084019190915260348084018690528451808503909101815260549093018452825192810192909220908601939093526000838152600590915281812080546001600160a01b0319166001600160a01b038816179055905182917fb5e6e01e79f91267dc17b4e6314d5d4d03593d2ceee0fbb452b750bd70ea5af991a26002546001600160a01b0316634000aea0858461140d87611af1565b6040518463ffffffff1660e01b815260040180846001600160a01b03166001600160a01b0316815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b8381101561147757818101518382015260200161145f565b50505050905090810190601f1680156114a45780820380516001836020036101000a031916815260200191505b50945050505050602060405180830381600087803b1580156114c557600080fd5b505af11580156114d9573d6000803e3d6000fd5b505050506040513d60208110156114ef57600080fd5b505161152c5760405162461bcd60e51b8152600401808060200182810382526023815260200180611fe36023913960400191505060405180910390fd5b6004805460010190559392505050565b60085460015b818110156109af576008818154811061155757fe5b90600052602060002001547fd8d7ecc4800d25fa53ce0372f13a416d98907a7ef3d8d3bdd79cf4fe75529c658b8b8b8b8b8b8b8b60405180896001600160a01b03166001600160a01b03168152602001888152602001878152602001866001600160a01b03166001600160a01b03168152602001856001600160e01b0319166001600160e01b031916815260200184815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b8381101561162a578181015183820152602001611612565b50505050905090810190601f1680156116575780820380516001836020036101000a031916815260200191505b50995050505050505050505060405180910390a2600101611542565b6002546001600160a01b031690565b604080517f19457468657265756d205369676e6564204d6573736167653a0a333200000000602080830191909152603c8083019490945282518083039094018452605c909101909152815191012090565b6000815160411461172b576040805162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e67746800604482015290519081900360640190fd5b60208201516040830151606084015160001a7f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a082111561179c5760405162461bcd60e51b81526004018080602001828103825260228152602001806120066022913960400191505060405180910390fd5b8060ff16601b14806117b157508060ff16601c145b6117ec5760405162461bcd60e51b81526004018080602001828103825260228152602001806120286022913960400191505060405180910390fd5b60408051600080825260208083018085528a905260ff85168385015260608301879052608083018690529251909260019260a080820193601f1981019281900390910190855afa158015611844573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b0381166118ac576040805162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e61747572650000000000000000604482015290519081900360640190fd5b93505050505b92915050565b3390565b60006118c6611f4c565b6118d760086000815481106106dd57fe5b90506119356040518060400160405280600c81526020016b726571756573745f7479706560a01b8152506040518060400160405280600b81526020016a736574526573706f6e736560a81b815250836113369092919063ffffffff16565b60408051808201909152600d81526c18dbdc9c995b185d1a5bdb9259609a1b602082015261196b9082908563ffffffff61133616565b60075461198a906001600160a01b031682670de0b6b3a7640000611365565b91506119b93083670de0b6b3a764000060010230637862629360e01b61012c600188608001516000015161153c565b50919050565b60008481526005602052604080822080546001600160a01b0319811690915590516001600160a01b039091169186917fe1fe3afa0f7f761ff0a8b89086790efd5140d2907ebd5b7ff6bfcb5e075fd4c59190a260408051636ee4d55360e01b815260048101879052602481018690526001600160e01b0319851660448201526064810184905290516001600160a01b03831691636ee4d55391608480830192600092919082900301818387803b158015611a7857600080fd5b505af1158015611a8c573d6000803e3d6000fd5b505050505050505050565b611a9f611f4c565b611aaf8560800151610100611c31565b50509183526001600160a01b031660208301526001600160e01b031916604082015290565b611ae18260038351611c6b565b611360828263ffffffff611d7516565b6060634042994660e01b6000808460000151856020015186604001518760600151600189608001516000015160405160240180896001600160a01b03166001600160a01b03168152602001888152602001878152602001866001600160a01b03166001600160a01b03168152602001856001600160e01b0319166001600160e01b031916815260200184815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b83811015611bbf578181015183820152602001611ba7565b50505050905090810190601f168015611bec5780820380516001836020036101000a031916815260200191505b5060408051601f198184030181529190526020810180516001600160e01b03166001600160e01b0319909d169c909c17909b5250989950505050505050505050919050565b611c39611f81565b6020820615611c4e5760208206602003820191505b506020828101829052604080518085526000815290920101905290565b60178111611c9257611c8c8360e0600585901b16831763ffffffff611d9616565b50611360565b60ff8111611cc857611cb5836018611fe0600586901b161763ffffffff611d9616565b50611c8c8382600163ffffffff611dae16565b61ffff8111611cff57611cec836019611fe0600586901b161763ffffffff611d9616565b50611c8c8382600263ffffffff611dae16565b63ffffffff8111611d3857611d2583601a611fe0600586901b161763ffffffff611d9616565b50611c8c8382600463ffffffff611dae16565b67ffffffffffffffff811161136057611d6283601b611fe0600586901b161763ffffffff611d9616565b5061120c8382600863ffffffff611dae16565b611d7d611f81565b611d8f83846000015151848551611dcf565b9392505050565b611d9e611f81565b611d8f8384600001515184611e7b565b611db6611f81565b611dc7848560000151518585611ec6565b949350505050565b611dd7611f81565b8251821115611de557600080fd5b84602001518285011115611e0f57611e0f85611e078760200151878601611f24565b600202611f35565b600080865180518760208301019350808887011115611e2e5787860182525b505050602084015b60208410611e555780518252601f199093019260209182019101611e36565b51815160001960208690036101000a019081169019919091161790525083949350505050565b611e83611f81565b83602001518310611e9f57611e9f848560200151600202611f35565b835180516020858301018481535080851415611ebc576001810182525b5093949350505050565b611ece611f81565b84602001518483011115611eeb57611eeb85858401600202611f35565b60006001836101000a039050855183868201018583198251161781525080518487011115611f195783860181525b509495945050505050565b6000818311156119b95750816118b2565b8151611f418383611c31565b5061120c8382611d75565b6040805160a081018252600080825260208201819052918101829052606081019190915260808101611f7c611f81565b905290565b60405180604001604052806060815260200160008152509056fe5345435552495459204556454e543a2042414420434f5252454c4154494f4e49444f776e61626c653a206e6577206f776e657220697320746865207a65726f2061646472657373756e61626c6520746f207472616e73666572416e6443616c6c20746f206f7261636c6545434453413a20696e76616c6964207369676e6174757265202773272076616c756545434453413a20696e76616c6964207369676e6174757265202776272076616c75654f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572536f75726365206d75737420626520746865206f7261636c65206f66207468652072657175657374a26469706673582212206d212d682255665dbfabecb59500927264d93f1f8a72d6910ab0d0ce58645fcb64736f6c63430006090033"

// DeployMyContract deploys a new Ethereum contract, binding an instance of MyContract to it.
func DeployMyContract(auth *bind.TransactOpts, backend bind.ContractBackend, _link common.Address, _oracle common.Address) (common.Address, *types.Transaction, *MyContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MyContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MyContractBin), backend, _link, _oracle)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MyContract{MyContractCaller: MyContractCaller{contract: contract}, MyContractTransactor: MyContractTransactor{contract: contract}, MyContractFilterer: MyContractFilterer{contract: contract}}, nil
}

// DeployMyContractSync deploys a new Ethereum contract and waits for receipt, binding an instance of MyContractSession to it.
func DeployMyContractSync(session *bind.TransactSession, backend bind.ContractBackend, _link common.Address, _oracle common.Address) (*types.Transaction, *types.Receipt, *MyContractSession, error) {
	parsed, err := abi.JSON(strings.NewReader(MyContractABI))
	if err != nil {
		return nil, nil, nil, err
	}
	session.Lock()
	address, tx, _, err := bind.DeployContract(session.TransactOpts, parsed, common.FromHex(MyContractBin), backend, _link, _oracle)
	receipt, err := session.WaitTransaction(tx)
	if err != nil {
		session.Unlock()
		return nil, nil, nil, err
	}
	session.TransactOpts.Nonce.Add(session.TransactOpts.Nonce, big.NewInt(1))
	session.Unlock()
	contractSession, err := NewMyContractSession(address, backend, session)
	return tx, receipt, contractSession, err
}

// MyContract is an auto generated Go binding around an Ethereum contract.
type MyContract struct {
	MyContractCaller     // Read-only binding to the contract
	MyContractTransactor // Write-only binding to the contract
	MyContractFilterer   // Log filterer for contract events
}

// MyContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type MyContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MyContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MyContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MyContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MyContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MyContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MyContractSession struct {
	Contract           *MyContract // Generic contract binding to set the session for
	transactionSession *bind.TransactSession
	Address            common.Address
}

// MyContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MyContractCallerSession struct {
	Contract *MyContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// MyContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MyContractTransactorSession struct {
	Contract     *MyContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// MyContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type MyContractRaw struct {
	Contract *MyContract // Generic contract binding to access the raw methods on
}

// MyContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MyContractCallerRaw struct {
	Contract *MyContractCaller // Generic read-only contract binding to access the raw methods on
}

// MyContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MyContractTransactorRaw struct {
	Contract *MyContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMyContract creates a new instance of MyContract, bound to a specific deployed contract.
func NewMyContract(address common.Address, backend bind.ContractBackend) (*MyContract, error) {
	contract, err := bindMyContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MyContract{MyContractCaller: MyContractCaller{contract: contract}, MyContractTransactor: MyContractTransactor{contract: contract}, MyContractFilterer: MyContractFilterer{contract: contract}}, nil
}

// NewMyContractCaller creates a new read-only instance of MyContract, bound to a specific deployed contract.
func NewMyContractCaller(address common.Address, caller bind.ContractCaller) (*MyContractCaller, error) {
	contract, err := bindMyContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MyContractCaller{contract: contract}, nil
}

// NewMyContractTransactor creates a new write-only instance of MyContract, bound to a specific deployed contract.
func NewMyContractTransactor(address common.Address, transactor bind.ContractTransactor) (*MyContractTransactor, error) {
	contract, err := bindMyContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MyContractTransactor{contract: contract}, nil
}

// NewMyContractFilterer creates a new log filterer instance of MyContract, bound to a specific deployed contract.
func NewMyContractFilterer(address common.Address, filterer bind.ContractFilterer) (*MyContractFilterer, error) {
	contract, err := bindMyContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MyContractFilterer{contract: contract}, nil
}

func NewMyContractSession(address common.Address, backend bind.ContractBackend, transactionSession *bind.TransactSession) (*MyContractSession, error) {
	MyContractInstance, err := NewMyContract(address, backend)
	if err != nil {
		return nil, err
	}
	return &MyContractSession{
		Contract:           MyContractInstance,
		transactionSession: transactionSession,
		Address:            address,
	}, nil
}

// bindMyContract binds a generic wrapper to an already deployed contract.
func bindMyContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MyContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MyContract *MyContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MyContract.Contract.MyContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MyContract *MyContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MyContract.Contract.MyContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MyContract *MyContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MyContract.Contract.MyContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MyContract *MyContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MyContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MyContract *MyContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MyContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MyContract *MyContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MyContract.Contract.contract.Transact(opts, method, params...)
}

// EXPIRYTIME is a free data retrieval call binding the contract method 0x4b602282.
//
// Solidity: function EXPIRY_TIME() view returns(uint256)
func (_MyContract *MyContractCaller) EXPIRYTIME(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MyContract.contract.Call(opts, out, "EXPIRY_TIME")
	return *ret0, err
}

// EXPIRYTIME is a free data retrieval call binding the contract method 0x4b602282.
//
// Solidity: function EXPIRY_TIME() view returns(uint256)
func (_MyContract *MyContractSession) EXPIRYTIME() (*big.Int, error) {
	return _MyContract.Contract.EXPIRYTIME(_MyContract.transactionSession.CallOpts)
}

// EXPIRYTIME is a free data retrieval call binding the contract method 0x4b602282.
//
// Solidity: function EXPIRY_TIME() view returns(uint256)
func (_MyContract *MyContractCallerSession) EXPIRYTIME() (*big.Int, error) {
	return _MyContract.Contract.EXPIRYTIME(&_MyContract.CallOpts)
}

// GetChainlinkToken is a free data retrieval call binding the contract method 0x165d35e1.
//
// Solidity: function getChainlinkToken() view returns(address)
func (_MyContract *MyContractCaller) GetChainlinkToken(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MyContract.contract.Call(opts, out, "getChainlinkToken")
	return *ret0, err
}

// GetChainlinkToken is a free data retrieval call binding the contract method 0x165d35e1.
//
// Solidity: function getChainlinkToken() view returns(address)
func (_MyContract *MyContractSession) GetChainlinkToken() (common.Address, error) {
	return _MyContract.Contract.GetChainlinkToken(_MyContract.transactionSession.CallOpts)
}

// GetChainlinkToken is a free data retrieval call binding the contract method 0x165d35e1.
//
// Solidity: function getChainlinkToken() view returns(address)
func (_MyContract *MyContractCallerSession) GetChainlinkToken() (common.Address, error) {
	return _MyContract.Contract.GetChainlinkToken(&_MyContract.CallOpts)
}

// GetPermissionJobId is a free data retrieval call binding the contract method 0xacdc1098.
//
// Solidity: function getPermissionJobId() view returns(bytes32[])
func (_MyContract *MyContractCaller) GetPermissionJobId(opts *bind.CallOpts) ([][32]byte, error) {
	var (
		ret0 = new([][32]byte)
	)
	out := ret0
	err := _MyContract.contract.Call(opts, out, "getPermissionJobId")
	return *ret0, err
}

// GetPermissionJobId is a free data retrieval call binding the contract method 0xacdc1098.
//
// Solidity: function getPermissionJobId() view returns(bytes32[])
func (_MyContract *MyContractSession) GetPermissionJobId() ([][32]byte, error) {
	return _MyContract.Contract.GetPermissionJobId(_MyContract.transactionSession.CallOpts)
}

// GetPermissionJobId is a free data retrieval call binding the contract method 0xacdc1098.
//
// Solidity: function getPermissionJobId() view returns(bytes32[])
func (_MyContract *MyContractCallerSession) GetPermissionJobId() ([][32]byte, error) {
	return _MyContract.Contract.GetPermissionJobId(&_MyContract.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_MyContract *MyContractCaller) Oracle(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MyContract.contract.Call(opts, out, "oracle")
	return *ret0, err
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_MyContract *MyContractSession) Oracle() (common.Address, error) {
	return _MyContract.Contract.Oracle(_MyContract.transactionSession.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_MyContract *MyContractCallerSession) Oracle() (common.Address, error) {
	return _MyContract.Contract.Oracle(&_MyContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MyContract *MyContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MyContract.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MyContract *MyContractSession) Owner() (common.Address, error) {
	return _MyContract.Contract.Owner(_MyContract.transactionSession.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MyContract *MyContractCallerSession) Owner() (common.Address, error) {
	return _MyContract.Contract.Owner(&_MyContract.CallOpts)
}

// Callback is a paid mutator transaction binding the contract method 0x78626293.
//
// Solidity: function callback(bytes32 _requestId, bytes32 _data) returns()
func (_MyContract *MyContractTransactor) Callback(opts *bind.TransactOpts, _requestId [32]byte, _data [32]byte) (*types.Transaction, error) {
	return _MyContract.contract.Transact(opts, "callback", _requestId, _data)
}

func (_MyContract *MyContractTransactor) CallbackRawTx(opts *bind.TransactOpts, _requestId [32]byte, _data [32]byte) (*types.Transaction, error) {
	return _MyContract.contract.RawTx(opts, "callback", _requestId, _data)
}

// Callback is a paid mutator transaction binding the contract method 0x78626293.
// Will wait for tx receipt
//
// Solidity: function callback(bytes32 _requestId, bytes32 _data) returns()
func (_MyContract *MyContractSession) Callback(_requestId [32]byte, _data [32]byte) (*types.Transaction, *types.Receipt, error) {
	_MyContract.transactionSession.Lock()
	tx, err := _MyContract.Contract.Callback(_MyContract.transactionSession.TransactOpts, _requestId, _data)
	if err != nil {
		_MyContract.transactionSession.Unlock()
		return nil, nil, err
	}
	_MyContract.transactionSession.TransactOpts.Nonce.Add(_MyContract.transactionSession.TransactOpts.Nonce, big.NewInt(1))
	_MyContract.transactionSession.Unlock()
	receipt, err := _MyContract.transactionSession.WaitTransaction(tx)
	return tx, receipt, err
}

// Callback returns raw transaction bound to the contract method 0x78626293.
//
// Solidity: function callback(bytes32 _requestId, bytes32 _data) returns()
func (_MyContract *MyContractSession) CallbackRawTx(_requestId [32]byte, _data [32]byte) (*types.Transaction, error) {
	tx, err := _MyContract.Contract.CallbackRawTx(_MyContract.transactionSession.TransactOpts, _requestId, _data)
	return tx, err
}

// Callback is a paid mutator transaction binding the contract method 0x78626293.
// Will not wait for tx, but put it to ch
//
// Solidity: function callback(bytes32 _requestId, bytes32 _data) returns()
func (_MyContract *MyContractSession) CallbackAsync(receiptCh chan *types.ReceiptResult, _requestId [32]byte, _data [32]byte) (*types.Transaction, error) {
	_MyContract.transactionSession.Lock()
	tx, err := _MyContract.Contract.Callback(_MyContract.transactionSession.TransactOpts, _requestId, _data)
	if err != nil {
		_MyContract.transactionSession.Unlock()
		return nil, err
	}
	_MyContract.transactionSession.TransactOpts.Nonce.Add(_MyContract.transactionSession.TransactOpts.Nonce, big.NewInt(1))
	_MyContract.transactionSession.Unlock()
	go func() {
		receipt, err := _MyContract.transactionSession.WaitTransaction(tx)
		receiptCh <- &types.ReceiptResult{
			Receipt: *receipt,
			Err:     err,
		}
	}()
	return tx, err
}

// Callback is a paid mutator transaction binding the contract method 0x78626293.
//
// Solidity: function callback(bytes32 _requestId, bytes32 _data) returns()
func (_MyContract *MyContractTransactorSession) Callback(wait bool, _requestId [32]byte, _data [32]byte) (*types.Transaction, error) {
	return _MyContract.Contract.Callback(&_MyContract.TransactOpts, _requestId, _data)
}

// CancelRequest is a paid mutator transaction binding the contract method 0xec65d0f8.
//
// Solidity: function cancelRequest(bytes32 _requestId, uint256 _payment, bytes4 _callbackFunctionId, uint256 _expiration) returns()
func (_MyContract *MyContractTransactor) CancelRequest(opts *bind.TransactOpts, _requestId [32]byte, _payment *big.Int, _callbackFunctionId [4]byte, _expiration *big.Int) (*types.Transaction, error) {
	return _MyContract.contract.Transact(opts, "cancelRequest", _requestId, _payment, _callbackFunctionId, _expiration)
}

func (_MyContract *MyContractTransactor) CancelRequestRawTx(opts *bind.TransactOpts, _requestId [32]byte, _payment *big.Int, _callbackFunctionId [4]byte, _expiration *big.Int) (*types.Transaction, error) {
	return _MyContract.contract.RawTx(opts, "cancelRequest", _requestId, _payment, _callbackFunctionId, _expiration)
}

// CancelRequest is a paid mutator transaction binding the contract method 0xec65d0f8.
// Will wait for tx receipt
//
// Solidity: function cancelRequest(bytes32 _requestId, uint256 _payment, bytes4 _callbackFunctionId, uint256 _expiration) returns()
func (_MyContract *MyContractSession) CancelRequest(_requestId [32]byte, _payment *big.Int, _callbackFunctionId [4]byte, _expiration *big.Int) (*types.Transaction, *types.Receipt, error) {
	_MyContract.transactionSession.Lock()
	tx, err := _MyContract.Contract.CancelRequest(_MyContract.transactionSession.TransactOpts, _requestId, _payment, _callbackFunctionId, _expiration)
	if err != nil {
		_MyContract.transactionSession.Unlock()
		return nil, nil, err
	}
	_MyContract.transactionSession.TransactOpts.Nonce.Add(_MyContract.transactionSession.TransactOpts.Nonce, big.NewInt(1))
	_MyContract.transactionSession.Unlock()
	receipt, err := _MyContract.transactionSession.WaitTransaction(tx)
	return tx, receipt, err
}

// CancelRequest returns raw transaction bound to the contract method 0xec65d0f8.
//
// Solidity: function cancelRequest(bytes32 _requestId, uint256 _payment, bytes4 _callbackFunctionId, uint256 _expiration) returns()
func (_MyContract *MyContractSession) CancelRequestRawTx(_requestId [32]byte, _payment *big.Int, _callbackFunctionId [4]byte, _expiration *big.Int) (*types.Transaction, error) {
	tx, err := _MyContract.Contract.CancelRequestRawTx(_MyContract.transactionSession.TransactOpts, _requestId, _payment, _callbackFunctionId, _expiration)
	return tx, err
}

// CancelRequest is a paid mutator transaction binding the contract method 0xec65d0f8.
// Will not wait for tx, but put it to ch
//
// Solidity: function cancelRequest(bytes32 _requestId, uint256 _payment, bytes4 _callbackFunctionId, uint256 _expiration) returns()
func (_MyContract *MyContractSession) CancelRequestAsync(receiptCh chan *types.ReceiptResult, _requestId [32]byte, _payment *big.Int, _callbackFunctionId [4]byte, _expiration *big.Int) (*types.Transaction, error) {
	_MyContract.transactionSession.Lock()
	tx, err := _MyContract.Contract.CancelRequest(_MyContract.transactionSession.TransactOpts, _requestId, _payment, _callbackFunctionId, _expiration)
	if err != nil {
		_MyContract.transactionSession.Unlock()
		return nil, err
	}
	_MyContract.transactionSession.TransactOpts.Nonce.Add(_MyContract.transactionSession.TransactOpts.Nonce, big.NewInt(1))
	_MyContract.transactionSession.Unlock()
	go func() {
		receipt, err := _MyContract.transactionSession.WaitTransaction(tx)
		receiptCh <- &types.ReceiptResult{
			Receipt: *receipt,
			Err:     err,
		}
	}()
	return tx, err
}

// CancelRequest is a paid mutator transaction binding the contract method 0xec65d0f8.
//
// Solidity: function cancelRequest(bytes32 _requestId, uint256 _payment, bytes4 _callbackFunctionId, uint256 _expiration) returns()
func (_MyContract *MyContractTransactorSession) CancelRequest(wait bool, _requestId [32]byte, _payment *big.Int, _callbackFunctionId [4]byte, _expiration *big.Int) (*types.Transaction, error) {
	return _MyContract.Contract.CancelRequest(&_MyContract.TransactOpts, _requestId, _payment, _callbackFunctionId, _expiration)
}

// ReceiveRequest is a paid mutator transaction binding the contract method 0x8f94bb87.
//
// Solidity: function receiveRequest(string reqId, bytes signature, bytes b, bytes32 tx, address receiveSide) returns()
func (_MyContract *MyContractTransactor) ReceiveRequest(opts *bind.TransactOpts, reqId string, signature []byte, b []byte, tx [32]byte, receiveSide common.Address) (*types.Transaction, error) {
	return _MyContract.contract.Transact(opts, "receiveRequest", reqId, signature, b, tx, receiveSide)
}

func (_MyContract *MyContractTransactor) ReceiveRequestRawTx(opts *bind.TransactOpts, reqId string, signature []byte, b []byte, tx [32]byte, receiveSide common.Address) (*types.Transaction, error) {
	return _MyContract.contract.RawTx(opts, "receiveRequest", reqId, signature, b, tx, receiveSide)
}

// ReceiveRequest is a paid mutator transaction binding the contract method 0x8f94bb87.
// Will wait for tx receipt
//
// Solidity: function receiveRequest(string reqId, bytes signature, bytes b, bytes32 tx, address receiveSide) returns()
func (_MyContract *MyContractSession) ReceiveRequest(reqId string, signature []byte, b []byte, tx [32]byte, receiveSide common.Address) (*types.Transaction, *types.Receipt, error) {
	_MyContract.transactionSession.Lock()
	tx, err := _MyContract.Contract.ReceiveRequest(_MyContract.transactionSession.TransactOpts, reqId, signature, b, tx, receiveSide)
	if err != nil {
		_MyContract.transactionSession.Unlock()
		return nil, nil, err
	}
	_MyContract.transactionSession.TransactOpts.Nonce.Add(_MyContract.transactionSession.TransactOpts.Nonce, big.NewInt(1))
	_MyContract.transactionSession.Unlock()
	receipt, err := _MyContract.transactionSession.WaitTransaction(tx)
	return tx, receipt, err
}

// ReceiveRequest returns raw transaction bound to the contract method 0x8f94bb87.
//
// Solidity: function receiveRequest(string reqId, bytes signature, bytes b, bytes32 tx, address receiveSide) returns()
func (_MyContract *MyContractSession) ReceiveRequestRawTx(reqId string, signature []byte, b []byte, tx [32]byte, receiveSide common.Address) (*types.Transaction, error) {
	tx, err := _MyContract.Contract.ReceiveRequestRawTx(_MyContract.transactionSession.TransactOpts, reqId, signature, b, tx, receiveSide)
	return tx, err
}

// ReceiveRequest is a paid mutator transaction binding the contract method 0x8f94bb87.
// Will not wait for tx, but put it to ch
//
// Solidity: function receiveRequest(string reqId, bytes signature, bytes b, bytes32 tx, address receiveSide) returns()
func (_MyContract *MyContractSession) ReceiveRequestAsync(receiptCh chan *types.ReceiptResult, reqId string, signature []byte, b []byte, tx [32]byte, receiveSide common.Address) (*types.Transaction, error) {
	_MyContract.transactionSession.Lock()
	tx, err := _MyContract.Contract.ReceiveRequest(_MyContract.transactionSession.TransactOpts, reqId, signature, b, tx, receiveSide)
	if err != nil {
		_MyContract.transactionSession.Unlock()
		return nil, err
	}
	_MyContract.transactionSession.TransactOpts.Nonce.Add(_MyContract.transactionSession.TransactOpts.Nonce, big.NewInt(1))
	_MyContract.transactionSession.Unlock()
	go func() {
		receipt, err := _MyContract.transactionSession.WaitTransaction(tx)
		receiptCh <- &types.ReceiptResult{
			Receipt: *receipt,
			Err:     err,
		}
	}()
	return tx, err
}

// ReceiveRequest is a paid mutator transaction binding the contract method 0x8f94bb87.
//
// Solidity: function receiveRequest(string reqId, bytes signature, bytes b, bytes32 tx, address receiveSide) returns()
func (_MyContract *MyContractTransactorSession) ReceiveRequest(wait bool, reqId string, signature []byte, b []byte, tx [32]byte, receiveSide common.Address) (*types.Transaction, error) {
	return _MyContract.Contract.ReceiveRequest(&_MyContract.TransactOpts, reqId, signature, b, tx, receiveSide)
}

// ReceiveResponse is a paid mutator transaction binding the contract method 0x47bc6e9a.
//
// Solidity: function receiveResponse(bytes32 correlationId, bytes signature, bytes32 tx, bytes32 reqId) returns()
func (_MyContract *MyContractTransactor) ReceiveResponse(opts *bind.TransactOpts, correlationId [32]byte, signature []byte, tx [32]byte, reqId [32]byte) (*types.Transaction, error) {
	return _MyContract.contract.Transact(opts, "receiveResponse", correlationId, signature, tx, reqId)
}

func (_MyContract *MyContractTransactor) ReceiveResponseRawTx(opts *bind.TransactOpts, correlationId [32]byte, signature []byte, tx [32]byte, reqId [32]byte) (*types.Transaction, error) {
	return _MyContract.contract.RawTx(opts, "receiveResponse", correlationId, signature, tx, reqId)
}

// ReceiveResponse is a paid mutator transaction binding the contract method 0x47bc6e9a.
// Will wait for tx receipt
//
// Solidity: function receiveResponse(bytes32 correlationId, bytes signature, bytes32 tx, bytes32 reqId) returns()
func (_MyContract *MyContractSession) ReceiveResponse(correlationId [32]byte, signature []byte, tx [32]byte, reqId [32]byte) (*types.Transaction, *types.Receipt, error) {
	_MyContract.transactionSession.Lock()
	tx, err := _MyContract.Contract.ReceiveResponse(_MyContract.transactionSession.TransactOpts, correlationId, signature, tx, reqId)
	if err != nil {
		_MyContract.transactionSession.Unlock()
		return nil, nil, err
	}
	_MyContract.transactionSession.TransactOpts.Nonce.Add(_MyContract.transactionSession.TransactOpts.Nonce, big.NewInt(1))
	_MyContract.transactionSession.Unlock()
	receipt, err := _MyContract.transactionSession.WaitTransaction(tx)
	return tx, receipt, err
}

// ReceiveResponse returns raw transaction bound to the contract method 0x47bc6e9a.
//
// Solidity: function receiveResponse(bytes32 correlationId, bytes signature, bytes32 tx, bytes32 reqId) returns()
func (_MyContract *MyContractSession) ReceiveResponseRawTx(correlationId [32]byte, signature []byte, tx [32]byte, reqId [32]byte) (*types.Transaction, error) {
	tx, err := _MyContract.Contract.ReceiveResponseRawTx(_MyContract.transactionSession.TransactOpts, correlationId, signature, tx, reqId)
	return tx, err
}

// ReceiveResponse is a paid mutator transaction binding the contract method 0x47bc6e9a.
// Will not wait for tx, but put it to ch
//
// Solidity: function receiveResponse(bytes32 correlationId, bytes signature, bytes32 tx, bytes32 reqId) returns()
func (_MyContract *MyContractSession) ReceiveResponseAsync(receiptCh chan *types.ReceiptResult, correlationId [32]byte, signature []byte, tx [32]byte, reqId [32]byte) (*types.Transaction, error) {
	_MyContract.transactionSession.Lock()
	tx, err := _MyContract.Contract.ReceiveResponse(_MyContract.transactionSession.TransactOpts, correlationId, signature, tx, reqId)
	if err != nil {
		_MyContract.transactionSession.Unlock()
		return nil, err
	}
	_MyContract.transactionSession.TransactOpts.Nonce.Add(_MyContract.transactionSession.TransactOpts.Nonce, big.NewInt(1))
	_MyContract.transactionSession.Unlock()
	go func() {
		receipt, err := _MyContract.transactionSession.WaitTransaction(tx)
		receiptCh <- &types.ReceiptResult{
			Receipt: *receipt,
			Err:     err,
		}
	}()
	return tx, err
}

// ReceiveResponse is a paid mutator transaction binding the contract method 0x47bc6e9a.
//
// Solidity: function receiveResponse(bytes32 correlationId, bytes signature, bytes32 tx, bytes32 reqId) returns()
func (_MyContract *MyContractTransactorSession) ReceiveResponse(wait bool, correlationId [32]byte, signature []byte, tx [32]byte, reqId [32]byte) (*types.Transaction, error) {
	return _MyContract.Contract.ReceiveResponse(&_MyContract.TransactOpts, correlationId, signature, tx, reqId)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MyContract *MyContractTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MyContract.contract.Transact(opts, "renounceOwnership")
}

func (_MyContract *MyContractTransactor) RenounceOwnershipRawTx(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MyContract.contract.RawTx(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
// Will wait for tx receipt
//
// Solidity: function renounceOwnership() returns()
func (_MyContract *MyContractSession) RenounceOwnership() (*types.Transaction, *types.Receipt, error) {
	_MyContract.transactionSession.Lock()
	tx, err := _MyContract.Contract.RenounceOwnership(_MyContract.transactionSession.TransactOpts)
	if err != nil {
		_MyContract.transactionSession.Unlock()
		return nil, nil, err
	}
	_MyContract.transactionSession.TransactOpts.Nonce.Add(_MyContract.transactionSession.TransactOpts.Nonce, big.NewInt(1))
	_MyContract.transactionSession.Unlock()
	receipt, err := _MyContract.transactionSession.WaitTransaction(tx)
	return tx, receipt, err
}

// RenounceOwnership returns raw transaction bound to the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MyContract *MyContractSession) RenounceOwnershipRawTx() (*types.Transaction, error) {
	tx, err := _MyContract.Contract.RenounceOwnershipRawTx(_MyContract.transactionSession.TransactOpts)
	return tx, err
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
// Will not wait for tx, but put it to ch
//
// Solidity: function renounceOwnership() returns()
func (_MyContract *MyContractSession) RenounceOwnershipAsync(receiptCh chan *types.ReceiptResult) (*types.Transaction, error) {
	_MyContract.transactionSession.Lock()
	tx, err := _MyContract.Contract.RenounceOwnership(_MyContract.transactionSession.TransactOpts)
	if err != nil {
		_MyContract.transactionSession.Unlock()
		return nil, err
	}
	_MyContract.transactionSession.TransactOpts.Nonce.Add(_MyContract.transactionSession.TransactOpts.Nonce, big.NewInt(1))
	_MyContract.transactionSession.Unlock()
	go func() {
		receipt, err := _MyContract.transactionSession.WaitTransaction(tx)
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
func (_MyContract *MyContractTransactorSession) RenounceOwnership(wait bool) (*types.Transaction, error) {
	return _MyContract.Contract.RenounceOwnership(&_MyContract.TransactOpts)
}

// SetControl is a paid mutator transaction binding the contract method 0x5ac8260c.
//
// Solidity: function setControl(address a) returns()
func (_MyContract *MyContractTransactor) SetControl(opts *bind.TransactOpts, a common.Address) (*types.Transaction, error) {
	return _MyContract.contract.Transact(opts, "setControl", a)
}

func (_MyContract *MyContractTransactor) SetControlRawTx(opts *bind.TransactOpts, a common.Address) (*types.Transaction, error) {
	return _MyContract.contract.RawTx(opts, "setControl", a)
}

// SetControl is a paid mutator transaction binding the contract method 0x5ac8260c.
// Will wait for tx receipt
//
// Solidity: function setControl(address a) returns()
func (_MyContract *MyContractSession) SetControl(a common.Address) (*types.Transaction, *types.Receipt, error) {
	_MyContract.transactionSession.Lock()
	tx, err := _MyContract.Contract.SetControl(_MyContract.transactionSession.TransactOpts, a)
	if err != nil {
		_MyContract.transactionSession.Unlock()
		return nil, nil, err
	}
	_MyContract.transactionSession.TransactOpts.Nonce.Add(_MyContract.transactionSession.TransactOpts.Nonce, big.NewInt(1))
	_MyContract.transactionSession.Unlock()
	receipt, err := _MyContract.transactionSession.WaitTransaction(tx)
	return tx, receipt, err
}

// SetControl returns raw transaction bound to the contract method 0x5ac8260c.
//
// Solidity: function setControl(address a) returns()
func (_MyContract *MyContractSession) SetControlRawTx(a common.Address) (*types.Transaction, error) {
	tx, err := _MyContract.Contract.SetControlRawTx(_MyContract.transactionSession.TransactOpts, a)
	return tx, err
}

// SetControl is a paid mutator transaction binding the contract method 0x5ac8260c.
// Will not wait for tx, but put it to ch
//
// Solidity: function setControl(address a) returns()
func (_MyContract *MyContractSession) SetControlAsync(receiptCh chan *types.ReceiptResult, a common.Address) (*types.Transaction, error) {
	_MyContract.transactionSession.Lock()
	tx, err := _MyContract.Contract.SetControl(_MyContract.transactionSession.TransactOpts, a)
	if err != nil {
		_MyContract.transactionSession.Unlock()
		return nil, err
	}
	_MyContract.transactionSession.TransactOpts.Nonce.Add(_MyContract.transactionSession.TransactOpts.Nonce, big.NewInt(1))
	_MyContract.transactionSession.Unlock()
	go func() {
		receipt, err := _MyContract.transactionSession.WaitTransaction(tx)
		receiptCh <- &types.ReceiptResult{
			Receipt: *receipt,
			Err:     err,
		}
	}()
	return tx, err
}

// SetControl is a paid mutator transaction binding the contract method 0x5ac8260c.
//
// Solidity: function setControl(address a) returns()
func (_MyContract *MyContractTransactorSession) SetControl(wait bool, a common.Address) (*types.Transaction, error) {
	return _MyContract.Contract.SetControl(&_MyContract.TransactOpts, a)
}

// SetOracle is a paid mutator transaction binding the contract method 0x7adbf973.
//
// Solidity: function setOracle(address val) returns()
func (_MyContract *MyContractTransactor) SetOracle(opts *bind.TransactOpts, val common.Address) (*types.Transaction, error) {
	return _MyContract.contract.Transact(opts, "setOracle", val)
}

func (_MyContract *MyContractTransactor) SetOracleRawTx(opts *bind.TransactOpts, val common.Address) (*types.Transaction, error) {
	return _MyContract.contract.RawTx(opts, "setOracle", val)
}

// SetOracle is a paid mutator transaction binding the contract method 0x7adbf973.
// Will wait for tx receipt
//
// Solidity: function setOracle(address val) returns()
func (_MyContract *MyContractSession) SetOracle(val common.Address) (*types.Transaction, *types.Receipt, error) {
	_MyContract.transactionSession.Lock()
	tx, err := _MyContract.Contract.SetOracle(_MyContract.transactionSession.TransactOpts, val)
	if err != nil {
		_MyContract.transactionSession.Unlock()
		return nil, nil, err
	}
	_MyContract.transactionSession.TransactOpts.Nonce.Add(_MyContract.transactionSession.TransactOpts.Nonce, big.NewInt(1))
	_MyContract.transactionSession.Unlock()
	receipt, err := _MyContract.transactionSession.WaitTransaction(tx)
	return tx, receipt, err
}

// SetOracle returns raw transaction bound to the contract method 0x7adbf973.
//
// Solidity: function setOracle(address val) returns()
func (_MyContract *MyContractSession) SetOracleRawTx(val common.Address) (*types.Transaction, error) {
	tx, err := _MyContract.Contract.SetOracleRawTx(_MyContract.transactionSession.TransactOpts, val)
	return tx, err
}

// SetOracle is a paid mutator transaction binding the contract method 0x7adbf973.
// Will not wait for tx, but put it to ch
//
// Solidity: function setOracle(address val) returns()
func (_MyContract *MyContractSession) SetOracleAsync(receiptCh chan *types.ReceiptResult, val common.Address) (*types.Transaction, error) {
	_MyContract.transactionSession.Lock()
	tx, err := _MyContract.Contract.SetOracle(_MyContract.transactionSession.TransactOpts, val)
	if err != nil {
		_MyContract.transactionSession.Unlock()
		return nil, err
	}
	_MyContract.transactionSession.TransactOpts.Nonce.Add(_MyContract.transactionSession.TransactOpts.Nonce, big.NewInt(1))
	_MyContract.transactionSession.Unlock()
	go func() {
		receipt, err := _MyContract.transactionSession.WaitTransaction(tx)
		receiptCh <- &types.ReceiptResult{
			Receipt: *receipt,
			Err:     err,
		}
	}()
	return tx, err
}

// SetOracle is a paid mutator transaction binding the contract method 0x7adbf973.
//
// Solidity: function setOracle(address val) returns()
func (_MyContract *MyContractTransactorSession) SetOracle(wait bool, val common.Address) (*types.Transaction, error) {
	return _MyContract.Contract.SetOracle(&_MyContract.TransactOpts, val)
}

// SetPermissionJobId is a paid mutator transaction binding the contract method 0x65f4916c.
//
// Solidity: function setPermissionJobId(bytes32 jobId) returns()
func (_MyContract *MyContractTransactor) SetPermissionJobId(opts *bind.TransactOpts, jobId [32]byte) (*types.Transaction, error) {
	return _MyContract.contract.Transact(opts, "setPermissionJobId", jobId)
}

func (_MyContract *MyContractTransactor) SetPermissionJobIdRawTx(opts *bind.TransactOpts, jobId [32]byte) (*types.Transaction, error) {
	return _MyContract.contract.RawTx(opts, "setPermissionJobId", jobId)
}

// SetPermissionJobId is a paid mutator transaction binding the contract method 0x65f4916c.
// Will wait for tx receipt
//
// Solidity: function setPermissionJobId(bytes32 jobId) returns()
func (_MyContract *MyContractSession) SetPermissionJobId(jobId [32]byte) (*types.Transaction, *types.Receipt, error) {
	_MyContract.transactionSession.Lock()
	tx, err := _MyContract.Contract.SetPermissionJobId(_MyContract.transactionSession.TransactOpts, jobId)
	if err != nil {
		_MyContract.transactionSession.Unlock()
		return nil, nil, err
	}
	_MyContract.transactionSession.TransactOpts.Nonce.Add(_MyContract.transactionSession.TransactOpts.Nonce, big.NewInt(1))
	_MyContract.transactionSession.Unlock()
	receipt, err := _MyContract.transactionSession.WaitTransaction(tx)
	return tx, receipt, err
}

// SetPermissionJobId returns raw transaction bound to the contract method 0x65f4916c.
//
// Solidity: function setPermissionJobId(bytes32 jobId) returns()
func (_MyContract *MyContractSession) SetPermissionJobIdRawTx(jobId [32]byte) (*types.Transaction, error) {
	tx, err := _MyContract.Contract.SetPermissionJobIdRawTx(_MyContract.transactionSession.TransactOpts, jobId)
	return tx, err
}

// SetPermissionJobId is a paid mutator transaction binding the contract method 0x65f4916c.
// Will not wait for tx, but put it to ch
//
// Solidity: function setPermissionJobId(bytes32 jobId) returns()
func (_MyContract *MyContractSession) SetPermissionJobIdAsync(receiptCh chan *types.ReceiptResult, jobId [32]byte) (*types.Transaction, error) {
	_MyContract.transactionSession.Lock()
	tx, err := _MyContract.Contract.SetPermissionJobId(_MyContract.transactionSession.TransactOpts, jobId)
	if err != nil {
		_MyContract.transactionSession.Unlock()
		return nil, err
	}
	_MyContract.transactionSession.TransactOpts.Nonce.Add(_MyContract.transactionSession.TransactOpts.Nonce, big.NewInt(1))
	_MyContract.transactionSession.Unlock()
	go func() {
		receipt, err := _MyContract.transactionSession.WaitTransaction(tx)
		receiptCh <- &types.ReceiptResult{
			Receipt: *receipt,
			Err:     err,
		}
	}()
	return tx, err
}

// SetPermissionJobId is a paid mutator transaction binding the contract method 0x65f4916c.
//
// Solidity: function setPermissionJobId(bytes32 jobId) returns()
func (_MyContract *MyContractTransactorSession) SetPermissionJobId(wait bool, jobId [32]byte) (*types.Transaction, error) {
	return _MyContract.Contract.SetPermissionJobId(&_MyContract.TransactOpts, jobId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MyContract *MyContractTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _MyContract.contract.Transact(opts, "transferOwnership", newOwner)
}

func (_MyContract *MyContractTransactor) TransferOwnershipRawTx(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _MyContract.contract.RawTx(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
// Will wait for tx receipt
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MyContract *MyContractSession) TransferOwnership(newOwner common.Address) (*types.Transaction, *types.Receipt, error) {
	_MyContract.transactionSession.Lock()
	tx, err := _MyContract.Contract.TransferOwnership(_MyContract.transactionSession.TransactOpts, newOwner)
	if err != nil {
		_MyContract.transactionSession.Unlock()
		return nil, nil, err
	}
	_MyContract.transactionSession.TransactOpts.Nonce.Add(_MyContract.transactionSession.TransactOpts.Nonce, big.NewInt(1))
	_MyContract.transactionSession.Unlock()
	receipt, err := _MyContract.transactionSession.WaitTransaction(tx)
	return tx, receipt, err
}

// TransferOwnership returns raw transaction bound to the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MyContract *MyContractSession) TransferOwnershipRawTx(newOwner common.Address) (*types.Transaction, error) {
	tx, err := _MyContract.Contract.TransferOwnershipRawTx(_MyContract.transactionSession.TransactOpts, newOwner)
	return tx, err
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
// Will not wait for tx, but put it to ch
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MyContract *MyContractSession) TransferOwnershipAsync(receiptCh chan *types.ReceiptResult, newOwner common.Address) (*types.Transaction, error) {
	_MyContract.transactionSession.Lock()
	tx, err := _MyContract.Contract.TransferOwnership(_MyContract.transactionSession.TransactOpts, newOwner)
	if err != nil {
		_MyContract.transactionSession.Unlock()
		return nil, err
	}
	_MyContract.transactionSession.TransactOpts.Nonce.Add(_MyContract.transactionSession.TransactOpts.Nonce, big.NewInt(1))
	_MyContract.transactionSession.Unlock()
	go func() {
		receipt, err := _MyContract.transactionSession.WaitTransaction(tx)
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
func (_MyContract *MyContractTransactorSession) TransferOwnership(wait bool, newOwner common.Address) (*types.Transaction, error) {
	return _MyContract.Contract.TransferOwnership(&_MyContract.TransactOpts, newOwner)
}

// TransmitRequest is a paid mutator transaction binding the contract method 0x1503aef9.
//
// Solidity: function transmitRequest(string rqt, string _selector, string receiveSide) returns(bytes32 requestId)
func (_MyContract *MyContractTransactor) TransmitRequest(opts *bind.TransactOpts, rqt string, _selector string, receiveSide string) (*types.Transaction, error) {
	return _MyContract.contract.Transact(opts, "transmitRequest", rqt, _selector, receiveSide)
}

func (_MyContract *MyContractTransactor) TransmitRequestRawTx(opts *bind.TransactOpts, rqt string, _selector string, receiveSide string) (*types.Transaction, error) {
	return _MyContract.contract.RawTx(opts, "transmitRequest", rqt, _selector, receiveSide)
}

// TransmitRequest is a paid mutator transaction binding the contract method 0x1503aef9.
// Will wait for tx receipt
//
// Solidity: function transmitRequest(string rqt, string _selector, string receiveSide) returns(bytes32 requestId)
func (_MyContract *MyContractSession) TransmitRequest(rqt string, _selector string, receiveSide string) (*types.Transaction, *types.Receipt, error) {
	_MyContract.transactionSession.Lock()
	tx, err := _MyContract.Contract.TransmitRequest(_MyContract.transactionSession.TransactOpts, rqt, _selector, receiveSide)
	if err != nil {
		_MyContract.transactionSession.Unlock()
		return nil, nil, err
	}
	_MyContract.transactionSession.TransactOpts.Nonce.Add(_MyContract.transactionSession.TransactOpts.Nonce, big.NewInt(1))
	_MyContract.transactionSession.Unlock()
	receipt, err := _MyContract.transactionSession.WaitTransaction(tx)
	return tx, receipt, err
}

// TransmitRequest returns raw transaction bound to the contract method 0x1503aef9.
//
// Solidity: function transmitRequest(string rqt, string _selector, string receiveSide) returns(bytes32 requestId)
func (_MyContract *MyContractSession) TransmitRequestRawTx(rqt string, _selector string, receiveSide string) (*types.Transaction, error) {
	tx, err := _MyContract.Contract.TransmitRequestRawTx(_MyContract.transactionSession.TransactOpts, rqt, _selector, receiveSide)
	return tx, err
}

// TransmitRequest is a paid mutator transaction binding the contract method 0x1503aef9.
// Will not wait for tx, but put it to ch
//
// Solidity: function transmitRequest(string rqt, string _selector, string receiveSide) returns(bytes32 requestId)
func (_MyContract *MyContractSession) TransmitRequestAsync(receiptCh chan *types.ReceiptResult, rqt string, _selector string, receiveSide string) (*types.Transaction, error) {
	_MyContract.transactionSession.Lock()
	tx, err := _MyContract.Contract.TransmitRequest(_MyContract.transactionSession.TransactOpts, rqt, _selector, receiveSide)
	if err != nil {
		_MyContract.transactionSession.Unlock()
		return nil, err
	}
	_MyContract.transactionSession.TransactOpts.Nonce.Add(_MyContract.transactionSession.TransactOpts.Nonce, big.NewInt(1))
	_MyContract.transactionSession.Unlock()
	go func() {
		receipt, err := _MyContract.transactionSession.WaitTransaction(tx)
		receiptCh <- &types.ReceiptResult{
			Receipt: *receipt,
			Err:     err,
		}
	}()
	return tx, err
}

// TransmitRequest is a paid mutator transaction binding the contract method 0x1503aef9.
//
// Solidity: function transmitRequest(string rqt, string _selector, string receiveSide) returns(bytes32 requestId)
func (_MyContract *MyContractTransactorSession) TransmitRequest(wait bool, rqt string, _selector string, receiveSide string) (*types.Transaction, error) {
	return _MyContract.Contract.TransmitRequest(&_MyContract.TransactOpts, rqt, _selector, receiveSide)
}

// WithdrawLink is a paid mutator transaction binding the contract method 0x8dc654a2.
//
// Solidity: function withdrawLink() returns()
func (_MyContract *MyContractTransactor) WithdrawLink(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MyContract.contract.Transact(opts, "withdrawLink")
}

func (_MyContract *MyContractTransactor) WithdrawLinkRawTx(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MyContract.contract.RawTx(opts, "withdrawLink")
}

// WithdrawLink is a paid mutator transaction binding the contract method 0x8dc654a2.
// Will wait for tx receipt
//
// Solidity: function withdrawLink() returns()
func (_MyContract *MyContractSession) WithdrawLink() (*types.Transaction, *types.Receipt, error) {
	_MyContract.transactionSession.Lock()
	tx, err := _MyContract.Contract.WithdrawLink(_MyContract.transactionSession.TransactOpts)
	if err != nil {
		_MyContract.transactionSession.Unlock()
		return nil, nil, err
	}
	_MyContract.transactionSession.TransactOpts.Nonce.Add(_MyContract.transactionSession.TransactOpts.Nonce, big.NewInt(1))
	_MyContract.transactionSession.Unlock()
	receipt, err := _MyContract.transactionSession.WaitTransaction(tx)
	return tx, receipt, err
}

// WithdrawLink returns raw transaction bound to the contract method 0x8dc654a2.
//
// Solidity: function withdrawLink() returns()
func (_MyContract *MyContractSession) WithdrawLinkRawTx() (*types.Transaction, error) {
	tx, err := _MyContract.Contract.WithdrawLinkRawTx(_MyContract.transactionSession.TransactOpts)
	return tx, err
}

// WithdrawLink is a paid mutator transaction binding the contract method 0x8dc654a2.
// Will not wait for tx, but put it to ch
//
// Solidity: function withdrawLink() returns()
func (_MyContract *MyContractSession) WithdrawLinkAsync(receiptCh chan *types.ReceiptResult) (*types.Transaction, error) {
	_MyContract.transactionSession.Lock()
	tx, err := _MyContract.Contract.WithdrawLink(_MyContract.transactionSession.TransactOpts)
	if err != nil {
		_MyContract.transactionSession.Unlock()
		return nil, err
	}
	_MyContract.transactionSession.TransactOpts.Nonce.Add(_MyContract.transactionSession.TransactOpts.Nonce, big.NewInt(1))
	_MyContract.transactionSession.Unlock()
	go func() {
		receipt, err := _MyContract.transactionSession.WaitTransaction(tx)
		receiptCh <- &types.ReceiptResult{
			Receipt: *receipt,
			Err:     err,
		}
	}()
	return tx, err
}

// WithdrawLink is a paid mutator transaction binding the contract method 0x8dc654a2.
//
// Solidity: function withdrawLink() returns()
func (_MyContract *MyContractTransactorSession) WithdrawLink(wait bool) (*types.Transaction, error) {
	return _MyContract.Contract.WithdrawLink(&_MyContract.TransactOpts)
}

// MyContractChainlinkCancelledIterator is returned from FilterChainlinkCancelled and is used to iterate over the raw logs and unpacked data for ChainlinkCancelled events raised by the MyContract contract.
type MyContractChainlinkCancelledIterator struct {
	Event *MyContractChainlinkCancelled // Event containing the contract specifics and raw log

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
func (it *MyContractChainlinkCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MyContractChainlinkCancelled)
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
		it.Event = new(MyContractChainlinkCancelled)
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
func (it *MyContractChainlinkCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MyContractChainlinkCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MyContractChainlinkCancelled represents a ChainlinkCancelled event raised by the MyContract contract.
type MyContractChainlinkCancelled struct {
	Id  [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterChainlinkCancelled is a free log retrieval operation binding the contract event 0xe1fe3afa0f7f761ff0a8b89086790efd5140d2907ebd5b7ff6bfcb5e075fd4c5.
//
// Solidity: event ChainlinkCancelled(bytes32 indexed id)
func (_MyContract *MyContractFilterer) FilterChainlinkCancelled(opts *bind.FilterOpts, id [][32]byte) (*MyContractChainlinkCancelledIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _MyContract.contract.FilterLogs(opts, "ChainlinkCancelled", idRule)
	if err != nil {
		return nil, err
	}
	return &MyContractChainlinkCancelledIterator{contract: _MyContract.contract, event: "ChainlinkCancelled", logs: logs, sub: sub}, nil
}

// WatchChainlinkCancelled is a free log subscription operation binding the contract event 0xe1fe3afa0f7f761ff0a8b89086790efd5140d2907ebd5b7ff6bfcb5e075fd4c5.
//
// Solidity: event ChainlinkCancelled(bytes32 indexed id)
func (_MyContract *MyContractFilterer) WatchChainlinkCancelled(opts *bind.WatchOpts, sink chan<- *MyContractChainlinkCancelled, id [][32]byte) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _MyContract.contract.WatchLogs(opts, "ChainlinkCancelled", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MyContractChainlinkCancelled)
				if err := _MyContract.contract.UnpackLog(event, "ChainlinkCancelled", log); err != nil {
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

// ParseChainlinkCancelled is a log parse operation binding the contract event 0xe1fe3afa0f7f761ff0a8b89086790efd5140d2907ebd5b7ff6bfcb5e075fd4c5.
//
// Solidity: event ChainlinkCancelled(bytes32 indexed id)
func (_MyContract *MyContractFilterer) ParseChainlinkCancelled(log types.Log) (*MyContractChainlinkCancelled, error) {
	event := new(MyContractChainlinkCancelled)
	if err := _MyContract.contract.UnpackLog(event, "ChainlinkCancelled", log); err != nil {
		return nil, err
	}
	return event, nil
}

// MyContractChainlinkFulfilledIterator is returned from FilterChainlinkFulfilled and is used to iterate over the raw logs and unpacked data for ChainlinkFulfilled events raised by the MyContract contract.
type MyContractChainlinkFulfilledIterator struct {
	Event *MyContractChainlinkFulfilled // Event containing the contract specifics and raw log

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
func (it *MyContractChainlinkFulfilledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MyContractChainlinkFulfilled)
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
		it.Event = new(MyContractChainlinkFulfilled)
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
func (it *MyContractChainlinkFulfilledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MyContractChainlinkFulfilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MyContractChainlinkFulfilled represents a ChainlinkFulfilled event raised by the MyContract contract.
type MyContractChainlinkFulfilled struct {
	Id  [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterChainlinkFulfilled is a free log retrieval operation binding the contract event 0x7cc135e0cebb02c3480ae5d74d377283180a2601f8f644edf7987b009316c63a.
//
// Solidity: event ChainlinkFulfilled(bytes32 indexed id)
func (_MyContract *MyContractFilterer) FilterChainlinkFulfilled(opts *bind.FilterOpts, id [][32]byte) (*MyContractChainlinkFulfilledIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _MyContract.contract.FilterLogs(opts, "ChainlinkFulfilled", idRule)
	if err != nil {
		return nil, err
	}
	return &MyContractChainlinkFulfilledIterator{contract: _MyContract.contract, event: "ChainlinkFulfilled", logs: logs, sub: sub}, nil
}

// WatchChainlinkFulfilled is a free log subscription operation binding the contract event 0x7cc135e0cebb02c3480ae5d74d377283180a2601f8f644edf7987b009316c63a.
//
// Solidity: event ChainlinkFulfilled(bytes32 indexed id)
func (_MyContract *MyContractFilterer) WatchChainlinkFulfilled(opts *bind.WatchOpts, sink chan<- *MyContractChainlinkFulfilled, id [][32]byte) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _MyContract.contract.WatchLogs(opts, "ChainlinkFulfilled", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MyContractChainlinkFulfilled)
				if err := _MyContract.contract.UnpackLog(event, "ChainlinkFulfilled", log); err != nil {
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

// ParseChainlinkFulfilled is a log parse operation binding the contract event 0x7cc135e0cebb02c3480ae5d74d377283180a2601f8f644edf7987b009316c63a.
//
// Solidity: event ChainlinkFulfilled(bytes32 indexed id)
func (_MyContract *MyContractFilterer) ParseChainlinkFulfilled(log types.Log) (*MyContractChainlinkFulfilled, error) {
	event := new(MyContractChainlinkFulfilled)
	if err := _MyContract.contract.UnpackLog(event, "ChainlinkFulfilled", log); err != nil {
		return nil, err
	}
	return event, nil
}

// MyContractChainlinkRequestedIterator is returned from FilterChainlinkRequested and is used to iterate over the raw logs and unpacked data for ChainlinkRequested events raised by the MyContract contract.
type MyContractChainlinkRequestedIterator struct {
	Event *MyContractChainlinkRequested // Event containing the contract specifics and raw log

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
func (it *MyContractChainlinkRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MyContractChainlinkRequested)
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
		it.Event = new(MyContractChainlinkRequested)
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
func (it *MyContractChainlinkRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MyContractChainlinkRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MyContractChainlinkRequested represents a ChainlinkRequested event raised by the MyContract contract.
type MyContractChainlinkRequested struct {
	Id  [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterChainlinkRequested is a free log retrieval operation binding the contract event 0xb5e6e01e79f91267dc17b4e6314d5d4d03593d2ceee0fbb452b750bd70ea5af9.
//
// Solidity: event ChainlinkRequested(bytes32 indexed id)
func (_MyContract *MyContractFilterer) FilterChainlinkRequested(opts *bind.FilterOpts, id [][32]byte) (*MyContractChainlinkRequestedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _MyContract.contract.FilterLogs(opts, "ChainlinkRequested", idRule)
	if err != nil {
		return nil, err
	}
	return &MyContractChainlinkRequestedIterator{contract: _MyContract.contract, event: "ChainlinkRequested", logs: logs, sub: sub}, nil
}

// WatchChainlinkRequested is a free log subscription operation binding the contract event 0xb5e6e01e79f91267dc17b4e6314d5d4d03593d2ceee0fbb452b750bd70ea5af9.
//
// Solidity: event ChainlinkRequested(bytes32 indexed id)
func (_MyContract *MyContractFilterer) WatchChainlinkRequested(opts *bind.WatchOpts, sink chan<- *MyContractChainlinkRequested, id [][32]byte) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _MyContract.contract.WatchLogs(opts, "ChainlinkRequested", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MyContractChainlinkRequested)
				if err := _MyContract.contract.UnpackLog(event, "ChainlinkRequested", log); err != nil {
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

// ParseChainlinkRequested is a log parse operation binding the contract event 0xb5e6e01e79f91267dc17b4e6314d5d4d03593d2ceee0fbb452b750bd70ea5af9.
//
// Solidity: event ChainlinkRequested(bytes32 indexed id)
func (_MyContract *MyContractFilterer) ParseChainlinkRequested(log types.Log) (*MyContractChainlinkRequested, error) {
	event := new(MyContractChainlinkRequested)
	if err := _MyContract.contract.UnpackLog(event, "ChainlinkRequested", log); err != nil {
		return nil, err
	}
	return event, nil
}

// MyContractOracleRequestIterator is returned from FilterOracleRequest and is used to iterate over the raw logs and unpacked data for OracleRequest events raised by the MyContract contract.
type MyContractOracleRequestIterator struct {
	Event *MyContractOracleRequest // Event containing the contract specifics and raw log

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
func (it *MyContractOracleRequestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MyContractOracleRequest)
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
		it.Event = new(MyContractOracleRequest)
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
func (it *MyContractOracleRequestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MyContractOracleRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MyContractOracleRequest represents a OracleRequest event raised by the MyContract contract.
type MyContractOracleRequest struct {
	SpecId             [32]byte
	Requester          common.Address
	RequestId          [32]byte
	Payment            *big.Int
	CallbackAddr       common.Address
	CallbackFunctionId [4]byte
	CancelExpiration   *big.Int
	DataVersion        *big.Int
	Data               []byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterOracleRequest is a free log retrieval operation binding the contract event 0xd8d7ecc4800d25fa53ce0372f13a416d98907a7ef3d8d3bdd79cf4fe75529c65.
//
// Solidity: event OracleRequest(bytes32 indexed specId, address requester, bytes32 requestId, uint256 payment, address callbackAddr, bytes4 callbackFunctionId, uint256 cancelExpiration, uint256 dataVersion, bytes data)
func (_MyContract *MyContractFilterer) FilterOracleRequest(opts *bind.FilterOpts, specId [][32]byte) (*MyContractOracleRequestIterator, error) {

	var specIdRule []interface{}
	for _, specIdItem := range specId {
		specIdRule = append(specIdRule, specIdItem)
	}

	logs, sub, err := _MyContract.contract.FilterLogs(opts, "OracleRequest", specIdRule)
	if err != nil {
		return nil, err
	}
	return &MyContractOracleRequestIterator{contract: _MyContract.contract, event: "OracleRequest", logs: logs, sub: sub}, nil
}

// WatchOracleRequest is a free log subscription operation binding the contract event 0xd8d7ecc4800d25fa53ce0372f13a416d98907a7ef3d8d3bdd79cf4fe75529c65.
//
// Solidity: event OracleRequest(bytes32 indexed specId, address requester, bytes32 requestId, uint256 payment, address callbackAddr, bytes4 callbackFunctionId, uint256 cancelExpiration, uint256 dataVersion, bytes data)
func (_MyContract *MyContractFilterer) WatchOracleRequest(opts *bind.WatchOpts, sink chan<- *MyContractOracleRequest, specId [][32]byte) (event.Subscription, error) {

	var specIdRule []interface{}
	for _, specIdItem := range specId {
		specIdRule = append(specIdRule, specIdItem)
	}

	logs, sub, err := _MyContract.contract.WatchLogs(opts, "OracleRequest", specIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MyContractOracleRequest)
				if err := _MyContract.contract.UnpackLog(event, "OracleRequest", log); err != nil {
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

// ParseOracleRequest is a log parse operation binding the contract event 0xd8d7ecc4800d25fa53ce0372f13a416d98907a7ef3d8d3bdd79cf4fe75529c65.
//
// Solidity: event OracleRequest(bytes32 indexed specId, address requester, bytes32 requestId, uint256 payment, address callbackAddr, bytes4 callbackFunctionId, uint256 cancelExpiration, uint256 dataVersion, bytes data)
func (_MyContract *MyContractFilterer) ParseOracleRequest(log types.Log) (*MyContractOracleRequest, error) {
	event := new(MyContractOracleRequest)
	if err := _MyContract.contract.UnpackLog(event, "OracleRequest", log); err != nil {
		return nil, err
	}
	return event, nil
}

// MyContractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the MyContract contract.
type MyContractOwnershipTransferredIterator struct {
	Event *MyContractOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MyContractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MyContractOwnershipTransferred)
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
		it.Event = new(MyContractOwnershipTransferred)
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
func (it *MyContractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MyContractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MyContractOwnershipTransferred represents a OwnershipTransferred event raised by the MyContract contract.
type MyContractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MyContract *MyContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MyContractOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MyContract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MyContractOwnershipTransferredIterator{contract: _MyContract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MyContract *MyContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MyContractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MyContract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MyContractOwnershipTransferred)
				if err := _MyContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_MyContract *MyContractFilterer) ParseOwnershipTransferred(log types.Log) (*MyContractOwnershipTransferred, error) {
	event := new(MyContractOwnershipTransferred)
	if err := _MyContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}
