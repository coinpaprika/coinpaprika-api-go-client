package coinpaprika

import (
	"encoding/json"
	"fmt"
)

// CoinsService is used to get coins information
type CoinsService service

// Coin stores basic currency information
type Coin struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Symbol   string `json:"symbol"`
	Rank     int64  `json:"rank"`
	IsNew    bool   `json:"is_new"`
	IsActive bool   `json:"is_active"`
}

// List returns list of all active coins listed on coinpaprika.
func (s *CoinsService) List() (coins []*Coin, err error) {
	url := fmt.Sprintf("%s/coins", baseURL)

	body, err := sendGET(s.httpClient, url)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(body, &coins); err != nil {
		return coins, err
	}

	return coins, err
}
