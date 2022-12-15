package main

import (
	"binancebot/order"
	"context"
	"fmt"

	"github.com/adshao/go-binance/v2"
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
	order.MarketEnterPosition(client)
	order.MarketExitPosition(client) // Exit

	// Place Limit Order
	order.LimitEnterPosition(client)
}
