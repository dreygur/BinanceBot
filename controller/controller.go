package controller

import (
	// "crypto/tls"

	"binancebot/service"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type RouteController interface {
	RouteHandler(writer http.ResponseWriter, request *http.Request)
}
type controller struct{}

var (
	orderService service.OrderService
)

func NewRouteController(orders service.OrderService) RouteController {
	orderService = orders
	return &controller{}
}

func (*controller) RouteHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-type", "application/json")
	if request.Method != "GET" {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	query := request.URL.Query()
	fmt.Println(query["token"])

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
