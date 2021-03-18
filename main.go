package main

import (
	"context"
	"github.com/linkpoolio/bridges"
	"github.com/sirupsen/logrus"
	adapter_service "github.com/sivo4kin/digiu-cross-chain/adapter-service"
	"github.com/sivo4kin/digiu-cross-chain/config"
	"github.com/sivo4kin/digiu-cross-chain/ticker-adapter"
	"log"
	"net/http"
)

func main() {
	err := config.LoadConfig(".")
	if err != nil {
		log.Fatal(err)
	}
	logrus.Print("CONFIG\n", config.Config, "\nEND")
	logrus.Print("STARTING")
	var bridgesList []bridges.Bridge
	ad, err := ticker_adapter.NewTicker(config.Config)
	if err != nil {
		logrus.Fatal(err)
	}
	err = ad.StartChainTicker()
	if err != nil {
		logrus.Fatal(err)
	}
	bridgesList = append(bridgesList, ad)
	chainId, err := ad.CLient.ChainID(context.Background())
	logrus.Printf("CHAIN ID %d", chainId)
	bridgesList = append(bridgesList, ad)
	//bridgesList = append(bridgesList, &adapters.Adapter3{})
	srv := bridges.NewServer(bridgesList...)
	logrus.Print("STARTED NEW BRIDGE")
	srv.Start(config.Config.PORT)
	logrus.Print("STARTED")
	router := adapter_service.NewRouter()
	log.Fatal(http.ListenAndServe(":3000", router))
}






/*
func (r *Router) HandleFunc(path string, f func(http.ResponseWriter,
	*http.Request)) *Route {
	return r.NewRoute().Path(path).HandlerFunc(f)
}

*/

/*func main() {
	r := mux.NewRouter()
	r.HandleFunc("/home", HomeHandler)

	http.Handle("/", r)
}
*/

// HandleFunc registers a new route with a matcher for the URL path.
// See Route.Path() and Route.HandlerFunc().
//func (r *mux.Router) HandleFunc(path string, f func(http.ResponseWriter,
//	*http.Request)) *mux.Route {
//	return r.NewRoute().Path(path).HandlerFunc(f)
//}
