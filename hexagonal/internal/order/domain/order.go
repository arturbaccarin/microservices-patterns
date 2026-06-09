package domain

import "errors"

// Order is the business entity at the center of the hexagonal architecture.
//
// The domain model should not know anything about HTTP, databases, or any other
// delivery mechanism. It only contains the rules that belong to the business.
type Order struct {
	ID       string
	Customer string
	Product  string
	Quantity int
	Status   string
}

// NewOrder creates a new order using only business rules.
//
// This function belongs to the core because it validates the minimum data the
// business needs before the order can exist.
func NewOrder(id, customer, product string, quantity int) (*Order, error) {
	if id == "" {
		return nil, errors.New("order id is required")
	}
	if customer == "" {
		return nil, errors.New("customer is required")
	}
	if product == "" {
		return nil, errors.New("product is required")
	}
	if quantity <= 0 {
		return nil, errors.New("quantity must be greater than zero")
	}

	return &Order{
		ID:       id,
		Customer: customer,
		Product:  product,
		Quantity: quantity,
		Status:   "CREATED",
	}, nil
}
