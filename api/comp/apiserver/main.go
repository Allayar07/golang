package main

import (
	"api/internal/app/apiserver"
	"flag"
	"log"

	"github.com/BurntSushi/toml"
)

var (
	pathConfig string
)

func init() {
	flag.StringVar(&pathConfig, "config-path", "configs/api.toml", "path to config")
}

func main() {

	flag.Parse()

	config := apiserver.NewConfigs()

	_, err := toml.DecodeFile(pathConfig, config)
	if err != nil{
		log.Fatal(err)
	}

	if err := apiserver.Start(config); err != nil {
		log.Fatal(err)
	}
}