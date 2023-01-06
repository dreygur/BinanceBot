package tests

import (
	"binancebot/service"
	"encoding/json"
	"io/ioutil"
	"log"
	"testing"
)

var (
	usdtSize     string = "300.00"
	tradeSide    string = "BUY"
	currencyPair string = "ETHUSDT"
)

func getSettings() service.OrderService {
	// Read the settings file
	content, err := ioutil.ReadFile("../settings.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	// Unmarshall the settings data into `settings`
	var o service.Service
	err = json.Unmarshal(content, &o)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	return service.NewClient(o.ApiKey, o.ApiKey, o.UseTestnet)
}

func TestGetMarketOrderLotSize(t *testing.T) {
	client := getSettings()

	_, err := client.GetMarketOrderLotSize(currencyPair, usdtSize)
	if err != nil {
		t.Error(err)
	}
}

func TestMarketEnterPosition(t *testing.T) {
	client := getSettings()

	lotSize, err := client.GetMarketOrderLotSize(currencyPair, usdtSize)
	if err != nil {
		t.Error(err)
	}

	_, err = client.MarketEnterPosition(currencyPair, tradeSide, lotSize)
	if err != nil {
		t.Error(err)
	}
}

func TestLimitEnterPosition(t *testing.T) {
	client := getSettings()

	lotSize, err := client.GetMarketOrderLotSize(currencyPair, usdtSize)
	if err != nil {
		t.Error(err)
	}

	_, err = client.LimitEnterPosition(currencyPair, tradeSide, lotSize, "18000")
	if err != nil {
		t.Error(err)
	}
}
