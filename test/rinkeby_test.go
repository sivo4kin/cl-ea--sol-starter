package test

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
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
