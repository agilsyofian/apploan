package models

import (
	"log"

	"github.com/agilsyofian/kreditplus/config"
	"gorm.io/gorm"

	"gorm.io/driver/mysql"
	"gorm.io/gorm/logger"
)

type Database struct {
	KreditPlus *gorm.DB
}

func New(cfg config.Config) *Database {
	conDBKreditPlus := NewConKreditPlus(cfg)

	return &Database{
		KreditPlus: conDBKreditPlus,
	}
}

func NewConKreditPlus(cfg config.Config) *gorm.DB {
	var err error

	logStatus := logger.Silent
	if cfg.Environment == "development" {
		logStatus = logger.Info
	}

	db, err := gorm.Open(mysql.Open(cfg.DBSource), &gorm.Config{
		Logger: logger.Default.LogMode(logStatus),
	})
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	return db
}
