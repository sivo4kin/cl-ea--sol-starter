package quic

import (
	"bytes"
	"context"
	"crypto/rand"
	"errors"
	"fmt"
	"github.com/libp2p/go-libp2p"
	circuit "github.com/libp2p/go-libp2p-circuit"
	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/transport"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	noise "github.com/libp2p/go-libp2p-noise"
	libp2pquic "github.com/libp2p/go-libp2p-quic-transport"
	"github.com/libp2p/go-libp2p-quic-transport/integrationtests/conn"
	"github.com/libp2p/go-libp2p-quic-transport/integrationtests/stream"
	ma "github.com/multiformats/go-multiaddr"
	"github.com/sivo4kin/ea-starter/keys"
	"golang.org/x/sync/errgroup"
	"io/ioutil"
	"strings"
	"time"
)

func RunWithRole(role, test string) (err error) {
	var addr ma.Multiaddr
	hostKey, peerPubKey, err := keys.ReadKeys("../keys/cl1-ecdsa.key", "../keys/cl1-rsa.key")

	if err != nil {
		return
	}
	addr, err = ma.NewMultiaddr("/ip4/127.0.0.1/udp/8182")
	if err != nil {
		return
	}

	switch role {
	case "server":
		if err = runServer(hostKey, peerPubKey, addr, test); err != nil {
			return
		}
	case "client":
		if err = runClient(hostKey, peerPubKey, addr, test); err != nil {
			return
		}
	}
	return
}

func runServer(hostKey crypto.PrivKey, peerKey crypto.PubKey, addr ma.Multiaddr, test string) error {
	fmt.Print("run server !!!")
	tr, err := libp2pquic.NewTransport(hostKey, nil, nil)
	if err != nil {
		return err
	}
	ln, err := tr.Listen(addr)
	if err != nil {
		fmt.Print(err)
		return err
	}
	conn, err := ln.Accept()
	if err != nil {
		fmt.Print(err)
		return err
	}
	fmt.Print("accepeed listener  !!!")

	if test == "handshake-failure" {
		return errors.New("didn't expect to accept a connection in the handshake-failure test")
	}
	defer ln.Close()
	if !conn.RemotePublicKey().Equals(peerKey) {
		return errors.New("mismatching public keys")
	}
	clientPeerID, err := peer.IDFromPublicKey(peerKey)
	if err != nil {
		return err
	}
	if conn.RemotePeer() != clientPeerID {
		return fmt.Errorf("remote Peer ID mismatch. Got %s, expected %s", conn.RemotePeer().Pretty(), clientPeerID.Pretty())
	}
	var g errgroup.Group
	for {
		st, err := conn.AcceptStream()
		if err != nil {
			break
		}
		fmt.Print("Accepted Stream !!!")
		str := stream.WrapStream(st)
		g.Go(func() error {
			data, err := ioutil.ReadAll(str)
			if err != nil {
				return err
			}
			if err := str.CloseRead(); err != nil {
				return err
			}
			if _, err := str.Write(data); err != nil {
				return err
			}
			return str.CloseWrite()
		})
	}
	fmt.Print("STARTED !!!")
	return g.Wait()

}

func runClient(hostKey crypto.PrivKey, serverKey crypto.PubKey, addr ma.Multiaddr, test string) error {
	tr, err := libp2pquic.NewTransport(hostKey, nil, nil)
	if err != nil {
		return err
	}
	switch test {
	case "single-transfer":
		return testSingleFileTransfer(tr, serverKey, addr)
	case "multi-transfer":
		return testMultipleFileTransfer(tr, serverKey, addr)
	case "handshake-failure":
		return testHandshakeFailure(tr, serverKey, addr)
	default:
		return errors.New("unknown test")
	}
}

