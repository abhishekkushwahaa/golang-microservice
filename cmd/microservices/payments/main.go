package main

import (
	"fmt"
	"log"
	"os"

	"github.com/abhishekkushwahaa/golang-microservice/pkg/common/cmd"
	payments_app "github.com/abhishekkushwahaa/golang-microservice/pkg/payments/application"
	payments_infra_orders "github.com/abhishekkushwahaa/golang-microservice/pkg/payments/infrastructure/orders"
	"github.com/abhishekkushwahaa/golang-microservice/pkg/payments/interfaces/amqp"
)

func main() {
	log.Println("Starting the payments microservices!")

	defer log.Println("Closing payments microservice!")

	ctx := cmd.Context()

	paymentsInterface := createPaymentsMicrodervice()

	if err := paymentsInterface.Run(ctx); err != nil {
		panic(err)
	}
}

func createPaymentsMicrodervice() amqp.PaymentsInterface {
	cmd.WaitForService(os.Getenv("SHOP_RABBITMQ_ADDR"))

	paymentsService := payments_app.NewPaymentsService(
		payments_infra_orders.NewHTTPClient(os.Getenv("SHOP_PRODUCT_SERVICE_ADDR")),
	)

	paymentsInterface, err := amqp.NewPaymentsInterface(
		fmt.Sprintf("amqp://%s", os.Getenv("SHOP_RABBITMQ_ADDR")),
		os.Getenv("SHOP_RABBITMQ_ORDERS_TO_PAY_QUEUE"),
		paymentsService,
	)
	if err != nil {
		panic(err)
	}

	return paymentsInterface
}
