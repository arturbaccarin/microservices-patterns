package outbound

import "hexagonal/internal/order/domain"

// OrderRepository is the outbound port.
//
// The business logic depends on this interface to persist or retrieve orders.
// The implementation lives outside the core, in an adapter.
type OrderRepository interface {
	Save(order *domain.Order) error
	FindByID(id string) (*domain.Order, bool)
}
