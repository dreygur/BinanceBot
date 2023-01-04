package tests

import (
	"binancebot/order"
	"testing"
)

func TestLimitEnterPosition(t *testing.T) {
	currencyPair := "ETHUSDT"

	lotSize, err := order.GetMarketOrderLotSize(client, currencyPair, "300")
	if err == nil {
		_, err := order.LimitEnterPosition(client, currencyPair, "BUY", lotSize, "100")
		if err != nil {
			t.Error(err)
		}
	}
}
