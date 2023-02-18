package application

import (
	"errors"

	"github.com/forceattack012/reservationroom/domain"
	"github.com/forceattack012/reservationroom/helper"
)

type PersonService struct {
	repo domain.PersonRepository
}

func NewPersonService(repo domain.PersonRepository) domain.PersonService {
	return &PersonService{repo: repo}
}

func (s *PersonService) GetAllPerson() ([]domain.Person, error) {
	personList, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	return personList, nil
}

func (s *PersonService) CreatePerson(person *domain.Person) error {
	if isPhone := helper.IsPhone(person.Phone); !isPhone {
		return errors.New("moblie invalid")
	}
	if err := s.repo.SavePerson(person); err != nil {
		return err
	}
	return nil
}

func (s *PersonService) DeletePersonById(id int) error {
	if err := s.repo.RemovePersonById(id); err != nil {
		return err
	}
	return nil
}
