package coinpaprika

import (
	"encoding/json"
	"fmt"
)

// PeopleService is used to get data about people.
type PeopleService service

// Position represents person position with relation to specific coin.
type Position struct {
	CoinID   *string `json:"coin_id"`
	CoinName *string `json:"coin_name"`
	Position *string `json:"position"`
}

// SocialLink represents a link to social media profile with followers count.
type SocialLink struct {
	URL       *string `json:"url"`
	Followers *int64  `json:"followers"`
}

// Person stores information about a person.
type Person struct {
	ID          *string                 `json:"id"`
	Name        *string                 `json:"name"`
	Position    *string                 `json:"position"`
	TeamsCount  *int64                  `json:"teams_count"`
	Description *string                 `json:"description"`
	Positions   []Position              `json:"positions"`
	Links       map[string][]SocialLink `json:"links"`
}

// GetByID gets person by person id (eg. vitalik-buterin).
func (s *PeopleService) GetByID(personID string) (person *Person, err error) {
	url := fmt.Sprintf("%s/people/%s", baseURL, personID)

	body, err := sendGET(s.httpClient, url)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &person)
	return person, err
}
