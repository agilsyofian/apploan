package main

import (
	"fmt"

	"github.com/agilsyofian/golang/util"
	"github.com/agilsyofian/kreditplus/config"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type Migration struct {
	migrate *migrate.Migrate
}

func main() {

	cfg, err := config.LoadConfig(".")
	if err != nil {
		fmt.Println("cannot load config")
	}

	m, err := migrate.New(cfg.MigrationURL, cfg.DBDriver+"://"+cfg.DBSource)
	if err != nil {
		fmt.Println("cannot create new migrate instance")
	}

	migration := &Migration{
		migrate: m,
	}

	menu := util.NewMenu("Which migration do you want?")
	menu.AddItem("up", "up")
	menu.AddItem("down", "down")
	choice := menu.Display()

	switch choice {
	case "up":
		migration.Up()
	case "down":
		migration.Down()
	default:
		fmt.Println("migration doesn't exist. abort migration.")
		return
	}
}

func (m *Migration) Up() {
	if err := m.migrate.Up(); err != nil && err != migrate.ErrNoChange {
		fmt.Println("failed to run migrate up")
		return
	}
	fmt.Println("db migrated successfully")
}

func (m *Migration) Down() {
	if err := m.migrate.Down(); err != nil && err != migrate.ErrNoChange {
		fmt.Println(err)
		return
	}
	fmt.Println("db migrated successfully")
}
