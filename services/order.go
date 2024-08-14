package services

import (
	"impl-ddd/aggregate"
	"impl-ddd/domain/customer"
	"impl-ddd/domain/customer/memory"
	"impl-ddd/domain/product"
	"log"

	prodmem "impl-ddd/domain/product/memory"

	"github.com/google/uuid"
)

type OrderConfiguration func(os *OrderService) error

type OrderService struct {
	customers customer.CustomerRepository
	products  product.ProdectRepository
}

func NewOrderService(cfgs ...OrderConfiguration) (*OrderService, error) {
	os := &OrderService{}
	// Loop throw all the cfgs and apply them

	for _, cfg := range cfgs {
		err := cfg(os)
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

func WithMemoryCustomerRepository() OrderConfiguration {
	cr := memory.New()
	return WithCustomerRepository(cr)
}

func WithMemoryProductRepository(products []aggregate.Product) OrderConfiguration {
	return func(os *OrderService) error {
		pr := prodmem.New()
		for _, p := range products {
			if err := pr.Add(p); err != nil {
				return err
			}
		}
		os.products = pr
		return nil
	}
}

func (o *OrderService) CreateOrder(customerID uuid.UUID, productsIDs []uuid.UUID) (float64, error) {
	// Fetch Customer
	c, err := o.customers.Get(customerID)
	if err != nil {
		return 0, nil
	}

	// get each product
	var products []aggregate.Product
	var total float64

	for _, id := range productsIDs {
		p, err := o.products.GetByID(id)
		if err != nil {
			return 0, err
		}

		products = append(products, p)
		total += p.GetPrice()
	}

	log.Printf("customer: %s has orderd %d products", c.GetID(), len(products))
	return total, nil
}
