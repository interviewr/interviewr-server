package repository

import (
	github.com/interviewr/interviewr-server/domain
)

type OrganizationRepository interface {
	Fetch() ([]*Organization, error)
	GetById(id int64) (*Organization, error)
	Update(org *Organization) (*Organization, error)
	Store(org *Organization) (int64, error)
	Delete(id int64) (bool, error)
}