package usecases

import (
	"interviewr-server/domain"
	"interviewr-server/repository"
)

type OrganizationUsecase interface {
	Fetch() ([]*domain.Organization, error)
	GetById(id string) (*domain.Organization, error)
	Update(org *domain.Organization) (*domain.Organization, error)
	Store(org *domain.Organization) (string, error)
	Delete(id string) (bool, error)
}

type organizationUsecase struct {
	orgRepository repository.OrganizationRepository
}

func NewOrganizationUsecase(repo repository.OrganizationRepository) OrganizationUsecase {
	return &organizationUsecase{repo}
}

func (uc *organizationUsecase) Fetch() ([]*domain.Organization, error) {
	orgs, err := uc.orgRepository.Fetch()
	if err != nil {
		return nil, err
	}
	return orgs, nil
}

func (uc *organizationUsecase) GetById(id string) (*domain.Organization, error) {
	org, err := uc.orgRepository.GetById(id)
	if err != nil {
		return nil, err
	}
	return org, nil
}

func (uc *organizationUsecase) Update(org *domain.Organization) (*domain.Organization, error) {
	return nil, nil
}

func (uc *organizationUsecase) Store(org *domain.Organization) (string, error) {
	return "nil", nil
}

func (uc *organizationUsecase) Delete(id string) (bool, error) {
	// Validation logic here (check if item by given id exists)
	return uc.orgRepository.Delete(id)
}
