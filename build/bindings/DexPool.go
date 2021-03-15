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

// DexPoolABI is the input ABI used to generate the binding from.
const DexPoolABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenOfPool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_myContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_util\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"curBlock\",\"type\":\"uint256\"}],\"name\":\"AddLiquidity\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"name\":\"Receive\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"curBlock\",\"type\":\"uint256\"}],\"name\":\"SwapDeposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"SwapWithdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"tx_fromNet2\",\"type\":\"bytes32\"}],\"name\":\"TxCompleteBothChain\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"myContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"test\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"util\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"luqidityProvider\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"_addLiquidity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountNet1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountNet2\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"luqidityProviderNet2\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balancePoolNet2\",\"type\":\"uint256\"}],\"name\":\"addLiquidity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount2\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipientOnNet2\",\"type\":\"address\"}],\"name\":\"swapDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"tx_fromNet2\",\"type\":\"bytes32\"}],\"name\":\"setPendingRequestsDone\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"swapWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balancePoolNet2\",\"type\":\"uint256\"}],\"name\":\"calculateLP\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"val\",\"type\":\"uint256\"}],\"name\":\"_setTest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"_getTest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"b\",\"type\":\"bytes\"}],\"name\":\"lowLevelGet\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"adr\",\"type\":\"address\"}],\"name\":\"setSecondPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// DexPoolBin is the compiled bytecode used for deploying new contracts.
var DexPoolBin = "0x608060405234801561001057600080fd5b5060405161206e38038061206e8339818101604052606081101561003357600080fd5b508051602082015160409092015190919060006100576001600160e01b036100e716565b600380546001600160a01b0319166001600160a01b038316908117909155604051919250906000907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a350600980546001600160a01b039485166001600160a01b0319918216179091556005805493851693821693909317909255600680549190931691161790556100eb565b3390565b611f74806100fa6000396000f3fe608060405234801561001057600080fd5b50600436106101425760003560e01c806370a08231116100b8578063bc33657a1161007c578063bc33657a14610416578063d92099841461041e578063dd62ed3e1461044a578063f2fde38b14610478578063f8a8fd6d1461049e578063fec10280146104a657610142565b806370a08231146103b4578063715018a6146103da5780638da5cb5b146103e2578063956138751461040657806395d89b411461040e57610142565b8063313ce5671161010a578063313ce567146102e157806349eba8f7146102ff57806350842bd01461030757806355ee964e14610333578063578b107b1461035c57806360477a881461038257610142565b806301d74a551461014757806306fdde031461018157806311038e46146101fe57806318160ddd146102b65780633098b954146102be575b600080fd5b61017f6004803603608081101561015d57600080fd5b508035906020810135906001600160a01b0360408201351690606001356104c3565b005b610189610a57565b6040805160208082528351818301528351919283929083019185019080838360005b838110156101c35781810151838201526020016101ab565b50505050905090810190601f1680156101f05780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6102a46004803603602081101561021457600080fd5b81019060208101813564010000000081111561022f57600080fd5b82018360208201111561024157600080fd5b8035906020019184600183028401116401000000008311171561026357600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610a78945050505050565b60408051918252519081900360200190f35b6102a4610b5e565b61017f600480360360408110156102d457600080fd5b5080359060200135610b64565b6102e9610c49565b6040805160ff9092168252519081900360200190f35b6102a4610c4e565b61017f6004803603604081101561031d57600080fd5b506001600160a01b038135169060200135610ca5565b6102a46004803603606081101561034957600080fd5b5080359060208101359060400135610dc8565b61017f6004803603602081101561037257600080fd5b50356001600160a01b0316610f4f565b61017f6004803603606081101561039857600080fd5b50803590602081013590604001356001600160a01b0316610fdb565b6102a4600480360360208110156103ca57600080fd5b50356001600160a01b0316611548565b61017f61155a565b6103ea61160e565b604080516001600160a01b039092168252519081900360200190f35b6103ea61161d565b61018961162c565b6103ea61164b565b61017f6004803603604081101561043457600080fd5b506001600160a01b03813516906020013561165a565b6102a46004803603604081101561046057600080fd5b506001600160a01b038135811691602001351661181b565b61017f6004803603602081101561048e57600080fd5b50356001600160a01b0316611838565b6102a4611943565b61017f600480360360208110156104bc57600080fd5b5035611949565b6001600160a01b03821661050d576040805162461bcd60e51b815260206004820152600c60248201526b4e554c4c204144445245535360a01b604482015290519081900360640190fd5b6007546001600160a01b0316610558576040805162461bcd60e51b815260206004820152600b60248201526a424144204144445245535360a81b604482015290519081900360640190fd5b600954604080516323b872dd60e01b81523360048201523060248201526044810187905290516001600160a01b03909216916323b872dd916064808201926020929091908290030181600087803b1580156105b257600080fd5b505af11580156105c6573d6000803e3d6000fd5b505050506040513d60208110156105dc57600080fd5b5050604080518082018252601e81527f5f6164644c697175696469747928616464726573732c75696e7432353629000060209182015281516001600160a01b0380861660248084019190915260448084018990528551808503820181526064909401865283850180516001600160e01b031663050842bd60e41b17815260055487518089018952600a8152691cd95d14995c5d595cdd60b21b818901526006549851633ea9061160e11b8152600481019889528751958101959095528651969860009892871697631503aef99792969190921694637d520c22948b94938493909201918083838e5b838110156106dc5781810151838201526020016106c4565b50505050905090810190601f1680156107095780820380516001836020036101000a031916815260200191505b509250505060006040518083038186803b15801561072657600080fd5b505afa15801561073a573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052602081101561076357600080fd5b810190808051604051939291908464010000000082111561078357600080fd5b90830190602082018581111561079857600080fd5b82516401000000008111828201881017156107b257600080fd5b82525081516020918201929091019080838360005b838110156107df5781810151838201526020016107c7565b50505050905090810190601f16801561080c5780820380516001836020036101000a031916815260200191505b50604052505060075461082891506001600160a01b031661199b565b6040518463ffffffff1660e01b815260040180806020018060200180602001848103845287818151815260200191508051906020019080838360005b8381101561087c578181015183820152602001610864565b50505050905090810190601f1680156108a95780820380516001836020036101000a031916815260200191505b50848103835286518152865160209182019188019080838360005b838110156108dc5781810151838201526020016108c4565b50505050905090810190601f1680156109095780820380516001836020036101000a031916815260200191505b50848103825285518152855160209182019187019080838360005b8381101561093c578181015183820152602001610924565b50505050905090810190601f1680156109695780820380516001836020036101000a031916815260200191505b509650505050505050602060405180830381600087803b15801561098c57600080fd5b505af11580156109a0573d6000803e3d6000fd5b505050506040513d60208110156109b657600080fd5b505160008181526008602090815260409182902080546001600160a01b03191633178155600181018a905560028101899055600381018790556203078360ec1b60048201556b6164644c697175696469747960a01b6005820155825184815243928101929092528251939450927fd29686626e0fa84b4c56eccda422e2b6cb56e8ed5e7d0688ddfaaa9ff0067358929181900390910190a150505050505050565b60405180604001604052806005815260200164446967695560d81b81525081565b6000806060306001600160a01b0316846040518082805190602001908083835b60208310610ab75780518252601f199092019160209182019101610a98565b6001836020036101000a038019825116818451168082178552505050505050905001915050600060405180830381855afa9150503d8060008114610b17576040519150601f19603f3d011682016040523d82523d6000602084013e610b1c565b606091505b5091509150811580610b2d57508051155b15610b3d57600092505050610b59565b808060200190516020811015610b5257600080fd5b5051925050505b919050565b60005481565b6005546001600160a01b03163314610bb1576040805162461bcd60e51b815260206004820152601d6024820152600080516020611f1f833981519152604482015290519081900360640190fd5b60008281526008602052604090206004810182905560058101546b6164644c697175696469747960a01b1415610c09578054600182015460028301546003840154610c07936001600160a01b0316929190611a8f565b505b604080518481526020810184905281517f740b289fa0ffa3e3808a813571dbb7dc63be083427b360ed6e7b834d76d8b364929181900390910190a1505050565b601281565b6005546000906001600160a01b03163314610c9e576040805162461bcd60e51b815260206004820152601d6024820152600080516020611f1f833981519152604482015290519081900360640190fd5b5060045490565b6005546001600160a01b03163314610cf2576040805162461bcd60e51b815260206004820152601d6024820152600080516020611f1f833981519152604482015290519081900360640190fd5b6001600160a01b038216610d3c576040805162461bcd60e51b815260206004820152600c60248201526b4e554c4c204144445245535360a01b604482015290519081900360640190fd5b600954604080516323b872dd60e01b81526001600160a01b03858116600483015230602483015260448201859052915191909216916323b872dd9160648083019260209291908290030181600087803b158015610d9857600080fd5b505af1158015610dac573d6000803e3d6000fd5b505050506040513d6020811015610dc257600080fd5b50505050565b600954604080516370a0823160e01b8152306004820152905160009283926001600160a01b03909116916370a0823191602480820192602092909190829003018186803b158015610e1857600080fd5b505afa158015610e2c573d6000803e3d6000fd5b505050506040513d6020811015610e4257600080fd5b5051600954604080516370a0823160e01b81523060048201529051929350859260009289926001600160a01b03909116916370a0823191602480820192602092909190829003018186803b158015610e9957600080fd5b505afa158015610ead573d6000803e3d6000fd5b505050506040513d6020811015610ec357600080fd5b5051600054910191508587019080610f0657610eff6103e8610ef3610eee8c8c63ffffffff611c7116565b611cda565b9063ffffffff611d2b16565b9550610f43565b610f4085610f1a8b8463ffffffff611c7116565b81610f2157fe5b0485610f338b8563ffffffff611c7116565b81610f3a57fe5b04611d7b565b95505b50505050509392505050565b610f57611d93565b6003546001600160a01b03908116911614610fb9576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b600780546001600160a01b0319166001600160a01b0392909216919091179055565b6001600160a01b038116611025576040805162461bcd60e51b815260206004820152600c60248201526b4e554c4c204144445245535360a01b604482015290519081900360640190fd5b6007546001600160a01b0316611070576040805162461bcd60e51b815260206004820152600b60248201526a424144204144445245535360a81b604482015290519081900360640190fd5b600954604080516323b872dd60e01b81523360048201523060248201526044810186905290516001600160a01b03909216916323b872dd916064808201926020929091908290030181600087803b1580156110ca57600080fd5b505af11580156110de573d6000803e3d6000fd5b505050506040513d60208110156110f457600080fd5b5050604080518082018252601d81527f73776170576974686472617728616464726573732c75696e743235362900000060209182015281516001600160a01b0380851660248084019190915260448084018890528551808503820181526064909401865283850180516001600160e01b0316633648266160e21b17815260055487518089018952600a8152691cd95d14995c5d595cdd60b21b818901526006549851633ea9061160e11b8152600481019889528751958101959095528651969860009892871697631503aef99792969190921694637d520c22948b94938493909201918083838e5b838110156111f45781810151838201526020016111dc565b50505050905090810190601f1680156112215780820380516001836020036101000a031916815260200191505b509250505060006040518083038186803b15801561123e57600080fd5b505afa158015611252573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052602081101561127b57600080fd5b810190808051604051939291908464010000000082111561129b57600080fd5b9083019060208201858111156112b057600080fd5b82516401000000008111828201881017156112ca57600080fd5b82525081516020918201929091019080838360005b838110156112f75781810151838201526020016112df565b50505050905090810190601f1680156113245780820380516001836020036101000a031916815260200191505b50604052505060075461134091506001600160a01b031661199b565b6040518463ffffffff1660e01b815260040180806020018060200180602001848103845287818151815260200191508051906020019080838360005b8381101561139457818101518382015260200161137c565b50505050905090810190601f1680156113c15780820380516001836020036101000a031916815260200191505b50848103835286518152865160209182019188019080838360005b838110156113f45781810151838201526020016113dc565b50505050905090810190601f1680156114215780820380516001836020036101000a031916815260200191505b50848103825285518152855160209182019187019080838360005b8381101561145457818101518382015260200161143c565b50505050905090810190601f1680156114815780820380516001836020036101000a031916815260200191505b509650505050505050602060405180830381600087803b1580156114a457600080fd5b505af11580156114b8573d6000803e3d6000fd5b505050506040513d60208110156114ce57600080fd5b50516000818152600860209081526040918290206203078360ec1b60048201556a1cddd85c11195c1bdcda5d60aa1b6005820155825184815243928101929092528251939450927fb1b2b63a76b38c6a6cd31f51715af205031ee4ca6afa123a87f22e103675218a929181900390910190a1505050505050565b60016020526000908152604090205481565b611562611d93565b6003546001600160a01b039081169116146115c4576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b6003546040516000916001600160a01b0316907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3600380546001600160a01b0319169055565b6003546001600160a01b031690565b6006546001600160a01b031681565b6040518060400160405280600381526020016244475560e81b81525081565b6005546001600160a01b031681565b600954604080516370a0823160e01b8152306004820152905183926001600160a01b0316916370a08231916024808301926020929190829003018186803b1580156116a457600080fd5b505afa1580156116b8573d6000803e3d6000fd5b505050506040513d60208110156116ce57600080fd5b5051101561170d5760405162461bcd60e51b8152600401808060200182810382526023815260200180611eae6023913960400191505060405180910390fd5b6005546001600160a01b0316331461175a576040805162461bcd60e51b815260206004820152601d6024820152600080516020611f1f833981519152604482015290519081900360640190fd5b6009546040805163a9059cbb60e01b81526001600160a01b038581166004830152602482018590529151919092169163a9059cbb9160448083019260209291908290030181600087803b1580156117b057600080fd5b505af11580156117c4573d6000803e3d6000fd5b505050506040513d60208110156117da57600080fd5b5050604080513381526020810183905281517f144b92d506514ad735da4651fee6aeb598d9eb2a83c4072689216d05a3bd6b66929181900390910190a15050565b600260209081526000928352604080842090915290825290205481565b611840611d93565b6003546001600160a01b039081169116146118a2576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b6001600160a01b0381166118e75760405162461bcd60e51b8152600401808060200182810382526026815260200180611ed16026913960400191505060405180910390fd5b6003546040516001600160a01b038084169216907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a3600380546001600160a01b0319166001600160a01b0392909216919091179055565b60045481565b6005546001600160a01b03163314611996576040805162461bcd60e51b815260206004820152601d6024820152600080516020611f1f833981519152604482015290519081900360640190fd5b600455565b604080516028808252606082810190935282919060208201818036833701905050905060005b6014811015611a885760008160130360080260020a856001600160a01b0316816119e757fe5b0460f81b9050600060108260f81c60ff16816119ff57fe5b0460f81b905060008160f81c6010028360f81c0360f81b9050611a2182611d97565b858560020281518110611a3057fe5b60200101906001600160f81b031916908160001a905350611a5081611d97565b858560020260010181518110611a6257fe5b60200101906001600160f81b031916908160001a90535050600190920191506119c19050565b5092915050565b600954604080516370a0823160e01b81523060048201529051600092839287926001600160a01b03909216916370a0823191602480820192602092909190829003018186803b158015611ae157600080fd5b505afa158015611af5573d6000803e3d6000fd5b505050506040513d6020811015611b0b57600080fd5b5051600954604080516370a0823160e01b8152306004820152905193909203935085926000926001600160a01b03909216916370a08231916024808301926020929190829003018186803b158015611b6257600080fd5b505afa158015611b76573d6000803e3d6000fd5b505050506040513d6020811015611b8c57600080fd5b50516000549091508587019080611bca57611bb66103e8610ef3610eee8c8c63ffffffff611c7116565b9550611bc560006103e8611dc8565b611be1565b611bde85610f1a8b8463ffffffff611c7116565b95505b60008611611c205760405162461bcd60e51b8152600401808060200182810382526028815260200180611ef76028913960400191505060405180910390fd5b611c2a8a87611dc8565b604080518a8152602081018a9052815133927f4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f928290030190a25050505050949350505050565b6000811580611c8c57505080820282828281611c8957fe5b04145b611cd4576040805162461bcd60e51b815260206004820152601460248201527364732d6d6174682d6d756c2d6f766572666c6f7760601b604482015290519081900360640190fd5b92915050565b60006003821115611d1d575080600160028204015b81811015611d1757809150600281828581611d0657fe5b040181611d0f57fe5b049050611cef565b50610b59565b8115610b5957506001919050565b80820382811115611cd4576040805162461bcd60e51b815260206004820152601560248201527464732d6d6174682d7375622d756e646572666c6f7760581b604482015290519081900360640190fd5b6000818310611d8a5781611d8c565b825b9392505050565b3390565b6000600a60f883901c1015611db7578160f81c60300160f81b9050610b59565b8160f81c60570160f81b9050610b59565b600054611ddb908263ffffffff611e5e16565b60009081556001600160a01b038316815260016020526040902054611e06908263ffffffff611e5e16565b6001600160a01b03831660008181526001602090815260408083209490945583518581529351929391927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9281900390910190a35050565b80820182811015611cd4576040805162461bcd60e51b815260206004820152601460248201527364732d6d6174682d6164642d6f766572666c6f7760601b604482015290519081900360640190fdfe494e53554646494349454e5420414d4f554e5420494e205357415057495448445241574f776e61626c653a206e6577206f776e657220697320746865207a65726f2061646472657373556e697377617056323a20494e53554646494349454e545f4c49515549444954595f4d494e5445444f4e4c59204345525441494e20434841494e4c494e4b20434c49454e54000000a26469706673582212202752c60b380d665b68e1295936add9f661bda84d978abf700681e4eba2cf38b964736f6c63430006090033"

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

