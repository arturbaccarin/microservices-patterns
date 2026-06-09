package inbound

// CheckoutCommand is the input for the checkout use case.
//
// It bundles the order and payment data so a higher-level application service
// can coordinate both business capabilities in one workflow.
type CheckoutCommand struct {
	OrderID       string
	Customer      string
	Product       string
	Quantity      int
	PaymentID     string
	PaymentAmount float64
}

// CheckoutResult is the output of the checkout use case.
type CheckoutResult struct {
	OrderID   string
	PaymentID string
}

// CheckoutUseCase is the inbound port for the checkout workflow.
type CheckoutUseCase interface {
	Checkout(cmd CheckoutCommand) (CheckoutResult, error)
}
