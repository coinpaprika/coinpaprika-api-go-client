package coinpaprika

import (
	"encoding/json"
	"fmt"
	"time"
)

// TickersService is used to get ticker information
type TickersService service

// Ticker represents ticker information.
type Ticker struct {
	ID                *string          `json:"id"`
	Name              *string          `json:"name"`
	Symbol            *string          `json:"symbol"`
	Rank              *int64           `json:"rank"`
	CirculatingSupply *int64           `json:"circulating_supply"`
	TotalSupply       *int64           `json:"total_supply"`
	MaxSupply         *int64           `json:"max_supply"`
	BetaValue         *float64         `json:"beta_value"`
	LastUpdated       *string          `json:"last_updated"`
	Quotes            map[string]Quote `json:"quotes"`
}

// Quote represents coin price data in quote currency.
type Quote struct {
	Price               *float64 `json:"price"`
	Volume24h           *float64 `json:"volume_24h"`
	Volume24hChange24h  *float64 `json:"volume_24h_change_24h"`
	MarketCap           *float64 `json:"market_cap"`
	MarketCapChange24h  *float64 `json:"market_cap_change_24h"`
	PercentChange1h     *float64 `json:"percent_change_1h"`
	PercentChange12h    *float64 `json:"percent_change_12h"`
	PercentChange24h    *float64 `json:"percent_change_24h"`
	PercentChange7d     *float64 `json:"percent_change_7d"`
	PercentChange30d    *float64 `json:"percent_change_30d"`
	PercentChange1y     *float64 `json:"percent_change_1y"`
	ATHPrice            *float64 `json:"ath_price"`
	ATHDate             *string  `json:"ath_date"`
	PercentFromPriceATH *float64 `json:"percent_from_price_ath"`
}

// TickerHistorical represents historical ticker information.
type TickerHistorical struct {
	Price     *float64   `json:"price"`
	Volume24h *float64   `json:"volume_24h"`
	MarketCap *float64   `json:"market_cap"`
	Timestamp *time.Time `json:"timestamp"`
}

// TickersOptions specifies optional parameters for tickers endpoints.
type TickersOptions struct {
	Quotes string `url:"quotes,omitempty"`
}

// TickersHistoricalOptions specifies optional parameters for tickers historical endpoint.
type TickersHistoricalOptions struct {
	Start    time.Time `url:"start"`
	End      time.Time `url:"end,omitempty"`
	Limit    int       `url:"limit,omitempty"`
	Quote    string    `url:"quote,omitempty"`
	Interval string    `url:"interval,omitempty"`
}

// List gets ticker information for all coins listed on coinpaprika.
func (s *TickersService) List(options *TickersOptions) (tickers []*Ticker, err error) {
	url := fmt.Sprintf("%s/tickers", baseURL)
	url, err = constructURL(url, options)
	if err != nil {
		return nil, err
	}

	body, err := sendGET(s.httpClient, url)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(body, &tickers); err != nil {
		return tickers, err
	}

	return tickers, err
}

// GetByID gets ticker information for specific coin by id (eg. btc-bitcoin).
func (s *TickersService) GetByID(coinID string, options *TickersOptions) (ticker *Ticker, err error) {
	url := fmt.Sprintf("%s/tickers/%s", baseURL, coinID)
	url, err = constructURL(url, options)
	if err != nil {
		return nil, err
	}

	body, err := sendGET(s.httpClient, url)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(body, &ticker); err != nil {
		return ticker, err
	}

	return ticker, err
}

// GetHistoricalTickersByID gets historical ticker information for specific coin by id (eg. btc-bitcoin).
func (s *TickersService) GetHistoricalTickersByID(coinID string, options *TickersHistoricalOptions) (tickersHistorical []*TickerHistorical, err error) {
	url := fmt.Sprintf("%s/tickers/%s/historical", baseURL, coinID)
	url, err = constructURL(url, options)
	if err != nil {
		return nil, err
	}

	body, err := sendGET(s.httpClient, url)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(body, &tickersHistorical); err != nil {
		return tickersHistorical, err
	}

	return tickersHistorical, err
}
