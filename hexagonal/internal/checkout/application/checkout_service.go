package application

import (
	checkoutinbound "hexagonal/internal/checkout/ports/inbound"
	"hexagonal/internal/order/ports/inbound"
	paymentinbound "hexagonal/internal/payment/ports/inbound"
)

// OrderUseCase represents the order capability needed by checkout.
type OrderUseCase interface {
	PlaceOrder(cmd inbound.PlaceOrderCommand) (string, error)
}

// PaymentUseCase represents the payment capability needed by checkout.
type PaymentUseCase interface {
	AuthorizePayment(cmd paymentinbound.AuthorizePaymentCommand) (string, error)
}

// CheckoutService orchestrates the order and payment use cases.
//
// This is an application-level workflow. It coordinates two independent cores
// instead of merging the order and payment business rules into one domain.
type CheckoutService struct {
	orderUseCase   OrderUseCase
	paymentUseCase PaymentUseCase
}

// NewCheckoutService wires the checkout workflow to the two use cases it needs.
func NewCheckoutService(orderUseCase OrderUseCase, paymentUseCase PaymentUseCase) *CheckoutService {
	return &CheckoutService{
		orderUseCase:   orderUseCase,
		paymentUseCase: paymentUseCase,
	}
}

// Checkout executes the higher-level workflow.
//
// The checkout service first places the order and then authorizes payment. In a
// real system, if payment fails after order creation, this workflow would usually
// trigger compensation or publish an event for async handling.
func (s *CheckoutService) Checkout(cmd checkoutinbound.CheckoutCommand) (checkoutinbound.CheckoutResult, error) {
	orderID, err := s.orderUseCase.PlaceOrder(inbound.PlaceOrderCommand{
		ID:       cmd.OrderID,
		Customer: cmd.Customer,
		Product:  cmd.Product,
		Quantity: cmd.Quantity,
	})
	if err != nil {
		return checkoutinbound.CheckoutResult{}, err
	}

	paymentID, err := s.paymentUseCase.AuthorizePayment(paymentinbound.AuthorizePaymentCommand{
		ID:      cmd.PaymentID,
		OrderID: orderID,
		Amount:  cmd.PaymentAmount,
	})
	if err != nil {
		return checkoutinbound.CheckoutResult{}, err
	}

	return checkoutinbound.CheckoutResult{
		OrderID:   orderID,
		PaymentID: paymentID,
	}, nil
}

// Ensure the service satisfies the checkout inbound port.
var _ checkoutinbound.CheckoutUseCase = (*CheckoutService)(nil)
