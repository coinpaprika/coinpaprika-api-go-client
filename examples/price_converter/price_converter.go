package main

import (
	"fmt"

	"github.com/coinpaprika/coinpaprika-api-go-client/v2/coinpaprika"
)

func main() {
	paprikaClient := coinpaprika.NewClient(nil)

	// Print the current value of 1.5 BTC expressed in USD.
	opts := &coinpaprika.PriceConverterOptions{
		BaseCurrencyID: "btc-bitcoin", QuoteCurrencyID: "usd-us-dollars", Amount: 1.5,
	}
	result, err := paprikaClient.PriceConverter.PriceConverter(opts)
	if err != nil {
		panic(err)
	}

	fmt.Printf("1.5 BTC is worth %v US Dollars\n", *result.Price)
}
