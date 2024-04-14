package database

import (
	"database/sql"
	"fmt"
	"help-your-city-backend/internal/config"
)

func ConnectDB(config *config.Config) (*sql.DB, error) {
	connectStr := fmt.Sprintf(
		"postgresql://%s:%s@%s/%s?sslmode=disable",
		config.PostgresConfig.Username,
		config.PostgresConfig.Password,
		config.PostgresConfig.Host,
		config.PostgresConfig.Name,
	)
	db, err := sql.Open("postgres", connectStr)
	if err != nil {
		return nil, fmt.Errorf("error opening database connection: %v", err)
	}

	defer db.Close()

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error pinging database: %v", err)
	}

	return db, nil
}
