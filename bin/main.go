package main

import (
	"binancebot/order"
	"binancebot/utils"
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/common-nighthawk/go-figure"
)

func main() {
	// Stop printing error stack
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
			os.Exit(0)
		}
	}()

	// Print the logo
	fmt.Println()
	myFigure := figure.NewColorFigure("BINANCE BOT", "digital", "green", true)
	myFigure.Print()
	fmt.Println()

	// Read the settings file
	content, err := ioutil.ReadFile("./settings.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	// Unmarshall the settings data into `settings`
	var o order.Order
	err = json.Unmarshal(content, &o)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	client := order.NewClient(o.ApiKey, o.Secretkey, o.UseTestnet)

	for {
		fmt.Print("\033[32m", "~# > ", "\033[0m")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan() // use `for scanner.Scan()` to keep reading
		rawString := scanner.Text()
		start := time.Now()
		utils.ProcessCommand(client, rawString)
		fmt.Printf("Time taken %v\n", time.Since(start))
	}
}
