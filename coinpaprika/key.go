package coinpaprika

import (
	"encoding/json"
	"fmt"
)

// KeyService is used to get API key information.
type KeyService service

// UsagePeriod represents API key usage data for a period.
type UsagePeriod struct {
	RequestsMade *int64 `json:"requests_made"`
	RequestsLeft *int64 `json:"requests_left"`
}

// Usage represents API key usage information.
type Usage struct {
	Message      *string      `json:"message"`
	CurrentMonth *UsagePeriod `json:"current_month"`
}

// KeyInfo represents API key information.
type KeyInfo struct {
	Plan          *string `json:"plan"`
	PlanStartedAt *string `json:"plan_started_at"`
	PlanStatus    *string `json:"plan_status"`
	PortalURL     *string `json:"portal_url"`
	Usage         *Usage  `json:"usage"`
}

// GetInfo returns information about the API key used for authentication.
func (s *KeyService) GetInfo() (keyInfo *KeyInfo, err error) {
	url := fmt.Sprintf("%s/key/info", baseURL)

	body, err := sendGET(s.httpClient, url)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &keyInfo)
	return keyInfo, err
}
