package main

import (
	"fmt"
	"github.com/ivanterekh/go-skeleton/server"
	"github.com/spf13/viper"
	"github.com/xlab/closer"
	"log"
)

func main() {
	err := initConfig()
	if err != nil {
		log.Fatal(err)
	}

	server.Start(":" + viper.GetString("port"))
	closer.Hold()
}

func initConfig() error {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		return fmt.Errorf("could not read config file: %v", err)
	}
	return nil
}
