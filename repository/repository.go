package repository

import (
	"errors"
	domain "interviewr-server/domain"
)

type OrganizationRepository interface {
	Fetch() ([]*domain.Organization, error)
	GetById(id string) (*domain.Organization, error)
	Update(org *domain.Organization) (*domain.Organization, error)
	Store(org *domain.Organization) (string, error)
	Delete(id string) (bool, error)
}

var (
	INTERNAL_SERVER_ERROR = errors.New("Internal Server Error")
	NOT_FOUND_ERROR = errors.New("Requested Item not found")
	CONFLICT_ERROR = errors.New("Item already exist")
)
