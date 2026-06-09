package main

import (
	"fmt"

	checkoutapplication "hexagonal/internal/checkout/application"
	checkoutinbound "hexagonal/internal/checkout/ports/inbound"
	"hexagonal/internal/order/adapters/memory"
	orderapplication "hexagonal/internal/order/application"
	paymentmemory "hexagonal/internal/payment/adapters/memory"
	paymentapplication "hexagonal/internal/payment/application"
)

// main acts as an inbound adapter in this example.
//
// In a real application, this could be an HTTP handler, a message consumer, or a
// CLI entry point. Its job is to translate outside input into a command for the
// inbound port.
func main() {
	repository := memory.NewRepository()
	paymentGateway := paymentmemory.NewGateway()
	orderService := orderapplication.NewPlaceOrderService(repository)
	paymentService := paymentapplication.NewAuthorizePaymentService(paymentGateway)
	checkoutService := checkoutapplication.NewCheckoutService(orderService, paymentService)

	// The adapter creates the command that will be sent into the checkout core.
	command := checkoutinbound.CheckoutCommand{
		OrderID:       "order-1",
		Customer:      "Ada Lovelace",
		Product:       "Mechanical Keyboard",
		Quantity:      2,
		PaymentID:     "payment-1",
		PaymentAmount: 249.98,
	}

	result, err := checkoutService.Checkout(command)
	if err != nil {
		fmt.Println("failed to checkout:", err)
		return
	}

	fmt.Println("checkout completed successfully")
	fmt.Println("order id:", result.OrderID)
	fmt.Println("payment id:", result.PaymentID)
}
