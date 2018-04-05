package repository

import (
	"interviewr-server/domain"

	_ "github.com/lib/pq"
	gorp "gopkg.in/gorp.v1"
)

type postgresOrganizationRepository struct {
	connection *gorp.DbMap
}

func NewPostgresOrganizationRepository(connection *gorp.DbMap) OrganizationRepository {
	return &postgresOrganizationRepository{connection}
}

func (p *postgresOrganizationRepository) Fetch() ([]*domain.Organization, error) {
	orgs := make([]*domain.Organization, 0)
	_, err := p.connection.Select(&orgs, "select * from organization order by name")
	if err != nil {
		return nil, err
	}

	return orgs, nil
}

func (p *postgresOrganizationRepository) GetById(id string) (*domain.Organization, error) {
	org := &domain.Organization{}
	err := p.connection.SelectOne("select * from organization where id=?", id)
	if err != nil {
		return nil, err
	}

	return org, nil
}

func (p *postgresOrganizationRepository) Update(org *domain.Organization) (*domain.Organization, error) {
	return nil, nil
}

func (p *postgresOrganizationRepository) Store(org *domain.Organization) (string, error) {
	return "empty", nil
}

func (p *postgresOrganizationRepository) Delete(id string) (bool, error) {
	return true, nil
}
