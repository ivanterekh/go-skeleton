package main

import (
	"fmt"
	"log"

	envcfg "github.com/caarlos0/env"
	"github.com/ivanterekh/go-skeleton/env"
	"github.com/joho/godotenv"

	"github.com/ivanterekh/go-skeleton/server"
)

type config struct {
	Address string `env:"ADDRESS" envDefault:":8080"`
}

func main() {
	cfg, err := initConfig()
	if err != nil {
		log.Fatalf("could not init config: %v\n", err) // TODO: change logging
	}

	log.Printf("starting server on %v", cfg.Address) // TODO: change logging
	server.Run(cfg.Address)
}

func initConfig() (*config, error) {
	var cfgName string
	switch {
	case env.IsProd():
		cfgName = ".env.prod"
	case env.IsStaging():
		cfgName = ".env.staging"
	default:
		cfgName = ".env.dev"
	}
	if err := godotenv.Load(cfgName); err != nil {
		log.Printf("could not load config from file %v, will read from env vars\n", cfgName) // TODO: change logging
	}

	cfg := new(config)
	err := envcfg.Parse(cfg)
	if err != nil {
		return nil, fmt.Errorf("could not parse config: %v", err)
	}

	return cfg, nil
}
