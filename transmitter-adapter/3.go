package transmitter_adapter

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/linkpoolio/bridges"
	"github.com/sivo4kin/digiu-cross-chain/config"
	"log"
)

type Adapter3 struct {
	CLient *ethclient.Client
	Config config.AppConfig

	//Session         *bindings.Session
	//Filterer *bindings.Filterer
}

func (ap *Adapter3) Run(helper *bridges.Helper) (interface{}, error) {
	return GetBlocks(), nil
}

func (ap *Adapter3) Opts() *bridges.Opts {
	return &bridges.Opts{
		Name:   "post",
		Lambda: true,
		Path:   "/3",
	}
}

func GetBlocks() error {
	header, err := Adapter2{}.CLient.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
		return err
	}
	fmt.Printf("current block %s", header.Number.String())
	return nil
}
