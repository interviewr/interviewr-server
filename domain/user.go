package domain

import "errors"

// User entity
type User struct {
	ID        string `json:"id"`
	Login     string `json:"login"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Location  string `json:"location"`
	Bio       string `json:"bio"`
	AvatarURL string `json:"avatarUrl"`
}

// Validate user instance
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

	if ValidateEmail(u.Email) {
		return errors.New("Email is invalid")
	}

	if u.AvatarURL == "" && ValidateURL(u.AvatarURL) {
		return errors.New("Avatar URL is invalid")
	}

	return nil
}
