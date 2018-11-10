package domain

type Organization struct {
	ID          string `json:"id"`
	Name        string `json:"name" validate:"required"`
	Email       string `json:"email" validate:"required,email"`
	Description string `json:"description"`
	Location    string `json:"location"`
	AvatarURL   string `json:"avatarUrl" validate:"url"`
}

// func (r Organization) Validate() error {
// 	if r.ID == "" {
// 		return errors.New("ID is empty")
// 	}

// 	if r.Name == "" {
// 		return errors.New("Name is empty")
// 	}

// 	// TODO: add proper email validation
// 	if r.Email == "" {
// 		return errors.New("Email is empty")
// 	}

// 	return nil
// }
