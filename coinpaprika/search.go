package coinpaprika

import (
	"encoding/json"
	"fmt"
)

// SearchService is used for search requests
type SearchService service

// SearchCurrency stores basic currency information
type SearchCurrency struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Symbol   string `json:"symbol"`
	Rank     int64  `json:"rank"`
	IsNew    bool   `json:"is_new"`
	IsActive bool   `json:"is_active"`
}

// SearchPerson represents a person in search result.
type SearchPerson struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	TeamsCount int    `json:"teams_count"`
}

// SearchTag represents a tag in search result.
type SearchTag struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	CoinCounter int    `json:"coin_counter"`
	ICOCounter  int    `json:"ico_counter"`
}

// SearchExchange represents an exchange in search result.
type SearchExchange struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Rank int    `json:"rank"`
}

// SearchICO represents an ICO in search result.
type SearchICO struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
	IsNew  bool   `json:"is_new"`
}

// SearchResult represents a result of search API endpoint.
type SearchResult struct {
	Currencies []*SearchCurrency `json:"currencies"`
	ICOS       []*SearchICO      `json:"icos"`
	Exchanges  []*SearchExchange `json:"exchanges"`
	People     []*SearchPerson   `json:"people"`
	Tags       []*SearchTag      `json:"tags"`
}

func constructSearchURL(query string, options *SearchOptions) string {
	url := fmt.Sprintf("%s/search?q=%s", baseURL, query)

	if options == nil {
		return url
	}

	if options.Categories != "" {
		url = fmt.Sprintf("%s&c=%s", url, options.Categories)
	}
	if options.Limit != 0 {
		url = fmt.Sprintf("%s&limit=%v", url, options.Limit)
	}

	return url
}

// SearchOptions specifies optional parameters for search endpoint.
type SearchOptions struct {
	// Comma separated categories to include in search results.
	// Available options: currencies|exchanges|icos|people|tags.
	// Eg. "currencies,exchanges"
	Categories string

	// The number of results per category.
	Limit int
}

// Search returns a list of currencies, exchanges, icos, people and tags for given query.
func (s *SearchService) Search(query string, options *SearchOptions) (searchResult *SearchResult, err error) {
	url := constructSearchURL(query, options)

	body, err := sendGET(s.httpClient, url)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(body, &searchResult); err != nil {
		return searchResult, err
	}

	return searchResult, nil
}
