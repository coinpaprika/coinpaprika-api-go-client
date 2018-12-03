package coinpaprika

import (
	"encoding/json"
	"fmt"
)

// SearchService is used for search requests
type SearchService service

// SearchResult represents a result of search API endpoint.
type SearchResult struct {
	Currencies []*Coin     `json:"currencies"`
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
