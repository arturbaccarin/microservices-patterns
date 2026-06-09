package inbound

// PlaceOrderCommand is the input data that an external actor sends into the use case.
//
// In hexagonal architecture, inbound adapters translate outside requests into a
// structure like this so the core stays independent from transport details.
type PlaceOrderCommand struct {
	ID       string
	Customer string
	Product  string
	Quantity int
}

// PlaceOrderUseCase is the inbound port.
//
// A port defines what the business logic offers to the outside world. External
// adapters call this interface, not the other way around.
type PlaceOrderUseCase interface {
	PlaceOrder(cmd PlaceOrderCommand) (string, error)
}
