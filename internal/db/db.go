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

func connect() *sql.DB {
	config := dbConfig{"postgres", "pg@123", "ms_assessment_app", "disable"}
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s", config.User, config.Password, config.DbName, config.SslMode)
	conn, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("failed to connect to DB: %v", err)
	}
	return conn
}

func close(conn *sql.DB) {
	conn.Close()
}
