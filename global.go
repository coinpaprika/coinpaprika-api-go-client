package coinpaprika

import (
	"encoding/json"
	"fmt"
)

type GlobalService service

type GlobalStats struct {
	MarketCapUSD               int64   `json:"market_cap_usd"`
	Volume24hUSD               int64   `json:"volume_24h_usd"`
	BitcoinDominancePercentage float64 `json:"bitcoin_dominance_percentage"`
	CryptocurrenciesNumber     int64   `json:"cryptocurrencies_number"`
	LastUpdated                int64   `json:"last_updated"`
}

// GetGlobalStats gets market overview data.
func (s *GlobalService) GetGlobalStats() (*GlobalStats, error) {
	url := fmt.Sprintf("%s/global", baseURL)

	body, err := sendGET(s.httpClient, url)
	if err != nil {
		return nil, err
	}

	var g GlobalStats
	if err := json.Unmarshal(body, &g); err != nil {
		return nil, err
	}

	return &g, nil
}
