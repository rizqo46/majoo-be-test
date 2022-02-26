package main

import (
	"majoo-be-test/config"
	"majoo-be-test/server"
)

func main() {
	cfg, err := config.InitConfig()
	if err != nil {
		panic("fail to initialize config")
	}
	server := server.InitServer(cfg)
	server.Run()
}
