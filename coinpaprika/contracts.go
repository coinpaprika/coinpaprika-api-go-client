package coinpaprika

import (
	"encoding/json"
	"fmt"
)

// ContractsService is used to get contract platforms and ticker data by contract address.
type ContractsService service

// Contract represents a contract address mapped to a coin.
type Contract struct {
	Address *string `json:"address"`
	Type    *string `json:"type"`
	ID      *string `json:"id"`
	Active  *bool   `json:"active"`
}

// ListPlatforms returns list of all contract platform IDs.
func (s *ContractsService) ListPlatforms() (platforms []string, err error) {
	url := fmt.Sprintf("%s/contracts", baseURL)

	body, err := sendGET(s.httpClient, url)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &platforms)
	return platforms, err
}

// GetByPlatform returns list of contracts for a given platform.
func (s *ContractsService) GetByPlatform(platformID string) (contracts []*Contract, err error) {
	url := fmt.Sprintf("%s/contracts/%s", baseURL, platformID)

	body, err := sendGET(s.httpClient, url)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &contracts)
	return contracts, err
}

// GetTickerByContractAddress returns ticker data for a contract address on a given platform.
func (s *ContractsService) GetTickerByContractAddress(platformID, contractAddress string, options *TickersOptions) (ticker *Ticker, err error) {
	url := fmt.Sprintf("%s/contracts/%s/%s", baseURL, platformID, contractAddress)
	url, err = constructURL(url, options)
	if err != nil {
		return nil, err
	}

	body, err := sendGET(s.httpClient, url)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &ticker)
	return ticker, err
}

// GetHistoricalTickerByContractAddress returns historical ticker data for a contract address on a given platform.
func (s *ContractsService) GetHistoricalTickerByContractAddress(platformID, contractAddress string, options *TickersHistoricalOptions) (tickersHistorical []*TickerHistorical, err error) {
	url := fmt.Sprintf("%s/contracts/%s/%s/historical", baseURL, platformID, contractAddress)
	url, err = constructURL(url, options)
	if err != nil {
		return nil, err
	}

	body, err := sendGET(s.httpClient, url)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &tickersHistorical)
	return tickersHistorical, err
}
