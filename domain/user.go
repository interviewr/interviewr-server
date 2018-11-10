package domain

import "errors"

type User struct {
	ID        string `json:"id"`
	Login     string `json:"login"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Location  string `json:"location"`
	Bio       string `json:"bio"`
	AvatarURL string `json:"avatarUrl"`
}

// TODO: add proper fields validation
func (u User) Validate() error {
	if u.ID == "" {
		return errors.New("ID is empty")
	}

	if u.Login == "" {
		return errors.New("Login is empty")
	}

	if u.Name == "" {
		return errors.New("Name is empty")
	}

	if u.Email == "" {
		return errors.New("Email is empty")
	}

	return nil
}
