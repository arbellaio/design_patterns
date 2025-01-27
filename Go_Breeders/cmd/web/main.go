package main

import (
	"database/sql"
	"flag"
	"fmt"
	"go-breeders/models"
	"html/template"
	"log"
	"net/http"
	"time"
)

const address = "localhost:5001"

type application struct {
	templateMap map[string]*template.Template
	config      appConfig
	DB          *sql.DB
	Models      models.Models
}

type appConfig struct {
	useCache bool
	dsn      string
}

func main() {

	app := application{
		templateMap: make(map[string]*template.Template),
	}
	flag.BoolVar(&app.config.useCache, "cache", false, "use template cache")
	flag.StringVar(&app.config.dsn, "dsn", "mariadb:mysql@tcp(localhost:3307)/breeders?parseTime=true&tls=false&collation=utf8_unicode_ci&timeout=5s", "DSN")
	flag.Parse()

	// get database
	db, err := initMySQLDB(app.config.dsn)
	if err != nil {
		log.Panic(err)
	}
	app.DB = db

	//we are connecting database by passing db
	app.Models = *models.New(db)

	srv := &http.Server{
		Addr:              address,
		Handler:           app.routes(),
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       30 * time.Second,
		ReadHeaderTimeout: 30 * time.Second,
		WriteTimeout:      30 * time.Second,
	}
	fmt.Println("starting web application on port: ", address)
	err = srv.ListenAndServe()
	if err != nil {
		log.Printf("error while starting server %v", err.Error())
	}
}
