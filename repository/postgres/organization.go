package postgres

import (
	"errors"
	"database/sql"
	"github.com/lib/pq"
	"github.com/interviewr/interviewr-server/domain"
	"github.com/interviewr/interviewr-server/repository"
)

type organizationRepository struct {
	Conn *sql.DB
}

func NewOrganizationRepository(Conn *sql.DB) OrganizationRepository {
	return &organizationRepository{Conn}
}

func (repo *organizationRepository) Fetch() ([]*Organization, error) {
	orgs := make([]*Organization, 0)
	rows, err := repo.Conn.Query(`SELECT id, name, email, description, location FROM organization WHERE id > ?`)
	if err != nil {
		return nil, repository.INTERNAL_SERVER_ERROR
	}
	defer rows.Close()
	for rows.Next() {
		t := new(Organization)
		err = rows.Scan(&org.ID, &org.Name, &org.Email, &org.Description, &org.Location)
		if err != nil {
			return nil, repository.INTERNAL_SERVER_ERROR
		}
		orgs = append(orgs, t)
	}

	return orgs, nil
}

func (repo *organizationRepository) GetById(id string) (*Organization, error) {
	org := &Organization{}
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

func (repo *organizationRepository) Update(org *Organization) (*Organization, error) {}

func (repo *organizationRepository) Store(org *Organization) (string, error) {}

func (repo *organizationRepository) Delete(id string) (bool, error) {}