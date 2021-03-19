package ticker

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/linkpoolio/bridges"
	"github.com/sivo4kin/ea-starter/config"
)

type Ticker struct {
	CLient *ethclient.Client
	Config config.AppConfig
}

func (ap *Ticker) Opts() *bridges.Opts {
	return &bridges.Opts{
		Name:   "Asset Price",
		Lambda: true,
		Path:   "/tick",
	}
}

func (ap *Ticker) Run(helper *bridges.Helper) (interface{}, error) {
	return ap.GetChainIDAndBlock()
}
