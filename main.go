package main

import (
	"context"
	"fmt"

	"github.com/adshao/go-binance/v2"
)

var (
	// Spot
	apiKey    = "PUkxJSbyWYzRYISEXfGodacu7uo59PIvyLvs9Z6Vkcy8gtNbicbsdJ5o36zhA2r8"
	secretKey = "VJxEllH0S02gDOQAYTYBMFKkX9dKfEoCqZ4Zrbaix835m9FdXv1JN4LXnTX9sQHe"

	// // Features
	// apiKey    = "b762687ec7d72cdf03b97a42e4d9ef24e99b082efb63b1d2d431a7aa394c573b"
	// secretKey = "ffe69c27b2315aeed69cf44594e3d360f73bb66b4dfc633f9076e19995ef1efc"
)

func main() {
	binance.UseTestnet = true
	client := binance.NewClient(apiKey, secretKey)
	// Enable RateLimit
	client.NewRateLimitService().Do(context.Background())

	price, err := client.NewListPricesService().Symbol("BTCUSDT").Do(context.Background())
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(price[0].Price)

	// currencyPair := "BTCUSDT"
	// tradeSide := "BUY"
	// entryPrice := "10000.0"
	// lotSize := order.GetLimitOrderLotSize("500.0", entryPrice)
	/*


		res, err := order.LimitEnterPosition(client, currencyPair, tradeSide, lotSize, entryPrice)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(res)
		res, err = order.MarketEnterPosition(client, currencyPair, tradeSide, lotSize)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(res)
		res, err = order.GetMarketOrderLotSize(client, currencyPair, "500.0")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(res)
	*/
}
