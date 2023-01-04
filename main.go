package main

import (
	"binancebot/utils"
	"bufio"
	"fmt"
	"os"

	"github.com/adshao/go-binance/v2"
	"github.com/adshao/go-binance/v2/futures"
)

var (
	// Spot
	// apiKey    = "PUkxJSbyWYzRYISEXfGodacu7uo59PIvyLvs9Z6Vkcy8gtNbicbsdJ5o36zhA2r8"
	// secretKey = "VJxEllH0S02gDOQAYTYBMFKkX9dKfEoCqZ4Zrbaix835m9FdXv1JN4LXnTX9sQHe"

	// // Features
	// apiKey    = "067ae6966fecb63c04926eee4233dfa29fa9defe1a40e621dd85a91941c6bf91"
	// secretKey = "582b349eba9d2d6fba35749c620c9f89d24292bd2b2871eaad1fa8f39274575d"

	// apiKey    = "8fb794b2bfc74c3380c1080db4693d47eaa2d3a7af2263727667b8f35894e943"
	// secretKey = "9d798723f64a381d3aead74733f34e2e7e9236db612295deb0e752b274d842ec"

	// Future
	apiKey    = "mnPeHnXenLvZ3SykyY6GUoQ0nzRr18Mgon2v5kBgL5O9gmtPGlwA3NQGdH4UsU2A"
	secretKey = "JhAcbVDcrvLp6bvWLhKoLXhS7ThEyDfaiMTENES644pw9bJM9IkbT9Ij6QR7RG7C"

	// Future Client
	client *futures.Client
)

func main() {
	// Stop printing error stack
	// defer func() {
	// 	if r := recover(); r != nil {
	// 		fmt.Println(r)
	// 		os.Exit(0)
	// 	}
	// }()

	binance.UseTestnet = true
	client = binance.NewFuturesClient(apiKey, secretKey)
	// Enable RateLimit
	// client.NewRateLimitService().Do(context.Background())

	fmt.Printf("__________WELCOME__________\n\n")
	for {
		// start := time.Now()
		fmt.Print("> ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan() // use `for scanner.Scan()` to keep reading
		rawString := scanner.Text()
		utils.ProcessCommand(client, rawString)
		// fmt.Printf("Time taken %v\n", time.Since(start))
	}

	// price, err := client.NewListPricesService().Symbol("BTCUSDT").Do(context.Background())
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(price[0].Price)

	// currencyPair := "BTCUSDT"
	// tradeSide := "BUY"
	// entryPrice := "1.0"
	// lotSize := order.GetLimitOrderLotSize("0.5", entryPrice)

	// res, err := order.LimitEnterPosition(client, currencyPair, tradeSide, lotSize, entryPrice)
	// if err != nil {
	// 	fmt.Println("Limit Enter Position", err)
	// }
	// fmt.Println(res)
	// res, err = order.MarketEnterPosition(client, currencyPair, tradeSide, lotSize)
	// if err != nil {
	// 	fmt.Println("Market Enter Position", err)
	// }
	// fmt.Println(res)
	// r, err := order.GetMarketOrderLotSize(client, currencyPair, "0.5")
	// if err != nil {
	// 	fmt.Println("Get Market order Lot Size", err)
	// }
	// fmt.Println(r)

}
