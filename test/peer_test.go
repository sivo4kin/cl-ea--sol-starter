package test

//
//import (
//	"github.com/libp2p/go-libp2p-core/peer"
//	"github.com/sivo4kin/ea-starter/libp2p/quic"
//	tls2 "github.com/sivo4kin/ea-starter/libp2p/tls"
//	"github.com/stretchr/testify/require"
//	"testing"
//)
//
//func TestRunWithServerRole(t *testing.T) {
//	err := quic.RunWithRole("server", "handshake-failure")
//	require.Nil(t, err)
//}
//
//func TestRunClientWithSRole(t *testing.T) {
//	quic.RunWithRole("client", "handshake-failure")
//}
//
//func TestRunTLSServer(t *testing.T) {
//	_, err := tls2.StartServer()
//	require.Nil(t, err)
//
//	//RunTLSClient(id.Pretty())
//
//}
//
//func TestRunTLSClient(t *testing.T) {
//
//	err := RunTLSClient("QmPrsWXhaABAFnHWhTWK3SXYque1P7EM9AiieBoR7sYEDZ")
//	require.Nil(t, err)
//}
//
//func RunTLSClient(id string) error {
//	return tls2.StartClient(id)
//}
//
////var peerid = ""
//
//func TestSimplePeer(t *testing.T) {
//	knockingtls.RunKnokingTLSPeer()
//}
//
//func TestSrv(t *testing.T) {
//	quic.Srv()
//
//}
//
//func TestQUICServer(t *testing.T) {
//	peerId := make(chan peer.ID)
//	defer close(peerId)
//	go knockingtls.RunTLSPeer(peerId)
//	go runQUICClient(peerId)
//}
//
//func runQUICClient(peerId chan peer.ID) {
//	q := <-peerId
//	go knockingtls.StartClient(q.Pretty())
//}
//
///*func TestTLSPeer(t *testing.T) {
//	//var id string = ""
//	done := make(chan struct{})
//	go func() {
//		_= p2p.RunTLSPeer()
//		done <- struct{}{} // перед завершением сообщаем об этом
//	}()
//	<- done // ожидание завершения горутины
//
///*	go func() {
//		err := p2p.StartClient(id)
//		require.Nil(t, err)
//		done <- struct{}{} // перед завершением сообщаем об этом
//	}()
//	<- done*/
//
////go p2p.StartClient("12D3KooWAxE6juLtabiaByhveuCjM74YWJSVKKChuNLnM1pZJxvx")
////require.Nil(t, err)
//
////}
