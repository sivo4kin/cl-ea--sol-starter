package bridge

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/linkpoolio/bridges"
	"github.com/sirupsen/logrus"
)

type EthHealth struct {
	CLient *ethclient.Client
	Name   string
	Url    string
	Lambda bool
	Func   func() (*Output, error)
}

type Output struct {
	ChainId  string `json:"chainId"`
	BlockNum string `json:"blockNum"`
}

func (ap *EthHealth) GetChainIDAndBlock() (*Output, error) {

	//chainID, err := ap.CLient.ChainID(context.Background())
	//if err != nil {
	//	return nil, err
	//}
	//log.Printf("chainID %d", chainID)
	//num, err := ap.GetLatestBlock(*ap.CLient)
	//if err != nil {
	//	return nil, err
	//}
	//o := Output{
	//	ChainId:  fmt.Sprintf("%s", chainID.String()),
	//	BlockNum: fmt.Sprintf("%d", num),
	//}
	o := Output{
		ChainId:  fmt.Sprintf("%s", "QWE"),
		BlockNum: fmt.Sprintf("%d", 999),
	}
	return &o, nil
}

func First() (*Output, error) {
	o := Output{
		ChainId:  fmt.Sprintf("%s", "QWE"),
		BlockNum: fmt.Sprintf("%d", 999),
	}
	return &o, nil
}

func Second() (*Output, error) {
	o := Output{
		ChainId:  fmt.Sprintf("%s", "SECOND"),
		BlockNum: fmt.Sprintf("%d", 1111),
	}
	return &o, nil
}

func GetChainIDAndBlock(qwe string) (*Output, error) {
	//ctxutil.Client(ctx)
	//chainID, err := client.ChainID(context.Background())
	//if err != nil {
	//	return nil, err
	//}
	//log.Printf("chainID %d", chainID)
	//num, err := GetLatestBlock(client)
	//if err != nil {
	//	return nil, err
	//}
	o := Output{
		ChainId:  fmt.Sprintf("%s", qwe),
		BlockNum: fmt.Sprintf("%d", 1),
	}
	return &o, nil
}

func GetLatestBlock(client ethclient.Client) (n uint64, err error) {
	n, err = client.BlockNumber(context.Background())
	return
}

func NewEthHealth(ethClient *ethclient.Client, name, url string, f func() (*Output, error)) (a *EthHealth, err error) {
	a = &EthHealth{}
	a.CLient = ethClient
	a.Name = name
	a.Url = url
	a.Func = f
	return
}

func (ap *EthHealth) Opts() *bridges.Opts {
	return &bridges.Opts{
		Name:   ap.Name,
		Lambda: ap.Lambda,
		Path:   "/" + ap.Url,
	}
}

func (ap *EthHealth) Run(helper *bridges.Helper) (interface{}, error) {

	res, err := ap.Func()

	if err != nil {
		return nil, err
	}
	logrus.Print(res)
	return res, nil

}
