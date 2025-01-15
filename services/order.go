package services

import (
	"domain/domain/customer"

	"github.com/google/uuid"
)

type OrderConfiguration func(os *OrderService) error

type OrderService struct {
	customers customer.CustomerRepository
}

// factory for order service
func NewOrderService(configs ...OrderConfiguration) (*OrderService, error) {

	os := &OrderService{}
	for _, config := range configs {
		err := config(os)
		if err != nil {
			return nil, err
		}
	}

	return os, nil
}

func WithCustomerRepository(cr customer.CustomerRepository) OrderConfiguration {

	return func(os *OrderService) error {
		os.customers = cr
		return nil
	}
}

// func WithMemoryCustomerRepository() OrderConfiguration {
// 	cr := memory.New()
// 	return WithCustomerRepository(cr)
// }

func (o *OrderService) CreateOrder(customerID uuid.UUID, products []uuid.UUID) error {
	_, err := o.customers.Get(customerID)
	if err != nil {
		return customer.ErrCustomerNotFound
	}

	return nil
}
