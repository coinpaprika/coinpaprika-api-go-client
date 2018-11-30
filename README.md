# Coinpaprika API Go Client

[![Build Status](https://travis-ci.org/coinpaprika/coinpaprika-api-go-client.svg?branch=master)](https://travis-ci.org/coinpaprika/coinpaprika-api-go-client)
[![go-doc](https://godoc.org/github.com/coinpaprika/coinpaprika-api-go-client?status.svg)](https://godoc.org/github.com/coinpaprika/coinpaprika-api-go-client)
[![Go Report Card](https://goreportcard.com/badge/github.com/coinpaprika/coinpaprika-api-go-client)](https://goreportcard.com/report/github.com/coinpaprika/coinpaprika-api-go-client)


## Usage

This library provides convenient way to use [coinpaprika.com API](https://api.coinpaprika.com/) in Go.

[Coinpaprika](https://coinpaprika.com) delivers full market data to the world of crypto: coin prices, volumes, market caps, ATHs, return rates and more.

## Install

```bash
go get -u github.com/coinpaprika/coinpaprika-api-go-client
```

## Getting started

```go
func main() {
	paprikaClient, err := coinpaprika.NewClient()
	if err != nil {
		panic(err)
	}

	coinTickers, err := paprikaClient.GetTickers()
	if err != nil {
		panic(err)
	}

	for _, c := range coinTickers {
		fmt.Println("Name", c.Name)
		fmt.Println("Symbol", c.Symbol)
		fmt.Println("Rank", c.Rank)
	}
}

```

## Setting custom HTTP client

```go
customClient := &http.Client{Timeout: 10 * time.Second}

paprikaClient, err := coinpaprika.NewClient(coinpaprika.SetHTTPClient(customClient))
if err != nil {
    panic(err)
}
```


## Examples

Check out the [`./examples`](./examples) directory.


## API Status (version 1.3)

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
- [ ] List tags
- [ ] Get tag by ID

### Tickers (new)
- [ ] Get tickers for all coins
- [ ] Get ticker information for specific coin
- [ ] Get historical tickers for specific coin

### Exchanges
- [ ] List exchanges
- [ ] Get exchange by ID
- [ ] List markets by exchange ID

### Search
- [x] Search tool

### Ticker (deprecated)
- [x] Get tickers for all coins
- [x] Get ticker information for specific coin

## License

CoinpaprikaAPI is available under the MIT license. See the [LICENSE file](./LICENSE.md) for more info.
