package order

import (
	"context"
	"fmt"

	"github.com/adshao/go-binance/v2/futures"
)

// Calculate Market Order Lot Size
// :Incomplete
func GetMarketOrderLotSize(client *futures.Client) {
	prices, err := client.NewListPricesService().Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, p := range prices {
		fmt.Println(p)
	}
}

// Calculate Limit Order Lot Size
// :Untested
func GetLimitOrderLotSize(usdt, limit int64) float64 {
	return float64(usdt) / float64(limit)
}
