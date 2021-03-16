package test

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/sirupsen/logrus"
	"github.com/sivo4kin/digiu-cross-chain/build/bindings"
	"github.com/sivo4kin/digiu-cross-chain/accounts/abi/bind"
	"github.com/stretchr/testify/require"
	"math/big"
	"testing"
)

func Test(t *testing.T){
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
