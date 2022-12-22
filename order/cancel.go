package order

import (
	"context"
	"fmt"
	"strings"

	"github.com/adshao/go-binance/v2"
)

func MarketExitPosition(client *binance.Client, currencyPair, tradeSide, lotSize string) (*binance.CreateOrderResponse, error) {
	var side binance.SideType
	fmt.Println(lotSize)

	if strings.ToLower(tradeSide) == "buy" {
		side = binance.SideTypeBuy
	} else {
		side = binance.SideTypeSell
	}

	order, err := client.NewCreateOrderService().
		Symbol(currencyPair).
		Side(side).
		Type(binance.OrderTypeMarket).
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
