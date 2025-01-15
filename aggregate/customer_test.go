package aggregate_test

import (
	"testing"

	"domain/aggregate"
)

func TestCustomer_NewCustomer(t *testing.T) {

	type testCase struct {
		test string
		name string
		expectedErr  error
	}

	testCases := []testCase{
		{
			test: "empty name validation test",
			name: "",
			expectedErr: aggregate.ErrInvalidPerson,
		},
		{
			test: "valid name test",
			name: "Sagnik Pal",
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			_, err := aggregate.NewCustomer(tc.name)
			if err != tc.expectedErr {
				t.Errorf("expected error %v, got %v", tc.expectedErr, err)
			}
		})
	}
}
