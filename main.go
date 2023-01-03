package main

import (
	"binancebot/order"
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/adshao/go-binance/v2"
)

var (
	// Spot
	// apiKey    = "PUkxJSbyWYzRYISEXfGodacu7uo59PIvyLvs9Z6Vkcy8gtNbicbsdJ5o36zhA2r8"
	// secretKey = "VJxEllH0S02gDOQAYTYBMFKkX9dKfEoCqZ4Zrbaix835m9FdXv1JN4LXnTX9sQHe"

	// // Features
	// apiKey    = "b762687ec7d72cdf03b97a42e4d9ef24e99b082efb63b1d2d431a7aa394c573b"
	// secretKey = "ffe69c27b2315aeed69cf44594e3d360f73bb66b4dfc633f9076e19995ef1efc"

	// apiKey    = "8fb794b2bfc74c3380c1080db4693d47eaa2d3a7af2263727667b8f35894e943"
	// secretKey = "9d798723f64a381d3aead74733f34e2e7e9236db612295deb0e752b274d842ec"

	// Future
	apiKey    = "067ae6966fecb63c04926eee4233dfa29fa9defe1a40e621dd85a91941c6bf91"
	secretKey = "582b349eba9d2d6fba35749c620c9f89d24292bd2b2871eaad1fa8f39274575d"

	helpString string = `
Valid Command Examples:
...............................
1> buy 300 eth   :-   Buy 300 USDT Worth Of ETH.

2> sell 500 xrp  :-   Sell 500 USDT Worth Of XRP At Market Price.

3> exit doge :-   Exit  Currently Open DOGE position.

4> cancel ada :- Cancel All Pending Orders For ADA

5> buy 500 btc 19000 :-   Buy 500 USDT Worth Of BTC At Limit Price of 19000.

6> sell 200 bnb 350 :-  Sell 500 USDT Worth Of BNB At Limit Price of 350.

...............................

-Command Can Be sent In Uppercase or Lowercase-
-Command Must Match Its Format To Process It Properly-
`
)

func init() {
	binance.UseTestnet = true
}

func processCommand(cmd string) {
	// client instance
	client := binance.NewFuturesClient(apiKey, secretKey)

	parsedCmd := strings.Split(strings.ToLower(cmd), " ")
	var dataList []string
	for _, v := range parsedCmd {
		dataList = append(dataList, strings.TrimSpace(v))
	}

	if len(dataList) == 0 {
		fmt.Printf("\n\n__________Invalid Command__________\n\n")
	}

	currencyPair := strings.ToUpper(dataList[1]) + "USDT"
	// Cancel all order
	if dataList[0] == "cancel" {
		err := client.NewCancelAllOpenOrdersService().Symbol(currencyPair).Do(context.Background())
		if err != nil {
			fmt.Println(err)
		}
	}

	// Exit the service
	if dataList[0] == "exit" {
		os.Exit(0)
	}

	// Help Message
	if dataList[0] == "help" {
		fmt.Println(helpString)
	}

	// Exit Position (incomplete)
	if len(dataList) == 2 {
		if dataList[0] == "exit" {
			order.MarketEnterPosition(client, currencyPair, "BUY", "500")
		}
	}
}

func main() {
	// Stop printing error stack
	// defer func() {
	// 	if r := recover(); r != nil {
	// 		fmt.Println(r)
	// 		os.Exit(0)
	// 	}
	// }()

	fmt.Printf("__________WELCOME__________\n\n")
	for {
		start := time.Now()
		var rawString string
		fmt.Print("> ")
		_, err := fmt.Scanln(&rawString)
		if err != nil {
			// fmt.Println(err)
			fmt.Printf("\n\n__________Invalid Command__________\n\n")
			fmt.Println(helpString)
			continue
		}
		processCommand(rawString)
		fmt.Printf("Time taken %v\n", time.Since(start))
	}

	// binance.UseTestnet = true
	// client := binance.NewFuturesClient(apiKey, secretKey)
	// Enable RateLimit
	// client.NewRateLimitService().Do(context.Background())

	// price, err := client.NewListPricesService().Symbol("BTCUSDT").Do(context.Background())
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(price[0].Price)

	// currencyPair := "BTCUSDT"
	// tradeSide := "BUY"
	// entryPrice := "10000.0"
	// lotSize := order.GetLimitOrderLotSize("500.0", entryPrice)

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
	// r, err := order.GetMarketOrderLotSize(client, currencyPair, "500.0")
	// if err != nil {
	// 	fmt.Println("Get Market order Lot Size", err)
	// }
	// fmt.Println(r)

}
