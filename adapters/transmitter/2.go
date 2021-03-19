package transmitter

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/linkpoolio/bridges"
	"github.com/sivo4kin/ea-starter/config"
	"log"
	"net/http"
)

type Adapter2 struct {
	CLient *ethclient.Client
	Config config.AppConfig

	//Session         *wrappers.Session
	//Filterer *wrappers.Filterer
}

func (ap *Adapter2) Run(helper *bridges.Helper) (interface{}, error) {
	return GetBlocks2, nil
}

func NewAdapter2(cfg config.AppConfig) (a *Adapter2, err error) {
	a = &Adapter2{}
	a.Config = cfg
	a.CLient, err = ethclient.Dial(config.Config.INFURA_URL)
	return
}

func (ap *Adapter2) Opts() *bridges.Opts {
	return &bridges.Opts{
		Name:   "post",
		Lambda: true,
		Path:   "/control",
	}
}

func GetBlocks2() error {
	header, err := Adapter2{}.CLient.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
		return err
	}
	fmt.Printf("current block %s", header.Number.String())
	return nil
}

func Handler2(g http.ResponseWriter, r *http.Request) error {
	log.Print("handler 3")
	return nil
	//bridges.NewServer(da.Handle(w, r))
}
