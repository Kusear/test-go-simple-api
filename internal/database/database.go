package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type DatabaseConnector struct {
	Host           string
	Port           int
	Username       string
	Password       string
	DBName         string
	ConnectTimeout int
}

func (dbConnector DatabaseConnector) Connect() (*sql.DB, error) {

	dbConfig := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?%s",
		dbConnector.Username,
		dbConnector.Password,
		dbConnector.Host,
		dbConnector.Port,
		dbConnector.DBName,
		"sslmode=disable",
	)

	db, err := sql.Open("postgres", dbConfig)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
