package customer

import (
	"domain/aggregate"
	"errors"

	"github.com/google/uuid"
)

var (
	ErrCustomerNotFound       = errors.New("the customer was not found in the repository")
	ErrFailedToAddCustomer    = errors.New("could not add customer to the repository")
	ErrFailedToUpdateCustomer = errors.New("could not update customer details")
)

type CustomerRepository interface {
	Get(uuid.UUID) (aggregate.Customer, error)
	Update(uuid.UUID) error
	Add(aggregate.Customer) error
}
