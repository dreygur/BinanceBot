package order

import (
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

func NewClient(apikey, secretkey string, useTestnet bool) OrderInterface {

	// Use testnet?
	if useTestnet {
		futures.UseTestnet = true
	}

	return &Order{
		ApiKey:     apikey,
		Secretkey:  secretkey,
		UseTestnet: useTestnet,
		// Future Client
		Client: binance.NewFuturesClient(apikey, secretkey),
	}
}
