package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/coinpaprika/coinpaprika-api-go-client"
)

func main() {
	customClient := &http.Client{Timeout: 10 * time.Second}

	paprikaClient, err := coinpaprika.NewClient(coinpaprika.SetHTTPClient(customClient))
	if err != nil {
		panic(err)
	}

	coins, err := paprikaClient.Global.Get()
	if err != nil {
		panic(err)
	}

	fmt.Println("Current MarketCapUSD:", coins.MarketCapUSD)
}
