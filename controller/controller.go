package controller

import (
	// "crypto/tls"

	"binancebot/service"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type ProxyController interface {
	ProxyHandler(writer http.ResponseWriter, request *http.Request)
}
type controller struct{}

var (
	orderService service.OrderService
)

func NewProxyController(orders service.OrderService) ProxyController {
	orderService = orders
	return &controller{}
}

func (*controller) ProxyHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-type", "application/json")

	orders, err := orderService.GetOpenPosition("ETHUSDT")
	if err != nil {
		log.Print(fmt.Sprint(err), "error")
		return
	}

	orderData, err := json.Marshal(orders)
	if err != nil {
		log.Print(fmt.Sprint(err), "error")
		return
	}

	writer.Write(orderData)
}
