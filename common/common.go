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

func Health(helper *bridges.Helper, rpcUrl string) (out *adapters.Output, err error) {
	out = &adapters.Output{}
	client, err := Connect(rpcUrl)
	if err != nil {
		return
	}
	block, err := client.BlockNumber(context.Background())
	if err != nil {
		return
	}
	chainId, err := client.ChainID(context.Background())
	if err != nil {
		return
	}
	out.ChainId = fmt.Sprintf("%d", chainId)
	out.BlockNum = fmt.Sprintf("%d", block)
	return
}

func HealthFirst(helper *bridges.Helper) (out *adapters.Output, err error) {
	return Health(helper, config.Config.NETWORK_RPC_1)
}

func HealthSecond(helper *bridges.Helper) (*adapters.Output, error) {
	return Health(helper, config.Config.NETWORK_RPC_2)
}

func ToECDSAFromHex(hexString string) (pk *ecdsa.PrivateKey, err error) {
	pk, err = crypto.HexToECDSA(strings.TrimPrefix(hexString, "0x"))
	return
}

func SetMockPoolTestRequest(helper *bridges.Helper) (o *adapters.Output, err error) {
	o = &adapters.Output{}
	reqId := helper.GetIntParam("id")
	client1, err := Connect(config.Config.NETWORK_RPC_1)
	if err != nil {
		return

	}

	pKey1, err := ToECDSAFromHex(os.Getenv("SK1"))
	if err != nil {
		return
	}

	txOpts1 := bind.NewKeyedTransactor(pKey1)

	mockDexPoolContract1, err := wrappers.NewMockDexPool(common.HexToAddress(config.Config.TOKENPOOL_ADDRESS_1), client1)
	if err != nil {
		return
	}

	tx, err := mockDexPoolContract1.SendRequestTest(txOpts1, big.NewInt(reqId), common.HexToAddress(config.Config.TOKENPOOL_ADDRESS_2))

	if err != nil {
		return
	}

	logrus.Printf("TX HASH %x", tx.Hash())
	o.ChainId = fmt.Sprintf("%s", tx.ChainId())
	o.TxHash = tx.Hash().Hex()
	return
}

func ChainlinkData(helper *bridges.Helper) (o *adapters.Output, err error) {
	o = &adapters.Output{}
	fmt.Print(helper.Data)
	o.Data2 = *helper.Data
	return
}
