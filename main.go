package main

import (
	"github.com/victorlui/gopportunities/config"
	"github.com/victorlui/gopportunities/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	//Initialize Configs
	err := config.Init()

	if err != nil {
		logger.ErrF("confgi inicialization error: %v", err)
		panic(err)
	}

	router.Initialize()
}
