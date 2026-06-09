package application

import (
	"fmt"

	"hexagonal/internal/payment/domain"
	"hexagonal/internal/payment/ports/inbound"
	"hexagonal/internal/payment/ports/outbound"
)

// AuthorizePaymentService is the application service for the payment domain.
type AuthorizePaymentService struct {
	gateway outbound.PaymentGateway
}

// NewAuthorizePaymentService wires the payment use case to an outbound port.
func NewAuthorizePaymentService(gateway outbound.PaymentGateway) *AuthorizePaymentService {
	return &AuthorizePaymentService{gateway: gateway}
}

// AuthorizePayment executes the payment workflow.
func (s *AuthorizePaymentService) AuthorizePayment(cmd inbound.AuthorizePaymentCommand) (string, error) {
	if existing, found := s.gateway.FindByID(cmd.ID); found && existing != nil {
		return "", fmt.Errorf("payment %s already exists", cmd.ID)
	}

	payment, err := domain.NewPayment(cmd.ID, cmd.OrderID, cmd.Amount)
	if err != nil {
		return "", err
	}

	if err := s.gateway.Save(payment); err != nil {
		return "", err
	}

	return payment.ID, nil
}

// Ensure the service satisfies the inbound port.
var _ inbound.AuthorizePaymentUseCase = (*AuthorizePaymentService)(nil)
