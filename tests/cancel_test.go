package tests

import (
	"binancebot/order"
	"testing"

	"github.com/adshao/go-binance/v2"
	"github.com/adshao/go-binance/v2/futures"
)

var (
	// Future
	apiKey    = "067ae6966fecb63c04926eee4233dfa29fa9defe1a40e621dd85a91941c6bf91"
	secretKey = "582b349eba9d2d6fba35749c620c9f89d24292bd2b2871eaad1fa8f39274575d"

	// Future Client
	client *futures.Client = binance.NewFuturesClient(apiKey, secretKey)
)

func init() {
	futures.UseTestnet = true
}

func TestMarketExitPosition(t *testing.T) {
	_, err := order.MarketExitPosition(client, "ETHUSDT")
	if err != nil {
		t.Error(err)
	}

}
