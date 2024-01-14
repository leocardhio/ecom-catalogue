package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/elastic/go-elasticsearch/v7"
	_ "github.com/lib/pq"
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
	Database string
	Driver SQLType
}

type Database struct {
	sql *sql.DB
	es *elasticsearch.Client
}

func NewDatabase() *Database {
	return &Database{}
}

func (db *Database) Connect(sqlcfg SQLConfig, escfg elasticsearch.Config) *Database {
	sqlDB, err := connectSql(sqlcfg)
	if err!= nil {
		log.Fatalln("failed to connect to sql", err)
	}
	
	// TODO: Uncomment when ready to connect to elasticsearch
	// esClient, err := connectElasticsearch(escfg)
	// if err!= nil {
	// 	log.Fatalln("failed to connect to elasticsearch", err)
	// }

	db.sql = sqlDB
	// db.es = esClient

	return db
}

func (db *Database) Close() {
	db.sql.Close()
}

func (db *Database) GetSQL() *sql.DB {
	return db.sql
}

func (db *Database) GetES() *elasticsearch.Client {
	return db.es
}

func connectSql(cfg SQLConfig) (*sql.DB, error) {
	url := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", cfg.User, cfg.Password,cfg.Host, cfg.Port, cfg.Database)
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func connectElasticsearch(cfg elasticsearch.Config) (*elasticsearch.Client, error) {
	return elasticsearch.NewClient(cfg)
}
