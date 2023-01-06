package service

import (
	"context"
	"fmt"
	"strconv"

	"github.com/adshao/go-binance/v2/futures"
)

// Calculate Market Order Lot Size
func (o *Service) GetMarketOrderLotSize(currencyPair, usdtSize string) (string, error) {
	prices, err := client.NewListPricesService().
		Symbol(currencyPair).
		Do(context.Background())
	if err != nil {
		return "", err
	}

	return o.GetLimitOrderLotSize(usdtSize, prices[0].Price), nil
}

// Calculate Limit Order Lot Size
func (o *Service) GetLimitOrderLotSize(usdt, limit string) string {
	usd, err := strconv.ParseFloat(usdt, 32)
	if err != nil {
		fmt.Println(err)
	}
	lim, err := strconv.ParseFloat(limit, 32)
	if err != nil {
		fmt.Println(err)
	}
	return fmt.Sprintf("%.2f", usd/lim)
}

// Fetch Open Position Data
func (o *Service) GetOpenPosition(currencyPair string) (*futures.PositionRisk, error) {
	res, err := client.NewGetPositionRiskService().Symbol(currencyPair).Do(context.Background())
	if err != nil {
		return nil, err
	}

	return res[0], nil
}
