package order

import (
	"context"
	"fmt"
	"strconv"

	"github.com/adshao/go-binance/v2/futures"
)

// Calculate Market Order Lot Size
func GetMarketOrderLotSize(client *futures.Client, currencyPair, usdtSize string) (string, error) {
	prices, err := client.NewListPricesService().
		Symbol(currencyPair).
		Do(context.Background())
	if err != nil {
		return "", err
	}
	return GetLimitOrderLotSize(usdtSize, prices[0].Price), nil
}

// Calculate Limit Order Lot Size
func GetLimitOrderLotSize(usdt, limit string) string {
	usd, err := strconv.ParseFloat(usdt, 32)
	if err != nil {
		fmt.Println(err)
	}
	lim, err := strconv.ParseFloat(limit, 32)
	if err != nil {
		fmt.Println(err)
	}
	return strconv.FormatFloat(usd/lim, 'G', -1, 32)
}

// Fetch Open Position Data
// func GetOpenPosition(client *futures.Client, currencyPair string) {
// 	positionData := client.N
// }
