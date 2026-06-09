# Hexagonal Architecture in Go

This folder contains a small Go example that demonstrates hexagonal architecture.

The example uses three small business workflows:

- The business logic lives in the center.
- An inbound port defines how the outside world asks the business to place an order.
- An outbound port defines how the business stores an order.
- A second payment domain shows that each business capability can have its own ports and adapters.
- A checkout workflow coordinates order and payment without merging the domains.
- A CLI-style `main` function acts as an inbound adapter.
- An in-memory repository acts as an outbound adapter.

## Folder Map

- `cmd/hexagonal/main.go`: inbound adapter that starts the example.
- `internal/checkout/application`: checkout workflow that orchestrates order and payment.
- `internal/checkout/ports/inbound`: checkout use case interface and command.
- `internal/order/domain`: core business entity and rules.
- `internal/order/ports/inbound`: interface used to enter the core.
- `internal/order/ports/outbound`: interface used by the core to reach outside systems.
- `internal/order/application`: application service that coordinates the use case.
- `internal/order/adapters/memory`: outbound adapter implementation.
- `internal/payment/domain`: core payment entity and rules.
- `internal/payment/ports/inbound`: payment use case interface.
- `internal/payment/ports/outbound`: payment gateway interface.
- `internal/payment/application`: payment application service.
- `internal/payment/adapters/memory`: payment outbound adapter.

## What To Notice

- The domain does not import any adapter package.
- The application service depends on interfaces, not concrete infrastructure.
- Adapters are replaceable without changing the core business rules.
- Each domain can evolve independently while still following the same hexagonal pattern.
- Cross-domain coordination belongs in an application workflow like checkout, not inside the core domains.
