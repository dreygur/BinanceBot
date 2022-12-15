package order

import (
	"context"
	"fmt"

	"github.com/adshao/go-binance/v2/futures"
)

func MarketEnterPosition(client *futures.Client) {
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

func LimitEnterPosition(client *futures.Client) {
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
