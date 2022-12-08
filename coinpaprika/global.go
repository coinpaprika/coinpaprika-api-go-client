package coinpaprika

import (
	"encoding/json"
	"fmt"
	"time"
)

// GlobalService is used to get global market overview data
type GlobalService service

// GlobalStats stores global market overview data
type GlobalStats struct {
	MarketCapUSD               *float64   `json:"market_cap_usd"`
	Volume24hUSD               *float64   `json:"volume_24h_usd"`
	BitcoinDominancePercentage *float64   `json:"bitcoin_dominance_percentage"`
	CryptocurrenciesNumber     *int64     `json:"cryptocurrencies_number"`
	MarketCapATHValue          *float64   `json:"market_cap_ath_value"`
	MarketCapATHDate           *time.Time `json:"market_cap_ath_date"`
	Volume24hATHValue          *float64   `json:"volume_24h_ath_value"`
	Volume24hATHDate           *time.Time `json:"volume_24h_ath_date"`
	MarketCapChange24h         *float64   `json:"market_cap_change_24h"`
	Volume24hChange24h         *float64   `json:"volume_24h_change_24h"`
	LastUpdated                *int64     `json:"last_updated"`
}

// Get gets market overview data.
func (s *GlobalService) Get() (g *GlobalStats, err error) {
	url := fmt.Sprintf("%s/global", baseURL)

	body, err := sendGET(s.httpClient, url)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &g)
	return g, err
}
