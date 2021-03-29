package dht

import (
	"context"
	"flag"
	"fmt"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/routing"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	ddht "github.com/libp2p/go-libp2p-kad-dht/dual"
	"github.com/sirupsen/logrus"
	"github.com/sivo4kin/ea-starter/keys"
	"time"
)

func NewDHTBootPeer(key string) (err error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var (
		port int
	)
	setFlags(ctx, &port)

	//настроить host который option
	listenAddr := libp2p.ListenAddrStrings(fmt.Sprintf("/ip4/0.0.0.0/tcp/%d", port))
	privKey, pubkey, err := keys.ReadKeys("./keys/srv3-ecdsa.key", "./keys/srv3-rsa.key")
	if err != nil {
		logrus.Errorf("ERROR GETTING CERT %v", err)
		//panic(err)
		return
	}
	logrus.Printf("pubkey %v", pubkey)
	identify := libp2p.Identity(privKey)
	routing := libp2p.Routing(func(h host.Host) (routing.PeerRouting, error) {
		dualDHT, err := ddht.New(ctx, h, ddht.DHTOption(dht.Mode(dht.ModeServer), dht.ProtocolPrefix("/myapp"))) //в качестве dhtServer
		//,dht.ProtocolPrefix("/myapp")
		_ = dualDHT.Bootstrap(ctx)

		go func() {
			ticker := time.NewTicker(time.Second * 15)
			time.Sleep(time.Second)
			for {
				logrus.Printf("***** Вывод синхронизации RoutingTable ******\n")
				dualDHT.LAN.RoutingTable().Print()
				<-ticker.C
			}
		}()
		return dualDHT, err
	})

	host, err := libp2p.New(
		ctx,
		identify,
		routing,
		listenAddr,
		//libp2p.NATPortMap(),
	)
	if err != nil {
		//panic(err)
		return
	}

	for _, addr := range host.Addrs() {
		logrus.Printf("Addr: %s/p2p/%s\n", addr, host.ID().Pretty())
	}

	host.SetStreamHandler(dht.ProtocolDHT, func(stream network.Stream) {
		logrus.Printf("handling %s\n", stream)
	})

	select {}
}

func setFlags(ctx context.Context, port *int) {
	flag.IntVar(port, "port", 6666, "")
}
