package main

import (
	"fmt"

	"github.com/coinpaprika/coinpaprika-api-go-client/coinpaprika"
)

func main() {
	paprikaClient := coinpaprika.NewClient(nil)

	// Search currencies and exchanges with "bitcoin" in name.
	searchOpts := &coinpaprika.SearchOptions{Categories: "currencies,exchanges"}
	searchResult, err := paprikaClient.Search.Search("bitcoin", searchOpts)
	if err != nil {
		panic(err)
	}

	fmt.Println("Currencies:")
	for _, c := range searchResult.Currencies {
		fmt.Printf("Rank: %-4d Name: %-20s Symbol: %-4s\n", *c.Rank, *c.Name, *c.Symbol)
	}

	fmt.Println("\nExchanges:")
	for _, c := range searchResult.Exchanges {
		fmt.Printf("Rank: %-4d Name: %-20s\n", *c.Rank, *c.Name)
	}
}
