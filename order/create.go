package order

import (
	"context"
	"strings"

	"github.com/adshao/go-binance/v2/futures"
)

func (o *Order) MarketEnterPosition(currencyPair, tradeSide, lotSize string) (*futures.CreateOrderResponse, error) {
	var side futures.SideType

	if strings.ToLower(tradeSide) == "buy" {
		side = futures.SideTypeBuy
	} else {
		side = futures.SideTypeSell
	}

	order, err := o.Client.NewCreateOrderService().
		Symbol(currencyPair).
		Side(side).
		Type(futures.OrderTypeMarket).
		Quantity(lotSize).
		Do(context.Background())
	if err != nil {
		return nil, err
	}

	return order, nil
}

func (o *Order) LimitEnterPosition(currencyPair, tradeSide, lotSize, entryPrice string) (*futures.CreateOrderResponse, error) {
	var side futures.SideType

	if strings.ToLower(tradeSide) == "buy" {
		side = futures.SideTypeBuy
	} else {
		side = futures.SideTypeSell
	}

	order, err := o.Client.NewCreateOrderService().
		Symbol(currencyPair).
		Side(side).
		Type(futures.OrderTypeLimit).
		Quantity(lotSize).
		TimeInForce(futures.TimeInForceTypeGTC).
		Price(entryPrice).Do(context.Background())
	if err != nil {
		return nil, err
	}

	return order, nil
}
