package order

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/adshao/go-binance/v2"
	"github.com/adshao/go-binance/v2/futures"
)

type Order struct {
	ApiKey     string `json:"apiKey"`
	Secretkey  string `json:"secretkey"`
	UseTestnet bool   `json:"testnet" default:"false"`
	Client     *futures.Client
}

type OrderInterface interface {
	MarketEnterPosition(currencyPair, tradeSide, lotSize string) (*futures.CreateOrderResponse, error)
	LimitEnterPosition(currencyPair, tradeSide, lotSize, entryPrice string) (*futures.CreateOrderResponse, error)
	MarketExitPosition(currencyPair string) (*futures.CreateOrderResponse, error)
	GetMarketOrderLotSize(currencyPair, usdtSize string) (string, error)
	GetLimitOrderLotSize(usdt, limit string) string
	GetOpenPosition(currencyPair string) (*futures.PositionRisk, error)
	CancelOrders(currencyPair string) error
}

func NewClient() OrderInterface {
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
	o.Client = binance.NewFuturesClient(o.ApiKey, o.Secretkey)

	return &o
}
