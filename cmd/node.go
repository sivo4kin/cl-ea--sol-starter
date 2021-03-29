package main

import (
	"context"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gorilla/mux"
	"github.com/linkpoolio/bridges"
	"github.com/multiformats/go-multiaddr"
	"github.com/sirupsen/logrus"
	"github.com/sivo4kin/ea-starter/adapters/chainlink-integration"
	"github.com/sivo4kin/ea-starter/adapters/linkpoolio_bridges/getblock"
	"github.com/sivo4kin/ea-starter/config"
	"github.com/sivo4kin/ea-starter/libp2p/dht"
	"github.com/sivo4kin/ea-starter/libp2p/knockingtls"
	myrpc "github.com/sivo4kin/ea-starter/libp2p/myrpc"
	"net/http"
	"os"
)

type Node struct {
	Config            config.AppConfig
	EthClient         *ethclient.Client
	Ctx               context.Context
	Srv               chainlink_integration.Chainlink
	Router            *mux.Router
	Server            bridges.Server
	SecurePeer        knockingtls.СoncretePeer
	DiscoveryPeers    addrList
	CurrentRendezvous string
}

type addrList []multiaddr.Multiaddr

func NewNode() (n *Node, err error) {
	dir, err := os.Getwd()
	if err != nil {
		logrus.Warn(err)
	}
	logrus.Printf("started in directory %s", dir)

	err = config.LoadConfig(".")
	if err != nil {
		return
	}
	ethClient, err := ethclient.Dial(config.Config.INFURA_URL)
	if err != nil {
		return
	}
	n = &Node{
		Config:            config.Config,
		EthClient:         ethClient,
		Ctx:               context.Background(),
		CurrentRendezvous: "FirstRun",
	}
	server := n.NewBridge()
	n.Server = *server
	go n.Server.Start(n.Config.PORT)
	n.Router = NewRouter()

	go func() {
		err = dht.NewDHTBootPeer(dir + config.Config.ECDSA_KEY)
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

func (n Node) strtKnokinkTLS() {
	go knockingtls.RunKnokingTLSPeer()
}

func (n Node) NewPeer() (err error) {
	h, err := myrpc.NewRSAKeysHost(n.Ctx, "./keys/srv1-rsa.key", 3456)
	logrus.Printf("Host ID: %s", h.ID().Pretty())
	logrus.Printf("Connect to me on:")
	for _, addr := range h.Addrs() {
		logrus.Printf("  %s/p2p/%s", addr, h.ID().Pretty())
	}
	dht, err := myrpc.NewDHT(n.Ctx, h, n.DiscoveryPeers)
	if err != nil {
		return
	}
	go myrpc.Discover(n.Ctx, h, dht, n.CurrentRendezvous)
	return
}

func (n Node) NewBridge() (srv *bridges.Server) {
	var bridgesList []bridges.Bridge
	ad, err := getblock.NewTicker(n.Config)
	if err != nil {
		logrus.Fatal(err)
		return
	}
	bridgesList = append(bridgesList, ad)
	srv = bridges.NewServer(bridgesList...)
	return
}

func NewRouter() (router *mux.Router) {
	router = chainlink_integration.NewRouter()
	return
}

func main() {
	//logrus.WithField("port", port).Info("Starting the bridge server")
	//logrus.Fatal(http.ListenAndServe(logrus.Sprintf(":%d", port), s.Mux()))
	node, err := NewNode()
	if err != nil {
		logrus.Errorf("%v", err)
		panic(err)
	}
	logrus.Print(node)

}
