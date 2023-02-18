package domain

type Person struct {
	Id    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email" validate:"email,omitempty"`
	Phone string `json:"phone"`
}

type PersonService interface {
	GetAllPerson() ([]Person, error)
	CreatePerson(*Person) error
	DeletePersonById(int) error
}

type PersonRepository interface {
	GetAll() ([]Person, error)
	SavePerson(*Person) error
	RemovePersonById(int) error
}
