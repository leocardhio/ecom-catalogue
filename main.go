package main

import (
	"database/sql"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/leocardhio/ecom-catalogue/db"
)

var (
	sqldb *sql.DB
	esdb *elasticsearch.Client
)


func main() {
	sqlcfg := db.SQLConfig{
		Host:     "localhost",
		Port:     "5432",
		User:     "postgres",
		Password: "password",
		Database: db.POSTGRES,
	}

	escfg := elasticsearch.Config{
		Addresses: []string{"http://localhost:9200"},
	}

	sqldb, esdb = db.ConnectDB(sqlcfg, escfg)
	defer sqldb.Close()
}