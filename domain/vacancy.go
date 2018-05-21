package domain

import "errors"

type Vacancy struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Location    string `json:"location"`
}

func (v Vacancy) Validate() error {
	if v.ID == "" {
		return errors.New("ID is empty")
	}

	if v.Title == "" {
		return errors.New("Title is empty")
	}

	return nil
}
