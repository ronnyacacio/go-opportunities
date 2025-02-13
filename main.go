package main

import (
	"github.com/ronnyacacio/go-opportunities/config"
	"github.com/ronnyacacio/go-opportunities/routes"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")

	err := config.Init()
	if err != nil {
		logger.Error("Erro")
		return
	}

	routes.Initialize()
}
