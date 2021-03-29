package config

import (
	"flag"
	//"context"
	"fmt"
	"github.com/spf13/viper"
	"time"
)

// Config is global object that holds all application level variables.
var Config AppConfig

type AppConfig struct {
	// Example Variable
	PORT                    int
	LISTEN_NETWORK          string
	POOL_ADDRESS            string
	ORACLE_CONTRACT_ADDRESS string
	TOKENPOOL_ADDRESS       string
	BRIDGE_ADDRESS          string
	CHAIN_1_URL             string
	CHAIN_2_URL             string
	INFURA_URL              string
	TickerInterval          time.Duration
	ECDSA_KEY               string
}

func ReadConfig() error {
	viper.SetConfigName("my") // name of config file (without extension)
	viper.SetEnvPrefix("cross-chain")
	viper.SetConfigType("yaml") // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("../config")
	viper.AddConfigPath("../../config") // path to look for the config file in
	viper.AddConfigPath("$HOME")        // call multiple times to add many search paths
	viper.AddConfigPath(".")            // optionally look for config in the working directory
	err := viper.ReadInConfig()         // Find and read the config file
	if err != nil {                     // Handle errors reading the config file
		//panic(fmt.Errorf("Fatal error config file: %s \n", err))
		return err
	}
	return nil
}

// LoadConfig loads config from files
func LoadConfig(configPaths ...string) error {
	v := viper.New()
	v.SetConfigName("my")
	v.SetConfigType("yaml")
	v.SetEnvPrefix("cross-chain")
	v.AutomaticEnv()
	v.AddConfigPath("../")
	v.AddConfigPath("./config")
	v.AddConfigPath("../config")
	v.AddConfigPath("../../config")
	for _, path := range configPaths {
		v.AddConfigPath(path)
		fmt.Print("path", path, "\n")
	}
	err := v.ReadInConfig()
	if err != nil {
		return err
	}

	return v.Unmarshal(&Config)
}

func LoadConfigAndArgs() (cfg *AppConfig) {
	cfg = NewConfig()
	LoadConfig(".")
	return cfg
}

func NewConfig() *AppConfig {
	c := AppConfig{}
	flag.IntVar(&c.PORT, "p", 8080, "Port number to serve")
	flag.DurationVar(
		&c.TickerInterval,
		"t",
		time.Minute,
		"Ticker interval for the adaptor to refresh supported trading pairs, suggested units: s, m, h")
	flag.Parse()
	return &c
}
