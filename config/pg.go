package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func ConnectDB() (*sql.DB, error) {
	const (
		PORT        = 5432
		DB_PASSWORD = "adrine2009"
		DB_NAME     = "help_city_data"
		DB_USERNAME = "postgres"
		DB_HOST     = "0.0.0.0"
	)
	connectStr := fmt.Sprintf(
		"postgresql://%s:%s@%s/%s?sslmode=disable",
		DB_USERNAME,
		DB_PASSWORD,
		DB_HOST,
		DB_NAME,
	)

	db, err := sql.Open("postgres", connectStr)
	if err != nil {
		return nil, fmt.Errorf("error opening database connection: %v", err)
	}

	if err := db.Ping(); err != nil {
		defer db.Close()
		return nil, fmt.Errorf("error pinging database: %v", err)
	}

	log.Println("Connected to the database")
	return db, nil
}
