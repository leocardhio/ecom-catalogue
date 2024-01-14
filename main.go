package main

import (
	"fmt"
	"log"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/leocardhio/ecom-catalogue/config"
	"github.com/leocardhio/ecom-catalogue/db"
	"github.com/leocardhio/ecom-catalogue/router"
)


func main() {
	cfg := config.NewConfig().Load()
	
	sqlcfg := db.SQLConfig{
		Host:     cfg.SQLHost,
		Port:     cfg.SQLPort,
		User:     cfg.SQLUser,
		Password: cfg.SQLPassword,
		Database: cfg.SQLDatabase,
		Driver	: db.SQLType(cfg.SQLDriver),
	}

	escfg := elasticsearch.Config{
		Addresses: []string{fmt.Sprintf("http://%s:%s", cfg.ESHost, cfg.ESPort)},
	}

	dbs := db.NewDatabase().Connect(sqlcfg, escfg)
	defer dbs.Close()

	r := router.NewRouter(*dbs)
	if err := r.Run(); err != nil {
		log.Fatal("failed to run server", err)
	}
}