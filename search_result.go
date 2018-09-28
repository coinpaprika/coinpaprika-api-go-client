package coinpaprika

// SearchResult represents a result of search API endpoint.
type SearchResult struct {
	Currencies []*CoinInfo `json:"currencies"`
	ICOS       []*ICO      `json:"icos"`
	Exchanges  []*Exchange `json:"exchanges"`
	People     []*Person   `json:"people"`
	Tags       []*Tag      `json:"tags"`
}
