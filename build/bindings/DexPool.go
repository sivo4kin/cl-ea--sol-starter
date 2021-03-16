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

// DexPoolABI is the input ABI used to generate the binding from.
const DexPoolABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenOfPool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_myContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_util\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"curBlock\",\"type\":\"uint256\"}],\"name\":\"AddLiquidity\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"name\":\"Receive\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"curBlock\",\"type\":\"uint256\"}],\"name\":\"SwapDeposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"SwapWithdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"tx_fromNet2\",\"type\":\"bytes32\"}],\"name\":\"TxCompleteBothChain\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"myContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"test\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"util\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"luqidityProvider\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"addLiquidityInternal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountNet1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountNet2\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"luqidityProviderNet2\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balancePoolNet2\",\"type\":\"uint256\"}],\"name\":\"addLiquidity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount2\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipientOnNet2\",\"type\":\"address\"}],\"name\":\"swapDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"tx_fromNet2\",\"type\":\"bytes32\"}],\"name\":\"setPendingRequestsDone\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"swapWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balancePoolNet2\",\"type\":\"uint256\"}],\"name\":\"calculateLP\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"val\",\"type\":\"uint256\"}],\"name\":\"_setTest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_getTest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"b\",\"type\":\"bytes\"}],\"name\":\"lowLevelGet\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"adr\",\"type\":\"address\"}],\"name\":\"setSecondPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// DexPoolBin is the compiled bytecode used for deploying new contracts.
var DexPoolBin = "0x608060405234801561001057600080fd5b506040516120893803806120898339818101604052606081101561003357600080fd5b508051602082015160409092015190919060006100576001600160e01b036100e716565b600380546001600160a01b0319166001600160a01b038316908117909155604051919250906000907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a350600980546001600160a01b039485166001600160a01b0319918216179091556005805493851693821693909317909255600680549190931691161790556100eb565b3390565b611f8f806100fa6000396000f3fe608060405234801561001057600080fd5b50600436106101425760003560e01c806370a08231116100b8578063bc33657a1161007c578063bc33657a14610416578063d92099841461041e578063dd62ed3e1461044a578063f2fde38b14610478578063f8a8fd6d1461049e578063fec10280146104a657610142565b806370a08231146103b4578063715018a6146103da5780638da5cb5b146103e2578063956138751461040657806395d89b411461040e57610142565b80633098b9541161010a5780633098b954146102ea578063313ce5671461030d57806349eba8f71461032b57806355ee964e14610333578063578b107b1461035c57806360477a881461038257610142565b806301d74a551461014757806306fdde03146101815780630828637a146101fe57806311038e461461022a57806318160ddd146102e2575b600080fd5b61017f6004803603608081101561015d57600080fd5b508035906020810135906001600160a01b0360408201351690606001356104c3565b005b610189610a4d565b6040805160208082528351818301528351919283929083019185019080838360005b838110156101c35781810151838201526020016101ab565b50505050905090810190601f1680156101f05780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b61017f6004803603604081101561021457600080fd5b506001600160a01b038135169060200135610a6e565b6102d06004803603602081101561024057600080fd5b81019060208101813564010000000081111561025b57600080fd5b82018360208201111561026d57600080fd5b8035906020019184600183028401116401000000008311171561028f57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610b91945050505050565b60408051918252519081900360200190f35b6102d0610c77565b61017f6004803603604081101561030057600080fd5b5080359060200135610c7d565b610315610d62565b6040805160ff9092168252519081900360200190f35b6102d0610d67565b6102d06004803603606081101561034957600080fd5b5080359060208101359060400135610dbe565b61017f6004803603602081101561037257600080fd5b50356001600160a01b0316610f45565b61017f6004803603606081101561039857600080fd5b50803590602081013590604001356001600160a01b0316610fd1565b6102d0600480360360208110156103ca57600080fd5b50356001600160a01b031661153e565b61017f611550565b6103ea611604565b604080516001600160a01b039092168252519081900360200190f35b6103ea611613565b610189611622565b6103ea611641565b61017f6004803603604081101561043457600080fd5b506001600160a01b038135169060200135611650565b6102d06004803603604081101561046057600080fd5b506001600160a01b0381358116916020013516611811565b61017f6004803603602081101561048e57600080fd5b50356001600160a01b031661182e565b6102d0611939565b61017f600480360360208110156104bc57600080fd5b503561193f565b6001600160a01b03821661050d576040805162461bcd60e51b815260206004820152600c60248201526b4e554c4c204144445245535360a01b604482015290519081900360640190fd5b6007546001600160a01b0316610558576040805162461bcd60e51b815260206004820152600b60248201526a424144204144445245535360a81b604482015290519081900360640190fd5b600954604080516323b872dd60e01b81523360048201523060248201526044810187905290516001600160a01b03909216916323b872dd916064808201926020929091908290030181600087803b1580156105b257600080fd5b505af11580156105c6573d6000803e3d6000fd5b505050506040513d60208110156105dc57600080fd5b50506040805160608181019092526025808252611ec760208301398051602091820120604080516001600160a01b0380881660248084019190915260448084018b90528451808503820181526064909401855283870180516001600160e01b03166001600160e01b031990971696909617865260055485518087018752600a8152691cd95d14995c5d595cdd60b21b818a01526006549651633ea9061160e11b815260048101998a52865194810194909452855195995060009891851697631503aef99791969190951694637d520c22948b948493910191908083838e5b838110156106d25781810151838201526020016106ba565b50505050905090810190601f1680156106ff5780820380516001836020036101000a031916815260200191505b509250505060006040518083038186803b15801561071c57600080fd5b505afa158015610730573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052602081101561075957600080fd5b810190808051604051939291908464010000000082111561077957600080fd5b90830190602082018581111561078e57600080fd5b82516401000000008111828201881017156107a857600080fd5b82525081516020918201929091019080838360005b838110156107d55781810151838201526020016107bd565b50505050905090810190601f1680156108025780820380516001836020036101000a031916815260200191505b50604052505060075461081e91506001600160a01b0316611991565b6040518463ffffffff1660e01b815260040180806020018060200180602001848103845287818151815260200191508051906020019080838360005b8381101561087257818101518382015260200161085a565b50505050905090810190601f16801561089f5780820380516001836020036101000a031916815260200191505b50848103835286518152865160209182019188019080838360005b838110156108d25781810151838201526020016108ba565b50505050905090810190601f1680156108ff5780820380516001836020036101000a031916815260200191505b50848103825285518152855160209182019187019080838360005b8381101561093257818101518382015260200161091a565b50505050905090810190601f16801561095f5780820380516001836020036101000a031916815260200191505b509650505050505050602060405180830381600087803b15801561098257600080fd5b505af1158015610996573d6000803e3d6000fd5b505050506040513d60208110156109ac57600080fd5b505160008181526008602090815260409182902080546001600160a01b03191633178155600181018a905560028101899055600381018790556203078360ec1b60048201556b6164644c697175696469747960a01b6005820155825184815243928101929092528251939450927fd29686626e0fa84b4c56eccda422e2b6cb56e8ed5e7d0688ddfaaa9ff0067358929181900390910190a150505050505050565b60405180604001604052806005815260200164446967695560d81b81525081565b6005546001600160a01b03163314610abb576040805162461bcd60e51b815260206004820152601d6024820152600080516020611f3a833981519152604482015290519081900360640190fd5b6001600160a01b038216610b05576040805162461bcd60e51b815260206004820152600c60248201526b4e554c4c204144445245535360a01b604482015290519081900360640190fd5b600954604080516323b872dd60e01b81526001600160a01b03858116600483015230602483015260448201859052915191909216916323b872dd9160648083019260209291908290030181600087803b158015610b6157600080fd5b505af1158015610b75573d6000803e3d6000fd5b505050506040513d6020811015610b8b57600080fd5b50505050565b6000806060306001600160a01b0316846040518082805190602001908083835b60208310610bd05780518252601f199092019160209182019101610bb1565b6001836020036101000a038019825116818451168082178552505050505050905001915050600060405180830381855afa9150503d8060008114610c30576040519150601f19603f3d011682016040523d82523d6000602084013e610c35565b606091505b5091509150811580610c4657508051155b15610c5657600092505050610c72565b808060200190516020811015610c6b57600080fd5b5051925050505b919050565b60005481565b6005546001600160a01b03163314610cca576040805162461bcd60e51b815260206004820152601d6024820152600080516020611f3a833981519152604482015290519081900360640190fd5b60008281526008602052604090206004810182905560058101546b6164644c697175696469747960a01b1415610d22578054600182015460028301546003840154610d20936001600160a01b0316929190611a85565b505b604080518481526020810184905281517f740b289fa0ffa3e3808a813571dbb7dc63be083427b360ed6e7b834d76d8b364929181900390910190a1505050565b601281565b6005546000906001600160a01b03163314610db7576040805162461bcd60e51b815260206004820152601d6024820152600080516020611f3a833981519152604482015290519081900360640190fd5b5060045490565b600954604080516370a0823160e01b8152306004820152905160009283926001600160a01b03909116916370a0823191602480820192602092909190829003018186803b158015610e0e57600080fd5b505afa158015610e22573d6000803e3d6000fd5b505050506040513d6020811015610e3857600080fd5b5051600954604080516370a0823160e01b81523060048201529051929350859260009289926001600160a01b03909116916370a0823191602480820192602092909190829003018186803b158015610e8f57600080fd5b505afa158015610ea3573d6000803e3d6000fd5b505050506040513d6020811015610eb957600080fd5b5051600054910191508587019080610efc57610ef56103e8610ee9610ee48c8c63ffffffff611c6716565b611cd0565b9063ffffffff611d2116565b9550610f39565b610f3685610f108b8463ffffffff611c6716565b81610f1757fe5b0485610f298b8563ffffffff611c6716565b81610f3057fe5b04611d71565b95505b50505050509392505050565b610f4d611d89565b6003546001600160a01b03908116911614610faf576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b600780546001600160a01b0319166001600160a01b0392909216919091179055565b6001600160a01b03811661101b576040805162461bcd60e51b815260206004820152600c60248201526b4e554c4c204144445245535360a01b604482015290519081900360640190fd5b6007546001600160a01b0316611066576040805162461bcd60e51b815260206004820152600b60248201526a424144204144445245535360a81b604482015290519081900360640190fd5b600954604080516323b872dd60e01b81523360048201523060248201526044810186905290516001600160a01b03909216916323b872dd916064808201926020929091908290030181600087803b1580156110c057600080fd5b505af11580156110d4573d6000803e3d6000fd5b505050506040513d60208110156110ea57600080fd5b5050604080518082018252601d81527f73776170576974686472617728616464726573732c75696e743235362900000060209182015281516001600160a01b0380851660248084019190915260448084018890528551808503820181526064909401865283850180516001600160e01b0316633648266160e21b17815260055487518089018952600a8152691cd95d14995c5d595cdd60b21b818901526006549851633ea9061160e11b8152600481019889528751958101959095528651969860009892871697631503aef99792969190921694637d520c22948b94938493909201918083838e5b838110156111ea5781810151838201526020016111d2565b50505050905090810190601f1680156112175780820380516001836020036101000a031916815260200191505b509250505060006040518083038186803b15801561123457600080fd5b505afa158015611248573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052602081101561127157600080fd5b810190808051604051939291908464010000000082111561129157600080fd5b9083019060208201858111156112a657600080fd5b82516401000000008111828201881017156112c057600080fd5b82525081516020918201929091019080838360005b838110156112ed5781810151838201526020016112d5565b50505050905090810190601f16801561131a5780820380516001836020036101000a031916815260200191505b50604052505060075461133691506001600160a01b0316611991565b6040518463ffffffff1660e01b815260040180806020018060200180602001848103845287818151815260200191508051906020019080838360005b8381101561138a578181015183820152602001611372565b50505050905090810190601f1680156113b75780820380516001836020036101000a031916815260200191505b50848103835286518152865160209182019188019080838360005b838110156113ea5781810151838201526020016113d2565b50505050905090810190601f1680156114175780820380516001836020036101000a031916815260200191505b50848103825285518152855160209182019187019080838360005b8381101561144a578181015183820152602001611432565b50505050905090810190601f1680156114775780820380516001836020036101000a031916815260200191505b509650505050505050602060405180830381600087803b15801561149a57600080fd5b505af11580156114ae573d6000803e3d6000fd5b505050506040513d60208110156114c457600080fd5b50516000818152600860209081526040918290206203078360ec1b60048201556a1cddd85c11195c1bdcda5d60aa1b6005820155825184815243928101929092528251939450927fb1b2b63a76b38c6a6cd31f51715af205031ee4ca6afa123a87f22e103675218a929181900390910190a1505050505050565b60016020526000908152604090205481565b611558611d89565b6003546001600160a01b039081169116146115ba576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b6003546040516000916001600160a01b0316907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3600380546001600160a01b0319169055565b6003546001600160a01b031690565b6006546001600160a01b031681565b6040518060400160405280600381526020016244475560e81b81525081565b6005546001600160a01b031681565b600954604080516370a0823160e01b8152306004820152905183926001600160a01b0316916370a08231916024808301926020929190829003018186803b15801561169a57600080fd5b505afa1580156116ae573d6000803e3d6000fd5b505050506040513d60208110156116c457600080fd5b505110156117035760405162461bcd60e51b8152600401808060200182810382526023815260200180611ea46023913960400191505060405180910390fd5b6005546001600160a01b03163314611750576040805162461bcd60e51b815260206004820152601d6024820152600080516020611f3a833981519152604482015290519081900360640190fd5b6009546040805163a9059cbb60e01b81526001600160a01b038581166004830152602482018590529151919092169163a9059cbb9160448083019260209291908290030181600087803b1580156117a657600080fd5b505af11580156117ba573d6000803e3d6000fd5b505050506040513d60208110156117d057600080fd5b5050604080513381526020810183905281517f144b92d506514ad735da4651fee6aeb598d9eb2a83c4072689216d05a3bd6b66929181900390910190a15050565b600260209081526000928352604080842090915290825290205481565b611836611d89565b6003546001600160a01b03908116911614611898576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b6001600160a01b0381166118dd5760405162461bcd60e51b8152600401808060200182810382526026815260200180611eec6026913960400191505060405180910390fd5b6003546040516001600160a01b038084169216907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a3600380546001600160a01b0319166001600160a01b0392909216919091179055565b60045481565b6005546001600160a01b0316331461198c576040805162461bcd60e51b815260206004820152601d6024820152600080516020611f3a833981519152604482015290519081900360640190fd5b600455565b604080516028808252606082810190935282919060208201818036833701905050905060005b6014811015611a7e5760008160130360080260020a856001600160a01b0316816119dd57fe5b0460f81b9050600060108260f81c60ff16816119f557fe5b0460f81b905060008160f81c6010028360f81c0360f81b9050611a1782611d8d565b858560020281518110611a2657fe5b60200101906001600160f81b031916908160001a905350611a4681611d8d565b858560020260010181518110611a5857fe5b60200101906001600160f81b031916908160001a90535050600190920191506119b79050565b5092915050565b600954604080516370a0823160e01b81523060048201529051600092839287926001600160a01b03909216916370a0823191602480820192602092909190829003018186803b158015611ad757600080fd5b505afa158015611aeb573d6000803e3d6000fd5b505050506040513d6020811015611b0157600080fd5b5051600954604080516370a0823160e01b8152306004820152905193909203935085926000926001600160a01b03909216916370a08231916024808301926020929190829003018186803b158015611b5857600080fd5b505afa158015611b6c573d6000803e3d6000fd5b505050506040513d6020811015611b8257600080fd5b50516000549091508587019080611bc057611bac6103e8610ee9610ee48c8c63ffffffff611c6716565b9550611bbb60006103e8611dbe565b611bd7565b611bd485610f108b8463ffffffff611c6716565b95505b60008611611c165760405162461bcd60e51b8152600401808060200182810382526028815260200180611f126028913960400191505060405180910390fd5b611c208a87611dbe565b604080518a8152602081018a9052815133927f4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f928290030190a25050505050949350505050565b6000811580611c8257505080820282828281611c7f57fe5b04145b611cca576040805162461bcd60e51b815260206004820152601460248201527364732d6d6174682d6d756c2d6f766572666c6f7760601b604482015290519081900360640190fd5b92915050565b60006003821115611d13575080600160028204015b81811015611d0d57809150600281828581611cfc57fe5b040181611d0557fe5b049050611ce5565b50610c72565b8115610c7257506001919050565b80820382811115611cca576040805162461bcd60e51b815260206004820152601560248201527464732d6d6174682d7375622d756e646572666c6f7760581b604482015290519081900360640190fd5b6000818310611d805781611d82565b825b9392505050565b3390565b6000600a60f883901c1015611dad578160f81c60300160f81b9050610c72565b8160f81c60570160f81b9050610c72565b600054611dd1908263ffffffff611e5416565b60009081556001600160a01b038316815260016020526040902054611dfc908263ffffffff611e5416565b6001600160a01b03831660008181526001602090815260408083209490945583518581529351929391927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9281900390910190a35050565b80820182811015611cca576040805162461bcd60e51b815260206004820152601460248201527364732d6d6174682d6164642d6f766572666c6f7760601b604482015290519081900360640190fdfe494e53554646494349454e5420414d4f554e5420494e205357415057495448445241576164644c6971756964697479496e7465726e616c28616464726573732c75696e74323536294f776e61626c653a206e6577206f776e657220697320746865207a65726f2061646472657373556e697377617056323a20494e53554646494349454e545f4c49515549444954595f4d494e5445444f4e4c59204345525441494e20434841494e4c494e4b20434c49454e54000000a26469706673582212202a9a2d1e1e6e1ddd456d95b8a7e71680e6aee1667341178833a81d0584dc2ec564736f6c63430006090033"

