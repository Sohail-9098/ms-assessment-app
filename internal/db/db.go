package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type dbConfig struct {
	User, Password, DbName, SslMode string
}

func Connect() *sql.DB {
	config := dbConfig{"postgres", "pg@123", "ms_assessment_app", "disable"}
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s", config.User, config.Password, config.DbName, config.SslMode)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("failed to connect to DB: %v", err)
	}
	return db
}

func Close(db *sql.DB) {
	db.Close()
}
