package repository

import "interviewr-server/domain"

type UserRepository interface {
	Fetch() ([]*domain.User, error)
	GetById(id string) (*domain.User, error)
	Update(person *domain.User) (*domain.User, error)
	Store(person *domain.User) (string, error)
	Delete(id string) (bool, error)
}

type OrganizationRepository interface {
	Fetch() ([]*domain.Organization, error)
	GetById(id int64) (*domain.Organization, error)
	Update(org *domain.Organization) (*domain.Organization, error)
	Store(org *domain.Organization) (int64, error)
	Delete(id int64) (bool, error)
	GetMembers(id int64) ([]*domain.User, error)
	CreateVacancy(companyId int64) (int64, error)
}

type VacancyRepository interface {
	Fetch() ([]*domain.Vacancy, error)
	GetById(id string) (*domain.Vacancy, error)
	GetFollowers(id string) ([]*domain.User, error)
	// GetTest(vacancyId string) (*domain.Test, error)
}

// var (
// 	INTERNAL_SERVER_ERROR = errors.New("Internal Server Error")
// 	NOT_FOUND_ERROR       = errors.New("Requested Item not found")
// 	CONFLICT_ERROR        = errors.New("Item already exist")
// )
