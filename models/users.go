package models

type User struct {
	id        int    `json:"id"`
	firstName string `json:"firstName"`
	lastName  string `json:"lastName"`
	email     string `json:"email"`
	birthday  string `json:"birthday"`
	phone     int    `json:"phone"`
	password  string `json:"password"`
}
