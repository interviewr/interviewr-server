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

func (p *postgresOrganizationRepository) GetById(id int64) (*domain.Organization, error) {
	org := &domain.Organization{}
	err := p.connection.SelectOne(org, "select * from organization where id = $1", id)
	if err != nil {
		return nil, err
	}

	return org, nil
}

func (p *postgresOrganizationRepository) Update(org *domain.Organization) (*domain.Organization, error) {
	_, err := p.connection.Exec(
		`update organization set name = $1, email = $2, description = $3, location = $4, avatarUrl = $5 where id = $6`,
		org.Name, org.Email, org.Description, org.Location, org.AvatarURL, org.ID,
	)

	if err != nil {
		return nil, err
	}

	return org, nil
}

func (p *postgresOrganizationRepository) Store(org *domain.Organization) (int64, error) {
	res, err := p.connection.Exec(
		`insert organization set name = $1, email = $2, description = $3, location = $4, avatarUrl = $5 where id = $6`,
		org.Name, org.Email, org.Description, org.Location, org.AvatarURL, org.ID,
	)

	if err != nil {
		return 0, err
	}

	return res.LastInsertId()
}

func (p *postgresOrganizationRepository) Delete(id int64) (bool, error) {
	_, err := p.connection.Exec("delete from organization where id = $1", id)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (p *postgresOrganizationRepository) GetMembers(id int64) ([]*domain.User, error) {
	members := make([]*domain.User, 0)
	_, err := p.connection.Select(
		&members,
		`select p.*
		from person p
		left join membership m on m.person_id = p.id
		left join organization o on m.organization_id = o.id
		where o.id = $1`,
		id,
	)
	if err != nil {
		return nil, err
	}

	return members, nil
}

func (p *postgresOrganizationRepository) CreateVacancy(companyId int64) (int64, error) {
	return 0, nil
}
