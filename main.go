package main

import (
	"fmt"
	"log"

	"github.com/agilsyofian/golang/pasetomaker"
	"github.com/agilsyofian/kreditplus/api"
	"github.com/agilsyofian/kreditplus/config"
	"github.com/agilsyofian/kreditplus/models"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {

	cfg, err := config.LoadConfig(".")
	if err != nil {
		fmt.Println("cannot load config")
	}

	go runDBMigration(cfg.MigrationURL, cfg.DBDriver+"://"+cfg.DBSource)

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

func runDBMigration(migrationURL string, dbSource string) {
	fmt.Println(migrationURL)
	fmt.Println(dbSource)
	migration, err := migrate.New(migrationURL, dbSource)
	if err != nil {
		fmt.Println("cannot create new migrate instance")
	}
	migration.Up()
}
