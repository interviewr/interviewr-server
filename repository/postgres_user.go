package repository

import (
	"interviewr-server/domain"

	_ "github.com/lib/pq"
	gorp "gopkg.in/gorp.v1"
)

type postgresUserRepository struct {
	connection *gorp.DbMap
}

func NewPostgresUserRepository(connection *gorp.DbMap) UserRepository {
	return &postgresUserRepository{connection}
}

func (p *postgresUserRepository) Fetch() ([]*domain.User, error) {
	users := make([]*domain.User, 0)
	_, err := p.connection.Select(&users, "select * from person order by name")
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (p *postgresUserRepository) GetById(id string) (*domain.User, error) {
	user := &domain.User{}
	err := p.connection.SelectOne(user, "select * from person where id = $1", id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (p *postgresUserRepository) Update(user *domain.User) (*domain.User, error) {
	return nil, nil
}

func (p *postgresUserRepository) Store(user *domain.User) (string, error) {
	return "empty", nil
}

func (p *postgresUserRepository) Delete(id string) (bool, error) {
	return true, nil
}
