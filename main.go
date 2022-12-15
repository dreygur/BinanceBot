package main

import (
	"context"
	"fmt"

	"github.com/adshao/go-binance/v2"
	"github.com/adshao/go-binance/v2/futures"
)

var (
	apiKey    = "fpvCUJQoGImqQShS36kowwRt15SDeaXOosjrUePQr3Ux9R7njuV1MbvIzaKYtaQY"
	secretKey = "ei15acV7HPr6lln6e36muSXANsKr19U7ULJcQqMXG9RONPDroBV9deZtWuIENc8s"
)

func main() {
	binance.UseTestnet = true
	client := binance.NewFuturesClient(apiKey, secretKey)

	orders, err := client.NewListOrdersService().Symbol("BNBETH").
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, o := range orders {
		fmt.Println(o)
	}

	// Place order
	marketEnterPosition(client)
	marketExitPosition(client) // Exit

	// Place Limit Order
	limitEnterPosition(client)
}

func marketEnterPosition(client *futures.Client) {
	order, err := client.NewCreateOrderService().Symbol("BNBETH").
		Side(futures.SideTypeBuy).Type(futures.OrderTypeMarket).
		TimeInForce(futures.TimeInForceTypeGTC).Quantity("5").
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(order)

	// Use Test() instead of Do() for testing.
}

func marketExitPosition(client *futures.Client) {
	order, err := client.NewCreateOrderService().Symbol("BNBETH").
		Side(futures.SideTypeBuy).Type(futures.OrderTypeMarket).
		TimeInForce(futures.TimeInForceTypeGTC).Quantity("5").
		ReduceOnly(true).
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(order)

	// Use Test() instead of Do() for testing.
}

func limitEnterPosition(client *futures.Client) {
	order, err := client.NewCreateOrderService().Symbol("BNBETH").
		Side(futures.SideTypeBuy).Type(futures.OrderTypeMarket).
		TimeInForce(futures.TimeInForceTypeGTC).Quantity("5").
		Price("0.0030000").Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(order)

	// Use Test() instead of Do() for testing.
}
