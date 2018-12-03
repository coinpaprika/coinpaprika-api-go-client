package coinpaprika

// Exchange represents an exchange.
type Exchange struct {
	ID   *string `json:"id"`
	Name *string `json:"name"`
	Rank *int64  `json:"rank"`
}
