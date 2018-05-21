package usecases

import (
	"interviewr-server/domain"
	"interviewr-server/repository"
)

type UserUsecase interface {
	Fetch() ([]*domain.User, error)
	GetById(id string) (*domain.User, error)
	Update(person *domain.User) (*domain.User, error)
	Store(person *domain.User) (string, error)
	Delete(id string) (bool, error)
}

type personUsecase struct {
	repository repository.UserRepository
}

func NewUserUsecase(repo repository.UserRepository) *personUsecase {
	return &personUsecase{repository: repo}
}

func (interactor *personUsecase) Fetch() ([]*domain.User, error) {
	users, err := interactor.repository.Fetch()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (interactor *personUsecase) GetById(id string) (*domain.User, error) {
	person, err := interactor.repository.GetById(id)
	if err != nil {
		return nil, err
	}

	return person, nil
}

func (interactor *personUsecase) Update(person *domain.User) (*domain.User, error) {
	return nil, nil
}

func (interactor *personUsecase) Store(person *domain.User) (string, error) {
	return "empty", nil
}

func (interactor *personUsecase) Delete(id string) (bool, error) {
	return interactor.repository.Delete(id)
}
