package main

import (
	"go-breeders/models"
	"log"
	"os"
	"testing"
)

var testApp application

func TestMain(m *testing.M) {
	dsn := "mariadb:mysql@tcp(localhost:3307)/breeders?parseTime=true&tls=false&collation=utf8_unicode_ci&timeout=5s"
	db, err := initMySQLDB(dsn)
	if err != nil {
		log.Panic(err)
	}

	testApp = application{
		DB:     db,
		Models: *models.New(db),
	}

	os.Exit(m.Run())
}
