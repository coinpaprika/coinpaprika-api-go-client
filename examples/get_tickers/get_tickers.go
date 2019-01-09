package main

import (
	"fmt"

	"github.com/coinpaprika/coinpaprika-api-go-client/v2/coinpaprika"
)

func main() {
	paprikaClient := coinpaprika.NewClient(nil)

	tickers, err := paprikaClient.Tickers.List(nil)
	if err != nil {
		panic(err)
	}

	for idx, t := range tickers {
		if t.Name == nil || t.Symbol == nil || t.Rank == nil {
			continue
		}

		fmt.Println("Name:", *t.Name)
		fmt.Println("Symbol:", *t.Symbol)
		fmt.Println("Rank:", *t.Rank)
		if quoteUSD, ok := t.Quotes["USD"]; ok {
			fmt.Printf("Price: %.2f USD\n\n", *quoteUSD.Price)
		}

		if idx >= 2 {
			break
		}
	}

}
