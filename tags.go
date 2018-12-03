package coinpaprika

import (
	"encoding/json"
	"fmt"
	"net/url"
)

// TagsService is used for listing and getting tags.
type TagsService service

// TagsOptions specifies optional parameters for tags endpoint.
type TagsOptions struct {
	AdditionalFields string
}

type Tag struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Type        string   `json:"type"`
	CoinCounter int      `json:"coin_counter"`
	ICOCounter  int      `json:"ico_counter"`
	Coins       []string `json:"coins"`
}

func constructTagsURL(tagID *string, options *TagsOptions) string {
	uri := fmt.Sprintf("%s/tags", baseURL)
	if tagID != nil {
		uri = fmt.Sprintf("%s/%s", uri, *tagID)
	}

	if options == nil {
		return uri
	}

	if options.AdditionalFields != "" {
		v := url.Values{}
		v.Set("additional_fields", options.AdditionalFields)
		uri = fmt.Sprintf("%s?%s", uri, v.Encode())
	}

	return uri
}

// ListTags returns a list of all tags.
func (s *TagsService) List(options *TagsOptions) (tags []*Tag, err error) {
	uri := constructTagsURL(nil, options)

	body, err := sendGET(s.httpClient, uri)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(body, &tags); err != nil {
		return tags, err
	}

	return tags, nil
}

func (s *TagsService) GetByID(tagID string, options *TagsOptions) (tag *Tag, err error) {
	uri := constructTagsURL(&tagID, options)

	body, err := sendGET(s.httpClient, uri)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(body, &tag); err != nil {
		return tag, err
	}

	return tag, nil
}
