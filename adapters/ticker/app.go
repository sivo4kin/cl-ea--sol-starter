package ticker

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sirupsen/logrus"
	"github.com/sivo4kin/ea-starter/config"
	"time"
)

type Output struct {
	ChainId  string `json:"chainId"`
	BlockNum string `json:"blockNum"`
}

func (ap *Ticker) GetChainIDAndBlock() (*Output, error) {

	chainID, err := ap.CLient.ChainID(context.Background())
	if err != nil {
		return nil, err
	}

	logrus.Printf("chainID %d", chainID)
	num, err := ap.GetLatestBlock()
	if err != nil {
		return nil, err
	}
	o := Output{
		ChainId:  fmt.Sprintf("%s", chainID.String()),
		BlockNum: fmt.Sprintf("%d", num),
	}
	return &o, nil
}

func (ap *Ticker) StartChainTicker() (err error) {
	logrus.Print("StartChainTicker")
	n, err := ap.GetLatestBlock()
	if err != nil {
		return
	}

	ticker := time.NewTicker(ap.Config.TickerInterval)
	go func() {
		for range ticker.C {
			n, err = ap.GetLatestBlock()
			logrus.Printf("Block %d", n)
		}
	}()
	return
}

func (ap *Ticker) GetLatestBlock() (n uint64, err error) {
	n, err = ap.CLient.BlockNumber(context.Background())
	return
}

func NewTicker(cfg config.AppConfig) (a *Ticker, err error) {
	a = &Ticker{}
	a.Config = cfg
	a.CLient, err = ethclient.Dial(config.Config.INFURA_URL)
	return
}
