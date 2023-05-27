package models

import (
	"fmt"
	"os"
	"testing"

	"github.com/agilsyofian/kreditplus/config"
)

var testQueries *Database

func TestMain(m *testing.M) {

	cfg, err := config.LoadConfig("../")
	if err != nil {
		fmt.Println("cannot load config")
	}

	testDB := New(cfg)

	testQueries = testDB

	os.Exit(m.Run())
}
