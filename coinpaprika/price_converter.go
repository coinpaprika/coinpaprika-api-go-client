package coinpaprika

import (
	"encoding/json"
	"fmt"
	"time"
)

// PriceConverterService is used for price converter requests
type PriceConverterService service

// PriceConverterResult represents a result of price converter API endpoint.
type PriceConverterResult struct {
	BaseCurrencyID        *string    `json:"base_currency_id"`
	BaseCurrencyName      *string    `json:"base_currency_name"`
	BasePriceLastUpdated  *time.Time `json:"base_price_last_updated"`
	QuoteCurrencyID       *string    `json:"quote_currency_id"`
	QuoteCurrencyName     *string    `json:"quote_currency_name"`
	QuotePriceLastUpdated *time.Time `json:"quote_price_last_updated"`
	Amount                *float64   `json:"amount"`
	Price                 *float64   `json:"price"`
}

// PriceConverterOptions specifies optional parameters for price converter endpoint.
type PriceConverterOptions struct {
	BaseCurrencyID  string  `url:"base_currency_id"`
	QuoteCurrencyID string  `url:"quote_currency_id"`
	Amount          float64 `url:"amount,omitempty"`
}

// PriceConverter returns a price of the base currency amount expressed in the quote currency
func (s *PriceConverterService) PriceConverter(options *PriceConverterOptions) (priceConverterResult *PriceConverterResult, err error) {
	url := fmt.Sprintf("%s/price-converter", baseURL)
	url, err = constructURL(url, options)
	if err != nil {
		return nil, err
	}

	body, err := sendGET(s.httpClient, url)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &priceConverterResult)
	return priceConverterResult, err
}
