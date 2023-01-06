package service

import (
	"github.com/adshao/go-binance/v2"
	"github.com/adshao/go-binance/v2/futures"
)

type Service struct {
	ApiKey     string `json:"apiKey"`
	Secretkey  string `json:"secretkey"`
	UseTestnet bool   `json:"testnet" default:"false"`
}

var (
	client *futures.Client
)

type OrderService interface {
	MarketEnterPosition(currencyPair, tradeSide, lotSize string) (*futures.CreateOrderResponse, error)
	LimitEnterPosition(currencyPair, tradeSide, lotSize, entryPrice string) (*futures.CreateOrderResponse, error)
	MarketExitPosition(currencyPair string) (*futures.CreateOrderResponse, error)
	GetMarketOrderLotSize(currencyPair, usdtSize string) (string, error)
	GetLimitOrderLotSize(usdt, limit string) string
	GetOpenPosition(currencyPair string) (*futures.PositionRisk, error)
	CancelOrders(currencyPair string) error
}

func NewClient(apikey, secretkey string, useTestnet bool) OrderService {

	// Use testnet?
	if useTestnet {
		futures.UseTestnet = true
	}

	// Future Client
	client = binance.NewFuturesClient(apikey, secretkey)

	return &Service{
		ApiKey:     apikey,
		Secretkey:  secretkey,
		UseTestnet: useTestnet,
	}
}
