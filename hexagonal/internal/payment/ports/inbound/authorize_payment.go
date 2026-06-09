package inbound

// AuthorizePaymentCommand is the input data for the payment use case.
type AuthorizePaymentCommand struct {
	ID      string
	OrderID string
	Amount  float64
}

// AuthorizePaymentUseCase is the inbound port for the payment domain.
type AuthorizePaymentUseCase interface {
	AuthorizePayment(cmd AuthorizePaymentCommand) (string, error)
}
