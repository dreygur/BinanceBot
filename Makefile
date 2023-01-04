run:
	go run bin/main.go

build:
	GOOS=linux GOARCH=amd64 go build -o BinanceBot bin/main.go