package domain

import (
	"errors"
)

// Organization entity
type Organization struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Description string `json:"description"`
	Location    string `json:"location"`
}

// Validate organization instance
func (o Organization) Validate() error {
	if o.ID == "" {
		return errors.New("ID is empty")
	}

	if o.Name == "" {
		return errors.New("Name is empty")
	}

	if ValidateEmail(o.Email) {
		return errors.New("Email is invalid")
	}

	return nil
}
