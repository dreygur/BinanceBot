package order

import (
	"context"
	"fmt"

	"github.com/adshao/go-binance/v2/futures"
)

func MarketExitPosition(client *futures.Client) (*futures.CreateOrderResponse, error) {
	order, err := client.NewCreateOrderService().Symbol("BNBETH").
		Side(futures.SideTypeBuy).Type(futures.OrderTypeMarket).
		TimeInForce(futures.TimeInForceTypeGTC).Quantity("5").
		ReduceOnly(true).
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println(order)

	return order, nil
	// Use Test() instead of Do() for testing.
}

func ExitPosition(client *futures.Client) {
	order, err := MarketExitPosition(client)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(order)
}
