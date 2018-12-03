package main

import (
	"fmt"

	"github.com/coinpaprika/coinpaprika-api-go-client/coinpaprika"
)

func main() {
	paprikaClient := coinpaprika.NewClient(nil)

	tickers, err := paprikaClient.Tickers.List()
	if err != nil {
		panic(err)
	}

	for idx, t := range tickers {
		fmt.Println("Name:", t.Name)
		fmt.Println("Symbol:", t.Symbol)
		fmt.Println("Rank:", t.Rank)
		if quoteUSD, ok := t.Quotes["USD"]; ok {
			fmt.Println("Price USD:", quoteUSD.Price)
		}
		fmt.Println()
		if idx >= 2 {
			break
		}
	}

}