// DeployDexPoolSync deploys a new Ethereum contract and waits for receipt, binding an instance of DexPoolSession to it.
func DeployDexPoolSync(session *bind.TransactSession, backend bind.ContractBackend, _tokenOfPool common.Address, _myContract common.Address, _util common.Address) (*types.Transaction, *types.Receipt, *DexPoolSession, error) {
	parsed, err := abi.JSON(strings.NewReader(DexPoolABI))
	if err != nil {
		return nil, nil, nil, err
	}
	session.Lock()
	address, tx, _, err := bind.DeployContract(session.TransactOpts, parsed, common.FromHex(DexPoolBin), backend, _tokenOfPool, _myContract, _util)
	receipt, err := session.WaitTransaction(tx)
	if err != nil {
		session.Unlock()
		return nil, nil, nil, err
	}
	session.TransactOpts.Nonce.Add(session.TransactOpts.Nonce, big.NewInt(1))
	session.Unlock()
	contractSession, err := NewDexPoolSession(address, backend, session)
	return tx, receipt, contractSession, err
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
	Contract           *DexPool // Generic contract binding to set the session for
	transactionSession *bind.TransactSession
	Address            common.Address
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

func NewDexPoolSession(address common.Address, backend bind.ContractBackend, transactionSession *bind.TransactSession) (*DexPoolSession, error) {
	DexPoolInstance, err := NewDexPool(address, backend)
	if err != nil {
		return nil, err
	}
	return &DexPoolSession{
		Contract:           DexPoolInstance,
		transactionSession: transactionSession,
		Address:            address,
	}, nil
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
func (_DexPool *DexPoolRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
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
func (_DexPool *DexPoolCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
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
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DexPool.contract.Call(opts, out, "_getTest")
	return *ret0, err
}

// GetTest is a free data retrieval call binding the contract method 0x49eba8f7.
//
// Solidity: function _getTest() view returns(uint256)
func (_DexPool *DexPoolSession) GetTest() (*big.Int, error) {
	return _DexPool.Contract.GetTest(_DexPool.transactionSession.CallOpts)
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
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DexPool.contract.Call(opts, out, "allowance", arg0, arg1)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_DexPool *DexPoolSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _DexPool.Contract.Allowance(_DexPool.transactionSession.CallOpts, arg0, arg1)
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
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DexPool.contract.Call(opts, out, "balanceOf", arg0)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_DexPool *DexPoolSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _DexPool.Contract.BalanceOf(_DexPool.transactionSession.CallOpts, arg0)
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
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DexPool.contract.Call(opts, out, "calculateLP", amount0, amount1, balancePoolNet2)
	return *ret0, err
}

// CalculateLP is a free data retrieval call binding the contract method 0x55ee964e.
//
// Solidity: function calculateLP(uint256 amount0, uint256 amount1, uint256 balancePoolNet2) view returns(uint256 liquidity)
func (_DexPool *DexPoolSession) CalculateLP(amount0 *big.Int, amount1 *big.Int, balancePoolNet2 *big.Int) (*big.Int, error) {
	return _DexPool.Contract.CalculateLP(_DexPool.transactionSession.CallOpts, amount0, amount1, balancePoolNet2)
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
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _DexPool.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_DexPool *DexPoolSession) Decimals() (uint8, error) {
	return _DexPool.Contract.Decimals(_DexPool.transactionSession.CallOpts)
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
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _DexPool.contract.Call(opts, out, "lowLevelGet", b)
	return *ret0, err
}

// LowLevelGet is a free data retrieval call binding the contract method 0x11038e46.
//
// Solidity: function lowLevelGet(bytes b) view returns(bytes32)
func (_DexPool *DexPoolSession) LowLevelGet(b []byte) ([32]byte, error) {
	return _DexPool.Contract.LowLevelGet(_DexPool.transactionSession.CallOpts, b)
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
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DexPool.contract.Call(opts, out, "myContract")
	return *ret0, err
}

// MyContract is a free data retrieval call binding the contract method 0xbc33657a.
//
// Solidity: function myContract() view returns(address)
func (_DexPool *DexPoolSession) MyContract() (common.Address, error) {
	return _DexPool.Contract.MyContract(_DexPool.transactionSession.CallOpts)
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
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _DexPool.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_DexPool *DexPoolSession) Name() (string, error) {
	return _DexPool.Contract.Name(_DexPool.transactionSession.CallOpts)
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
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DexPool.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DexPool *DexPoolSession) Owner() (common.Address, error) {
	return _DexPool.Contract.Owner(_DexPool.transactionSession.CallOpts)
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
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _DexPool.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_DexPool *DexPoolSession) Symbol() (string, error) {
	return _DexPool.Contract.Symbol(_DexPool.transactionSession.CallOpts)
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
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DexPool.contract.Call(opts, out, "test")
	return *ret0, err
}

// Test is a free data retrieval call binding the contract method 0xf8a8fd6d.
//
// Solidity: function test() view returns(uint256)
func (_DexPool *DexPoolSession) Test() (*big.Int, error) {
	return _DexPool.Contract.Test(_DexPool.transactionSession.CallOpts)
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
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DexPool.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_DexPool *DexPoolSession) TotalSupply() (*big.Int, error) {
	return _DexPool.Contract.TotalSupply(_DexPool.transactionSession.CallOpts)
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
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DexPool.contract.Call(opts, out, "util")
	return *ret0, err
}

// Util is a free data retrieval call binding the contract method 0x95613875.
//
// Solidity: function util() view returns(address)
func (_DexPool *DexPoolSession) Util() (common.Address, error) {
	return _DexPool.Contract.Util(_DexPool.transactionSession.CallOpts)
}

// Util is a free data retrieval call binding the contract method 0x95613875.
//
// Solidity: function util() view returns(address)
func (_DexPool *DexPoolCallerSession) Util() (common.Address, error) {
	return _DexPool.Contract.Util(&_DexPool.CallOpts)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x50842bd0.
//
// Solidity: function _addLiquidity(address luqidityProvider, uint256 amount) returns()
func (_DexPool *DexPoolTransactor) AddLiquidity(opts *bind.TransactOpts, luqidityProvider common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DexPool.contract.Transact(opts, "_addLiquidity", luqidityProvider, amount)
}

func (_DexPool *DexPoolTransactor) AddLiquidityRawTx(opts *bind.TransactOpts, luqidityProvider common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DexPool.contract.RawTx(opts, "_addLiquidity", luqidityProvider, amount)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x50842bd0.
// Will wait for tx receipt
//
// Solidity: function _addLiquidity(address luqidityProvider, uint256 amount) returns()
func (_DexPool *DexPoolSession) AddLiquidity(luqidityProvider common.Address, amount *big.Int) (*types.Transaction, *types.Receipt, error) {
	_DexPool.transactionSession.Lock()
	tx, err := _DexPool.Contract.AddLiquidity(_DexPool.transactionSession.TransactOpts, luqidityProvider, amount)
	if err != nil {
		_DexPool.transactionSession.Unlock()
		return nil, nil, err
	}
	_DexPool.transactionSession.TransactOpts.Nonce.Add(_DexPool.transactionSession.TransactOpts.Nonce, big.NewInt(1))
	_DexPool.transactionSession.Unlock()
	receipt, err := _DexPool.transactionSession.WaitTransaction(tx)
	return tx, receipt, err
}

// AddLiquidity returns raw transaction bound to the contract method 0x50842bd0.
//
// Solidity: function _addLiquidity(address luqidityProvider, uint256 amount) returns()
func (_DexPool *DexPoolSession) AddLiquidityRawTx(luqidityProvider common.Address, amount *big.Int) (*types.Transaction, error) {
	tx, err := _DexPool.Contract.AddLiquidityRawTx(_DexPool.transactionSession.TransactOpts, luqidityProvider, amount)
	return tx, err
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x50842bd0.
// Will not wait for tx, but put it to ch
//
// Solidity: function _addLiquidity(address luqidityProvider, uint256 amount) returns()
func (_DexPool *DexPoolSession) AddLiquidityAsync(receiptCh chan *types.ReceiptResult, luqidityProvider common.Address, amount *big.Int) (*types.Transaction, error) {
	_DexPool.transactionSession.Lock()
	tx, err := _DexPool.Contract.AddLiquidity(_DexPool.transactionSession.TransactOpts, luqidityProvider, amount)
	if err != nil {
		_DexPool.transactionSession.Unlock()
		return nil, err
	}
	_DexPool.transactionSession.TransactOpts.Nonce.Add(_DexPool.transactionSession.TransactOpts.Nonce, big.NewInt(1))
	_DexPool.transactionSession.Unlock()
	go func() {
		receipt, err := _DexPool.transactionSession.WaitTransaction(tx)
		receiptCh <- &types.ReceiptResult{
			Receipt: *receipt,
			Err:     err,
		}
	}()
	return tx, err
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x50842bd0.
//
// Solidity: function _addLiquidity(address luqidityProvider, uint256 amount) returns()
func (_DexPool *DexPoolTransactorSession) AddLiquidity(wait bool, luqidityProvider common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DexPool.Contract.AddLiquidity(&_DexPool.TransactOpts, luqidityProvider, amount)
}

// SetTest is a paid mutator transaction binding the contract method 0xfec10280.
//
// Solidity: function _setTest(uint256 val) returns()
func (_DexPool *DexPoolTransactor) SetTest(opts *bind.TransactOpts, val *big.Int) (*types.Transaction, error) {
	return _DexPool.contract.Transact(opts, "_setTest", val)
}

func (_DexPool *DexPoolTransactor) SetTestRawTx(opts *bind.TransactOpts, val *big.Int) (*types.Transaction, error) {
	return _DexPool.contract.RawTx(opts, "_setTest", val)
}

// SetTest is a paid mutator transaction binding the contract method 0xfec10280.
// Will wait for tx receipt
//
// Solidity: function _setTest(uint256 val) returns()
func (_DexPool *DexPoolSession) SetTest(val *big.Int) (*types.Transaction, *types.Receipt, error) {
	_DexPool.transactionSession.Lock()
	tx, err := _DexPool.Contract.SetTest(_DexPool.transactionSession.TransactOpts, val)
	if err != nil {
		_DexPool.transactionSession.Unlock()
		return nil, nil, err
	}
	_DexPool.transactionSession.TransactOpts.Nonce.Add(_DexPool.transactionSession.TransactOpts.Nonce, big.NewInt(1))
	_DexPool.transactionSession.Unlock()
	receipt, err := _DexPool.transactionSession.WaitTransaction(tx)
	return tx, receipt, err
}

// SetTest returns raw transaction bound to the contract method 0xfec10280.
//
// Solidity: function _setTest(uint256 val) returns()
func (_DexPool *DexPoolSession) SetTestRawTx(val *big.Int) (*types.Transaction, error) {
	tx, err := _DexPool.Contract.SetTestRawTx(_DexPool.transactionSession.TransactOpts, val)
	return tx, err
}

// SetTest is a paid mutator transaction binding the contract method 0xfec10280.
// Will not wait for tx, but put it to ch
//
// Solidity: function _setTest(uint256 val) returns()
func (_DexPool *DexPoolSession) SetTestAsync(receiptCh chan *types.ReceiptResult, val *big.Int) (*types.Transaction, error) {
	_DexPool.transactionSession.Lock()
	tx, err := _DexPool.Contract.SetTest(_DexPool.transactionSession.TransactOpts, val)
	if err != nil {
		_DexPool.transactionSession.Unlock()
		return nil, err
	}
	_DexPool.transactionSession.TransactOpts.Nonce.Add(_DexPool.transactionSession.TransactOpts.Nonce, big.NewInt(1))
	_DexPool.transactionSession.Unlock()
	go func() {
		receipt, err := _DexPool.transactionSession.WaitTransaction(tx)
		receiptCh <- &types.ReceiptResult{
			Receipt: *receipt,
			Err:     err,
		}
	}()
	return tx, err
}

// SetTest is a paid mutator transaction binding the contract method 0xfec10280.
//
// Solidity: function _setTest(uint256 val) returns()
func (_DexPool *DexPoolTransactorSession) SetTest(wait bool, val *big.Int) (*types.Transaction, error) {
	return _DexPool.Contract.SetTest(&_DexPool.TransactOpts, val)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x01d74a55.
//
// Solidity: function addLiquidity(uint256 amountNet1, uint256 amountNet2, address luqidityProviderNet2, uint256 balancePoolNet2) returns()
func (_DexPool *DexPoolTransactor) AddLiquidity(opts *bind.TransactOpts, amountNet1 *big.Int, amountNet2 *big.Int, luqidityProviderNet2 common.Address, balancePoolNet2 *big.Int) (*types.Transaction, error) {
	return _DexPool.contract.Transact(opts, "addLiquidity", amountNet1, amountNet2, luqidityProviderNet2, balancePoolNet2)
}

func (_DexPool *DexPoolTransactor) AddLiquidityRawTx(opts *bind.TransactOpts, amountNet1 *big.Int, amountNet2 *big.Int, luqidityProviderNet2 common.Address, balancePoolNet2 *big.Int) (*types.Transaction, error) {
	return _DexPool.contract.RawTx(opts, "addLiquidity", amountNet1, amountNet2, luqidityProviderNet2, balancePoolNet2)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x01d74a55.
// Will wait for tx receipt
//
// Solidity: function addLiquidity(uint256 amountNet1, uint256 amountNet2, address luqidityProviderNet2, uint256 balancePoolNet2) returns()
func (_DexPool *DexPoolSession) AddLiquidity(amountNet1 *big.Int, amountNet2 *big.Int, luqidityProviderNet2 common.Address, balancePoolNet2 *big.Int) (*types.Transaction, *types.Receipt, error) {
	_DexPool.transactionSession.Lock()
	tx, err := _DexPool.Contract.AddLiquidity(_DexPool.transactionSession.TransactOpts, amountNet1, amountNet2, luqidityProviderNet2, balancePoolNet2)
	if err != nil {
		_DexPool.transactionSession.Unlock()
		return nil, nil, err
	}
	_DexPool.transactionSession.TransactOpts.Nonce.Add(_DexPool.transactionSession.TransactOpts.Nonce, big.NewInt(1))
	_DexPool.transactionSession.Unlock()
	receipt, err := _DexPool.transactionSession.WaitTransaction(tx)
	return tx, receipt, err
}

// AddLiquidity returns raw transaction bound to the contract method 0x01d74a55.
//
// Solidity: function addLiquidity(uint256 amountNet1, uint256 amountNet2, address luqidityProviderNet2, uint256 balancePoolNet2) returns()
func (_DexPool *DexPoolSession) AddLiquidityRawTx(amountNet1 *big.Int, amountNet2 *big.Int, luqidityProviderNet2 common.Address, balancePoolNet2 *big.Int) (*types.Transaction, error) {
	tx, err := _DexPool.Contract.AddLiquidityRawTx(_DexPool.transactionSession.TransactOpts, amountNet1, amountNet2, luqidityProviderNet2, balancePoolNet2)
	return tx, err
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x01d74a55.
// Will not wait for tx, but put it to ch
//
// Solidity: function addLiquidity(uint256 amountNet1, uint256 amountNet2, address luqidityProviderNet2, uint256 balancePoolNet2) returns()
func (_DexPool *DexPoolSession) AddLiquidityAsync(receiptCh chan *types.ReceiptResult, amountNet1 *big.Int, amountNet2 *big.Int, luqidityProviderNet2 common.Address, balancePoolNet2 *big.Int) (*types.Transaction, error) {
	_DexPool.transactionSession.Lock()
	tx, err := _DexPool.Contract.AddLiquidity(_DexPool.transactionSession.TransactOpts, amountNet1, amountNet2, luqidityProviderNet2, balancePoolNet2)
	if err != nil {
		_DexPool.transactionSession.Unlock()
		return nil, err
	}
	_DexPool.transactionSession.TransactOpts.Nonce.Add(_DexPool.transactionSession.TransactOpts.Nonce, big.NewInt(1))
	_DexPool.transactionSession.Unlock()
	go func() {
		receipt, err := _DexPool.transactionSession.WaitTransaction(tx)
		receiptCh <- &types.ReceiptResult{
			Receipt: *receipt,
			Err:     err,
		}
	}()
	return tx, err
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x01d74a55.
//
// Solidity: function addLiquidity(uint256 amountNet1, uint256 amountNet2, address luqidityProviderNet2, uint256 balancePoolNet2) returns()
func (_DexPool *DexPoolTransactorSession) AddLiquidity(wait bool, amountNet1 *big.Int, amountNet2 *big.Int, luqidityProviderNet2 common.Address, balancePoolNet2 *big.Int) (*types.Transaction, error) {
	return _DexPool.Contract.AddLiquidity(&_DexPool.TransactOpts, amountNet1, amountNet2, luqidityProviderNet2, balancePoolNet2)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DexPool *DexPoolTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DexPool.contract.Transact(opts, "renounceOwnership")
}

func (_DexPool *DexPoolTransactor) RenounceOwnershipRawTx(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DexPool.contract.RawTx(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
// Will wait for tx receipt
//
// Solidity: function renounceOwnership() returns()
func (_DexPool *DexPoolSession) RenounceOwnership() (*types.Transaction, *types.Receipt, error) {
	_DexPool.transactionSession.Lock()
	tx, err := _DexPool.Contract.RenounceOwnership(_DexPool.transactionSession.TransactOpts)
	if err != nil {
		_DexPool.transactionSession.Unlock()
		return nil, nil, err
	}
	_DexPool.transactionSession.TransactOpts.Nonce.Add(_DexPool.transactionSession.TransactOpts.Nonce, big.NewInt(1))
	_DexPool.transactionSession.Unlock()
	receipt, err := _DexPool.transactionSession.WaitTransaction(tx)
	return tx, receipt, err
}

// RenounceOwnership returns raw transaction bound to the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DexPool *DexPoolSession) RenounceOwnershipRawTx() (*types.Transaction, error) {
	tx, err := _DexPool.Contract.RenounceOwnershipRawTx(_DexPool.transactionSession.TransactOpts)
	return tx, err
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
// Will not wait for tx, but put it to ch
//
// Solidity: function renounceOwnership() returns()
func (_DexPool *DexPoolSession) RenounceOwnershipAsync(receiptCh chan *types.ReceiptResult) (*types.Transaction, error) {
	_DexPool.transactionSession.Lock()
	tx, err := _DexPool.Contract.RenounceOwnership(_DexPool.transactionSession.TransactOpts)
	if err != nil {
		_DexPool.transactionSession.Unlock()
		return nil, err
	}
	_DexPool.transactionSession.TransactOpts.Nonce.Add(_DexPool.transactionSession.TransactOpts.Nonce, big.NewInt(1))
	_DexPool.transactionSession.Unlock()
	go func() {
		receipt, err := _DexPool.transactionSession.WaitTransaction(tx)
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
func (_DexPool *DexPoolTransactorSession) RenounceOwnership(wait bool) (*types.Transaction, error) {
	return _DexPool.Contract.RenounceOwnership(&_DexPool.TransactOpts)
}

// SetPendingRequestsDone is a paid mutator transaction binding the contract method 0x3098b954.
//
// Solidity: function setPendingRequestsDone(bytes32 requestId, bytes32 tx_fromNet2) returns()
func (_DexPool *DexPoolTransactor) SetPendingRequestsDone(opts *bind.TransactOpts, requestId [32]byte, tx_fromNet2 [32]byte) (*types.Transaction, error) {
	return _DexPool.contract.Transact(opts, "setPendingRequestsDone", requestId, tx_fromNet2)
}

func (_DexPool *DexPoolTransactor) SetPendingRequestsDoneRawTx(opts *bind.TransactOpts, requestId [32]byte, tx_fromNet2 [32]byte) (*types.Transaction, error) {
	return _DexPool.contract.RawTx(opts, "setPendingRequestsDone", requestId, tx_fromNet2)
}

// SetPendingRequestsDone is a paid mutator transaction binding the contract method 0x3098b954.
// Will wait for tx receipt
//
// Solidity: function setPendingRequestsDone(bytes32 requestId, bytes32 tx_fromNet2) returns()
func (_DexPool *DexPoolSession) SetPendingRequestsDone(requestId [32]byte, tx_fromNet2 [32]byte) (*types.Transaction, *types.Receipt, error) {
	_DexPool.transactionSession.Lock()
	tx, err := _DexPool.Contract.SetPendingRequestsDone(_DexPool.transactionSession.TransactOpts, requestId, tx_fromNet2)
	if err != nil {
		_DexPool.transactionSession.Unlock()
		return nil, nil, err
	}
	_DexPool.transactionSession.TransactOpts.Nonce.Add(_DexPool.transactionSession.TransactOpts.Nonce, big.NewInt(1))
	_DexPool.transactionSession.Unlock()
	receipt, err := _DexPool.transactionSession.WaitTransaction(tx)
	return tx, receipt, err
}

// SetPendingRequestsDone returns raw transaction bound to the contract method 0x3098b954.
//
// Solidity: function setPendingRequestsDone(bytes32 requestId, bytes32 tx_fromNet2) returns()
func (_DexPool *DexPoolSession) SetPendingRequestsDoneRawTx(requestId [32]byte, tx_fromNet2 [32]byte) (*types.Transaction, error) {
	tx, err := _DexPool.Contract.SetPendingRequestsDoneRawTx(_DexPool.transactionSession.TransactOpts, requestId, tx_fromNet2)
	return tx, err
}

// SetPendingRequestsDone is a paid mutator transaction binding the contract method 0x3098b954.
// Will not wait for tx, but put it to ch
//
// Solidity: function setPendingRequestsDone(bytes32 requestId, bytes32 tx_fromNet2) returns()
func (_DexPool *DexPoolSession) SetPendingRequestsDoneAsync(receiptCh chan *types.ReceiptResult, requestId [32]byte, tx_fromNet2 [32]byte) (*types.Transaction, error) {
	_DexPool.transactionSession.Lock()
	tx, err := _DexPool.Contract.SetPendingRequestsDone(_DexPool.transactionSession.TransactOpts, requestId, tx_fromNet2)
	if err != nil {
		_DexPool.transactionSession.Unlock()
		return nil, err
	}
	_DexPool.transactionSession.TransactOpts.Nonce.Add(_DexPool.transactionSession.TransactOpts.Nonce, big.NewInt(1))
	_DexPool.transactionSession.Unlock()
	go func() {
		receipt, err := _DexPool.transactionSession.WaitTransaction(tx)
		receiptCh <- &types.ReceiptResult{
			Receipt: *receipt,
			Err:     err,
		}
	}()
	return tx, err
}

// SetPendingRequestsDone is a paid mutator transaction binding the contract method 0x3098b954.
//
// Solidity: function setPendingRequestsDone(bytes32 requestId, bytes32 tx_fromNet2) returns()
func (_DexPool *DexPoolTransactorSession) SetPendingRequestsDone(wait bool, requestId [32]byte, tx_fromNet2 [32]byte) (*types.Transaction, error) {
	return _DexPool.Contract.SetPendingRequestsDone(&_DexPool.TransactOpts, requestId, tx_fromNet2)
}

// SetSecondPool is a paid mutator transaction binding the contract method 0x578b107b.
//
// Solidity: function setSecondPool(address adr) returns()
func (_DexPool *DexPoolTransactor) SetSecondPool(opts *bind.TransactOpts, adr common.Address) (*types.Transaction, error) {
	return _DexPool.contract.Transact(opts, "setSecondPool", adr)
}

func (_DexPool *DexPoolTransactor) SetSecondPoolRawTx(opts *bind.TransactOpts, adr common.Address) (*types.Transaction, error) {
	return _DexPool.contract.RawTx(opts, "setSecondPool", adr)
}

// SetSecondPool is a paid mutator transaction binding the contract method 0x578b107b.
// Will wait for tx receipt
//
// Solidity: function setSecondPool(address adr) returns()
func (_DexPool *DexPoolSession) SetSecondPool(adr common.Address) (*types.Transaction, *types.Receipt, error) {
	_DexPool.transactionSession.Lock()
	tx, err := _DexPool.Contract.SetSecondPool(_DexPool.transactionSession.TransactOpts, adr)
	if err != nil {
		_DexPool.transactionSession.Unlock()
		return nil, nil, err
	}
	_DexPool.transactionSession.TransactOpts.Nonce.Add(_DexPool.transactionSession.TransactOpts.Nonce, big.NewInt(1))
	_DexPool.transactionSession.Unlock()
	receipt, err := _DexPool.transactionSession.WaitTransaction(tx)
	return tx, receipt, err
}

// SetSecondPool returns raw transaction bound to the contract method 0x578b107b.
//
// Solidity: function setSecondPool(address adr) returns()
func (_DexPool *DexPoolSession) SetSecondPoolRawTx(adr common.Address) (*types.Transaction, error) {
	tx, err := _DexPool.Contract.SetSecondPoolRawTx(_DexPool.transactionSession.TransactOpts, adr)
	return tx, err
}

// SetSecondPool is a paid mutator transaction binding the contract method 0x578b107b.
// Will not wait for tx, but put it to ch
//
// Solidity: function setSecondPool(address adr) returns()
func (_DexPool *DexPoolSession) SetSecondPoolAsync(receiptCh chan *types.ReceiptResult, adr common.Address) (*types.Transaction, error) {
	_DexPool.transactionSession.Lock()
	tx, err := _DexPool.Contract.SetSecondPool(_DexPool.transactionSession.TransactOpts, adr)
	if err != nil {
		_DexPool.transactionSession.Unlock()
		return nil, err
	}
	_DexPool.transactionSession.TransactOpts.Nonce.Add(_DexPool.transactionSession.TransactOpts.Nonce, big.NewInt(1))
	_DexPool.transactionSession.Unlock()
	go func() {
		receipt, err := _DexPool.transactionSession.WaitTransaction(tx)
		receiptCh <- &types.ReceiptResult{
			Receipt: *receipt,
			Err:     err,
		}
	}()
	return tx, err
}

// SetSecondPool is a paid mutator transaction binding the contract method 0x578b107b.
//
// Solidity: function setSecondPool(address adr) returns()
func (_DexPool *DexPoolTransactorSession) SetSecondPool(wait bool, adr common.Address) (*types.Transaction, error) {
	return _DexPool.Contract.SetSecondPool(&_DexPool.TransactOpts, adr)
}

// SwapDeposit is a paid mutator transaction binding the contract method 0x60477a88.
//
// Solidity: function swapDeposit(uint256 amount1, uint256 amount2, address recipientOnNet2) returns()
func (_DexPool *DexPoolTransactor) SwapDeposit(opts *bind.TransactOpts, amount1 *big.Int, amount2 *big.Int, recipientOnNet2 common.Address) (*types.Transaction, error) {
	return _DexPool.contract.Transact(opts, "swapDeposit", amount1, amount2, recipientOnNet2)
}

func (_DexPool *DexPoolTransactor) SwapDepositRawTx(opts *bind.TransactOpts, amount1 *big.Int, amount2 *big.Int, recipientOnNet2 common.Address) (*types.Transaction, error) {
	return _DexPool.contract.RawTx(opts, "swapDeposit", amount1, amount2, recipientOnNet2)
}

// SwapDeposit is a paid mutator transaction binding the contract method 0x60477a88.
// Will wait for tx receipt
//
// Solidity: function swapDeposit(uint256 amount1, uint256 amount2, address recipientOnNet2) returns()
func (_DexPool *DexPoolSession) SwapDeposit(amount1 *big.Int, amount2 *big.Int, recipientOnNet2 common.Address) (*types.Transaction, *types.Receipt, error) {
	_DexPool.transactionSession.Lock()
	tx, err := _DexPool.Contract.SwapDeposit(_DexPool.transactionSession.TransactOpts, amount1, amount2, recipientOnNet2)
	if err != nil {
		_DexPool.transactionSession.Unlock()
		return nil, nil, err
	}
	_DexPool.transactionSession.TransactOpts.Nonce.Add(_DexPool.transactionSession.TransactOpts.Nonce, big.NewInt(1))
	_DexPool.transactionSession.Unlock()
	receipt, err := _DexPool.transactionSession.WaitTransaction(tx)
	return tx, receipt, err
}

// SwapDeposit returns raw transaction bound to the contract method 0x60477a88.
//
// Solidity: function swapDeposit(uint256 amount1, uint256 amount2, address recipientOnNet2) returns()
func (_DexPool *DexPoolSession) SwapDepositRawTx(amount1 *big.Int, amount2 *big.Int, recipientOnNet2 common.Address) (*types.Transaction, error) {
	tx, err := _DexPool.Contract.SwapDepositRawTx(_DexPool.transactionSession.TransactOpts, amount1, amount2, recipientOnNet2)
	return tx, err
}

// SwapDeposit is a paid mutator transaction binding the contract method 0x60477a88.
// Will not wait for tx, but put it to ch
//
// Solidity: function swapDeposit(uint256 amount1, uint256 amount2, address recipientOnNet2) returns()
func (_DexPool *DexPoolSession) SwapDepositAsync(receiptCh chan *types.ReceiptResult, amount1 *big.Int, amount2 *big.Int, recipientOnNet2 common.Address) (*types.Transaction, error) {
	_DexPool.transactionSession.Lock()
	tx, err := _DexPool.Contract.SwapDeposit(_DexPool.transactionSession.TransactOpts, amount1, amount2, recipientOnNet2)
	if err != nil {
		_DexPool.transactionSession.Unlock()
		return nil, err
	}
	_DexPool.transactionSession.TransactOpts.Nonce.Add(_DexPool.transactionSession.TransactOpts.Nonce, big.NewInt(1))
	_DexPool.transactionSession.Unlock()
	go func() {
		receipt, err := _DexPool.transactionSession.WaitTransaction(tx)
		receiptCh <- &types.ReceiptResult{
			Receipt: *receipt,
			Err:     err,
		}
	}()
	return tx, err
}

// SwapDeposit is a paid mutator transaction binding the contract method 0x60477a88.
//
// Solidity: function swapDeposit(uint256 amount1, uint256 amount2, address recipientOnNet2) returns()
func (_DexPool *DexPoolTransactorSession) SwapDeposit(wait bool, amount1 *big.Int, amount2 *big.Int, recipientOnNet2 common.Address) (*types.Transaction, error) {
	return _DexPool.Contract.SwapDeposit(&_DexPool.TransactOpts, amount1, amount2, recipientOnNet2)
}

// SwapWithdraw is a paid mutator transaction binding the contract method 0xd9209984.
//
// Solidity: function swapWithdraw(address recipient, uint256 amount) returns()
func (_DexPool *DexPoolTransactor) SwapWithdraw(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DexPool.contract.Transact(opts, "swapWithdraw", recipient, amount)
}

func (_DexPool *DexPoolTransactor) SwapWithdrawRawTx(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DexPool.contract.RawTx(opts, "swapWithdraw", recipient, amount)
}

// SwapWithdraw is a paid mutator transaction binding the contract method 0xd9209984.
// Will wait for tx receipt
//
// Solidity: function swapWithdraw(address recipient, uint256 amount) returns()
func (_DexPool *DexPoolSession) SwapWithdraw(recipient common.Address, amount *big.Int) (*types.Transaction, *types.Receipt, error) {
	_DexPool.transactionSession.Lock()
	tx, err := _DexPool.Contract.SwapWithdraw(_DexPool.transactionSession.TransactOpts, recipient, amount)
	if err != nil {
		_DexPool.transactionSession.Unlock()
		return nil, nil, err
	}
	_DexPool.transactionSession.TransactOpts.Nonce.Add(_DexPool.transactionSession.TransactOpts.Nonce, big.NewInt(1))
	_DexPool.transactionSession.Unlock()
	receipt, err := _DexPool.transactionSession.WaitTransaction(tx)
	return tx, receipt, err
}

// SwapWithdraw returns raw transaction bound to the contract method 0xd9209984.
//
// Solidity: function swapWithdraw(address recipient, uint256 amount) returns()
func (_DexPool *DexPoolSession) SwapWithdrawRawTx(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	tx, err := _DexPool.Contract.SwapWithdrawRawTx(_DexPool.transactionSession.TransactOpts, recipient, amount)
	return tx, err
}

// SwapWithdraw is a paid mutator transaction binding the contract method 0xd9209984.
// Will not wait for tx, but put it to ch
//
// Solidity: function swapWithdraw(address recipient, uint256 amount) returns()
func (_DexPool *DexPoolSession) SwapWithdrawAsync(receiptCh chan *types.ReceiptResult, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	_DexPool.transactionSession.Lock()
	tx, err := _DexPool.Contract.SwapWithdraw(_DexPool.transactionSession.TransactOpts, recipient, amount)
	if err != nil {
		_DexPool.transactionSession.Unlock()
		return nil, err
	}
	_DexPool.transactionSession.TransactOpts.Nonce.Add(_DexPool.transactionSession.TransactOpts.Nonce, big.NewInt(1))
	_DexPool.transactionSession.Unlock()
	go func() {
		receipt, err := _DexPool.transactionSession.WaitTransaction(tx)
		receiptCh <- &types.ReceiptResult{
			Receipt: *receipt,
			Err:     err,
		}
	}()
	return tx, err
}

// SwapWithdraw is a paid mutator transaction binding the contract method 0xd9209984.
//
// Solidity: function swapWithdraw(address recipient, uint256 amount) returns()
func (_DexPool *DexPoolTransactorSession) SwapWithdraw(wait bool, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DexPool.Contract.SwapWithdraw(&_DexPool.TransactOpts, recipient, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DexPool *DexPoolTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _DexPool.contract.Transact(opts, "transferOwnership", newOwner)
}

func (_DexPool *DexPoolTransactor) TransferOwnershipRawTx(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _DexPool.contract.RawTx(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
// Will wait for tx receipt
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DexPool *DexPoolSession) TransferOwnership(newOwner common.Address) (*types.Transaction, *types.Receipt, error) {
	_DexPool.transactionSession.Lock()
	tx, err := _DexPool.Contract.TransferOwnership(_DexPool.transactionSession.TransactOpts, newOwner)
	if err != nil {
		_DexPool.transactionSession.Unlock()
		return nil, nil, err
	}
	_DexPool.transactionSession.TransactOpts.Nonce.Add(_DexPool.transactionSession.TransactOpts.Nonce, big.NewInt(1))
	_DexPool.transactionSession.Unlock()
	receipt, err := _DexPool.transactionSession.WaitTransaction(tx)
	return tx, receipt, err
}

// TransferOwnership returns raw transaction bound to the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DexPool *DexPoolSession) TransferOwnershipRawTx(newOwner common.Address) (*types.Transaction, error) {
	tx, err := _DexPool.Contract.TransferOwnershipRawTx(_DexPool.transactionSession.TransactOpts, newOwner)
	return tx, err
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
// Will not wait for tx, but put it to ch
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DexPool *DexPoolSession) TransferOwnershipAsync(receiptCh chan *types.ReceiptResult, newOwner common.Address) (*types.Transaction, error) {
	_DexPool.transactionSession.Lock()
	tx, err := _DexPool.Contract.TransferOwnership(_DexPool.transactionSession.TransactOpts, newOwner)
	if err != nil {
		_DexPool.transactionSession.Unlock()
		return nil, err
	}
	_DexPool.transactionSession.TransactOpts.Nonce.Add(_DexPool.transactionSession.TransactOpts.Nonce, big.NewInt(1))
	_DexPool.transactionSession.Unlock()
	go func() {
		receipt, err := _DexPool.transactionSession.WaitTransaction(tx)
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
func (_DexPool *DexPoolTransactorSession) TransferOwnership(wait bool, newOwner common.Address) (*types.Transaction, error) {
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
	return event, nil
}
