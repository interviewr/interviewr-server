package models

type Vacancy struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
}