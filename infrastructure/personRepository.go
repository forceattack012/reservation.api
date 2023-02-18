package infrastructure

import (
	"github.com/forceattack012/reservationroom/domain"
)

type PersonRepository struct {
	db *GormDB
}

func NewPersonRepository(gormDb *GormDB) *PersonRepository {
	return &PersonRepository{db: gormDb}
}

func (repo *PersonRepository) GetAll() ([]domain.Person, error) {
	var personList []domain.Person
	if result := repo.db.Find(&personList); result.Error != nil {
		return nil, result.Error
	}

	return personList, nil
}

func (repo *PersonRepository) SavePerson(person *domain.Person) error {
	return repo.db.Save(person).Error
}

func (repo *PersonRepository) RemovePersonById(id int) error {
	return repo.db.Delete(domain.Person{}, id).Error
}
