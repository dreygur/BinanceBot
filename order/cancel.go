package order

import (
	"context"
	"fmt"
	"math"
	"strconv"

	"github.com/adshao/go-binance/v2/futures"
)

func MarketExitPosition(client *futures.Client, currencyPair string) (*futures.CreateOrderResponse, error) {
	var (
		reverseTradeSide string
		positionAmt      float64
	)
	position, err := GetOpenPosition(client, currencyPair)
	if err != nil {
		return nil, err
	}

	if position.PositionSide == "BOTH" {
		reverseTradeSide = "SELL"
		positionAmt, err = strconv.ParseFloat(position.PositionAmt, 32)
		if err != nil {
			return nil, err
		}
		if positionAmt < 0 {
			reverseTradeSide = "BUY"
		}
	}

	order, err := client.NewCreateOrderService().
		Symbol(currencyPair).
		Side(futures.SideType(reverseTradeSide)).
		Type(futures.OrderTypeMarket).
		Quantity(fmt.Sprintf("%.2f", math.Abs(positionAmt))).
		ReduceOnly(true).
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
