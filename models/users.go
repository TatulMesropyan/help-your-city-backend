package models

type User struct {
	Id        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Birthday  string `json:"birthday"`
	Phone     int    `json:"phone"`
	Password  string `json:"password"`
}
