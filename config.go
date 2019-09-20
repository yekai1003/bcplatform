package main

import (
	"fmt"
	"os"
)

var (
	Version   = "1.0.0"
	Commit    = "yekai"
	BuildTime = "2019-09-18"
)

type ServerConfig struct {
	Eth   *ChainConfig
	Fisco *ChainConfig
	Eos   *ChainConfig
}

type ChainConfig struct {
	Connstr string
}

func usage() {
	fmt.Printf("Usage: %s -c config_file [-v] [-h]\n", "bcplatform")
	os.Exit(1)
}

var Config *ServerConfig //引用配置文件结构

func init() {
	fmt.Println("进程初始化成功")
	Config = GetConfig()
}

func GetConfig() (config *ServerConfig) {

	// var configFile = flag.String("c", "", "Config file")

	// var ver = flag.Bool("v", false, "version")
	// var help = flag.Bool("h", false, "Help")

	// flag.Usage = usage
	// flag.Parse()

	// if *help {
	// 	usage()
	// 	return nil
	// }

	// if *ver {
	// 	fmt.Println("Version: ", Version)
	// 	fmt.Println("Commit: ", Commit)
	// 	fmt.Println("BuildTime: ", BuildTime)
	// 	return nil
	// }
	// // get server config
	// if *configFile != "" {
	// 	config = &ServerConfig{}
	// 	if _, err := toml.DecodeFile(*configFile, &config); err != nil {
	// 		panic(err)
	// 	}
	// } else {
	// 	usage()
	// 	return nil
	// }
	config = &ServerConfig{
		&ChainConfig{"http://localhost:8546"},
		&ChainConfig{"http://localhost:8545"},
		&ChainConfig{"http://localhost:8547"},
	}
	return config
}
