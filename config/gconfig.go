package config

import (
	"fmt"
)



// Cfg is config struct
type Cfg struct {
	runMod           string
	serverAddr       string
	serverPort       string
	databaseAddr     string
	databasePort     string
	databaseAdmin    string
	databasePassword string
}

var globalCfg *Cfg

// SetCfg is the function for setting the global config from outside
func (c *Cfg) SetCfg(key string, vlaue interface{}) {

}

// GetCfg is the function for getting the global config from outside
func (c *Cfg) GetCfg(key string) {

}

func init() {
	defaultCfg := &Cfg{
		runMod:           "debug",
		serverAddr:       "127.0.0.1",
		serverPort:       "8080",
		databaseAddr:     "127.0.0.1",
		databasePort:     "6379",
		databaseAdmin:    "",
		databasePassword: "",
	}
	globalCfg = defaultCfg
	fmt.Println(defaultCfg)
}
