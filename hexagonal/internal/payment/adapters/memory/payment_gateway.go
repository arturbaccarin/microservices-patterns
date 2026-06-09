package memory

import (
	"sync"

	"hexagonal/internal/payment/domain"
)

// Gateway is an outbound adapter for the payment domain.
type Gateway struct {
	mu       sync.RWMutex
	payments map[string]*domain.Payment
}

// NewGateway creates the in-memory adapter.
func NewGateway() *Gateway {
	return &Gateway{payments: make(map[string]*domain.Payment)}
}

// Save stores the payment in memory.
func (g *Gateway) Save(payment *domain.Payment) error {
	g.mu.Lock()
	defer g.mu.Unlock()

	g.payments[payment.ID] = payment
	return nil
}

// FindByID looks up a payment in memory.
func (g *Gateway) FindByID(id string) (*domain.Payment, bool) {
	g.mu.RLock()
	defer g.mu.RUnlock()

	payment, found := g.payments[id]
	return payment, found
}

// Ensure the adapter satisfies the outbound port.
var _ interface {
	Save(*domain.Payment) error
	FindByID(string) (*domain.Payment, bool)
} = (*Gateway)(nil)
