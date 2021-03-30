package ethealth

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/linkpoolio/bridges"
	"github.com/sirupsen/logrus"
	"log"
	"time"
)

type Output struct {
	ChainId  string `json:"chainId"`
	BlockNum string `json:"blockNum"`
}

func (ap *EthHealth) GetChainIDAndBlock(client ethclient.Client) (*Output, error) {

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

func (ap *EthHealth) StartChainTicker(client ethclient.Client, dur time.Duration) (err error) {
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

func (ap *EthHealth) GetLatestBlock(client ethclient.Client) (n uint64, err error) {
	n, err = client.BlockNumber(context.Background())
	return
}

func NewEthHealth(ethClient *ethclient.Client) (a *EthHealth, err error) {
	a = &EthHealth{}
	a.CLient = ethClient
	return
}

type EthHealth struct {
	CLient *ethclient.Client
}

func (ap *EthHealth) Opts() *bridges.Opts {
	return &bridges.Opts{
		Name:   "Health",
		Lambda: true,
		Path:   "/health",
	}
}

func (ap *EthHealth) Run(helper *bridges.Helper) (interface{}, error) {
	return ap.GetChainIDAndBlock(*ap.CLient)
}
