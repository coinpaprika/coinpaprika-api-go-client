package coinpaprika

type GlobalStats struct {
	MarketCapUSD               int64   `json:"market_cap_usd"`
	Volume24hUSD               int64   `json:"volume_24h_usd"`
	BitcoinDominancePercentage float64 `json:"bitcoin_dominance_percentage"`
	CryptocurrenciesNumber     int64   `json:"cryptocurrencies_number"`
	LastUpdated                int64   `json:"last_updated"`
}
