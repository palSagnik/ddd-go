package aggregate

import (
	"domain/entity"
	"domain/valueobjects"
	"errors"

	"github.com/google/uuid"
)

type Customer struct {
	person       *entity.Person
	items        []*entity.Item
	transactions []valueobjects.Transaction
}

var (
	ErrInvalidPerson = errors.New("a customer has to be a valid person")
)

// Following factory design for function implementation
func NewCustomer(name string) (Customer, error) {
	
	if len(name) == 0 {
		return Customer{}, ErrInvalidPerson
	}

	person := &entity.Person{
		ID: uuid.New(),
		Name: name,
	}

	return Customer{
		person: person,
		items: make([]*entity.Item, 0),
		transactions: make([]valueobjects.Transaction, 0),
	}, nil
}

func (c *Customer) GetID() uuid.UUID {
	return c.person.ID
}

func (c *Customer) SetID(id uuid.UUID) {
	if c.person == nil {
		c.person = &entity.Person{}
	}

	c.person.ID = id
}

func (c *Customer) GetName() string {
	return c.person.Name
}

func (c *Customer) SetName(name string) {
	if c.person == nil {
		c.person = &entity.Person{}
	}

	c.person.Name = name
}

