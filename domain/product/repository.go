package product

import (
	"errors"
	"impl-ddd/aggregate"

	"github.com/google/uuid"
)

var (
	ErrProductNotFound      = errors.New("product is not found")
	ErrProductAlreadyExists = errors.New("this is already such product")
)

type ProdectRepository interface {
	GetAll() ([]aggregate.Product, error)
	GetByID(id uuid.UUID) (aggregate.Product, error)
	Add(product aggregate.Product) error
	Update(product aggregate.Product) error
	Delete(id uuid.UUID) error
}
