package usecases

import (
	"github.com/interviewr/interviewr-server/domain"
	"github.com/interviewr/interviewr-server/repository"
)

type OrganizationUsecase interface {
	Fetch() ([]*Organization, error)
	GetById(id string) (*Organization, error)
	Update(org *Organization) (*Organization, error)
	Store(org *Organization) (string, error)
	Delete(id string) (bool, error)
}

type organizationUsecase struct {
	orgRepository OrganizationRepository
}

func NewOrganizationUsecase(repo OrganizationRepository) OrganizationUsecase {
	return &organizationUsecase{repo}
}

func (uc *organizationUsecase) Fetch() ([]*Organization, error) {
	orgs, err := uc.orgRepository.Fetch()
	if err != nil {
		return nil, err
	}
	return orgs, nil
}

func (uc *organizationUsecase) GetById(id string) (*Organization, error) {
	return uc.orgRepository.GetById(id)
}

func (uc *organizationUsecase) Update(org *Organization) (*Organization, error) {}

func (uc *organizationUsecase) Store(org *Organization) (string, error) {}

func (uc *organizationUsecase) Delete(id string) (bool, error) {
	// Validation logic here (check if item by given id exists)
	return uc.orgRepository.Delete(id)
}