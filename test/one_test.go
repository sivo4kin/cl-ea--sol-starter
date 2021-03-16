package test

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sirupsen/logrus"
	"github.com/sivo4kin/digiu-cross-chain/accounts/abi/bind"
	"github.com/sivo4kin/digiu-cross-chain/build/bindings"
	"github.com/stretchr/testify/require"
	"math/big"
	"testing"
)

func Test(t *testing.T) {
	ethereumKey, _ := crypto.GenerateKey()
	auth, err := bind.NewKeyedTransactorWithChainID(ethereumKey, big.NewInt(1337))
	require.NoError(t, err)
	var genesisData = core.GenesisAlloc{auth.From: {Balance: big.NewInt(1000000000)}}
	gasLimit := uint64(100000000)
	backend := backends.NewSimulatedBackend(genesisData, gasLimit)
	require.NotNil(t, backend)
	addr, _, dToken, err := bindings.DeployDigiUToken(auth, backend)
	require.NoError(t, err)
	require.NotNil(t, addr)
	logrus.Printf("Token deployed 0x%x", addr)
	backend.Commit()
	bal, err := dToken.BalanceOf(&bind.CallOpts{}, common.HexToAddress("0"))
	//_, err = verifier.RandomValueFromVRFProof(nil, proof[:])
	require.NoError(t, err)
	logrus.Printf("balance %d", bal)
}

func TestRinkebyPool(t *testing.T) {
	conn, err := ethclient.Dial("wss://rinkeby.infura.io/ws/v3/1f520f603c094850aafcb11291818e29")
	require.NoError(t, err)
	// Current rinkeby pool address 0x8C2e2b076ccd2d1654de5A094a8626ADa609b415
	poolContract, err := bindings.NewDexPool(common.HexToAddress("0x8C2e2b076ccd2d1654de5A094a8626ADa609b415"), conn)
	require.NoError(t, err)
	name, err := poolContract.Name(&bind.CallOpts{})
	require.NoError(t, err)
	myContratcAddress, err := poolContract.MyContract(&bind.CallOpts{})
	require.NoError(t, err)
	logrus.Printf("Pool %s myContract Address: %s", name, myContratcAddress)
	myContract, err := bindings.NewMyContract(myContratcAddress, conn)
	require.NoError(t, err)
	var job [32]byte
	var jobs [][32]byte
	jobs, err = myContract.GetPermissionJobId(&bind.CallOpts{})
	logrus.Printf("length %d", len(jobs))
	for _, job = range jobs {
		logrus.Printf("QWEE %s", fmt.Sprintf("%x", job))
	}

}