func testSingleFileTransfer(tr transport.Transport, serverKey crypto.PubKey, addr ma.Multiaddr) error {
	serverPeerID, err := peer.IDFromPublicKey(serverKey)
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	c, err := tr.Dial(ctx, addr, serverPeerID)
	if err != nil {
		return fmt.Errorf("Dial failed: %w", err)
	}
	defer c.Close()
	if !c.RemotePublicKey().Equals(serverKey) {
		return errors.New("mismatching public keys")
	}
	if c.RemotePeer() != serverPeerID {
		return fmt.Errorf("remote Peer ID mismatch. Got %s, expected %s", c.RemotePeer().Pretty(), serverPeerID.Pretty())
	}
	st, err := conn.OpenStream(context.Background(), c)
	if err != nil {
		return err
	}
	str := stream.WrapStream(st)
	data := make([]byte, 1<<15)
	rand.Read(data)
	if _, err := str.Write(data); err != nil {
		return err
	}
	if err := str.CloseWrite(); err != nil {
		return err
	}
	echoed, err := ioutil.ReadAll(str)
	if err != nil {
		return err
	}
	if !bytes.Equal(data, echoed) {
		return errors.New("echoed data does not match")
	}

	return nil

}

func testMultipleFileTransfer(tr transport.Transport, serverKey crypto.PubKey, addr ma.Multiaddr) error {
	serverPeerID, err := peer.IDFromPublicKey(serverKey)
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	c, err := tr.Dial(ctx, addr, serverPeerID)
	if err != nil {
		return fmt.Errorf("Dial failed: %w", err)
	}
	defer c.Close()
	if !c.RemotePublicKey().Equals(serverKey) {
		return errors.New("mismatching public keys")
	}
	if c.RemotePeer() != serverPeerID {
		return fmt.Errorf("remote Peer ID mismatch. Got %s, expected %s", c.RemotePeer().Pretty(), serverPeerID.Pretty())
	}
	var g errgroup.Group
	for i := 0; i < 2000; i++ {
		g.Go(func() error {
			st, err := conn.OpenStream(context.Background(), c)
			if err != nil {
				return err
			}
			str := stream.WrapStream(st)
			data := make([]byte, 1<<9)
			rand.Read(data)
			if _, err := str.Write(data); err != nil {
				return err
			}
			if err := str.CloseWrite(); err != nil {
				return err
			}
			echoed, err := ioutil.ReadAll(str)
			if err != nil {
				return err
			}
			if !bytes.Equal(data, echoed) {
				return errors.New("echoed data does not match")
			}
			return nil
		})
	}
	return g.Wait()
}

func testHandshakeFailure(tr transport.Transport, serverKey crypto.PubKey, addr ma.Multiaddr) error {
	serverPeerID, err := peer.IDFromPublicKey(serverKey)
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err = tr.Dial(ctx, addr, serverPeerID)
	if err == nil {
		return errors.New("expected the handshake to fail")
	}
	if !strings.Contains(err.Error(), "CRYPTO_ERROR") || !strings.Contains(err.Error(), "peer IDs don't match") {
		return fmt.Errorf("got unexpected error: %w", err)
	}
	return nil
}

func Srv() {
	ctx := context.Background()
	sourceMultiAddr, _ := ma.NewMultiaddr("/ip4/0.0.0.0/tcp/4000")
	prvKey, err := keys.ReadHostKey("../keys/srv2-ed25519.key")
	if err != nil {
		panic(err)
	}
	host, err := libp2p.New(
		ctx,
		libp2p.ListenAddrs(sourceMultiAddr),
		libp2p.Security(noise.ID, noise.New),
		libp2p.Identity(prvKey),
		libp2p.EnableRelay(circuit.OptHop),
	)
	if err != nil {
		panic(err)
	}

	fmt.Println("This node: ", host.ID().Pretty(), " ", host.Addrs())

	_, err = dht.New(ctx, host)
	if err != nil {
		panic(err)
	}

	select {}
}

/*
	libListenAddrStrings(
			"/ip4/0.0.0.0/tcp/9000",      // regular tcp connections
			"/ip4/0.0.0.0/udp/9000/quic", // a UDP endpoint for the QUIC transport
		),
*/
