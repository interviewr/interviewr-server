package repository

import (
	"interviewr-server/domain"

	_ "github.com/lib/pq"
	gorp "gopkg.in/gorp.v1"
)

type postgresVacancyRepository struct {
	connection *gorp.DbMap
}

func NewPostgresVacancyRepository(connection *gorp.DbMap) VacancyRepository {
	return &postgresVacancyRepository{connection}
}

func (p *postgresVacancyRepository) Fetch() ([]*domain.Vacancy, error) {
	vacancies := make([]*domain.Vacancy, 0)
	_, err := p.connection.Select(&vacancies, `
		select id, title, description, location
		from vacancy
		order by title
	`)
	if err != nil {
		return nil, err
	}

	return vacancies, nil
}

func (p *postgresVacancyRepository) GetById(id string) (*domain.Vacancy, error) {
	vacancy := &domain.Vacancy{}
	err := p.connection.SelectOne(vacancy, "select * from vacancy where id = $1", id)
	if err != nil {
		return nil, err
	}

	return vacancy, nil
}

func (p *postgresVacancyRepository) GetFollowers(id string) ([]*domain.User, error) {
	followers := make([]*domain.User, 0)
	_, err := p.connection.Select(&followers, `
		select p.*
		from person p
			left join applicant a on a.person_id = p.id
			left join vacancy v on a.vacancy_id = v.id
		where v.id = $1
	`, id)
	if err != nil {
		return nil, err
	}

	return followers, nil
}
