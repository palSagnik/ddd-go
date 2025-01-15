package aggregate_test

import (
	"domain/aggregate"
	"testing"
)

func TestProduct_NewProduct(t *testing.T) {
	type testcase struct {
		test        string
		name        string
		desc        string
		price       float64
		expectedErr error
	}

	testcases := []testcase{
		{
			test:        "missing value test",
			name:        "",
			desc:        "",
			expectedErr: aggregate.ErrMissingValues,
		},
		{
			test:        "valid values",
			name:        "prod1",
			desc:        "prod1 desc",
			price:       2.0,
			expectedErr: nil,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.test, func(t *testing.T) {
			_, err := aggregate.NewProduct(tc.name, tc.desc, tc.price)
			if err != tc.expectedErr {
				t.Errorf("expected error %v, got %v", tc.expectedErr, err)
			}
		})
	}
}
