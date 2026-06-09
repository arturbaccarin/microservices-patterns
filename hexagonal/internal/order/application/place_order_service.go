package application

import (
	"fmt"

	"hexagonal/internal/order/domain"
	"hexagonal/internal/order/ports/inbound"
	"hexagonal/internal/order/ports/outbound"
)

// PlaceOrderService is the application service that implements the inbound port.
//
// This is the business workflow. It coordinates domain rules and talks to the
// outbound port, but it does not know which adapter implements that port.
type PlaceOrderService struct {
	repository outbound.OrderRepository
}

// NewPlaceOrderService wires the core use case to an outbound port.
//
// Dependency injection happens here: the business service receives an interface
// instead of creating a concrete repository itself.
func NewPlaceOrderService(repository outbound.OrderRepository) *PlaceOrderService {
	return &PlaceOrderService{repository: repository}
}

// PlaceOrder executes the use case.
//
// The service validates the command through the domain, checks whether the order
// already exists, and then saves it through the outbound port.
func (s *PlaceOrderService) PlaceOrder(cmd inbound.PlaceOrderCommand) (string, error) {
	if existing, found := s.repository.FindByID(cmd.ID); found && existing != nil {
		return "", fmt.Errorf("order %s already exists", cmd.ID)
	}

	order, err := domain.NewOrder(cmd.ID, cmd.Customer, cmd.Product, cmd.Quantity)
	if err != nil {
		return "", err
	}

	if err := s.repository.Save(order); err != nil {
		return "", err
	}

	return order.ID, nil
}

// Ensure the service satisfies the inbound port.
var _ inbound.PlaceOrderUseCase = (*PlaceOrderService)(nil)
