package postgres

import (
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

func (repo *organizationRepository) Fetch() ([]*Organization, error) {}

func (repo *organizationRepository) GetById(id int64) (*Organization, error) {}

func (repo *organizationRepository) Update(org *Organization) (*Organization, error) {}

func (repo *organizationRepository) Store(org *Organization) (int64, error) {}

func (repo *organizationRepository) Delete(id int64) (bool, error) {}