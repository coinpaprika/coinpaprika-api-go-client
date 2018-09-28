package coinpaprika

// Tag represents a tag. Tag can be associated with coins or ICO projects.
type Tag struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	CoinCounter int    `json:"coin_counter"`
	ICOCounter  int    `json:"ico_counter"`
}
