package usecases

import (
	"github.com/interviewr/interviewr-server/domain"
	"github.com/interviewr/interviewr-server/repository"
)

type OrganizationUsecase interface {
	Fetch() ([]*Organization, error)
	GetById(id int64) (*Organization, error)
	Update(org *Organization) (*Organization, error)
	Store(org *Organization) (int64, error)
	Delete(id int64) (bool, error)
}

type organizationUsecase struct {
	orgRepository OrganizationRepository
}

func NewOrganizationUsecase(repo OrganizationRepository) OrganizationUsecase {
	return &organizationUsecase{repo}
}

func (uc *organizationUsecase) Fetch() ([]*Organization, error) {}

func (uc *organizationUsecase) GetById(id int64) (*Organization, error) {
	return uc.orgRepository.GetById(id)
}

func (uc *organizationUsecase) Update(org *Organization) (*Organization, error) {}

func (uc *organizationUsecase) Store(org *Organization) (int64, error) {}

func (uc *organizationUsecase) Delete(id int64) (bool, error) {
	// Validation logic here (check if item by given id exists)
	return uc.orgRepository.Delete(id)
}