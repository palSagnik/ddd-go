package memory

import (
	"domain/aggregate"
	"domain/domain/customer"
	"fmt"
	"sync"

	"github.com/google/uuid"
)

type MemoryRepository struct {
	customers map[uuid.UUID]aggregate.Customer
	sync.Mutex
}

func New() *MemoryRepository {
	return &MemoryRepository{
		customers: make(map[uuid.UUID]aggregate.Customer),
	}
}

// get customer by id
func (mem *MemoryRepository) Get(id uuid.UUID) (aggregate.Customer, error) {
	
	if customer, ok := mem.customers[id]; ok {
		return customer, nil
	}
	return aggregate.Customer{}, customer.ErrCustomerNotFound
}

// add customer to repository
func (mem *MemoryRepository) Add(c aggregate.Customer) error {
	if mem.customers == nil {
		mem.Lock()
		mem.customers = make(map[uuid.UUID]aggregate.Customer)
		mem.Unlock()
	}

	// check if customer already present
	if _, ok := mem.customers[c.GetID()]; ok {
		return fmt.Errorf("customer already exists: %v", customer.ErrFailedToAddCustomer)
	}

	mem.Lock()
	mem.customers[c.GetID()] = c
	mem.Unlock()
	return nil
}

// update an existing customer
func (mem *MemoryRepository) Update(c aggregate.Customer) error {
	
	// making sure if customer is present
	if _, ok := mem.customers[c.GetID()]; !ok {
		return fmt.Errorf("customer does not exist: %w", customer.ErrFailedToUpdateCustomer)
	}
	mem.Lock()
	mem.customers[c.GetID()] = c
	mem.Unlock()
	return nil
}
