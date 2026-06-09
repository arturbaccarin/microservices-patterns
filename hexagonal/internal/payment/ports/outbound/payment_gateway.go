package outbound

import "hexagonal/internal/payment/domain"

// PaymentGateway is the outbound port for payment operations.
//
// The application layer depends on this interface, while the adapter decides
// how authorization is actually stored or sent to an external system.
type PaymentGateway interface {
	Save(payment *domain.Payment) error
	FindByID(id string) (*domain.Payment, bool)
}
