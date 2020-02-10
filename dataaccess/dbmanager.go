package dataaccess

import (
	"database/sql"
	"errors"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

var logger *logrus.Entry

type PostgresClient struct {
	DB *sql.DB
}

func NewClient(log *logrus.Entry) (*PostgresClient, error) {
	logger = log

	db, err := sql.Open("postgres", "user=postgres dbname=sns host=localhost port=5432 sslmode=disable")
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return &PostgresClient{
		DB: db,
	}, nil
}

func (cl *PostgresClient) Close() (err error) {

	if cl.DB == nil {
		return
	}

	if err = cl.DB.Close(); err != nil {
		err = errors.New("Errored closing database connection")
	}
	return
}
