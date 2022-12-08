package main

import (
	"fmt"

	"github.com/coinpaprika/coinpaprika-api-go-client/v2/coinpaprika"
)

func main() {
	paprikaClient := coinpaprika.NewClient(nil, coinpaprika.WithAPIKey("your_api_key_goes_here"))

	tickers, err := paprikaClient.Tickers.List(nil)
	if err != nil {
		panic(err)
	}

	for _, t := range tickers {
		if t.Name == nil || t.Symbol == nil || t.Rank == nil {
			continue
		}

		fmt.Println("Name:", *t.Name)
		fmt.Println("Symbol:", *t.Symbol)
		fmt.Println("Rank:", *t.Rank)
		fmt.Println("----")
	}
}
