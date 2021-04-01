package test

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sirupsen/logrus"
	"github.com/sivo4kin/ea-starter/wrappers"
	"github.com/stretchr/testify/require"
	"math/big"
	"testing"
)

func TestRinkebyPool(t *testing.T) {
	conn, err := ethclient.Dial("wss://rinkeby.infura.io/ws/v3/1f520f603c094850aafcb11291818e29")
	require.NoError(t, err)
	// Current rinkeby pool address 0x8C2e2b076ccd2d1654de5A094a8626ADa609b415
	poolContract, err := wrappers.NewDexPool(common.HexToAddress("0x8C2e2b076ccd2d1654de5A094a8626ADa609b415"), conn)
	require.NoError(t, err)

	tx, err := poolContract.SetTest(&bind.TransactOpts{}, big.NewInt(99))
	require.NoError(t, err)
	logrus.Print("TRANSACTION", tx)

	// this is represents of bytes memory out = abi.encodeWithSelector(bytes4(keccak256(bytes('_setTest(uint256)'))), amount);
	data := "0xfec102800000000000000000000000000000000000000000000000008ac7230489e80000"
	qwe, err := poolContract.LowLevelGet(&bind.CallOpts{}, []byte(data))
	require.NoError(t, err)
	logrus.Print("--------------->", qwe)

	qwe, err = poolContract.LowLevelGet(&bind.CallOpts{}, []byte("0x49eba8f7"))
	require.NoError(t, err)

	logrus.Print("--------------->", qwe)
	name, err := poolContract.Name(&bind.CallOpts{})
	require.NoError(t, err)
	myContratcAddress, err := poolContract.MyContract(&bind.CallOpts{})
	require.NoError(t, err)
	logrus.Printf("Pool %s myContract Address: %s", name, myContratcAddress)
	myContract, err := wrappers.NewMyContract(myContratcAddress, conn)
	require.NoError(t, err)
	var job [32]byte
	var jobs [][32]byte
	jobs, err = myContract.GetPermissionJobId(&bind.CallOpts{})
	logrus.Printf("length %d", len(jobs))
	for _, job = range jobs {
		logrus.Printf("QWEE %s", fmt.Sprintf("%x", job))
	}
}

func Test_MockDexPool(t *testing.T) {
	/*	SK1 := "0x194e868506502e5ecae3e5b623801548125a748e6b2da15681312a7cf0283acc"
		SK2 := "0x0da632a0af66bc9748f4fe4e8261facffbaef084ae1c591b1d30889622975735"

		privateKey1 , err := common2.ToECDSAFromHex(SK1)
		require.Nil(t, err)

		privateKey2 , err := common2.ToECDSAFromHex(SK2)
		require.Nil(t, err)*/

	client1, err := ethclient.Dial("wss://rinkeby.infura.io/ws/v3/1f520f603c094850aafcb11291818e29")
	require.NoError(t, err)

	/*	currentBlock, _ := client1.BlockByNumber(context.Background(), nil)
		start := currentBlock.Number().Uint64()*/

	receipt, err := client1.TransactionReceipt(context.Background(), common.HexToHash("0xb737e23134508cddf6282e4a6772d315aaffae1cf3e3a3f1f7db0126ff7f6067"))

	require.Nil(t, err)

	logrus.Printf("RECEIPT %v", receipt)

	logrus.Println("receipt.Status", receipt.Status) // 1

	res, err := GetDexPoolEventByAction(*client1, 0, common.HexToAddress("0xd3bccAA30816Af1682e25cE6880F0Df50793aa4A"))
	require.Nil(t, err)

	logrus.Println("res", res) // ...

	//client2, err := common2.Connect("ws://95.217.104.54:8576")
	//require.NoError(t, err)

	//mock2, err := wrappers.NewMockDexPool(common.HexToAddress("0x5D2c2D8ac778185f118c0a399DF0c41f9ADC8070"), client2)
	//require.Nil(t, err)

	mock1, err := wrappers.NewMockDexPool(common.HexToAddress("0xd3bccAA30816Af1682e25cE6880F0Df50793aa4A"), client1)
	require.Nil(t, err)

	b1, b2, err := mock1.GetPendingRequests(&bind.CallOpts{}, res)
	require.Nil(t, err)

	logrus.Printf("b1 %v", b1)
	logrus.Printf("b2 %v", b2)

	logrus.Print("b1 ", ToHex(b1))
	logrus.Print("b1 ", ToHex(b1))
	logrus.Print("b1 ", common.BytesToHash(b1[:]))

	logrus.Printf("b2 %x", ToHex(b2))

	logrus.Printf(fmt.Sprintf("%x", 99))
}

func GetTxReceipt(c ethclient.Client) (*types.Receipt, error) {
	return c.TransactionReceipt(context.Background(), common.HexToHash("0xb737e23134508cddf6282e4a6772d315aaffae1cf3e3a3f1f7db0126ff7f6067"))
}

func GetDexPoolEventByAction(client ethclient.Client, start uint64, contractAddress common.Address) (res [32]byte, err error) {
	mockFilterer, err := wrappers.NewMockDexPoolFilterer(contractAddress, &client)
	require.Nil(&testing.T{}, err)
	it, err := mockFilterer.FilterRequestSended(&bind.FilterOpts{Start: start, Context: context.Background()})
	require.Nil(&testing.T{}, err)
	for it.Next() {
		logrus.Print("it.Event", it.Event.ReqId)
		if it.Event != nil {
			res = it.Event.ReqId
			return
		}
	}
	return
}
