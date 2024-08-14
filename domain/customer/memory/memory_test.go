package memory

import (
	"errors"
	"impl-ddd/aggregate"
	"impl-ddd/domain/customer"
	"testing"

	"github.com/google/uuid"
)

func TestMemory_GetCustom(t *testing.T) {
	cust, err := aggregate.NewCustomer("percy")
	if err != nil {
		t.Fatal(err)
	}

	id := cust.GetID()

	repo := MemoryRepository{
		customers: map[uuid.UUID]aggregate.Customer{
			id: cust,
		},
	}

	type testCase struct {
		name        string
		id          uuid.UUID
		expectedErr error
	}

	testCases := []testCase{
		{
			name:        "no customer by id",
			id:          uuid.MustParse("558f2d94-8895-44d2-b123-0aee2a1d35a4"),
			expectedErr: customer.ErrCustomerNotFound,
		},
		{
			name:        "cumtomer by id ",
			id:          id,
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := repo.Get(tc.id)
			if !errors.Is(err, tc.expectedErr) {
				t.Errorf("expected err:%v, got:%v", tc.expectedErr, err)
			}
		})
	}
}
