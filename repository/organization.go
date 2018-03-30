package repository

import (
	"interviewr-server/domain"

	"github.com/go-gorp/gorp"
	_ "github.com/lib/pq"
)

type OrganizationRepository interface {
	Fetch() ([]*domain.Organization, error)
	GetById(id string) (*domain.Organization, error)
	Update(org *domain.Organization) (*domain.Organization, error)
	Store(org *domain.Organization) (string, error)
	Delete(id string) (bool, error)
}

type organizationRepository struct {
	DbMap *gorp.DbMap
}

func NewOrganizationRepository(DbMap *gorp.DbMap) OrganizationRepository {
	return &organizationRepository{DbMap}
}

func (repo *organizationRepository) Fetch() ([]*domain.Organization, error) {
	orgs := make([]*domain.Organization, 0)
	_, err := repo.DbMap.Select(&orgs, "select * from organization order by name")
	if err != nil {
		return nil, err
	}

	return orgs, nil
}

func (repo *organizationRepository) GetById(id string) (*domain.Organization, error) {
	org := &domain.Organization{}
	err := repo.DbMap.SelectOne("select * from organization where id=?", id)
	if err != nil {
		return nil, err
	}

	return org, nil
}

func (repo *organizationRepository) Update(org *domain.Organization) (*domain.Organization, error) {
	return nil, nil
}

func (repo *organizationRepository) Store(org *domain.Organization) (string, error) {
	return "nil", nil
}

func (repo *organizationRepository) Delete(id string) (bool, error) {
	return true, nil
}
