package coinpaprika

import (
	"encoding/json"
	"fmt"
)

// TagsService is used for listing and getting tags.
type TagsService service

// TagsOptions specifies optional parameters for tags endpoint.
type TagsOptions struct {
	AdditionalFields string `url:"additional_fields,omitempty"`
}

// Tag represents a tag (related to currency or ico).
type Tag struct {
	ID          *string  `json:"id"`
	Name        *string  `json:"name"`
	CoinCounter *int64   `json:"coin_counter"`
	ICOCounter  *int64   `json:"ico_counter"`
	Description *string  `json:"description"`
	Type        *string  `json:"type"`
	Coins       []string `json:"coins"`
	ICOs        []string `json:"icos"`
}

// List returns a list of all tags.
func (s *TagsService) List(options *TagsOptions) (tags []*Tag, err error) {
	url := fmt.Sprintf("%s/tags", baseURL)
	url, err = constructURL(url, options)
	if err != nil {
		return nil, err
	}

	body, err := sendGET(s.httpClient, url)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &tags)
	return tags, err
}

// GetByID return a tag by id.
func (s *TagsService) GetByID(tagID string, options *TagsOptions) (tag *Tag, err error) {
	url := fmt.Sprintf("%s/tags/%s", baseURL, tagID)
	url, err = constructURL(url, options)
	if err != nil {
		return nil, err
	}

	body, err := sendGET(s.httpClient, url)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &tag)
	return tag, err
}
