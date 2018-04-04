package domain

import "errors"

type Organization struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Description string `json:"description"`
	Location    string `json:"location"`
}

func (r Organization) Validate() error {
	if r.ID == "" {
		return errors.New("ID is empty")
	}

	if r.Name == "" {
		return errors.New("Name is empty")
	}

	// TODO: add proper email validation
	if r.Email == "" {
		return errors.New("Email is empty")
	}

	return nil
}
