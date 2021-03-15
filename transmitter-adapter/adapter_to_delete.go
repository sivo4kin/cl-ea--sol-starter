package transmitter_adapter

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sivo4kin/digiu-cross-chain/config"
)

func NewBridge(cfg config.AppConfig) (a *Adapter, err error) {
	a = &Adapter{}
	a.Config = cfg
	a.CLient, err = ethclient.Dial(config.Config.INFURA_URL)
	return
}
