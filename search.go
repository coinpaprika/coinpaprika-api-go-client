package coinpaprika

import (
	"encoding/json"
	"fmt"
)

type SearchService service

// Person represents a person.
type Person struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	TeamsCount int    `json:"teams_count"`
}

// Tag represents a tag. Tag can be associated with coins or ICO projects.
type Tag struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	CoinCounter int    `json:"coin_counter"`
	ICOCounter  int    `json:"ico_counter"`
}

// Exchange represents an exchange.
type Exchange struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Rank int    `json:"rank"`
}

// ICO represents an ICO project.
type ICO struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
	IsNew  bool   `json:"is_new"`
}

// SearchResult represents a result of search API endpoint.
type SearchResult struct {
	Currencies []*CoinInfo `json:"currencies"`
	ICOS       []*ICO      `json:"icos"`
	Exchanges  []*Exchange `json:"exchanges"`
	People     []*Person   `json:"people"`
	Tags       []*Tag      `json:"tags"`
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
