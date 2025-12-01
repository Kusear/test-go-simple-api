package database

import (
	"fmt"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseConnector struct {
	Host           string
	Port           int
	Username       string
	Password       string
	DBName         string
	ConnectTimeout int
}

// func (dbConnector DatabaseConnector) Connect() (*sql.DB, error) {
// 	dbConfig := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?%s",
// 		dbConnector.Username,
// 		dbConnector.Password,
// 		dbConnector.Host,
// 		dbConnector.Port,
// 		dbConnector.DBName,
// 		"sslmode=disable",
// 	)
// 	db, err := sql.Open("postgres", dbConfig)
// 	if err != nil {
// 		return nil, err
// 	}
// 	err = db.Ping()
// 	if err != nil {
// 		return nil, err
// 	}
// 	return db, nil
// }

func (dbConnector DatabaseConnector) Connect2() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		dbConnector.Host,
		dbConnector.Username,
		dbConnector.Password,
		dbConnector.DBName,
		dbConnector.Port,
		"disable",
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
