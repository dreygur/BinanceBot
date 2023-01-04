package order

import (
	"context"
	"fmt"
	"math"
	"strconv"

	"github.com/adshao/go-binance/v2/futures"
)

func (o *Order) MarketExitPosition(currencyPair string) (*futures.CreateOrderResponse, error) {
	var (
		reverseTradeSide string
		positionAmt      float64
	)
	position, err := o.GetOpenPosition(currencyPair)
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

	order, err := o.Client.NewCreateOrderService().
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

func (o *Order) CancelOrders(currencyPair string) error {
	return o.Client.NewCancelAllOpenOrdersService().Symbol(currencyPair).Do(context.Background())
}
