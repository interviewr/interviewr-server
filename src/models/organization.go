package models

type Organization struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Description string `json:"description"`
	Location string `json:"location"`
}