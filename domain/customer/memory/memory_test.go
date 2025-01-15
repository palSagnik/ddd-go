package memory

import (
	"domain/aggregate"
	"domain/domain/customer"
	"testing"

	"github.com/google/uuid"
)

func TestMemory_GetCustomer(t *testing.T) {
	type testCase struct {
		test        string
		id          uuid.UUID
		expectedErr error
	}

	// fake customer for testing
	testCust, err := aggregate.NewCustomer("Sagnik")
	if err != nil {
		t.Fatal(err)
	}
	id := testCust.GetID()

	// test repo
	repo := MemoryRepository{
		customers: map[uuid.UUID]aggregate.Customer{
			id: testCust,
		},
	}

	testCases := []testCase{
		{
			test:        "no customer ID test",
			id:          uuid.MustParse("f47ac10b-58cc-0372-8567-0e02b2c3d479"),
			expectedErr: customer.ErrCustomerNotFound,
		},
		{
			test:        "valid ID test",
			id:          id,
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {

			_, err := repo.Get(tc.id)
			if err != tc.expectedErr {
				t.Errorf("expected error %v, got %v", tc.expectedErr, err)
			}
		})
	}
}

func TestMemory_AddCustomer(t *testing.T) {

	type testCase struct {
		test         string
		customerName string
		expectedErr  error
	}

	testCases := []testCase{
		{
			test: "Add customer",
			customerName: "Sagnik",
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			
			repo := MemoryRepository{
				customers: map[uuid.UUID]aggregate.Customer{},
			}
			
			cust, err := aggregate.NewCustomer(tc.customerName)
			if err != nil {
				t.Fatal(err)
			}

			if err := repo.Add(cust); err != tc.expectedErr {
				t.Errorf("expected error %v, got %v", tc.expectedErr, err)
			}

			// checking id
			found, err := repo.Get(cust.GetID())
			if err != nil {
				t.Fatal(err)
			}

			if found.GetID() != cust.GetID() {
				t.Errorf("Expected %v, got %v", cust.GetID(), found.GetID())
			}
		})
	}

}
