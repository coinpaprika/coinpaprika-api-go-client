package coinpaprika

import (
	"encoding/json"
	"fmt"
)

// ExchangesService is used to get exchanges and markets data.
type ExchangesService service

// ExchangeQuote represents exchange volume data in quote currency.
type ExchangeQuote struct {
	ReportedVolume24h *float64 `json:"reported_volume_24h"`
	AdjustedVolume24h *float64 `json:"adjusted_volume_24h"`
}

// ExchangeFiat represents fiat traded on exchange.
type ExchangeFiat struct {
	Name   *string `json:"name"`
	Symbol *string `json:"symbol"`
}

// Exchange represents an exchange.
type Exchange struct {
	ID                     *string                  `json:"id"`
	Name                   *string                  `json:"name"`
	Message                *string                  `json:"message"`
	Active                 *bool                    `json:"active"`
	MarketsDataFetched     *bool                    `json:"markets_data_fetched"`
	Rank                   *int64                   `json:"rank"`
	AdjustedRank           *int64                   `json:"adjusted_rank"`
	ReportedRank           *int64                   `json:"reported_rank"`
	Currencies             *int64                   `json:"currencies"`
	Markets                *int64                   `json:"markets"`
	AdjustedVolume24hShare *float64                 `json:"adjusted_volume_24h_share"`
	Fiats                  []ExchangeFiat           `json:"fiats"`
	Quotes                 map[string]ExchangeQuote `json:"quotes"`
	Links                  map[string][]string      `json:"links"`
	LastUpdated            *string                  `json:"last_updated"`
}

// MarketQuote represents market price and volume data in quote currency.
type MarketQuote struct {
	Price     *float64 `json:"price"`
	Volume24h *float64 `json:"volume_24h"`
}

// Market represents a market.
type Market struct {
	Pair                   *string                `json:"pair"`
	BaseCurrencyID         *string                `json:"base_currency_id"`
	BaseCurrencyName       *string                `json:"base_currency_name"`
	QuoteCurrencyID        *string                `json:"quote_currency_id"`
	QuoteCurrencyName      *string                `json:"quote_currency_name"`
	MarketURL              *string                `json:"market_url"`
	Category               *string                `json:"category"`
	FeeType                *string                `json:"fee_type"`
	Outlier                *bool                  `json:"outlier"`
	ReportedVolume24hShare *float64               `json:"reported_volume_24h_share"`
	Quotes                 map[string]MarketQuote `json:"quotes"`
	LastUpdated            *string                `json:"last_updated"`
	ExchangeID             *string                `json:"exchange_id"`
	ExchangeName           *string                `json:"exchange_name"`
	AdjustedVolume24hShare *float64               `json:"adjusted_volume_24h_share"`
}

// ExchangesOptions specifies optional parameters for exchanges endpoints.
type ExchangesOptions struct {
	Quote string `url:"quote,omitempty"`
}

// MarketsOptions specifies optional parameters for markets endpoint.
type MarketsOptions struct {
	Quote string `url:"quote,omitempty"`
}

// List returns list of all exchanges listed on coinpaprika.
func (s *ExchangesService) List(options *ExchangesOptions) (exchanges []*Exchange, err error) {
	url := fmt.Sprintf("%s/exchanges", baseURL)
	url, err = constructURL(url, options)
	if err != nil {
		return nil, err
	}

	body, err := sendGET(s.httpClient, url)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &exchanges)
	return exchanges, err
}

// GetByID gets exchange by ID.
func (s *ExchangesService) GetByID(exchangeID string, options *ExchangesOptions) (exchange *Exchange, err error) {
	url := fmt.Sprintf("%s/exchanges/%s", baseURL, exchangeID)
	url, err = constructURL(url, options)
	if err != nil {
		return nil, err
	}

	body, err := sendGET(s.httpClient, url)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &exchange)
	return exchange, err
}

// GetMarketsByExchangeID gets list of markets for an exchange.
func (s *ExchangesService) GetMarketsByExchangeID(exchangeID string, options *MarketsOptions) (markets []*Market, err error) {
	url := fmt.Sprintf("%s/exchanges/%s/markets", baseURL, exchangeID)
	url, err = constructURL(url, options)
	if err != nil {
		return nil, err
	}

	body, err := sendGET(s.httpClient, url)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &markets)
	return markets, err
}
