package transmitter

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sivo4kin/ea-starter/config"
)

func NewBridge(cfg config.AppConfig) (a *Adapter, err error) {
	a = &Adapter{}
	a.Config = cfg
	a.CLient, err = ethclient.Dial(config.Config.INFURA_URL)
	return
}
