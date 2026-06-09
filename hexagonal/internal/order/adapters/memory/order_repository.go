package memory

import (
	"sync"

	"hexagonal/internal/order/domain"
)

// Repository is an outbound adapter.
//
// It implements the outbound port using an in-memory map. In a real application,
// this adapter could talk to PostgreSQL, MongoDB, or another persistent store.
type Repository struct {
	mu     sync.RWMutex
	orders map[string]*domain.Order
}

// NewRepository creates the adapter state.
func NewRepository() *Repository {
	return &Repository{orders: make(map[string]*domain.Order)}
}

// Save stores the order in memory.
func (r *Repository) Save(order *domain.Order) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.orders[order.ID] = order
	return nil
}

// FindByID looks up an order in memory.
func (r *Repository) FindByID(id string) (*domain.Order, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	order, found := r.orders[id]
	return order, found
}

// Ensure the adapter satisfies the outbound port.
var _ interface {
	Save(*domain.Order) error
	FindByID(string) (*domain.Order, bool)
} = (*Repository)(nil)