// DeployDexPool deploys a new Ethereum contract, binding an instance of DexPool to it.
func DeployDexPool(auth *bind.TransactOpts, backend bind.ContractBackend, _tokenOfPool common.Address, _myContract common.Address, _util common.Address) (common.Address, *types.Transaction, *DexPool, error) {
	parsed, err := abi.JSON(strings.NewReader(DexPoolABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DexPoolBin), backend, _tokenOfPool, _myContract, _util)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DexPool{DexPoolCaller: DexPoolCaller{contract: contract}, DexPoolTransactor: DexPoolTransactor{contract: contract}, DexPoolFilterer: DexPoolFilterer{contract: contract}}, nil
}

// DexPool is an auto generated Go binding around an Ethereum contract.
type DexPool struct {
	DexPoolCaller     // Read-only binding to the contract
	DexPoolTransactor // Write-only binding to the contract
	DexPoolFilterer   // Log filterer for contract events
}

// DexPoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type DexPoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DexPoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DexPoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DexPoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DexPoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DexPoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DexPoolSession struct {
	Contract     *DexPool          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DexPoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DexPoolCallerSession struct {
	Contract *DexPoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// DexPoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DexPoolTransactorSession struct {
	Contract     *DexPoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// DexPoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type DexPoolRaw struct {
	Contract *DexPool // Generic contract binding to access the raw methods on
}

// DexPoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DexPoolCallerRaw struct {
	Contract *DexPoolCaller // Generic read-only contract binding to access the raw methods on
}

// DexPoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DexPoolTransactorRaw struct {
	Contract *DexPoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDexPool creates a new instance of DexPool, bound to a specific deployed contract.
func NewDexPool(address common.Address, backend bind.ContractBackend) (*DexPool, error) {
	contract, err := bindDexPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DexPool{DexPoolCaller: DexPoolCaller{contract: contract}, DexPoolTransactor: DexPoolTransactor{contract: contract}, DexPoolFilterer: DexPoolFilterer{contract: contract}}, nil
}

// NewDexPoolCaller creates a new read-only instance of DexPool, bound to a specific deployed contract.
func NewDexPoolCaller(address common.Address, caller bind.ContractCaller) (*DexPoolCaller, error) {
	contract, err := bindDexPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DexPoolCaller{contract: contract}, nil
}

// NewDexPoolTransactor creates a new write-only instance of DexPool, bound to a specific deployed contract.
func NewDexPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*DexPoolTransactor, error) {
	contract, err := bindDexPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DexPoolTransactor{contract: contract}, nil
}

// NewDexPoolFilterer creates a new log filterer instance of DexPool, bound to a specific deployed contract.
func NewDexPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*DexPoolFilterer, error) {
	contract, err := bindDexPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DexPoolFilterer{contract: contract}, nil
}

