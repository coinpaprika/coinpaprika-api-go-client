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

	coinTickers, err := paprikaClient.GetTickersUnconverted()
	if err != nil {
		panic(err)
	}

	for idx, c := range coinTickers {
		fmt.Println("Name", c.Name)
		fmt.Println("Symbol", c.Symbol)
		fmt.Println("Rank", c.Rank)
		fmt.Println("PriceUSD", c.PriceUSD)
		fmt.Println()

		if idx >= 2 {
			break
		}
	}

}
