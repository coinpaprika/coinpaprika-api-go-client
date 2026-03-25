package coinpaprika

import (
	"encoding/json"
	"fmt"
)

// ChangelogService is used to get changelog data.
type ChangelogService service

// ChangelogEntry represents a single ID change entry.
type ChangelogEntry struct {
	CurrencyID *string `json:"currency_id"`
	OldID      *string `json:"old_id"`
	NewID      *string `json:"new_id"`
	ChangedAt  *string `json:"changed_at"`
}

// ChangelogOptions specifies optional parameters for changelog endpoints.
type ChangelogOptions struct {
	Page  int `url:"page,omitempty"`
	Limit int `url:"limit,omitempty"`
}

// GetIDChanges returns list of ID changes in the coinpaprika system.
func (s *ChangelogService) GetIDChanges(options *ChangelogOptions) (entries []*ChangelogEntry, err error) {
	url := fmt.Sprintf("%s/changelog/ids", baseURL)
	url, err = constructURL(url, options)
	if err != nil {
		return nil, err
	}

	body, err := sendGET(s.httpClient, url)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &entries)
	return entries, err
}
