package main

import (
	"fmt"
	"log"

	"github.com/agilsyofian/golang/pasetomaker"
	"github.com/agilsyofian/kreditplus/api"
	"github.com/agilsyofian/kreditplus/config"
	"github.com/agilsyofian/kreditplus/models"
)

func main() {

	cfg, err := config.LoadConfig(".")
	if err != nil {
		fmt.Println("cannot load config")
	}

	tokenMaker, err := pasetomaker.NewPasetoMaker(cfg.TokenSymmetricKey)
	if err != nil {
		log.Fatal("cannot create token maker: %w", err)
	}

	store := models.New(cfg)
	server, err := api.NewServer(cfg, store, tokenMaker)
	if err != nil {
		log.Fatal("cannot start server")
	}

	err = server.Start(cfg.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server")
	}

}