// bindDexPool binds a generic wrapper to an already deployed contract.
func bindDexPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DexPoolABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DexPool *DexPoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DexPool.Contract.DexPoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DexPool *DexPoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DexPool.Contract.DexPoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DexPool *DexPoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DexPool.Contract.DexPoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DexPool *DexPoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DexPool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DexPool *DexPoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DexPool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DexPool *DexPoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DexPool.Contract.contract.Transact(opts, method, params...)
}

// GetTest is a free data retrieval call binding the contract method 0x49eba8f7.
//
// Solidity: function _getTest() view returns(uint256)
func (_DexPool *DexPoolCaller) GetTest(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DexPool.contract.Call(opts, &out, "_getTest")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTest is a free data retrieval call binding the contract method 0x49eba8f7.
//
// Solidity: function _getTest() view returns(uint256)
func (_DexPool *DexPoolSession) GetTest() (*big.Int, error) {
	return _DexPool.Contract.GetTest(&_DexPool.CallOpts)
}

// GetTest is a free data retrieval call binding the contract method 0x49eba8f7.
//
// Solidity: function _getTest() view returns(uint256)
func (_DexPool *DexPoolCallerSession) GetTest() (*big.Int, error) {
	return _DexPool.Contract.GetTest(&_DexPool.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_DexPool *DexPoolCaller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DexPool.contract.Call(opts, &out, "allowance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_DexPool *DexPoolSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _DexPool.Contract.Allowance(&_DexPool.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_DexPool *DexPoolCallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _DexPool.Contract.Allowance(&_DexPool.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_DexPool *DexPoolCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DexPool.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_DexPool *DexPoolSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _DexPool.Contract.BalanceOf(&_DexPool.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_DexPool *DexPoolCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _DexPool.Contract.BalanceOf(&_DexPool.CallOpts, arg0)
}

// CalculateLP is a free data retrieval call binding the contract method 0x55ee964e.
//
// Solidity: function calculateLP(uint256 amount0, uint256 amount1, uint256 balancePoolNet2) view returns(uint256 liquidity)
func (_DexPool *DexPoolCaller) CalculateLP(opts *bind.CallOpts, amount0 *big.Int, amount1 *big.Int, balancePoolNet2 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DexPool.contract.Call(opts, &out, "calculateLP", amount0, amount1, balancePoolNet2)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateLP is a free data retrieval call binding the contract method 0x55ee964e.
//
// Solidity: function calculateLP(uint256 amount0, uint256 amount1, uint256 balancePoolNet2) view returns(uint256 liquidity)
func (_DexPool *DexPoolSession) CalculateLP(amount0 *big.Int, amount1 *big.Int, balancePoolNet2 *big.Int) (*big.Int, error) {
	return _DexPool.Contract.CalculateLP(&_DexPool.CallOpts, amount0, amount1, balancePoolNet2)
}

// CalculateLP is a free data retrieval call binding the contract method 0x55ee964e.
//
// Solidity: function calculateLP(uint256 amount0, uint256 amount1, uint256 balancePoolNet2) view returns(uint256 liquidity)
func (_DexPool *DexPoolCallerSession) CalculateLP(amount0 *big.Int, amount1 *big.Int, balancePoolNet2 *big.Int) (*big.Int, error) {
	return _DexPool.Contract.CalculateLP(&_DexPool.CallOpts, amount0, amount1, balancePoolNet2)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_DexPool *DexPoolCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _DexPool.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_DexPool *DexPoolSession) Decimals() (uint8, error) {
	return _DexPool.Contract.Decimals(&_DexPool.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_DexPool *DexPoolCallerSession) Decimals() (uint8, error) {
	return _DexPool.Contract.Decimals(&_DexPool.CallOpts)
}

// LowLevelGet is a free data retrieval call binding the contract method 0x11038e46.
//
// Solidity: function lowLevelGet(bytes b) view returns(bytes32)
func (_DexPool *DexPoolCaller) LowLevelGet(opts *bind.CallOpts, b []byte) ([32]byte, error) {
	var out []interface{}
	err := _DexPool.contract.Call(opts, &out, "lowLevelGet", b)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// LowLevelGet is a free data retrieval call binding the contract method 0x11038e46.
//
// Solidity: function lowLevelGet(bytes b) view returns(bytes32)
func (_DexPool *DexPoolSession) LowLevelGet(b []byte) ([32]byte, error) {
	return _DexPool.Contract.LowLevelGet(&_DexPool.CallOpts, b)
}

// LowLevelGet is a free data retrieval call binding the contract method 0x11038e46.
//
// Solidity: function lowLevelGet(bytes b) view returns(bytes32)
func (_DexPool *DexPoolCallerSession) LowLevelGet(b []byte) ([32]byte, error) {
	return _DexPool.Contract.LowLevelGet(&_DexPool.CallOpts, b)
}

// MyContract is a free data retrieval call binding the contract method 0xbc33657a.
//
// Solidity: function myContract() view returns(address)
func (_DexPool *DexPoolCaller) MyContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DexPool.contract.Call(opts, &out, "myContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MyContract is a free data retrieval call binding the contract method 0xbc33657a.
//
// Solidity: function myContract() view returns(address)
func (_DexPool *DexPoolSession) MyContract() (common.Address, error) {
	return _DexPool.Contract.MyContract(&_DexPool.CallOpts)
}

// MyContract is a free data retrieval call binding the contract method 0xbc33657a.
//
// Solidity: function myContract() view returns(address)
func (_DexPool *DexPoolCallerSession) MyContract() (common.Address, error) {
	return _DexPool.Contract.MyContract(&_DexPool.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_DexPool *DexPoolCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _DexPool.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_DexPool *DexPoolSession) Name() (string, error) {
	return _DexPool.Contract.Name(&_DexPool.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_DexPool *DexPoolCallerSession) Name() (string, error) {
	return _DexPool.Contract.Name(&_DexPool.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DexPool *DexPoolCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DexPool.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DexPool *DexPoolSession) Owner() (common.Address, error) {
	return _DexPool.Contract.Owner(&_DexPool.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DexPool *DexPoolCallerSession) Owner() (common.Address, error) {
	return _DexPool.Contract.Owner(&_DexPool.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_DexPool *DexPoolCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _DexPool.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_DexPool *DexPoolSession) Symbol() (string, error) {
	return _DexPool.Contract.Symbol(&_DexPool.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_DexPool *DexPoolCallerSession) Symbol() (string, error) {
	return _DexPool.Contract.Symbol(&_DexPool.CallOpts)
}

// Test is a free data retrieval call binding the contract method 0xf8a8fd6d.
//
// Solidity: function test() view returns(uint256)
func (_DexPool *DexPoolCaller) Test(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DexPool.contract.Call(opts, &out, "test")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Test is a free data retrieval call binding the contract method 0xf8a8fd6d.
//
// Solidity: function test() view returns(uint256)
func (_DexPool *DexPoolSession) Test() (*big.Int, error) {
	return _DexPool.Contract.Test(&_DexPool.CallOpts)
}

// Test is a free data retrieval call binding the contract method 0xf8a8fd6d.
//
// Solidity: function test() view returns(uint256)
func (_DexPool *DexPoolCallerSession) Test() (*big.Int, error) {
	return _DexPool.Contract.Test(&_DexPool.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_DexPool *DexPoolCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DexPool.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_DexPool *DexPoolSession) TotalSupply() (*big.Int, error) {
	return _DexPool.Contract.TotalSupply(&_DexPool.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_DexPool *DexPoolCallerSession) TotalSupply() (*big.Int, error) {
	return _DexPool.Contract.TotalSupply(&_DexPool.CallOpts)
}

// Util is a free data retrieval call binding the contract method 0x95613875.
//
// Solidity: function util() view returns(address)
func (_DexPool *DexPoolCaller) Util(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DexPool.contract.Call(opts, &out, "util")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Util is a free data retrieval call binding the contract method 0x95613875.
//
// Solidity: function util() view returns(address)
func (_DexPool *DexPoolSession) Util() (common.Address, error) {
	return _DexPool.Contract.Util(&_DexPool.CallOpts)
}

// Util is a free data retrieval call binding the contract method 0x95613875.
//
// Solidity: function util() view returns(address)
func (_DexPool *DexPoolCallerSession) Util() (common.Address, error) {
	return _DexPool.Contract.Util(&_DexPool.CallOpts)
}

// SetTest is a paid mutator transaction binding the contract method 0xfec10280.
//
// Solidity: function _setTest(uint256 val) returns()
func (_DexPool *DexPoolTransactor) SetTest(opts *bind.TransactOpts, val *big.Int) (*types.Transaction, error) {
	return _DexPool.contract.Transact(opts, "_setTest", val)
}

// SetTest is a paid mutator transaction binding the contract method 0xfec10280.
//
// Solidity: function _setTest(uint256 val) returns()
func (_DexPool *DexPoolSession) SetTest(val *big.Int) (*types.Transaction, error) {
	return _DexPool.Contract.SetTest(&_DexPool.TransactOpts, val)
}

// SetTest is a paid mutator transaction binding the contract method 0xfec10280.
//
// Solidity: function _setTest(uint256 val) returns()
func (_DexPool *DexPoolTransactorSession) SetTest(val *big.Int) (*types.Transaction, error) {
	return _DexPool.Contract.SetTest(&_DexPool.TransactOpts, val)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x01d74a55.
//
// Solidity: function addLiquidity(uint256 amountNet1, uint256 amountNet2, address luqidityProviderNet2, uint256 balancePoolNet2) returns()
func (_DexPool *DexPoolTransactor) AddLiquidity(opts *bind.TransactOpts, amountNet1 *big.Int, amountNet2 *big.Int, luqidityProviderNet2 common.Address, balancePoolNet2 *big.Int) (*types.Transaction, error) {
	return _DexPool.contract.Transact(opts, "addLiquidity", amountNet1, amountNet2, luqidityProviderNet2, balancePoolNet2)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x01d74a55.
//
// Solidity: function addLiquidity(uint256 amountNet1, uint256 amountNet2, address luqidityProviderNet2, uint256 balancePoolNet2) returns()
func (_DexPool *DexPoolSession) AddLiquidity(amountNet1 *big.Int, amountNet2 *big.Int, luqidityProviderNet2 common.Address, balancePoolNet2 *big.Int) (*types.Transaction, error) {
	return _DexPool.Contract.AddLiquidity(&_DexPool.TransactOpts, amountNet1, amountNet2, luqidityProviderNet2, balancePoolNet2)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x01d74a55.
//
// Solidity: function addLiquidity(uint256 amountNet1, uint256 amountNet2, address luqidityProviderNet2, uint256 balancePoolNet2) returns()
func (_DexPool *DexPoolTransactorSession) AddLiquidity(amountNet1 *big.Int, amountNet2 *big.Int, luqidityProviderNet2 common.Address, balancePoolNet2 *big.Int) (*types.Transaction, error) {
	return _DexPool.Contract.AddLiquidity(&_DexPool.TransactOpts, amountNet1, amountNet2, luqidityProviderNet2, balancePoolNet2)
}

// AddLiquidityInternal is a paid mutator transaction binding the contract method 0x0828637a.
//
// Solidity: function addLiquidityInternal(address luqidityProvider, uint256 amount) returns()
func (_DexPool *DexPoolTransactor) AddLiquidityInternal(opts *bind.TransactOpts, luqidityProvider common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DexPool.contract.Transact(opts, "addLiquidityInternal", luqidityProvider, amount)
}

// AddLiquidityInternal is a paid mutator transaction binding the contract method 0x0828637a.
//
// Solidity: function addLiquidityInternal(address luqidityProvider, uint256 amount) returns()
func (_DexPool *DexPoolSession) AddLiquidityInternal(luqidityProvider common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DexPool.Contract.AddLiquidityInternal(&_DexPool.TransactOpts, luqidityProvider, amount)
}

// AddLiquidityInternal is a paid mutator transaction binding the contract method 0x0828637a.
//
// Solidity: function addLiquidityInternal(address luqidityProvider, uint256 amount) returns()
func (_DexPool *DexPoolTransactorSession) AddLiquidityInternal(luqidityProvider common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DexPool.Contract.AddLiquidityInternal(&_DexPool.TransactOpts, luqidityProvider, amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DexPool *DexPoolTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DexPool.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DexPool *DexPoolSession) RenounceOwnership() (*types.Transaction, error) {
	return _DexPool.Contract.RenounceOwnership(&_DexPool.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DexPool *DexPoolTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _DexPool.Contract.RenounceOwnership(&_DexPool.TransactOpts)
}

// SetPendingRequestsDone is a paid mutator transaction binding the contract method 0x3098b954.
//
// Solidity: function setPendingRequestsDone(bytes32 requestId, bytes32 tx_fromNet2) returns()
func (_DexPool *DexPoolTransactor) SetPendingRequestsDone(opts *bind.TransactOpts, requestId [32]byte, tx_fromNet2 [32]byte) (*types.Transaction, error) {
	return _DexPool.contract.Transact(opts, "setPendingRequestsDone", requestId, tx_fromNet2)
}

// SetPendingRequestsDone is a paid mutator transaction binding the contract method 0x3098b954.
//
// Solidity: function setPendingRequestsDone(bytes32 requestId, bytes32 tx_fromNet2) returns()
func (_DexPool *DexPoolSession) SetPendingRequestsDone(requestId [32]byte, tx_fromNet2 [32]byte) (*types.Transaction, error) {
	return _DexPool.Contract.SetPendingRequestsDone(&_DexPool.TransactOpts, requestId, tx_fromNet2)
}

// SetPendingRequestsDone is a paid mutator transaction binding the contract method 0x3098b954.
//
// Solidity: function setPendingRequestsDone(bytes32 requestId, bytes32 tx_fromNet2) returns()
func (_DexPool *DexPoolTransactorSession) SetPendingRequestsDone(requestId [32]byte, tx_fromNet2 [32]byte) (*types.Transaction, error) {
	return _DexPool.Contract.SetPendingRequestsDone(&_DexPool.TransactOpts, requestId, tx_fromNet2)
}

// SetSecondPool is a paid mutator transaction binding the contract method 0x578b107b.
//
// Solidity: function setSecondPool(address adr) returns()
func (_DexPool *DexPoolTransactor) SetSecondPool(opts *bind.TransactOpts, adr common.Address) (*types.Transaction, error) {
	return _DexPool.contract.Transact(opts, "setSecondPool", adr)
}

// SetSecondPool is a paid mutator transaction binding the contract method 0x578b107b.
//
// Solidity: function setSecondPool(address adr) returns()
func (_DexPool *DexPoolSession) SetSecondPool(adr common.Address) (*types.Transaction, error) {
	return _DexPool.Contract.SetSecondPool(&_DexPool.TransactOpts, adr)
}

// SetSecondPool is a paid mutator transaction binding the contract method 0x578b107b.
//
// Solidity: function setSecondPool(address adr) returns()
func (_DexPool *DexPoolTransactorSession) SetSecondPool(adr common.Address) (*types.Transaction, error) {
	return _DexPool.Contract.SetSecondPool(&_DexPool.TransactOpts, adr)
}

// SwapDeposit is a paid mutator transaction binding the contract method 0x60477a88.
//
// Solidity: function swapDeposit(uint256 amount1, uint256 amount2, address recipientOnNet2) returns()
func (_DexPool *DexPoolTransactor) SwapDeposit(opts *bind.TransactOpts, amount1 *big.Int, amount2 *big.Int, recipientOnNet2 common.Address) (*types.Transaction, error) {
	return _DexPool.contract.Transact(opts, "swapDeposit", amount1, amount2, recipientOnNet2)
}

// SwapDeposit is a paid mutator transaction binding the contract method 0x60477a88.
//
// Solidity: function swapDeposit(uint256 amount1, uint256 amount2, address recipientOnNet2) returns()
func (_DexPool *DexPoolSession) SwapDeposit(amount1 *big.Int, amount2 *big.Int, recipientOnNet2 common.Address) (*types.Transaction, error) {
	return _DexPool.Contract.SwapDeposit(&_DexPool.TransactOpts, amount1, amount2, recipientOnNet2)
}

// SwapDeposit is a paid mutator transaction binding the contract method 0x60477a88.
//
// Solidity: function swapDeposit(uint256 amount1, uint256 amount2, address recipientOnNet2) returns()
func (_DexPool *DexPoolTransactorSession) SwapDeposit(amount1 *big.Int, amount2 *big.Int, recipientOnNet2 common.Address) (*types.Transaction, error) {
	return _DexPool.Contract.SwapDeposit(&_DexPool.TransactOpts, amount1, amount2, recipientOnNet2)
}

// SwapWithdraw is a paid mutator transaction binding the contract method 0xd9209984.
//
// Solidity: function swapWithdraw(address recipient, uint256 amount) returns()
func (_DexPool *DexPoolTransactor) SwapWithdraw(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DexPool.contract.Transact(opts, "swapWithdraw", recipient, amount)
}

// SwapWithdraw is a paid mutator transaction binding the contract method 0xd9209984.
//
// Solidity: function swapWithdraw(address recipient, uint256 amount) returns()
func (_DexPool *DexPoolSession) SwapWithdraw(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DexPool.Contract.SwapWithdraw(&_DexPool.TransactOpts, recipient, amount)
}

// SwapWithdraw is a paid mutator transaction binding the contract method 0xd9209984.
//
// Solidity: function swapWithdraw(address recipient, uint256 amount) returns()
func (_DexPool *DexPoolTransactorSession) SwapWithdraw(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DexPool.Contract.SwapWithdraw(&_DexPool.TransactOpts, recipient, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DexPool *DexPoolTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _DexPool.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DexPool *DexPoolSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DexPool.Contract.TransferOwnership(&_DexPool.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DexPool *DexPoolTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DexPool.Contract.TransferOwnership(&_DexPool.TransactOpts, newOwner)
}

// DexPoolAddLiquidityIterator is returned from FilterAddLiquidity and is used to iterate over the raw logs and unpacked data for AddLiquidity events raised by the DexPool contract.
type DexPoolAddLiquidityIterator struct {
	Event *DexPoolAddLiquidity // Event containing the contract specifics and raw log

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
func (it *DexPoolAddLiquidityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DexPoolAddLiquidity)
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
		it.Event = new(DexPoolAddLiquidity)
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
func (it *DexPoolAddLiquidityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DexPoolAddLiquidityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DexPoolAddLiquidity represents a AddLiquidity event raised by the DexPool contract.
type DexPoolAddLiquidity struct {
	RequestId [32]byte
	CurBlock  *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAddLiquidity is a free log retrieval operation binding the contract event 0xd29686626e0fa84b4c56eccda422e2b6cb56e8ed5e7d0688ddfaaa9ff0067358.
//
// Solidity: event AddLiquidity(bytes32 requestId, uint256 curBlock)
func (_DexPool *DexPoolFilterer) FilterAddLiquidity(opts *bind.FilterOpts) (*DexPoolAddLiquidityIterator, error) {

	logs, sub, err := _DexPool.contract.FilterLogs(opts, "AddLiquidity")
	if err != nil {
		return nil, err
	}
	return &DexPoolAddLiquidityIterator{contract: _DexPool.contract, event: "AddLiquidity", logs: logs, sub: sub}, nil
}

// WatchAddLiquidity is a free log subscription operation binding the contract event 0xd29686626e0fa84b4c56eccda422e2b6cb56e8ed5e7d0688ddfaaa9ff0067358.
//
// Solidity: event AddLiquidity(bytes32 requestId, uint256 curBlock)
func (_DexPool *DexPoolFilterer) WatchAddLiquidity(opts *bind.WatchOpts, sink chan<- *DexPoolAddLiquidity) (event.Subscription, error) {

	logs, sub, err := _DexPool.contract.WatchLogs(opts, "AddLiquidity")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DexPoolAddLiquidity)
				if err := _DexPool.contract.UnpackLog(event, "AddLiquidity", log); err != nil {
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

// ParseAddLiquidity is a log parse operation binding the contract event 0xd29686626e0fa84b4c56eccda422e2b6cb56e8ed5e7d0688ddfaaa9ff0067358.
//
// Solidity: event AddLiquidity(bytes32 requestId, uint256 curBlock)
func (_DexPool *DexPoolFilterer) ParseAddLiquidity(log types.Log) (*DexPoolAddLiquidity, error) {
	event := new(DexPoolAddLiquidity)
	if err := _DexPool.contract.UnpackLog(event, "AddLiquidity", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DexPoolApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the DexPool contract.
type DexPoolApprovalIterator struct {
	Event *DexPoolApproval // Event containing the contract specifics and raw log

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
func (it *DexPoolApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DexPoolApproval)
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
		it.Event = new(DexPoolApproval)
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
func (it *DexPoolApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DexPoolApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DexPoolApproval represents a Approval event raised by the DexPool contract.
type DexPoolApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_DexPool *DexPoolFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*DexPoolApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _DexPool.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &DexPoolApprovalIterator{contract: _DexPool.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_DexPool *DexPoolFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *DexPoolApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _DexPool.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DexPoolApproval)
				if err := _DexPool.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_DexPool *DexPoolFilterer) ParseApproval(log types.Log) (*DexPoolApproval, error) {
	event := new(DexPoolApproval)
	if err := _DexPool.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DexPoolMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the DexPool contract.
type DexPoolMintIterator struct {
	Event *DexPoolMint // Event containing the contract specifics and raw log

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
func (it *DexPoolMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DexPoolMint)
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
		it.Event = new(DexPoolMint)
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
func (it *DexPoolMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DexPoolMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DexPoolMint represents a Mint event raised by the DexPool contract.
type DexPoolMint struct {
	Sender  common.Address
	Amount0 *big.Int
	Amount1 *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f.
//
// Solidity: event Mint(address indexed sender, uint256 amount0, uint256 amount1)
func (_DexPool *DexPoolFilterer) FilterMint(opts *bind.FilterOpts, sender []common.Address) (*DexPoolMintIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _DexPool.contract.FilterLogs(opts, "Mint", senderRule)
	if err != nil {
		return nil, err
	}
	return &DexPoolMintIterator{contract: _DexPool.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f.
//
// Solidity: event Mint(address indexed sender, uint256 amount0, uint256 amount1)
func (_DexPool *DexPoolFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *DexPoolMint, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _DexPool.contract.WatchLogs(opts, "Mint", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DexPoolMint)
				if err := _DexPool.contract.UnpackLog(event, "Mint", log); err != nil {
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

// ParseMint is a log parse operation binding the contract event 0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f.
//
// Solidity: event Mint(address indexed sender, uint256 amount0, uint256 amount1)
func (_DexPool *DexPoolFilterer) ParseMint(log types.Log) (*DexPoolMint, error) {
	event := new(DexPoolMint)
	if err := _DexPool.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DexPoolOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the DexPool contract.
type DexPoolOwnershipTransferredIterator struct {
	Event *DexPoolOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *DexPoolOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DexPoolOwnershipTransferred)
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
		it.Event = new(DexPoolOwnershipTransferred)
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
func (it *DexPoolOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DexPoolOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DexPoolOwnershipTransferred represents a OwnershipTransferred event raised by the DexPool contract.
type DexPoolOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DexPool *DexPoolFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*DexPoolOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DexPool.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DexPoolOwnershipTransferredIterator{contract: _DexPool.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DexPool *DexPoolFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DexPoolOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DexPool.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DexPoolOwnershipTransferred)
				if err := _DexPool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_DexPool *DexPoolFilterer) ParseOwnershipTransferred(log types.Log) (*DexPoolOwnershipTransferred, error) {
	event := new(DexPoolOwnershipTransferred)
	if err := _DexPool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DexPoolReceiveIterator is returned from FilterReceive and is used to iterate over the raw logs and unpacked data for Receive events raised by the DexPool contract.
type DexPoolReceiveIterator struct {
	Event *DexPoolReceive // Event containing the contract specifics and raw log

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
func (it *DexPoolReceiveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DexPoolReceive)
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
		it.Event = new(DexPoolReceive)
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
func (it *DexPoolReceiveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DexPoolReceiveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DexPoolReceive represents a Receive event raised by the DexPool contract.
type DexPoolReceive struct {
	Success bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterReceive is a free log retrieval operation binding the contract event 0xae6da96d2ad064f22619554707f11e8f4088acdbeaa51240a87ebd4285fdd9ba.
//
// Solidity: event Receive(bool success)
func (_DexPool *DexPoolFilterer) FilterReceive(opts *bind.FilterOpts) (*DexPoolReceiveIterator, error) {

	logs, sub, err := _DexPool.contract.FilterLogs(opts, "Receive")
	if err != nil {
		return nil, err
	}
	return &DexPoolReceiveIterator{contract: _DexPool.contract, event: "Receive", logs: logs, sub: sub}, nil
}

// WatchReceive is a free log subscription operation binding the contract event 0xae6da96d2ad064f22619554707f11e8f4088acdbeaa51240a87ebd4285fdd9ba.
//
// Solidity: event Receive(bool success)
func (_DexPool *DexPoolFilterer) WatchReceive(opts *bind.WatchOpts, sink chan<- *DexPoolReceive) (event.Subscription, error) {

	logs, sub, err := _DexPool.contract.WatchLogs(opts, "Receive")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DexPoolReceive)
				if err := _DexPool.contract.UnpackLog(event, "Receive", log); err != nil {
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

// ParseReceive is a log parse operation binding the contract event 0xae6da96d2ad064f22619554707f11e8f4088acdbeaa51240a87ebd4285fdd9ba.
//
// Solidity: event Receive(bool success)
func (_DexPool *DexPoolFilterer) ParseReceive(log types.Log) (*DexPoolReceive, error) {
	event := new(DexPoolReceive)
	if err := _DexPool.contract.UnpackLog(event, "Receive", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DexPoolSwapDepositIterator is returned from FilterSwapDeposit and is used to iterate over the raw logs and unpacked data for SwapDeposit events raised by the DexPool contract.
type DexPoolSwapDepositIterator struct {
	Event *DexPoolSwapDeposit // Event containing the contract specifics and raw log

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
func (it *DexPoolSwapDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DexPoolSwapDeposit)
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
		it.Event = new(DexPoolSwapDeposit)
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
func (it *DexPoolSwapDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DexPoolSwapDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DexPoolSwapDeposit represents a SwapDeposit event raised by the DexPool contract.
type DexPoolSwapDeposit struct {
	RequestId [32]byte
	CurBlock  *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSwapDeposit is a free log retrieval operation binding the contract event 0xb1b2b63a76b38c6a6cd31f51715af205031ee4ca6afa123a87f22e103675218a.
//
// Solidity: event SwapDeposit(bytes32 requestId, uint256 curBlock)
func (_DexPool *DexPoolFilterer) FilterSwapDeposit(opts *bind.FilterOpts) (*DexPoolSwapDepositIterator, error) {

	logs, sub, err := _DexPool.contract.FilterLogs(opts, "SwapDeposit")
	if err != nil {
		return nil, err
	}
	return &DexPoolSwapDepositIterator{contract: _DexPool.contract, event: "SwapDeposit", logs: logs, sub: sub}, nil
}

// WatchSwapDeposit is a free log subscription operation binding the contract event 0xb1b2b63a76b38c6a6cd31f51715af205031ee4ca6afa123a87f22e103675218a.
//
// Solidity: event SwapDeposit(bytes32 requestId, uint256 curBlock)
func (_DexPool *DexPoolFilterer) WatchSwapDeposit(opts *bind.WatchOpts, sink chan<- *DexPoolSwapDeposit) (event.Subscription, error) {

	logs, sub, err := _DexPool.contract.WatchLogs(opts, "SwapDeposit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DexPoolSwapDeposit)
				if err := _DexPool.contract.UnpackLog(event, "SwapDeposit", log); err != nil {
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

// ParseSwapDeposit is a log parse operation binding the contract event 0xb1b2b63a76b38c6a6cd31f51715af205031ee4ca6afa123a87f22e103675218a.
//
// Solidity: event SwapDeposit(bytes32 requestId, uint256 curBlock)
func (_DexPool *DexPoolFilterer) ParseSwapDeposit(log types.Log) (*DexPoolSwapDeposit, error) {
	event := new(DexPoolSwapDeposit)
	if err := _DexPool.contract.UnpackLog(event, "SwapDeposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DexPoolSwapWithdrawIterator is returned from FilterSwapWithdraw and is used to iterate over the raw logs and unpacked data for SwapWithdraw events raised by the DexPool contract.
type DexPoolSwapWithdrawIterator struct {
	Event *DexPoolSwapWithdraw // Event containing the contract specifics and raw log

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
func (it *DexPoolSwapWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DexPoolSwapWithdraw)
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
		it.Event = new(DexPoolSwapWithdraw)
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
func (it *DexPoolSwapWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DexPoolSwapWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DexPoolSwapWithdraw represents a SwapWithdraw event raised by the DexPool contract.
type DexPoolSwapWithdraw struct {
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSwapWithdraw is a free log retrieval operation binding the contract event 0x144b92d506514ad735da4651fee6aeb598d9eb2a83c4072689216d05a3bd6b66.
//
// Solidity: event SwapWithdraw(address recipient, uint256 amount)
func (_DexPool *DexPoolFilterer) FilterSwapWithdraw(opts *bind.FilterOpts) (*DexPoolSwapWithdrawIterator, error) {

	logs, sub, err := _DexPool.contract.FilterLogs(opts, "SwapWithdraw")
	if err != nil {
		return nil, err
	}
	return &DexPoolSwapWithdrawIterator{contract: _DexPool.contract, event: "SwapWithdraw", logs: logs, sub: sub}, nil
}

// WatchSwapWithdraw is a free log subscription operation binding the contract event 0x144b92d506514ad735da4651fee6aeb598d9eb2a83c4072689216d05a3bd6b66.
//
// Solidity: event SwapWithdraw(address recipient, uint256 amount)
func (_DexPool *DexPoolFilterer) WatchSwapWithdraw(opts *bind.WatchOpts, sink chan<- *DexPoolSwapWithdraw) (event.Subscription, error) {

	logs, sub, err := _DexPool.contract.WatchLogs(opts, "SwapWithdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DexPoolSwapWithdraw)
				if err := _DexPool.contract.UnpackLog(event, "SwapWithdraw", log); err != nil {
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

// ParseSwapWithdraw is a log parse operation binding the contract event 0x144b92d506514ad735da4651fee6aeb598d9eb2a83c4072689216d05a3bd6b66.
//
// Solidity: event SwapWithdraw(address recipient, uint256 amount)
func (_DexPool *DexPoolFilterer) ParseSwapWithdraw(log types.Log) (*DexPoolSwapWithdraw, error) {
	event := new(DexPoolSwapWithdraw)
	if err := _DexPool.contract.UnpackLog(event, "SwapWithdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DexPoolTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the DexPool contract.
type DexPoolTransferIterator struct {
	Event *DexPoolTransfer // Event containing the contract specifics and raw log

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
func (it *DexPoolTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DexPoolTransfer)
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
		it.Event = new(DexPoolTransfer)
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
func (it *DexPoolTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DexPoolTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DexPoolTransfer represents a Transfer event raised by the DexPool contract.
type DexPoolTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_DexPool *DexPoolFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*DexPoolTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _DexPool.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &DexPoolTransferIterator{contract: _DexPool.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_DexPool *DexPoolFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *DexPoolTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _DexPool.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DexPoolTransfer)
				if err := _DexPool.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_DexPool *DexPoolFilterer) ParseTransfer(log types.Log) (*DexPoolTransfer, error) {
	event := new(DexPoolTransfer)
	if err := _DexPool.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DexPoolTxCompleteBothChainIterator is returned from FilterTxCompleteBothChain and is used to iterate over the raw logs and unpacked data for TxCompleteBothChain events raised by the DexPool contract.
type DexPoolTxCompleteBothChainIterator struct {
	Event *DexPoolTxCompleteBothChain // Event containing the contract specifics and raw log

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
func (it *DexPoolTxCompleteBothChainIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DexPoolTxCompleteBothChain)
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
		it.Event = new(DexPoolTxCompleteBothChain)
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
func (it *DexPoolTxCompleteBothChainIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DexPoolTxCompleteBothChainIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DexPoolTxCompleteBothChain represents a TxCompleteBothChain event raised by the DexPool contract.
type DexPoolTxCompleteBothChain struct {
	RequestId  [32]byte
	TxFromNet2 [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTxCompleteBothChain is a free log retrieval operation binding the contract event 0x740b289fa0ffa3e3808a813571dbb7dc63be083427b360ed6e7b834d76d8b364.
//
// Solidity: event TxCompleteBothChain(bytes32 requestId, bytes32 tx_fromNet2)
func (_DexPool *DexPoolFilterer) FilterTxCompleteBothChain(opts *bind.FilterOpts) (*DexPoolTxCompleteBothChainIterator, error) {

	logs, sub, err := _DexPool.contract.FilterLogs(opts, "TxCompleteBothChain")
	if err != nil {
		return nil, err
	}
	return &DexPoolTxCompleteBothChainIterator{contract: _DexPool.contract, event: "TxCompleteBothChain", logs: logs, sub: sub}, nil
}

// WatchTxCompleteBothChain is a free log subscription operation binding the contract event 0x740b289fa0ffa3e3808a813571dbb7dc63be083427b360ed6e7b834d76d8b364.
//
// Solidity: event TxCompleteBothChain(bytes32 requestId, bytes32 tx_fromNet2)
func (_DexPool *DexPoolFilterer) WatchTxCompleteBothChain(opts *bind.WatchOpts, sink chan<- *DexPoolTxCompleteBothChain) (event.Subscription, error) {

	logs, sub, err := _DexPool.contract.WatchLogs(opts, "TxCompleteBothChain")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DexPoolTxCompleteBothChain)
				if err := _DexPool.contract.UnpackLog(event, "TxCompleteBothChain", log); err != nil {
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

// ParseTxCompleteBothChain is a log parse operation binding the contract event 0x740b289fa0ffa3e3808a813571dbb7dc63be083427b360ed6e7b834d76d8b364.
//
// Solidity: event TxCompleteBothChain(bytes32 requestId, bytes32 tx_fromNet2)
func (_DexPool *DexPoolFilterer) ParseTxCompleteBothChain(log types.Log) (*DexPoolTxCompleteBothChain, error) {
	event := new(DexPoolTxCompleteBothChain)
	if err := _DexPool.contract.UnpackLog(event, "TxCompleteBothChain", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
