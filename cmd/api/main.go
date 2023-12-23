package main

import (
	"log"

	"github.com/athunlal/Note-Taking-Application/pkg/config"
	"github.com/athunlal/Note-Taking-Application/pkg/di"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Unable to load the config file : ", err)
	}
	httpServer, err := di.InitApi(cfg)

	httpServer.Engine.Run(cfg.PORT)
}
