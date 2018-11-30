package coinpaprika

import (
	"encoding/json"
	"fmt"
)

// TagsService is used for listing and getting tags.
type TagsService service

// TagsOptions specifies optional parameters for tags endpoint.
type TagsOptions struct {
	AdditionalFields string
}

type TagExtended struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Type        string   `json:"type"`
	CoinCounter int      `json:"coin_counter"`
	ICOCounter  int      `json:"ico_counter"`
	Coins       []string `json:"coins"`
}

func constructTagsURL(tagID *string, options *TagsOptions) string {
	url := fmt.Sprintf("%s/tags", baseURL)
	if tagID != nil {
		url = fmt.Sprintf("%s/%s", url, *tagID)
	}

	if options == nil {
		return url
	}

	if options.AdditionalFields != "" {
		url = fmt.Sprintf("%s?additional_fields=%s", url, options.AdditionalFields)
	}

	return url
}

// GetTags returns a list of all tags.
func (s *TagsService) GetTags(options *TagsOptions) (tags []*TagExtended, err error) {
	url := constructTagsURL(nil, options)

	body, err := sendGET(s.httpClient, url)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(body, &tags); err != nil {
		return tags, err
	}

	return tags, nil
}

func (s *TagsService) GetTagByID(tagID string, options *TagsOptions) (tag *TagExtended, err error) {
	url := constructTagsURL(&tagID, options)

	body, err := sendGET(s.httpClient, url)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(body, &tag); err != nil {
		return tag, err
	}

	return tag, nil
}
