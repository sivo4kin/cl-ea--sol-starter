package test

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/sirupsen/logrus"
	"github.com/sivo4kin/ea-starter/wrappers"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"math/big"
	"testing"
	"unsafe"
)

type DexPoolTestSuite struct {
	suite.Suite
	Modifier    Bytes32
	Sim         *backends.SimulatedBackend
	MockDexPool *wrappers.MockDexPool
	//LinkToken		*wrappers.LinkToken
	Oracle           *wrappers.BridgeOracleRequest
	ContractAdresses map[string]common.Address
}

/*// newIdentity returns a go-ethereum abstraction of an ethereum account for
// interacting with contract golang wrappers
func newIdentity(t *testing.T) *bind.TransactOpts {
	key, err := crypto.GenerateKey()
	require.NoError(t, err, "failed to generate ethereum identity")
	return MustNewSimulatedBackendKeyedTransactor(t, key)
}*/

//func SetupTest(t *testing.T) (s *DexPoolTestSuite) {
//	s =  &DexPoolTestSuite{}
//	s.
//	return
//}

//func deployMockDexPoll(s *DexPoolTestSuite) {
//	mockDexpPool, _, rolemodel, err := wrappers.DeployDexPool(
//		s.MC.Admin.IdAccount,
//		s.Sim,
//		common.Address{},
//		2)
//	require.Nil(s.T(), err)
//	s.ContractAdresses["mockDexpPool"] = mockDexpPool
//
//	s.Sim.Commit()
//	s.RolemodelI = rolemodel
//	s.RolemodelMC = s.RolemodelI.(*wrappers.RoleModel)
//}

func Test(t *testing.T) {
	ethereumKey, _ := crypto.GenerateKey()
	auth, err := bind.NewKeyedTransactorWithChainID(ethereumKey, big.NewInt(1337))
	require.NoError(t, err)
	var genesisData = core.GenesisAlloc{auth.From: {Balance: big.NewInt(1000000000)}}
	gasLimit := uint64(100000000)
	backend := backends.NewSimulatedBackend(genesisData, gasLimit)
	require.NotNil(t, backend)
	addr, _, dToken, err := wrappers.DeployDigiUToken(auth, backend)
	require.NoError(t, err)
	require.NotNil(t, addr)
	logrus.Printf("Token deployed 0x%x", addr)
	backend.Commit()
	bal, err := dToken.BalanceOf(&bind.CallOpts{}, common.HexToAddress("0"))
	//_, err = verifier.RandomValueFromVRFProof(nil, proof[:])
	require.NoError(t, err)
	logrus.Printf("balance %d", bal)
	test, _, testContract, err := wrappers.DeployTest(auth, backend)
	require.NoError(t, err)
	require.NotNil(t, test)
	_, err = testContract.SetTest(auth, big.NewInt(30))
	require.NoError(t, err)
	backend.Commit()

	qwe, err := testContract.GetTest(&bind.CallOpts{})
	logrus.Print("GetTest --------------->", qwe)

	data, err := hex.DecodeString("a8cd0a80")
	require.NoError(t, err)
	ewq, err := testContract.LowLevelGet(&bind.CallOpts{}, data)
	require.NoError(t, err)
	ToString(ewq)
	logrus.Print("LowLevelGet ", ewq)
	slice := ewq[:]
	myInt32 := ReadIntNEW32Unsafe(ewq)
	//ReadInt32Unsafe
	logrus.Print("LowLevelGet slice ", myInt32)

	//myInt := binary.BigEndian.Uint64(slice)

	//buf := bytes.NewBuffer(slice) // b is []byte
	//myfirstint, err := binary.ReadVarint(buf)
	//anotherint, err := binary.ReadUvarint(buf)

	var pi float64
	//b := []byte{0x18, 0x2d, 0x44, 0x54, 0xfb, 0x21, 0x09, 0x40}
	buf := bytes.NewReader(slice)
	err = binary.Read(buf, binary.LittleEndian, &pi)
	require.NoError(t, err)
	logrus.Print(pi)
	logrus.Print("LowLevelGet ", pi)

}

func readInt32(b []byte) int32 {
	// equivalnt of return int32(binary.LittleEndian.Uint32(b))
	return int32(uint32(b[0]) | uint32(b[1])<<8 | uint32(b[2])<<16 | uint32(b[3])<<24)
}

// this is much faster and more efficient, however it won't work on appengine
// since it doesn't have the unsafe package.
// Also this would blow up silently if len(b) < 4.
func ReadInt32Unsafe(b []byte) int32 {
	return *(*int32)(unsafe.Pointer(&b[0]))
}

func ReadIntNEW32Unsafe(b [32]byte) int32 {
	return *(*int32)(unsafe.Pointer(&b[0]))
}

func read_int32(data []byte) (ret int32) {
	buf := bytes.NewBuffer(data)
	binary.Read(buf, binary.LittleEndian, &ret)
	return
}
