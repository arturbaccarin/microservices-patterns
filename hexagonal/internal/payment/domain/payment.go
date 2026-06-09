package domain

import "errors"

// Payment is the business entity for the payment domain.
//
// It stays independent from the order domain so each business capability can
// evolve through its own ports and adapters.
type Payment struct {
	ID      string
	OrderID string
	Amount  float64
	Status  string
}

// NewPayment creates a payment using only business rules.
func NewPayment(id, orderID string, amount float64) (*Payment, error) {
	if id == "" {
		return nil, errors.New("payment id is required")
	}
	if orderID == "" {
		return nil, errors.New("order id is required")
	}
	if amount <= 0 {
		return nil, errors.New("amount must be greater than zero")
	}

	return &Payment{
		ID:      id,
		OrderID: orderID,
		Amount:  amount,
		Status:  "AUTHORIZED",
	}, nil
}
