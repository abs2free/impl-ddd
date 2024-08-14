package customer

import (
	"errors"
	"impl-ddd/aggregate"

	"github.com/google/uuid"
)

var (
	ErrCustomerNotFound  = errors.New("the customer was not found in the repository")
	ErrFailedAddCustomer = errors.New("failed to add the customer")
	ErrUpdateCustomer    = errors.New("failed to update the customer")
)

type CustomerRepository interface {
	Get(uuid.UUID) (aggregate.Customer, error)
	Add(aggregate.Customer) error
	Update(aggregate.Customer) error
}
