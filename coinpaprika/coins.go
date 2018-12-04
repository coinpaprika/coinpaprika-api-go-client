package coinpaprika

import (
	"encoding/json"
	"fmt"
)

// CoinsService is used to get coins information.
type CoinsService service

// Parent represents coin parent information.
type Parent struct {
	ID     *string `json:"id"`
	Name   *string `json:"name"`
	Symbol *string `json:"symbol"`
}

// Whitepaper represents coin whitepaper.
type Whitepaper struct {
	Link      *string `json:"link"`
	Thumbnail *string `json:"thumbnail"`
}

// Coin stores basic currency information.
type Coin struct {
	ID       *string `json:"id"`
	Name     *string `json:"name"`
	Symbol   *string `json:"symbol"`
	Rank     *int64  `json:"rank"`
	IsNew    *bool   `json:"is_new"`
	IsActive *bool   `json:"is_active"`

	// Extended information
	Parent            *Parent             `json:"parent"`
	OpenSource        *bool               `json:"open_source"`
	HardwareWallet    *bool               `json:"hardware_wallet"`
	Description       *string             `json:"description"`
	Message           *string             `json:"message"`
	StartedAt         *string             `json:"started_at"`
	DevelopmentStatus *string             `json:"development_status"`
	ProofType         *string             `json:"proof_type"`
	OrgStructure      *string             `json:"org_structure"`
	HashAlgorithm     *string             `json:"hash_algorithm"`
	Whitepaper        *Whitepaper         `json:"whitepaper"`
	Links             map[string][]string `json:"links"`
	Tags              []Tag               `json:"tags"`
	Team              []Person            `json:"team"`
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

// GetByID gets coin by id (eg. btc-bitcoin).
func (s *CoinsService) GetByID(coinID string) (coin *Coin, err error) {
	url := fmt.Sprintf("%s/coins/%s", baseURL, coinID)

	body, err := sendGET(s.httpClient, url)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(body, &coin); err != nil {
		return coin, err
	}

	return coin, err
}
