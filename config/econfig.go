package config

import (
	"fmt"
	"os"
)

// CfgFile is the function of confgi file

// EnvCfg is the interface for multi-env function
type EnvCfg interface {
	ParseCfg(string)
}

// EnvList is map interface
type EnvList map[string]EnvCfg

var envList EnvList

// AppPath is the var of app path
var AppPath string

var runMod string

func init() {
	workPath, err := os.Getwd()
	if err != nil {
		panic("error")
	}
	fmt.Println(workPath)
	runMod = os.Getenv("GIN_MODE")
	if len(runMod) == 0 {
		runMod = "debug"
	}
	fmt.Println(runMod, "~~~~~")
	envList = make(EnvList)
}

//Register is the function for get different config
func Register(envFile interface{}) {
	switch v := envFile.(type) {
	case CustomCfg:
		envList[v.Filename] = v
	}
	fmt.Println("this is enviroment register")
}
