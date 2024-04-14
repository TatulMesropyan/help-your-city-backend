package service

import (
	userStore "help-your-city-backend/internal/database/pg/user"
	model "help-your-city-backend/internal/model"
)

type UserService interface {
	RegisterUser(firstName, lastName, email, birthday, password, phone string) (string, error)
	ChangePhone(id, phone string) (string, error)
	ChangePassword(id, password string) (string, error)
	SignIn(email, password string) (string, error)
}

type DBStorage struct {
	store *userStore.PostgresStore
}

func NewUserService(store *userStore.PostgresStore) *DBStorage {
	return &DBStorage{store: store}
}

func (db *DBStorage) RegisterUser(firstName, lastName, email, birthday, password, phone string) (string, error) {
	message, err := db.store.RegisterUser(&model.User{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Birthday:  birthday,
		Password:  password,
		Phone:     phone,
	})
	return message, err
}

func (db *DBStorage) ChangePhone(id, phone string) (string, error) {
	message, err := db.store.ChangePhone(id, phone)
	return message, err
}

func (db *DBStorage) ChangePassword(id, password string) (string, error) {
	message, err := db.store.ChangePassword(id, password)
	return message, err
}
func (db *DBStorage) SignIn(email, password string) (string, error) {
	message, err := db.store.SignInUser(email, password)
	return message, err
}
