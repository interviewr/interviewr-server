package usecases

import (
	"interviewr-server/domain"
	"interviewr-server/repository"
)

type VacancyUsecase interface {
	Fetch() ([]*domain.Vacancy, error)
	GetById(id string) (*domain.Vacancy, error)
	GetFollowers(id string) ([]*domain.User, error)
}

type vacancyUsecase struct {
	repository repository.VacancyRepository
}

func NewVacancyUsecase(repo repository.VacancyRepository) *vacancyUsecase {
	return &vacancyUsecase{repository: repo}
}

func (interactor *vacancyUsecase) Fetch() ([]*domain.Vacancy, error) {
	vacancies, err := interactor.repository.Fetch()
	if err != nil {
		return nil, err
	}

	return vacancies, nil
}

func (interactor *vacancyUsecase) GetById(id string) (*domain.Vacancy, error) {
	vacancy, err := interactor.repository.GetById(id)
	if err != nil {
		return nil, err
	}

	return vacancy, nil
}

func (interactor *vacancyUsecase) GetFollowers(id string) ([]*domain.User, error) {
	followers, err := interactor.repository.GetFollowers(id)
	if err != nil {
		return nil, err
	}

	return followers, nil
}
