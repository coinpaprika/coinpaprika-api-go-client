package coinpaprika

// CoinInfo stores basic currency information
type CoinInfo struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Symbol   string `json:"symbol"`
	Rank     int64  `json:"rank"`
	IsNew    bool   `json:"is_new"`
	IsActive bool   `json:"is_active"`
}
