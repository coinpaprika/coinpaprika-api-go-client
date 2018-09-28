package coinpaprika

// Person represents a person.
type Person struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	TeamsCount int    `json:"teams_count"`
}
