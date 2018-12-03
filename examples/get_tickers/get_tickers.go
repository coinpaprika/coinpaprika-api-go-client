package main

import (
	"fmt"

	"github.com/coinpaprika/coinpaprika-api-go-client"
)

func main() {
	paprikaClient, err := coinpaprika.NewClient()
	if err != nil {
		panic(err)
	}

	tickers, err := paprikaClient.Tickers.List()
	if err != nil {
		panic(err)
	}

	for idx, t := range tickers {
		fmt.Println("Name", t.Name)
		fmt.Println("Symbol", t.Symbol)
		fmt.Println("Rank", t.Rank)
		if t.PriceUSD != nil {
			fmt.Println("PriceUSD", *t.PriceUSD)
		}
		fmt.Println()
		if idx >= 2 {
			break
		}
	}

}
