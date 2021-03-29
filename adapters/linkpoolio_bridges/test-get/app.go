package test_get

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/linkpoolio/bridges"
	"github.com/sirupsen/logrus"
	"github.com/sivo4kin/ea-starter/config"
	"github.com/sivo4kin/ea-starter/wrappers"
)

/*
app.get('/testGet', async function (req, res) {
    console.log('/testGet', req.body);
    // the owner - his deployed smart-contract on Net2
    let ownerAdapter = (await worker.web3.eth.getAccounts())[0];
    // this is represents of bytes memory out = abi.encodeWithSelector(bytes4(keccak256(bytes('_getTest()'))));
    let data = '0x49eba8f7';
    //pass 'data' for call inside smart-contracts
    const getResult  = await bridge.methods.lowLevelGet(data).call();
    console.log('Result staticcall ', getResult);
    res.status(200).send({});
});
*/

type TestGet struct {
	CLient *ethclient.Client
	Config config.AppConfig
}

func (ap *TestGet) Opts() *bridges.Opts {
	return &bridges.Opts{
		Name:   "TestGet",
		Lambda: true,
		Path:   "/get",
	}
}

type Output struct {
	DataFromContract string `json:"data"`
}

func (brg *TestGet) Run(helper *bridges.Helper) (interface{}, error) {
	data := helper.GetParam("data")
	return brg.GeFromContract(data)
}

func (brg *TestGet) GeFromContract(data string) (out Output, err error) {
	log.Printf("START get from contract %s", data)
	//poolContract, err := wrappers.NewDexPool(common.HexToAddress(brg.Config.POOL_ADDRESS), brg.CLient)
	poolContract, err := wrappers.NewDexPool(common.HexToAddress("0x8C2e2b076ccd2d1654de5A094a8626ADa609b415"), brg.CLient)
	if err != nil {
		return
	}
	data = "0x49eba8f7"
	dataFrom, err := poolContract.LowLevelGet(&bind.CallOpts{}, []byte(data))
	if err != nil {
		return
	}
	log.Printf("get from contract %v", dataFrom)
	out = Output{
		DataFromContract: fmt.Sprintf("%v", dataFrom),
	}
	return
}

func NewTestGet(cfg config.AppConfig) (a *TestGet, err error) {
	a = &TestGet{}
	a.Config = cfg
	a.CLient, err = ethclient.Dial(config.Config.INFURA_URL)
	return
}
