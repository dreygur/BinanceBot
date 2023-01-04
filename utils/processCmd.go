package utils

import (
	"binancebot/order"
	"context"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var HelpString string = `
Valid Command Examples:
..........................................................................
1) buy 300 eth        : Buy 300 USDT Worth Of ETH
2) sell 500 eth       : Sell 500 USDT Worth Of ETH At Market Price
3) exit eth           : Exit  Currently Open ETH position
4) cancel eth         : Cancel All Pending Orders For ETH
5) buy 500 btc 19000  : Buy 500 USDT Worth Of BTC At Limit Price of 19000
6) sell 200 bnb 350   : Sell 500 USDT Worth Of BNB At Limit Price of 350
..........................................................................

* Command Can Be sent In Uppercase or Lowercase
* Command Must Match Its Format To Process It Properly
`

func ProcessCommand(cmd string) {
	var re = regexp.MustCompile(`(?m).*msg=(?P<Message>.*)`)

	var currencyPair string
	parsedCmd := strings.Split(strings.ToLower(cmd), " ")
	var dataList []string
	for _, v := range parsedCmd {
		dataList = append(dataList, strings.TrimSpace(v))
	}

	if len(dataList) == 0 {
		fmt.Printf("\n\n__________Invalid Command__________\n\n")
	}

	// Exit the service
	if dataList[0] == "close" {
		os.Exit(0)
	}

	// Help Message
	if dataList[0] == "help" {
		fmt.Println(HelpString)
	}

	if len(dataList) == 2 {
		// Exit Position
		if dataList[0] == "exit" {
			currencyPair = strings.ToUpper(dataList[1]) + "USDT"
			res, err := order.MarketExitPosition(client, currencyPair)
			if err != nil {
				fmt.Println("Error:", re.FindStringSubmatch(err.Error())[1])
			}

			if res != nil {
				fmt.Printf("\nPosition Closed For: %s\n", currencyPair)
			}
		}

		// Cancel all order
		if dataList[0] == "cancel" {
			currencyPair = strings.ToUpper(dataList[1]) + "USDT"
			err := client.NewCancelAllOpenOrdersService().Symbol(currencyPair).Do(context.Background())
			if err != nil {
				fmt.Println("Error:", re.FindStringSubmatch(err.Error())[1])
			}

			if err == nil {
				fmt.Printf("\nOrders Canceled For: %s\n", currencyPair)
			}
		}
	}

	// Market Order
	if len(dataList) == 3 {
		usdtSize := dataList[1]
		tradeSide := strings.ToUpper(dataList[0])
		currencyPair = strings.ToUpper(dataList[2]) + "USDT"

		lotSize, err := order.GetMarketOrderLotSize(client, currencyPair, usdtSize)
		if err != nil {
			fmt.Println("Error:", err)
		}

		res, err := order.MarketEnterPosition(client, currencyPair, tradeSide, lotSize)
		if err != nil {
			fmt.Println("Error:", re.FindStringSubmatch(err.Error())[1])
		}
		if res != nil {
			// fmt.Printf(
			// 	"\n*** Market Order Filled ***\nSymbol: %s\nSide: %s\nFill Price: %s\nSize: %s\nUSD Value: %s\n\n",
			// 	currencyPair,
			// 	tradeSide,
			// 	res.AvgPrice,
			// 	lotSize,
			// 	res.CumQuote,
			// )

			fmt.Printf(
				"\n*** Market Order Filled ***\nSymbol: %s\nSide: %s\nSize: %s\n\n",
				currencyPair,
				tradeSide,
				lotSize,
			)
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
			fmt.Println("Error:", re.FindStringSubmatch(err.Error())[1])
		}
		if res != nil {
			fmt.Printf(
				"\n*** Limit Order Filled ***\nSymbol: %s\nSide: %s\nSize: %s\n\n",
				currencyPair,
				tradeSide,
				lotSize,
			)
		}
	}
}
