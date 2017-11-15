package repository

import (
	"errors"
	"github.com/interviewr/interviewr-server/domain"
)

type OrganizationRepository interface {
	Fetch() ([]*Organization, error)
	GetById(id string) (*Organization, error)
	Update(org *Organization) (*Organization, error)
	Store(org *Organization) (string, error)
	Delete(id string) (bool, error)
}

var (
	INTERNAL_SERVER_ERROR = errors.New("Internal Server Error")
	NOT_FOUND_ERROR = errors.New("Requested Item not found")
	CONFLICT_ERROR = errors.New("Item already exist")
)