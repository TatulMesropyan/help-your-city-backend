package config

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"help-your-city-backend/models"

	_ "github.com/lib/pq"
)

type UserStorage interface {
	RegisterUser(*models.User) error
	SignInUser(email, password string) string
	DeleteUser(email, password string) error
	ChangePassword(email, password string) error
	ChangePhone(email, password string) error
}
type PostgresStore struct {
	db *sql.DB
}

func (store *PostgresStore) RegisterUser(user *models.User) error {
	query := `
        INSERT INTO users ("firstName", "lastName", email, birthday, password, phone)
        VALUES ($1, $2, $3, $4, $5, $6)
    `
	_, err := store.db.Exec(query, user.FirstName, user.LastName, user.Email, user.Birthday, user.Password, user.Phone)
	if err != nil {
		return err
	}
	return nil
}

func (store *PostgresStore) SignInUser(email, password string) (string, error) {
	var storedPassword string
	query := `SELECT password FROM users WHERE email = $1`
	err := store.db.QueryRow(query, email).Scan(&storedPassword)
	if err != nil {
		return "", err
	}

	if password != storedPassword {
		return "", errors.New("invalid password")
	}

	return "User authenticated successfully", nil
}

func (store *PostgresStore) ChangePassword(id, password string) (string, error) {
	var storedPassword string
	query := `SELECT password FROM users WHERE id = $1`
	err := store.db.QueryRow(query, id).Scan(&storedPassword)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	if password == storedPassword {
		return "", fmt.Errorf("old password and new password are the same")
	}
	changePassQuery := `UPDATE users SET password = $1 WHERE id = $2`
	_, err = store.db.Exec(changePassQuery, password, id)
	if err != nil {
		return "", err
	}
	return "Password changed successfully", nil
}
func (store *PostgresStore) ChangePhone(id string, phone string) (string, error) {
	var storedPhone string
	query := `SELECT password FROM users WHERE id = $1`
	err := store.db.QueryRow(query, id).Scan(&storedPhone)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	if phone == storedPhone {
		return "", fmt.Errorf("old phone and new phone are the same")
	}
	changePassQuery := `UPDATE users SET phone = $1 WHERE id = $2`
	_, err = store.db.Exec(changePassQuery, phone, id)
	if err != nil {
		return "", err
	}
	return "Phone successfully changed", nil
}
func ConnectDB() (*PostgresStore, error) {
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
	return &PostgresStore{
		db: db,
	}, nil
}

func (store *PostgresStore) createUserTable() error {
	query := `CREATE TABLE IF NOT EXISTS public.users
	(
		id SERIAL PRIMARY KEY,
		"firstName" character varying(100) COLLATE pg_catalog."default",
		"lastName" character varying(100) COLLATE pg_catalog."default",
		email character varying(100) COLLATE pg_catalog."default",
		birthday date,
		password character varying(100) COLLATE pg_catalog."default",
		phone character varying(25) COLLATE pg_catalog."default"
	)`

	_, err := store.db.Exec(query)
	return err
}
func InitUsers(store *PostgresStore) {
	store.createUserTable()
}
