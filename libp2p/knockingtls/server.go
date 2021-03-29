package knockingtls

import (
	"context"
	"flag"
	"fmt"
	"github.com/libp2p/go-libp2p"
	p2pcrypto "github.com/libp2p/go-libp2p-core/crypto"
	p2phost "github.com/libp2p/go-libp2p-core/host"
	p2ppeer "github.com/libp2p/go-libp2p-core/peer"
	p2ppeerstore "github.com/libp2p/go-libp2p-core/peerstore"
	mplex "github.com/libp2p/go-libp2p-mplex"
	peer "github.com/libp2p/go-libp2p-peer"
	libp2ptls "github.com/libp2p/go-libp2p-tls"
	"github.com/libp2p/go-tcp-transport"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/sivo4kin/ea-starter/keys"
	"github.com/smartcontractkit/chainlink/core/services/signatures/secp256k1"
	"github.com/smartcontractkit/libocr/offchainreporting/loghelper"
	"github.com/smartcontractkit/libocr/offchainreporting/types"
	"gitlab.com/thorchain/tss/go-tss/keygen"
	"gitlab.com/thorchain/tss/go-tss/keysign"
	"net"
	"sync"
	"time"
)

var secp256k1Suite = secp256k1.NewBlakeKeccackSecp256k1()

// Server define the necessary functionality should be provide by a TSS Server implementation
type IServer interface {
	Start() error
	Stop()
	GetLocalPeerID() string
	Keygen(req keygen.Request) (keygen.Response, error)
	KeySign(req keysign.Request) (keysign.Response, error)
}

type СoncretePeer struct {
	p2phost.Host
	tls    *KnockingTLSTransport
	logger loghelper.LoggerWithContext
	//endpointConfig EndpointConfig
	registrants   map[types.ConfigDigest]struct{}
	registrantsMu *sync.Mutex
}

type PeerConfig struct {
	PrivKey      p2pcrypto.PrivKey
	ListenIP     net.IP
	ListenPort   uint16
	AnnounceIP   net.IP
	AnnouncePort uint16
	Logger       types.Logger
	Peerstore    p2ppeerstore.Peerstore
}

func NewKnokingTLSPeer(c PeerConfig) (*СoncretePeer, error) {
	if c.ListenPort == 0 {
		return nil, errors.New("NewPeer requires a non-zero listen port")
	}

	peerID, err := p2ppeer.IDFromPrivateKey(c.PrivKey)
	if err != nil {
		return nil, errors.Wrap(err, "error extracting peer ID from private key")
	}

	listenAddr, err := makeMultiaddr(c.ListenIP, c.ListenPort)
	if err != nil {
		return nil, errors.Wrap(err, "could not make listen multiaddr")
	}

	logger := loghelper.MakeRootLoggerWithContext(c.Logger).MakeChild(types.LogFields{
		"id":         "Peer",
		"peerID":     peerID.Pretty(),
		"listenPort": c.ListenPort,
		"listenIP":   c.ListenIP.String(),
		"listenAddr": listenAddr.String(),
	})

	addrsFactory, err := makeAddrsFactory(c.AnnounceIP, c.AnnouncePort)
	if err != nil {
		return nil, errors.Wrap(err, "could not make addrs factory")
	}

	tlsID := ID
	tls, err := NewKnockingTLS(logger, c.PrivKey)
	if err != nil {
		return nil, errors.Wrap(err, "could not create knocking tls")
	}

	transportCon := tcp.NewTCPTransport

	opts := []libp2p.Option{
		libp2p.ListenAddrs(listenAddr),
		libp2p.Identity(c.PrivKey),
		libp2p.DisableRelay(),
		libp2p.Security(tlsID, tls),
		libp2p.Peerstore(c.Peerstore),
		libp2p.AddrsFactory(addrsFactory),
		libp2p.Transport(transportCon),
		libp2p.Muxer("/mplex/6.7.0", mplex.DefaultTransport),
	}

	basicHost, err := libp2p.New(context.Background(), opts...)
	if err != nil {
		logrus.Fatal(err)
		return nil, err
	}
	fmt.Println("This peer: ", basicHost.ID().Pretty(), " ", basicHost.Addrs())

	return &СoncretePeer{
		Host:          basicHost,
		tls:           tls,
		logger:        logger,
		registrants:   make(map[types.ConfigDigest]struct{}),
		registrantsMu: &sync.Mutex{},
	}, nil
}

func RunKnokingTLSPeer() {
	priv, err := keys.ReadHostKey("/home/syi/src/digiu-cross-chain/keys/srv1-ed25519.key")
	if err != nil {
		panic(err)
	}
	peerConfig := PeerConfig{
		PrivKey:    priv,
		ListenIP:   net.ParseIP("0.0.0.0"),
		ListenPort: 8989,
		Peerstore:  nil,
	}
	_, err = NewKnokingTLSPeer(peerConfig)
	if err != nil {
		panic(err)
	}

}

func handleConn(tp *libp2ptls.Transport, conn net.Conn) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	sconn, err := tp.SecureInbound(ctx, conn)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Authenticated client: %s\n", sconn.RemotePeer().Pretty())
	fmt.Fprintf(sconn, "Hello client!")
	fmt.Printf("Closing connection to %s\n", conn.RemoteAddr())

}

func RunTLSPeer(peerId chan peer.ID) {
	port := flag.Int("port", 5555, "port")
	priv, err := keys.ReadHostKey("../keys/tls1-ed25519.key")
	if err != nil {
		return
	}
	peerID, err := peer.IDFromPrivateKey(priv)
	if err != nil {
		panic(err)
	}
	peerId <- peerID

	tp, err := libp2ptls.New(priv)
	if err != nil {
		return
	}
	ln, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		return
	}
	fmt.Printf("Listening for new connections on %s\n", ln.Addr())
	fmt.Printf("Now run the following command in a separate terminal:\n")
	fmt.Printf("\tgo run cmd/tlsdiag.go client -p %d -id %s\n", *port, peerId)

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Errorf("ERROR ACCEPT", err)
		}
		fmt.Printf("Accepted raw connection from %s\n", conn.RemoteAddr())
		go handleConn(tp, conn)

	}
	return
}
