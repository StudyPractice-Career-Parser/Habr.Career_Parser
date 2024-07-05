package service

import (
	parser "habr-career"
	"habr-career/pkg/repository"
)

type Vacancies interface {
	InsertAll(vacancies []parser.Vacancy) error
}

type Resume interface {
}

type Service struct {
	Vacancies
	Resume
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Vacancies: NewVacService(repo.Vacancies),
	}
}
