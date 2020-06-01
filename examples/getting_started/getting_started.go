package main

import (
	"fmt"

	"github.com/coinpaprika/coinpaprika-api-go-client/coinpaprika"
)

func main() {
	paprikaClient := coinpaprika.NewClient(nil)

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
