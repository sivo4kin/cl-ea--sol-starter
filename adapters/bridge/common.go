package bridge

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sivo4kin/ea-starter/config"
)

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

func connect(string2 string) (*ethclient.Client, error) {
	return ethclient.Dial(string2)
}

func HealthFirst() (*Output, error) {
	client, err := connect(config.Config.CHAIN_1_URL)
	if err != nil {
		panic(err)
	}
	n, err := client.BlockNumber(context.Background())
	o := Output{
		ChainId:  fmt.Sprintf("%s", config.Config.CHAIN_1_URL),
		BlockNum: fmt.Sprintf("%d", n),
	}
	return &o, err
}

func HealthSecond() (*Output, error) {
	client, err := connect(config.Config.CHAIN_2_URL)
	if err != nil {
		panic(err)
	}
	n, err := client.BlockNumber(context.Background())
	o := Output{
		ChainId:  fmt.Sprintf("%s", config.Config.CHAIN_2_URL),
		BlockNum: fmt.Sprintf("%d", n),
	}
	return &o, err
}
