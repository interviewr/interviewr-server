package domain

import (
	"errors"
	"regexp"
)

// Organization entity
type Organization struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Description string `json:"description"`
	Location    string `json:"location"`
}

// Validate instance
func (o Organization) Validate() error {
	if o.ID == "" {
		return errors.New("ID is empty")
	}

	if o.Name == "" {
		return errors.New("Name is empty")
	}

	Re := regexp.MustCompile(`^[a-z0-9._+\-]+@[a-z0-9.\-]+\.[a-z]{2,8}$`)

	if !Re.MatchString(o.Email) {
		return errors.New("Email is invalid")
	}

	return nil
}
