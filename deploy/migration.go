package main

import (
	"fmt"

	"github.com/agilsyofian/kreditplus/config"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	cfg, err := config.LoadConfig(".")
	if err != nil {
		fmt.Println("cannot load config")
	}

	m, err := migrate.New(cfg.MigrationURL, cfg.DBDriver+"://root:root@tcp(127.0.0.1:3306)/kreditplus?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("cannot create new migrate instance")
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		fmt.Println("failed to run migrate up")
		return
	}
	fmt.Println("db migrated successfully")
}
