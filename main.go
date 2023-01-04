package main

import (
	"binancebot/utils"
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/adshao/go-binance/v2"
	"github.com/adshao/go-binance/v2/futures"
)

var (
	// Future
	apiKey    = "067ae6966fecb63c04926eee4233dfa29fa9defe1a40e621dd85a91941c6bf91"
	secretKey = "582b349eba9d2d6fba35749c620c9f89d24292bd2b2871eaad1fa8f39274575d"

	// Future Client
	client *futures.Client
)

func main() {
	// Stop printing error stack
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
			os.Exit(0)
		}
	}()

	futures.UseTestnet = true
	client = binance.NewFuturesClient(apiKey, secretKey)

	fmt.Printf("\n__________WELCOME__________\n\n")
	for {
		fmt.Print("> ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan() // use `for scanner.Scan()` to keep reading
		rawString := scanner.Text()
		start := time.Now()
		utils.ProcessCommand(client, rawString)
		fmt.Printf("Time taken %v\n", time.Since(start))
	}
}
