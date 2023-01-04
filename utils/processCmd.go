package utils

import (
	"binancebot/order"
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/adshao/go-binance/v2/futures"
)

var HelpString string = `
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

func ProcessCommand(client *futures.Client, cmd string) {
	var currencyPair string
	parsedCmd := strings.Split(strings.ToLower(cmd), " ")
	var dataList []string
	for _, v := range parsedCmd {
		dataList = append(dataList, strings.TrimSpace(v))
	}

	if len(dataList) == 0 {
		fmt.Printf("\n\n__________Invalid Command__________\n\n")
	}

	// Cancel all order
	if dataList[0] == "cancel" {
		currencyPair = strings.ToUpper(dataList[1]) + "USDT"
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
		fmt.Println(HelpString)
	}

	// Exit Position (incomplete)
	if len(dataList) == 2 {
		if dataList[0] == "exit" {
			currencyPair = strings.ToUpper(dataList[1]) + "USDT"
			order.MarketExitPosition(client, currencyPair, "BUY", "500")
		}
	}

	// Market Order
	if len(dataList) == 3 {
		usdtSize := dataList[1]
		tradeSide := strings.ToUpper(dataList[0])
		currencyPair = strings.ToUpper(dataList[2]) + "USDT"

		fmt.Println(currencyPair, tradeSide, usdtSize)

		res, err := order.MarketEnterPosition(client, currencyPair, tradeSide, usdtSize)
		if err != nil {
			fmt.Println("Error:", err)
		}
		if res != nil {
			fmt.Print("\nMarket Order Executed Successfully\n")
		}
	}

	// Limit Order
	if len(dataList) == 4 {
		usdtSize := dataList[1]
		entryPrice := dataList[3]
		tradeSide := strings.ToUpper(dataList[0])
		currencyPair = strings.ToUpper(dataList[2]) + "USDT"

		lotSize := order.GetLimitOrderLotSize(usdtSize, entryPrice)
		res, err := order.LimitEnterPosition(client, currencyPair, tradeSide, lotSize, entryPrice)
		if err != nil {
			fmt.Println("Error:", err)
		}
		if res != nil {
			fmt.Print("\nLimit Order Executed Successfully\n")
		}
	}
}
