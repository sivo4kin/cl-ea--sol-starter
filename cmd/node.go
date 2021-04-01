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
	"github.com/sivo4kin/ea-starter/adapters"
	common2 "github.com/sivo4kin/ea-starter/common"
	"github.com/sivo4kin/ea-starter/config"
	"github.com/sivo4kin/ea-starter/libp2p/dht"
	"net/http"
	"os"
	"strconv"
)

type Node struct {
	Config            config.AppConfig
	Ctx               context.Context
	Router            *mux.Router
	Server            bridges.Server
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

	n.pKey, err = common2.ToECDSAFromHex(os.Getenv("SK1"))
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

func (n Node) NewBridge() (srv *bridges.Server) {
	var bridgesList []bridges.Bridge
	ad, err := adapters.NewDBridge(n.EthClient_1, "Health Chain 1", "1", common2.HealthFirst)
	if err != nil {
		logrus.Fatal(err)
		return
	}

	ad2, err := adapters.NewDBridge(n.EthClient_2, "Health Chain 2", "2", common2.HealthSecond)
	if err != nil {
		logrus.Fatal(err)
		return
	}

	ad3, err := adapters.NewDBridge(n.EthClient_1, "SetMockPoolTestRequest", "test", common2.SetMockPoolTestRequest)
	if err != nil {
		logrus.Fatal(err)
		return
	}

	ad4, err := adapters.NewDBridge(n.EthClient_1, "ChainlinkData", "control", common2.ChainlinkData)

	if err != nil {
		logrus.Fatal(err)
		return
	}

	bridgesList = append(bridgesList, ad)
	bridgesList = append(bridgesList, ad2)
	bridgesList = append(bridgesList, ad3)
	bridgesList = append(bridgesList, ad4)
	srv = bridges.NewServer(bridgesList...)
	return
}

func main() {
	err := NewNode()
	if err != nil {
		logrus.Errorf(" NewNode %v", err)
	}
}
