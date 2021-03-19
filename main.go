package main

import (
	"context"
	"github.com/linkpoolio/bridges"
	"github.com/sirupsen/logrus"
	test_get "github.com/sivo4kin/ea-starter/adapters/test-get"
	"github.com/sivo4kin/ea-starter/adapters/ticker"
	"github.com/sivo4kin/ea-starter/config"
	"log"
)

func main() {
	//router := adapter_service.NewRouter()
	//log.Fatal(http.ListenAndServe(":3000", router))

	err := config.LoadConfig(".")
	if err != nil {
		log.Fatal(err)
	}
	logrus.Print("CONFIG\n", config.Config, "\nEND")
	logrus.Print("STARTING")
	var bridgesList []bridges.Bridge
	ad, err := ticker.NewTicker(config.Config)
	if err != nil {
		logrus.Fatal(err)
	}
	err = ad.StartChainTicker()
	if err != nil {
		logrus.Fatal(err)
	}

	tg, err := test_get.NewTestGet(config.Config)
	if err != nil {
		logrus.Fatal(err)
	}

	bridgesList = append(bridgesList, tg)
	bridgesList = append(bridgesList, ad)
	chainId, err := ad.CLient.ChainID(context.Background())
	logrus.Printf("CHAIN ID %d", chainId)
	bridgesList = append(bridgesList, ad)
	//bridgesList = append(bridgesList, &adapters.Adapter3{})
	srv := bridges.NewServer(bridgesList...)
	logrus.Print("STARTED NEW BRIDGE")
	srv.Start(config.Config.PORT)
	logrus.Print("STARTED")

}
