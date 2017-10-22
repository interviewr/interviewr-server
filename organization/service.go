package organization

import (
	"context"
	"errors"
	"sync"
	// "database/sql"
	// "github.com/lib/pq"
)

type Organization struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Description string `json:"description"`
	Location string `json:"location"`
}

type Service interface {
	GetOrganization(ctx context.Context, id string) (Organization, error)
	PostOrganization(ctx context.Context, org Organization) error
	PutOrganization(ctx context.Context, id string, org Organization) error
	PatchOrganization(ctx context.Context, id string, org Organization) error
	DeleteOrganization(ctx, context.Context, id string) error
}

var (
	ErrInconsistentIDs = errors.New("inconsistent IDs")
	ErrAlreadyExists = errors.New("already exists")
	ErrNotFound = errors.New("not found")
)

// type Service struct {
// 	DB *sql.DB
// }

// func (s *Service) GetOrganization(ctx context.Context, id string) (*Organization, error) {
// 	var org Organization
// 	row := db.QueryRow(`SELECT id, name FROM organizations WHERE id = $1`, id)
// 	if row.Scan(&org.ID, &org.Name); err != nil {
// 		return nil, err
// 	}
// 	return &org, nil
// }

type inmemService struct {
	mtx sync.RWMutex
	m map[string]Organization
}

func NewInmemService() Service {
	return &inmemService{
		m: map[string]Organization{},
	}
}

func (s *inmemService) PostOrganization(ctx context.Context, org Organization) error {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	if _, ok := s.m[org.ID]; ok {
		return ErrAlreadyExists
	}

	s.m[org.ID] = org
	return nil
}

func (s *inmemService) GetOrganization(ctx context.Context, id string) (Organization, error) {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	org, ok := s.m[id]
	if !ok {
		return Organization{}, ErrNotFound
	}

	return org, nil
}