package main

import (
	"log"

	"github.com/alwindoss/duka/internal/engine"
)

func main() {
	cfg := new(engine.Config)
	cfg.Addr = ":9090"
	log.Fatal(engine.Start(cfg))
}
