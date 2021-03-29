module github.com/sivo4kin/ea-starter

go 1.16

require (
	github.com/StackExchange/wmi v0.0.0-20190523213315-cbe66965904d // indirect
	github.com/btcsuite/btcd v0.21.0-beta
	github.com/ethereum/go-ethereum v1.10.1
	github.com/go-ole/go-ole v1.2.4 // indirect
	github.com/golang/groupcache v0.0.0-20200121045136-8c9f03a8e57e // indirect
	github.com/golang/protobuf v1.4.3
	github.com/google/go-cmp v0.5.2 // indirect
	github.com/gopherjs/gopherjs v0.0.0-20190812055157-5d271430af9f // indirect
	github.com/gorilla/mux v1.8.0
	github.com/ipfs/go-datastore v0.4.5
	github.com/ipfs/go-log/v2 v2.1.2-0.20200626104915-0016c0b4b3e4 // indirect
	github.com/libp2p/go-libp2p v0.13.0
	github.com/libp2p/go-libp2p-circuit v0.4.0
	github.com/libp2p/go-libp2p-core v0.8.5
	github.com/libp2p/go-libp2p-crypto v0.1.0
	github.com/libp2p/go-libp2p-discovery v0.5.0
	github.com/libp2p/go-libp2p-gorpc v0.1.2
	github.com/libp2p/go-libp2p-host v0.1.0
	github.com/libp2p/go-libp2p-kad-dht v0.11.1
	github.com/libp2p/go-libp2p-mplex v0.4.1
	github.com/libp2p/go-libp2p-noise v0.1.2
	github.com/libp2p/go-libp2p-peer v0.2.0
	github.com/libp2p/go-libp2p-peerstore v0.2.6
	github.com/libp2p/go-libp2p-pubsub v0.4.1
	github.com/libp2p/go-libp2p-quic-transport v0.10.0
	github.com/libp2p/go-libp2p-tls v0.1.3
	github.com/libp2p/go-tcp-transport v0.2.1
	github.com/libp2p/go-ws-transport v0.4.0
	github.com/linkpoolio/bridges v0.0.0-20200226172010-aa6f132d477e
	github.com/multiformats/go-multiaddr v0.3.1
	github.com/pkg/errors v0.9.1
	github.com/sirupsen/logrus v1.8.1
	github.com/smartcontractkit/chainlink v0.9.5-0.20201207211610-6c7fee37d5b7
	github.com/smartcontractkit/libocr v0.0.0-20210319202758-14aa50f869b7
	github.com/smartystreets/assertions v1.0.1 // indirect
	github.com/spf13/viper v1.7.1
	github.com/stretchr/testify v1.7.0
	gitlab.com/thorchain/tss/go-tss v1.3.0
	go.dedis.ch/kyber/v3 v3.0.13
	golang.org/x/crypto v0.0.0-20201221181555-eec23a3978ad
	golang.org/x/sync v0.0.0-20201207232520-09787c993a3a
	gopkg.in/guregu/null.v3 v3.5.0
	gopkg.in/urfave/cli.v1 v1.20.0
)

replace (
	github.com/binance-chain/tss-lib => gitlab.com/thorchain/tss/tss-lib v0.0.0-20201118045712-70b2cb4bf916
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.2-alpha.regen.4
)
