package models

type User struct {
	Id int `json:"id"`
	Username string `json:"username"`
	Name string `json:"name"`
	Surname string `json:"surname"`
	Email string `json:"email"`
}