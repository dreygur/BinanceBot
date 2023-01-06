package order

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/adshao/go-binance/v2"
	"github.com/adshao/go-binance/v2/futures"
	"github.com/common-nighthawk/go-figure"
)

type Order struct {
	ApiKey     string `json:"apiKey"`
	ApiSecret  string `json:"secretkey"`
	UseTestnet bool   `json:"testnet" default:"false"`
	client     *futures.Client
}

type OrderInterface interface {
	MarketEnterPosition(currencyPair, tradeSide, lotSize string) (*futures.CreateOrderResponse, error)
	LimitEnterPosition(currencyPair, tradeSide, lotSize, entryPrice string) (*futures.CreateOrderResponse, error)
	MarketExitPosition(currencyPair string) (*futures.CreateOrderResponse, error)
	GetMarketOrderLotSize(currencyPair, usdtSize string) (string, error)
	GetLimitOrderLotSize(usdt, limit string) string
	GetOpenPosition(currencyPair string) (*futures.PositionRisk, error)
}

func NewOrder() *Order {
	// Read the settings file
	content, err := ioutil.ReadFile("./settings.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	// Unmarshall the settings data into `settings`
	var o Order
	err = json.Unmarshal(content, &o)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}
	// Use testnet?
	if o.UseTestnet {
		futures.UseTestnet = true
	}

	// Future Client
	o.client = binance.NewFuturesClient(o.ApiKey, o.ApiSecret)

	// Print the logo
	fmt.Println()
	myFigure := figure.NewColorFigure("BINANCE BOT", "digital", "green", true)
	myFigure.Print()
	fmt.Println()

	return &o
}
