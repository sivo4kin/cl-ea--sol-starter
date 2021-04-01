package common

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/linkpoolio/bridges"
	"github.com/sirupsen/logrus"
	"github.com/sivo4kin/ea-starter/adapters"
	"github.com/sivo4kin/ea-starter/config"
	"github.com/sivo4kin/ea-starter/wrappers"
	"math/big"
	"os"
	"strings"
)

func Connect(string2 string) (*ethclient.Client, error) {
	return ethclient.Dial(string2)
}

func HealthFirst(helper *bridges.Helper) (*adapters.Output, error) {
	client, err := Connect(config.Config.CHAIN_1_URL)
	if err != nil {
		panic(err)
	}
	n, err := client.BlockNumber(context.Background())
	o := adapters.Output{
		ChainId:  fmt.Sprintf("%s", config.Config.CHAIN_1_URL),
		BlockNum: fmt.Sprintf("%d", n),
	}
	return &o, err
}

func HealthSecond(helper *bridges.Helper) (*adapters.Output, error) {
	client, err := Connect(config.Config.CHAIN_2_URL)
	if err != nil {
		panic(err)
	}
	n, err := client.BlockNumber(context.Background())
	o := adapters.Output{
		ChainId:  fmt.Sprintf("%s", config.Config.CHAIN_2_URL),
		BlockNum: fmt.Sprintf("%d", n),
	}
	return &o, err
}

func ToECDSAFromHex(hexString string) (pk *ecdsa.PrivateKey, err error) {
	pk, err = crypto.HexToECDSA(strings.TrimPrefix(hexString, "0x"))
	return
}

func SetMockPoolTestRequest(helper *bridges.Helper) (o *adapters.Output, err error) {
	logrus.Printf("%v", 1)
	client1, err := Connect(config.Config.CHAIN_1_URL)
	if err != nil {
		logrus.Errorf("%v", err)
	}

	pKey1, err := ToECDSAFromHex(os.Getenv("SK1"))
	if err != nil {
		logrus.Errorf("ToECDSAFromHex %v", err)
	}

	txOpts1 := bind.NewKeyedTransactor(pKey1)

	mockDexPoolContract1, err := wrappers.NewMockDexPool(common.HexToAddress(config.Config.POOL_1_ADDRESS), client1)
	if err != nil {
		logrus.Errorf("NewMockDexPool 1 %v", err)
	}

	tx, err := mockDexPoolContract1.SendRequestTest(txOpts1, big.NewInt(99), common.HexToAddress(config.Config.POOL_2_ADDRESS))

	if err != nil {
		logrus.Errorf(" SendRequestTest %v", err)
	}

	logrus.Printf("TX HASH %x", tx.Hash())
	o.ChainId = fmt.Sprintf("%s", tx.Hash().String())
	o.BlockNum = fmt.Sprintf("%d", 1)
	return
}
