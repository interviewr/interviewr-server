package domain

import (
	"context"
	"errors"
	"sync"
)

type Organization struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Description string `json:"description"`
	Location string `json:"location"`
}