package main

import (
	"context"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gorilla/mux"
	"github.com/linkpoolio/bridges"
	"github.com/multiformats/go-multiaddr"
	"github.com/sirupsen/logrus"
	"github.com/sivo4kin/ea-starter/adapters/bridge"
	chainlink_integration "github.com/sivo4kin/ea-starter/adapters/chainlink-integration"
	"github.com/sivo4kin/ea-starter/config"
	"github.com/sivo4kin/ea-starter/libp2p/dht"
	"github.com/sivo4kin/ea-starter/libp2p/knockingtls"
	"math/big"
	"net/http"
	"os"
	"strconv"
)

type Node struct {
	Config            config.AppConfig
	Ctx               context.Context
	Srv               chainlink_integration.Chainlink
	Router            *mux.Router
	Server            bridges.Server
	SecurePeer        knockingtls.СoncretePeer
	DiscoveryPeers    addrList
	CurrentRendezvous string
	EthClient_1       *ethclient.Client
	EthClient_2       *ethclient.Client
	BRIDGE_1_ADDRESS  common.Address
	ORACLE_1_ADDRESS  common.Address
	BRIDGE_2_ADDRESS  common.Address
	ORACLE_2_ADDRESS  common.Address
	pKey              *ecdsa.PrivateKey
}

type addrList []multiaddr.Multiaddr

func NewNode() (err error) {

	dir, err := os.Getwd()
	if err != nil {
		logrus.Warn(err)
	}
	logrus.Printf("started in directory %s", dir)

	err = config.LoadConfig(".")
	if err != nil {
		return
	}

	n := &Node{
		Config:            config.Config,
		Ctx:               context.Background(),
		CurrentRendezvous: "FirstRun",
		//BRIDGE_1_ADDRESS:          common.HexToAddress(os.Getenv("BRIDGE_1_ADDRESS")),
		//ORACLE_1_ADDRESS: common.HexToAddress(os.Getenv("ORACLE_1_ADDRESS")),
	}

	err = n.initKey()
	if err != nil {
		return
	}

	err = n.initEthClients()
	if err != nil {
		return
	}

	server := n.NewBridge()
	n.Server = *server
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = n.Config.PORT
	}

	go n.Server.Start(port)

	go func() {
		err = dht.NewDHTBootPeer(dir+config.Config.ECDSA_KEY, n.Config.P2P_PORT)
		if err != nil {
			return
		}
	}()

	err = http.ListenAndServe(":3000", n.Router)
	if err != nil {
		return
	}

	return
}

func (n Node) initEthClients() (err error) {
	n.EthClient_1, err = ethclient.Dial(config.Config.CHAIN_1_URL)
	if err != nil {
		return
	}

	n.EthClient_2, err = ethclient.Dial(config.Config.CHAIN_2_URL)
	if err != nil {
		return
	}
	return
}

func (n Node) strtKnokinkTLS() {
	go knockingtls.RunKnokingTLSPeer()
}

func (n Node) NewBridge() (srv *bridges.Server) {
	var bridgesList []bridges.Bridge
	ad, err := bridge.NewEthHealth(n.EthClient_1, "Health Chain 1", "health1", bridge.HealthFirst)
	if err != nil {
		logrus.Fatal(err)
		return
	}

	ad2, err := bridge.NewEthHealth(n.EthClient_2, "Health Chain 2", "health2", bridge.HealthSecond)
	if err != nil {
		logrus.Fatal(err)
		return
	}
	bridgesList = append(bridgesList, ad)
	bridgesList = append(bridgesList, ad2)
	srv = bridges.NewServer(bridgesList...)
	return
}

func (n Node) initKey() (err error) {
	n.pKey, err = ToECDSAFromHex(os.Getenv("SK"))
	if err != nil {
		return
	}
	return
}

func ToECDSAFromHex(hexString string) (*ecdsa.PrivateKey, error) {
	pk := new(ecdsa.PrivateKey)
	pk.D, _ = new(big.Int).SetString(hexString, 16)
	return pk, nil
}

func main() {
	logrus.Fatalf("NewNode %v", NewNode())
}
