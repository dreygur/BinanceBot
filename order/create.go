package order

import (
	"context"
	"fmt"
	"strings"

	"github.com/adshao/go-binance/v2/futures"
)

func MarketEnterPosition(client *futures.Client, currencyPair, tradeSide, lotSize string) (*futures.CreateOrderResponse, error) {
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
		Do(context.Background())
	if err != nil {
		return nil, err
	}

	return order, nil
}

func LimitEnterPosition(client *futures.Client, currencyPair, tradeSide, lotSize, entryPrice string) (*futures.CreateOrderResponse, error) {
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
		Type(futures.OrderTypeLimit).
		Quantity(lotSize).
		TimeInForce(futures.TimeInForceTypeGTC).
		Price(entryPrice).Do(context.Background())
	if err != nil {
		return nil, err
	}

	return order, nil
}
