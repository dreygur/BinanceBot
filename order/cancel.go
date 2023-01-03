package order

import (
	"context"
	"fmt"
	"strings"

	"github.com/adshao/go-binance/v2/futures"
)

func MarketExitPosition(client *futures.Client, currencyPair, tradeSide, lotSize string) (*futures.CreateOrderResponse, error) {
	var side futures.SideType
	fmt.Println(lotSize)

	if strings.ToLower(tradeSide) == "buy" {
		side = futures.SideTypeBuy
	} else {
		side = futures.SideTypeSell
	}

	order, err := client.NewCreateOrderService().
		Symbol(currencyPair).
		Side(side).
		Type(futures.OrderTypeMarket).
		Quantity(lotSize).
		// ReduceOnly(true).
		Do(context.Background())
	if err != nil {
		return nil, err
	}

	return order, nil
}

// func ExitPosition(client *futures.Client) {
// 	order, err := MarketExitPosition(client)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(order)
// }
