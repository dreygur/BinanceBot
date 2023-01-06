package main

import (
	"binancebot/controller"
	"binancebot/router"
	"binancebot/service"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/common-nighthawk/go-figure"
)

type Settings struct {
	ApiKey     string `json:"apiKey"`
	Secretkey  string `json:"secretkey"`
	SslCert    string `json:"sslCert"`
	SslKey     string `json:"sslKey"`
	UseTestnet bool   `json:"testnet" default:"false"`
	UseHTTP2   bool   `json:"useHTTP2" default:"false"`
}

var (
	settings      Settings
	orderService  service.OrderService
	botController controller.RouteController
	httpRouter    router.Router
)

func init() {
	// Print the logo
	fmt.Println()
	myFigure := figure.NewColorFigure("BINANCE BOT API SERVER", "digital", "green", true)
	myFigure.Print()
	fmt.Println()

	// Read the settings file
	content, err := ioutil.ReadFile("./settings.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	// Unmarshall the settings data into `settings`
	err = json.Unmarshal(content, &settings)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	// Server Settings
	httpSettings := router.Settings{
		UseHTTP2: settings.UseHTTP2,
		SslCert:  settings.SslCert,
		SslKey:   settings.SslKey,
	}

	// Initiate services
	orderService = service.NewClient(settings.ApiKey, settings.Secretkey, settings.UseTestnet)
	botController = controller.NewRouteController(orderService)
	httpRouter = router.NewStdRouter(httpSettings)
}

func main() {
	// Stop printing error stack
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
			os.Exit(0)
		}
	}()

	// Routes
	httpRouter.HandleReq("/", botController.RouteHandler)
	httpRouter.HandleReq("/openposition", botController.RouteHandler)

	// HTTP Server
	httpRouter.Serve("8080")
}
