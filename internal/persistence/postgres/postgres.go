package postgres

import (
	"database/sql"
	"fmt"
	"log"

	"go.uber.org/zap"

	"github.com/burkaydurdu/wish/config"
	_ "github.com/lib/pq"
)

const (
	wishTableName = "wish"
)

type Client struct {
	DB   *sql.DB
	conf *config.DatabaseConfig
}

func NewPostgresClient(conf *config.DatabaseConfig) (client *Client, err error) {
	connStr := fmt.Sprintf("postgresql://%s:%s@%s/%s?sslmode=%s",
		conf.Username, conf.Password, conf.Host, conf.Database, conf.SSLMode,
	)

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}

	client = &Client{
		DB:   db,
		conf: conf,
	}

	return
}

type repository struct {
	logger *zap.Logger
	client *Client
}
