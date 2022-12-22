package order

import (
	"context"
	"fmt"
	"strings"

	"github.com/adshao/go-binance/v2"
)

func MarketEnterPosition(client *binance.Client, currencyPair, tradeSide, lotSize string) (*binance.CreateOrderResponse, error) {
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
		Do(context.Background())
	if err != nil {
		return nil, err
	}

	return order, nil
}

func LimitEnterPosition(client *binance.Client, currencyPair, tradeSide, lotSize, entryPrice string) (*binance.CreateOrderResponse, error) {
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
		Type(binance.OrderTypeLimit).
		Quantity(lotSize).
		TimeInForce(binance.TimeInForceTypeGTC).
		Price(entryPrice).Do(context.Background())
	if err != nil {
		return nil, err
	}

	return order, nil
}
