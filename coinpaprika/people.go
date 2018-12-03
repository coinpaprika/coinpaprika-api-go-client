package coinpaprika

// Person stores information about a person.
type Person struct {
	ID         *string `json:"id"`
	Name       *string `json:"name"`
	TeamsCount *int64  `json:"teams_count"`
}
