package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/elastic/go-elasticsearch/v7"
)

type SQLType string

const (
	MYSQL SQLType = "mysql"
	POSTGRES SQLType = "postgres"
)

type SQLConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database SQLType
}

func ConnectDB(sqlcfg SQLConfig, escfg elasticsearch.Config) (*sql.DB, *elasticsearch.Client) {
	sqlDB, err := connectSql(sqlcfg)
	if err!= nil {
		log.Fatalln("failed to connect to sql", err)
	}
	
	esClient, err := connectElasticsearch(escfg)
	if err!= nil {
		log.Fatalln("failed to connect to elasticsearch", err)
	}

	return sqlDB, esClient
}

func connectSql(cfg SQLConfig) (*sql.DB, error) {
	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/", cfg.User, cfg.Password,cfg.Host, cfg.Port)
	db, err := sql.Open(string(cfg.Database), url)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func connectElasticsearch(cfg elasticsearch.Config) (*elasticsearch.Client, error) {
	return elasticsearch.NewClient(cfg)
}
