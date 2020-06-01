package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/coinpaprika/coinpaprika-api-go-client/coinpaprika"
)

func main() {
	customClient := &http.Client{Timeout: 10 * time.Second}

	paprikaClient := coinpaprika.NewClient(customClient)

	coins, err := paprikaClient.Global.Get()
	if err != nil {
		panic(err)
	}

	if coins.MarketCapUSD != nil {
		fmt.Printf("Current market cap: %.2f USD \n", *coins.MarketCapUSD)
	}
}
