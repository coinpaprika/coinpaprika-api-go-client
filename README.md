# Coinpaprika API Go Client

[![Build Status](https://travis-ci.org/coinpaprika/coinpaprika-api-go-client.svg?branch=master)](https://travis-ci.org/coinpaprika/coinpaprika-api-go-client)
[![go-doc](https://godoc.org/github.com/coinpaprika/coinpaprika-api-go-client?status.svg)](https://godoc.org/github.com/coinpaprika/coinpaprika-api-go-client/coinpaprika)
[![Go Report Card](https://goreportcard.com/badge/github.com/coinpaprika/coinpaprika-api-go-client)](https://goreportcard.com/report/github.com/coinpaprika/coinpaprika-api-go-client)


## Usage

This library provides convenient way to use [coinpaprika.com API](https://api.coinpaprika.com/) in Go.

[Coinpaprika](https://coinpaprika.com) delivers full market data to the world of crypto: coin prices, volumes, market caps, ATHs, return rates and more.

## Getting started

```go
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

	for _, t := range tickers {
		fmt.Println("Name:", *t.Name)
		fmt.Println("Symbol:", *t.Symbol)
		fmt.Println("Rank:", *t.Rank)
		fmt.Println("----")
	}
}
```

## Setting custom HTTP client

```
customClient := &http.Client{Timeout: 10 * time.Second}
paprikaClient := coinpaprika.NewClient(customClient)
```


## Examples

Check out the [`./examples`](./examples) directory.


## Implementation status (API v1.3.0)

### Global
- [x] Get market overview data

### Coins
- [x] List coins
- [ ] Get coin by ID
- [ ] Get twitter timeline for coin
- [ ] Get coin events by coin ID
- [ ] Get exchanges by coin ID
- [ ] Get markets by coin ID
- [ ] Get latest OHLC
- [ ] Get historical OHLC

### People
- [ ] Get people by ID

### Tags
- [x] List tags
- [x] Get tag by ID

### Tickers 
- [x] Get tickers for all coins
- [x] Get ticker information for specific coin
- [ ] Get historical ticker for specific coin

### Exchanges
- [ ] List exchanges
- [ ] Get exchange by ID
- [ ] List markets by exchange ID

### Search
- [x] Search tool


## License

CoinpaprikaAPI is available under the MIT license. See the [LICENSE file](./LICENSE.md) for more info.
