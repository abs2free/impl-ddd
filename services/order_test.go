package services

import (
	"impl-ddd/aggregate"
	"testing"

	"github.com/google/uuid"
)

func init_products(t *testing.T) ([]aggregate.Product, error) {
	berr, err := aggregate.NewProduct("Beer", "health", 1.22)
	if err != nil {
		return nil, err
	}

	peenuts, err := aggregate.NewProduct("Meanuts", "snacks", 0.99)
	if err != nil {
		return nil, err
	}

	wine, err := aggregate.NewProduct("Wine", "nasty drink", 0.99)
	if err != nil {
		return nil, err
	}
	return []aggregate.Product{berr, peenuts, wine}, nil
}

func TestOrder_NewOrder(t *testing.T) {
	products, err := init_products(t)
	if err != nil {
		t.Fatal(err)
	}

	os, err := NewOrderService(
		WithMemoryCustomerRepository(),
		WithMemoryProductRepository(products),
	)
	if err != nil {
		t.Fatal(err)
	}

	cust, err := aggregate.NewCustomer("nancy")
	if err != nil {
		t.Fatal(err)
	}

	err = os.customers.Add(cust)
	if err != nil {
		t.Fatal(err)
	}

	order := []uuid.UUID{
		products[0].GetID(),
		products[2].GetID(),
	}

	price, err := os.CreateOrder(cust.GetID(), order)
	if err != nil {
		t.Fatal(err)
	}

	if price != 2.21 {
		t.Errorf("has wrong price:%f", price)
	}
}
