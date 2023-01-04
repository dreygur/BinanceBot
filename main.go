package main

import (
	"binancebot/utils"
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/adshao/go-binance/v2"
	"github.com/adshao/go-binance/v2/futures"

	figure "github.com/common-nighthawk/go-figure"
)

type Settings struct {
	ApiKey     string `json:"apiKey"`
	Secretkey  string `json:"secretkey"`
	UseTestnet bool   `json:"testnet" default:"false"`
}

func main() {
	// Stop printing error stack
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
			os.Exit(0)
		}
	}()

	// Read the settings file
	content, err := ioutil.ReadFile("./settings.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	// Unmarshall the settings data into `settings`
	var settings Settings
	err = json.Unmarshal(content, &settings)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	// Use testnet?
	if settings.UseTestnet {
		futures.UseTestnet = true
	}
	// Future Client
	client := binance.NewFuturesClient(settings.ApiKey, settings.Secretkey)

	// Print the logo
	fmt.Println()
	myFigure := figure.NewColorFigure("BINANCE BOT", "digital", "green", true)
	myFigure.Print()
	fmt.Println()
	// fmt.Printf("\n__________WELCOME__________\n\n")

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
