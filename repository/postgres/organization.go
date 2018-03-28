package postgres

import (
	"database/sql"
	_ "github.com/lib/pq"
	domain "interviewr-server/domain"
	"interviewr-server/repository"
)

type organizationRepository struct {
	Conn *sql.DB
}

func NewOrganizationRepository(Conn *sql.DB) repository.OrganizationRepository {
	return &organizationRepository{Conn}
}

func (repo *organizationRepository) Fetch() ([]*domain.Organization, error) {
	orgs := make([]*domain.Organization, 0)
	rows, err := repo.Conn.Query(`SELECT id, name, email, description, location FROM organization`)
	if err != nil {
		return nil, repository.INTERNAL_SERVER_ERROR
	}
	defer rows.Close()
	for rows.Next() {
		org := new(domain.Organization)
		err = rows.Scan(&org.ID, &org.Name, &org.Email, &org.Description, &org.Location)
		if err != nil {
			return nil, repository.INTERNAL_SERVER_ERROR
		}
		orgs = append(orgs, org)
	}

	return orgs, nil
}

func (repo *organizationRepository) GetById(id string) (*domain.Organization, error) {
	org := &domain.Organization{}
	rows, err := repo.Conn.Query(`SELECT id, name, email, description, location FROM organization WHERE id = $1`, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if !rows.Next() {
		return nil, repository.NOT_FOUND_ERROR
	}
	err = rows.Scan(&org.ID, &org.Name, &org.Email, &org.Description, &org.Location)
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
