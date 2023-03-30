package main

import (
	"github.com/shrtkv/gopportunities/config"
	"github.com/shrtkv/gopportunities/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	err := config.Init()
	if err == nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}
	router.Initialize()
}
