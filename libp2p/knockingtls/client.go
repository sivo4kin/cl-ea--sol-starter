package knockingtls

import (
	"context"
	"fmt"
	"github.com/libp2p/go-libp2p-core/peer"
	libp2ptls "github.com/libp2p/go-libp2p-tls"
	"github.com/sivo4kin/ea-starter/keys"
	"io/ioutil"
	"net"
	"time"
)

func StartClient(peerIDString string) error {
	port := 5533

	priv, err := keys.ReadHostKey("../keys/tls2-ed25519.key")
	if err != nil {
		return err
	}

	peerID, err := peer.IDB58Decode(peerIDString)
	if err != nil {
		return err
	}

	fmt.Printf("peerID %s", peerID)

	id, err := peer.IDFromPrivateKey(priv)
	if err != nil {
		return err
	}
	fmt.Printf(" Peer ID: %s\n", id.Pretty())
	tp, err := libp2ptls.New(priv)
	if err != nil {
		return err
	}

	remoteAddr := fmt.Sprintf("localhost:%d", port)
	fmt.Printf("Dialing %s\n", remoteAddr)
	conn, err := net.Dial("tcp", remoteAddr)
	if err != nil {
		return err
	}
	fmt.Printf("Dialed raw connection to %s\n", conn.RemoteAddr())

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	sconn, err := tp.SecureOutbound(ctx, conn, peerID)
	if err != nil {
		return err
	}
	fmt.Printf("Authenticated server: %s\n", sconn.RemotePeer().Pretty())
	data, err := ioutil.ReadAll(sconn)
	if err != nil {
		return err
	}
	fmt.Printf("Received message from server: %s\n", string(data))
	return nil
}
