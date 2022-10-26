package main

import (
	"log"

	"github.com/youshimitsu/mole/config"
	"github.com/youshimitsu/mole/internal/app"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("TunnelConfig error: %s", err)
	}

	app.Run(cfg)
}
