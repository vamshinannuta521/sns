package dataaccess

import (
	"database/sql"
	"errors"
)

type PostgresClient struct {
    DB *sql.DB
}

func(pgClient *PostgresClient) OpenDb(dbname string, dbuser string) error{
	db, err := sql.Open("postgres", "user="+dbuser+" dbname="+dbname+" sslmode=disable")
	pgClient.DB = db
    return err
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