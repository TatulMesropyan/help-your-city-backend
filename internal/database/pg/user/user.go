package database

import (
	"database/sql"
	"fmt"

	model "help-your-city-backend/internal/model"

	_ "github.com/lib/pq"
)

type UserStorage interface {
	RegisterUser(*model.User) (string, error)
	SignInUser(email, password string) (string, error)
	ChangePassword(id, password string) (string, error)
	ChangePhone(id, phone string) (string, error)
}

type PostgresStore struct {
	db *sql.DB
}

func UserRepo(db *sql.DB) *PostgresStore {
	return &PostgresStore{db}
}

const (
	insertUser = `
	INSERT INTO users ("firstName", "lastName", email, birthday, password, phone)
	VALUES ($1, $2, $3, $4, $5, $6)
`
	getUser        = `SELECT password FROM users WHERE email = $1`
	getPassword    = `SELECT password FROM users WHERE id = $1`
	getPhone       = `SELECT phone FROM users WHERE id = $1`
	updatePassword = `UPDATE users SET password = $1 WHERE id = $2`
	updatePhone    = `UPDATE users SET phone = $1 WHERE id = $2`
)

func (store *PostgresStore) RegisterUser(user *model.User) (string, error) {
	_, err := store.db.Exec(insertUser, user.FirstName, user.LastName, user.Email, user.Birthday, user.Password, user.Phone)
	if err != nil {
		return "", err
	}
	return "Successfully registered", nil
}

func (store *PostgresStore) SignInUser(email, password string) (string, error) {
	var storedPassword string
	err := store.db.QueryRow(getUser, email).Scan(&storedPassword)
	if err != nil {
		return "", err

	}

	if password != storedPassword {
		return "", nil
	}

	return "Successfully signed in", nil
}

func (store *PostgresStore) ChangePassword(id, password string) (string, error) {
	var storedPassword string
	err := store.db.QueryRow(getPassword, id).Scan(&storedPassword)
	if err != nil {
		return "", err
	}
	if password == storedPassword {
		return "", fmt.Errorf("old password and new password are the same")
	}
	_, err = store.db.Exec(updatePassword, password, id)
	if err != nil {
		return "", err
	}
	return "Password changed successfully", nil
}
func (store *PostgresStore) ChangePhone(id string, phone string) (string, error) {
	var storedPhone string
	err := store.db.QueryRow(getPhone, id).Scan(&storedPhone)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	if phone == storedPhone {
		return "", fmt.Errorf("old phone and new phone are the same")
	}
	_, err = store.db.Exec(updatePhone, phone, id)
	if err != nil {
		return "", err
	}
	return "Phone successfully changed", nil
}

// func (store *PostgresStore) createUserTable() error {
// 	query := `CREATE TABLE IF NOT EXISTS public.users
// 	(
// 		id SERIAL PRIMARY KEY,
// 		"firstName" character varying(100) COLLATE pg_catalog."default",
// 		"lastName" character varying(100) COLLATE pg_catalog."default",
// 		email character varying(100) COLLATE pg_catalog."default",
// 		birthday date,
// 		password character varying(100) COLLATE pg_catalog."default",
// 		phone character varying(25) COLLATE pg_catalog."default"
// 	)`

// 	_, err := store.db.Exec(query)
// 	return err
// }

// func InitUsers(store *PostgresStore) {
// 	store.createUserTable()
// }
