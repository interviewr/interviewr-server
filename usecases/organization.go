package usecases

import (
	"interviewr-server/domain"
	"interviewr-server/repository"
)

type OrganizationUsecase interface {
	Fetch() ([]*domain.Organization, error)
	GetById(id int64) (*domain.Organization, error)
	Update(org *domain.Organization) (*domain.Organization, error)
	Store(org *domain.Organization) (int64, error)
	Delete(id int64) (bool, error)
	GetMembers(id int64) ([]*domain.User, error)
	CreateVacancy(companyId int64) (int64, error)
}

type organizationUsecase struct {
	repository repository.OrganizationRepository
}

func NewOrganizationUsecase(repo repository.OrganizationRepository) *organizationUsecase {
	return &organizationUsecase{repository: repo}
}

func (interactor *organizationUsecase) Fetch() ([]*domain.Organization, error) {
	orgs, err := interactor.repository.Fetch()
	if err != nil {
		return nil, err
	}

	return orgs, nil
}

func (interactor *organizationUsecase) GetById(id int64) (*domain.Organization, error) {
	org, err := interactor.repository.GetById(id)
	if err != nil {
		return nil, err
	}

	return org, nil
}

func (interactor *organizationUsecase) Update(org *domain.Organization) (*domain.Organization, error) {
	org, err := interactor.repository.Update(org)
	if err != nil {
		return nil, err
	}

	return org, nil
}

func (interactor *organizationUsecase) Store(org *domain.Organization) (int64, error) {
	id, err := interactor.repository.Store(org)
	if err != nil {
		return 0, nil
	}

	return id, nil
}

func (interactor *organizationUsecase) Delete(id int64) (bool, error) {
	// Validation logic here (check if item by given id exists)
	return interactor.repository.Delete(id)
}

func (interactor *organizationUsecase) GetMembers(id int64) ([]*domain.User, error) {
	members, err := interactor.repository.GetMembers(id)
	if err != nil {
		return nil, err
	}

	return members, nil
}

func (interactor *organizationUsecase) CreateVacancy(companyId int64) (int64, error) {
	return 0, nil
}
