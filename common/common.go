package common

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sirupsen/logrus"
	"github.com/sivo4kin/ea-starter/adapters/bridge"
	"github.com/sivo4kin/ea-starter/config"
	"github.com/sivo4kin/ea-starter/wrappers"
	"math/big"
	"os"
)

func connect(string2 string) (*ethclient.Client, error) {
	return ethclient.Dial(string2)
}

func HealthFirst() (*bridge.Output, error) {
	client, err := connect(config.Config.CHAIN_1_URL)
	if err != nil {
		panic(err)
	}
	n, err := client.BlockNumber(context.Background())
	o := bridge.Output{
		ChainId:  fmt.Sprintf("%s", config.Config.CHAIN_1_URL),
		BlockNum: fmt.Sprintf("%d", n),
	}
	return &o, err
}

func HealthSecond() (*bridge.Output, error) {
	client, err := connect(config.Config.CHAIN_2_URL)
	if err != nil {
		panic(err)
	}
	n, err := client.BlockNumber(context.Background())
	o := bridge.Output{
		ChainId:  fmt.Sprintf("%s", config.Config.CHAIN_2_URL),
		BlockNum: fmt.Sprintf("%d", n),
	}
	return &o, err
}

func ToECDSAFromHex(hexString string) (*ecdsa.PrivateKey, error) {
	pk := new(ecdsa.PrivateKey)
	pk.D, _ = new(big.Int).SetString(hexString, 16)
	return pk, nil
}

func SetMockPoolTestRequest() (o *bridge.Output, err error) {
	logrus.Printf("%v", 1)
	client, err := connect(config.Config.CHAIN_1_URL)
	if err != nil {
		logrus.Errorf("%v", err)
	}

	mockDexPoolContract, err := wrappers.NewMockDexPool(common.HexToAddress(config.Config.BRIDGE_1_ADDRESS), client)
	if err != nil {
		logrus.Errorf("%v", err)
	}

	pKey, err := ToECDSAFromHex(os.Getenv("SK"))
	if err != nil {
		logrus.Errorf("%v", err)
	}

	txOpts := bind.NewKeyedTransactor(pKey)
	if err != nil {
		logrus.Errorf("%v", err)
	}

	tx, err := mockDexPoolContract.SendRequestTest(txOpts, big.NewInt(99), common.HexToAddress(config.Config.POOL_ADDRESS))

	if err != nil {
		logrus.Errorf("%v", err)
	}
	logrus.Printf("TX HASH %x", tx.Hash())
	o.ChainId = fmt.Sprintf("%s", tx.Hash().String())
	o.BlockNum = fmt.Sprintf("%d", 1)
	return
}
