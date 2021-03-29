package getblock

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/linkpoolio/bridges"
	"github.com/sirupsen/logrus"
	"github.com/sivo4kin/ea-starter/config"
	"log"
	"time"
)

type Output struct {
	ChainId  string `json:"chainId"`
	BlockNum string `json:"blockNum"`
}

func (ap *Ticker) GetChainIDAndBlock(client ethclient.Client) (*Output, error) {

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return nil, err
	}

	log.Printf("chainID %d", chainID)
	num, err := ap.GetLatestBlock(client)
	if err != nil {
		return nil, err
	}
	o := Output{
		ChainId:  fmt.Sprintf("%s", chainID.String()),
		BlockNum: fmt.Sprintf("%d", num),
	}
	return &o, nil
}

func (ap *Ticker) StartChainTicker(client ethclient.Client, dur time.Duration) (err error) {
	logrus.Print("StartChainTicker")
	n, err := ap.GetLatestBlock(client)
	if err != nil {
		return
	}

	ticker := time.NewTicker(dur)
	go func() {
		for range ticker.C {
			n, err = ap.GetLatestBlock(client)
			log.Printf("Block %d", n)
		}
	}()
	return
}

func (ap *Ticker) GetLatestBlock(client ethclient.Client) (n uint64, err error) {
	n, err = client.BlockNumber(context.Background())
	return
}

func NewTicker(cfg config.AppConfig) (a *Ticker, err error) {
	a = &Ticker{}
	a.Config = cfg
	a.CLient, err = ethclient.Dial(config.Config.INFURA_URL)
	return
}

type Ticker struct {
	CLient *ethclient.Client
	Config config.AppConfig
}

func (ap *Ticker) Opts() *bridges.Opts {
	return &bridges.Opts{
		Name:   "Tick",
		Lambda: true,
		Path:   "/tick",
	}
}

func (ap *Ticker) Run(helper *bridges.Helper) (interface{}, error) {
	return ap.GetChainIDAndBlock(*ap.CLient)
}
