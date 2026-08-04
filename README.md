# Coinpaprika API Go Client

[![Build Status](https://github.com/coinpaprika/coinpaprika-api-go-client/actions/workflows/main.yml/badge.svg)](https://github.com/coinpaprika/coinpaprika-api-go-client/actions)
[![Go Reference](https://pkg.go.dev/badge/github.com/coinpaprika/coinpaprika-api-go-client/v2.svg)](https://pkg.go.dev/github.com/coinpaprika/coinpaprika-api-go-client/v2)
[![Go Report Card](https://goreportcard.com/badge/github.com/coinpaprika/coinpaprika-api-go-client)](https://goreportcard.com/report/github.com/coinpaprika/coinpaprika-api-go-client)


## Usage

This library provides convenient way to use [coinpaprika.com API](https://api.coinpaprika.com/) in Go.

[Coinpaprika](https://coinpaprika.com) delivers full market data to the world of crypto: coin prices, volumes, market caps, ATHs, return rates and more.

## Installation

```sh
go get github.com/coinpaprika/coinpaprika-api-go-client/v2
```

## Getting started

```go
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
```

## Setting custom HTTP client

```go
customClient := &http.Client{Timeout: 10 * time.Second}
paprikaClient := coinpaprika.NewClient(customClient)
```

## Setting API key for and enabling access to Coinpaprika Pro API
Key can be obtained from [Coinpaprika API](https://coinpaprika.com/api/)

```go
paprikaClient := coinpaprika.NewClient(nil, coinpaprika.WithAPIKey("your_api_key_goes_here"))
```

## Examples

Check out the [`./examples`](./examples) directory.


## Implementation status

### Global
- [x] Get market overview data

### Coins
- [x] List coins
- [x] Get coin by ID
- [x] ~~Get twitter timeline for coin~~ (deprecated)
- [x] Get coin events by coin ID
- [x] Get exchanges by coin ID
- [x] Get markets by coin ID
- [x] Get latest OHLCV
- [x] Get historical OHLCV
- [x] Get today OHLCV
- [x] Get ID mappings (Business+)

### People
- [x] Get people by ID

### Tags
- [x] List tags
- [x] Get tag by ID

### Tickers
- [x] Get tickers for all coins
- [x] Get ticker information for specific coin
- [x] Get historical tickers for specific coin

### Exchanges
- [x] List exchanges
- [x] Get exchange by ID
- [x] List markets by exchange ID

### Contracts
- [x] List contract platforms
- [x] Get contracts by platform
- [x] Get ticker by contract address
- [x] Get historical ticks by contract address

### Search
- [x] Search tool

### Price Converter
- [x] Price converter

### Key
- [x] Get API key info (Pro+)

### Changelog
- [x] Get ID changelog (Starter+)


## Other Coinpaprika Tools

Looking for other ways to integrate crypto data? Check out these tools:

- **[Coinpaprika CLI](https://github.com/coinpaprika/coinpaprika-cli)**: crypto market data from your terminal. 12,000+ cryptocurrencies, 350+ exchanges, prices, OHLCV. Free tier included.
- **[Coinpaprika MCP Server](https://github.com/coinpaprika/coinpaprika-mcp)**: plug crypto market data directly into AI assistants like Claude. 30+ tools, no API key needed to get started.
- **[DexPaprika](https://dexpaprika.com)**: DEX data across 36 chains and 230+ DEXes, over 96% of on-chain DEX volume. If you need on-chain trading data, pools, or token prices, start here.

## License

CoinpaprikaAPI is available under the MIT license. See the [LICENSE file](./LICENSE.md) for more info.
