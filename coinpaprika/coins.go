package coinpaprika

import (
	"encoding/json"
	"fmt"
	"time"
)

// CoinsService is used to get coins information.
type CoinsService service

// Parent represents coin parent information.
type Parent struct {
	ID     *string `json:"id"`
	Name   *string `json:"name"`
	Symbol *string `json:"symbol"`
}

// Whitepaper represents coin whitepaper.
type Whitepaper struct {
	Link      *string `json:"link"`
	Thumbnail *string `json:"thumbnail"`
}

// Coin stores basic currency information.
type Coin struct {
	ID       *string `json:"id"`
	Name     *string `json:"name"`
	Symbol   *string `json:"symbol"`
	Rank     *int64  `json:"rank"`
	IsNew    *bool   `json:"is_new"`
	IsActive *bool   `json:"is_active"`

	// Extended information
	Parent            *Parent             `json:"parent"`
	OpenSource        *bool               `json:"open_source"`
	HardwareWallet    *bool               `json:"hardware_wallet"`
	Description       *string             `json:"description"`
	Message           *string             `json:"message"`
	StartedAt         *string             `json:"started_at"`
	DevelopmentStatus *string             `json:"development_status"`
	ProofType         *string             `json:"proof_type"`
	OrgStructure      *string             `json:"org_structure"`
	HashAlgorithm     *string             `json:"hash_algorithm"`
	Whitepaper        *Whitepaper         `json:"whitepaper"`
	Links             map[string][]string `json:"links"`
	Tags              []Tag               `json:"tags"`
	Team              []Person            `json:"team"`
}

// Tweet represents twitter timeline entry.
type Tweet struct {
	Date        *time.Time `json:"date"`
	UserName    *string    `json:"user_name"`
	Status      *string    `json:"status"`
	IsRetweet   *bool      `json:"is_retweet"`
	StatusLink  *string    `json:"status_link"`
	MediaLink   *string    `json:"media_link,omitempty"`
	YoutubeLink *string    `json:"youtube_link,omitempty"`
}

// Event represents event on related to coin.
type Event struct {
	Date           *string `json:"date"`
	DateTo         *string `json:"date_to"`
	Name           *string `json:"name"`
	Description    *string `json:"description"`
	IsConference   *bool   `json:"is_conference"`
	Link           *string `json:"link"`
	ProofImageLink *string `json:"proof_image_link"`
}

// OHLCVEntry stores OHLCV (open, high, low, close, volume) values for cryptocurrency.
type OHLCVEntry struct {
	TimeOpen  *time.Time `json:"time_open"`
	TimeClose *time.Time `json:"time_close"`
	Open      *float64   `json:"open"`
	High      *float64   `json:"high"`
	Low       *float64   `json:"low"`
	Close     *float64   `json:"close"`
	Volume    *int64     `json:"volume"`
	MarketCap *int64     `json:"market_cap"`
}

// LatestOHLCVOptions specifies optional parameters for ohlcv latest endpoint.
type LatestOHLCVOptions struct {
	Quote string `url:"quote,omitempty"`
}

// HistoricalOHLCVOptions specifies optional parameters for ohlcv historical endpoint.
type HistoricalOHLCVOptions struct {
	Start time.Time `url:"start"`
	End   time.Time `url:"end,omitempty"`
	Limit int       `url:"limit,omitempty"`
	Quote string    `url:"quote,omitempty"`
}

// List returns list of all active coins listed on coinpaprika.
func (s *CoinsService) List() (coins []*Coin, err error) {
	url := fmt.Sprintf("%s/coins", baseURL)

	body, err := sendGET(s.httpClient, url)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &coins)
	return coins, err
}

// GetByID gets coin by id (eg. btc-bitcoin).
func (s *CoinsService) GetByID(coinID string) (coin *Coin, err error) {
	url := fmt.Sprintf("%s/coins/%s", baseURL, coinID)

	body, err := sendGET(s.httpClient, url)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &coin)
	return coin, err
}

// GetTwitterTimelineByCoinID gets twitter timeline for a coin by coin id (eg. btc-bitcoin).
func (s *CoinsService) GetTwitterTimelineByCoinID(coinID string) (timeline []*Tweet, err error) {
	url := fmt.Sprintf("%s/coins/%s/twitter", baseURL, coinID)

	body, err := sendGET(s.httpClient, url)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &timeline)
	return timeline, err
}

// GetEventsByCoinID gets events for a coin by coin id (eg. btc-bitcoin).
func (s *CoinsService) GetEventsByCoinID(coinID string) (events []*Event, err error) {
	url := fmt.Sprintf("%s/coins/%s/events", baseURL, coinID)

	body, err := sendGET(s.httpClient, url)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &events)
	return events, err
}

// GetExchangesByCoinID gets exchanges for a coin by coin id.
func (s *CoinsService) GetExchangesByCoinID(coinID string) (exchanges []*Exchange, err error) {
	url := fmt.Sprintf("%s/coins/%s/exchanges", baseURL, coinID)

	body, err := sendGET(s.httpClient, url)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &exchanges)
	return exchanges, err
}

// GetMarketsByCoinID gets markets for a coin by coin id.
func (s *CoinsService) GetMarketsByCoinID(coinID string) (markets []*Market, err error) {
	url := fmt.Sprintf("%s/coins/%s/markets", baseURL, coinID)

	body, err := sendGET(s.httpClient, url)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &markets)
	return markets, err
}

// GetLatestOHLCVByCoinID gets latest ohlcv values for a coin by coin id (eg. btc-bitcoin).
func (s *CoinsService) GetLatestOHLCVByCoinID(coinID string, options *LatestOHLCVOptions) (entries []*OHLCVEntry, err error) {
	url := fmt.Sprintf("%s/coins/%s/ohlcv/latest", baseURL, coinID)
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

// GetHistoricalOHLCVByCoinID gets historical ohlcv values for a coin by coin id (eg. btc-bitcoin).
func (s *CoinsService) GetHistoricalOHLCVByCoinID(coinID string, options *HistoricalOHLCVOptions) (entries []*OHLCVEntry, err error) {
	url := fmt.Sprintf("%s/coins/%s/ohlcv/historical", baseURL, coinID)
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
