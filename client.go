package coinpaprika

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	userAgent = "Coinpaprika API Client - Go"
	baseURL   = "https://api.coinpaprika.com/v1"
)

// OptionFunc is a function that is used to configure a Client.
type OptionFunc func(*Client) error

// Client can be used to get data from coinpaprika API.
type Client struct {
	httpClient *http.Client
}

// NewClient creates a new client to work with coinpaprika API.
func NewClient(options ...OptionFunc) (*Client, error) {
	c := &Client{
		httpClient: http.DefaultClient,
	}

	for _, o := range options {
		if err := o(c); err != nil {
			return nil, err
		}
	}

	return c, nil
}

// SetHTTPClient can be used to specify the http.Client to use when making HTTP requests.
func SetHTTPClient(httpClient *http.Client) OptionFunc {
	return func(c *Client) error {
		if httpClient != nil {
			c.httpClient = httpClient
		} else {
			c.httpClient = http.DefaultClient
		}
		return nil
	}
}

// GetGlobalStats gets market overview data.
func (c *Client) GetGlobalStats() (*GlobalStats, error) {
	url := fmt.Sprintf("%s/global", baseURL)

	body, err := sendGET(c.httpClient, url)
	if err != nil {
		return nil, err
	}

	var g GlobalStats
	if err := json.Unmarshal(body, &g); err != nil {
		return nil, err
	}

	return &g, nil
}

// GetTickersUnconverted gets ticker information for all coins listed on coinpaprika. Returned data is in original string format.
func (c *Client) GetTickersUnconverted() (tickersUnconverted []*CoinTickerUnconverted, err error) {
	url := fmt.Sprintf("%s/ticker", baseURL)

	body, err := sendGET(c.httpClient, url)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(body, &tickersUnconverted); err != nil {
		return tickersUnconverted, err
	}

	return tickersUnconverted, err
}

// GetTickers gets ticker information for all coins listed on coinpaprika. Returned data is automatically parsed.
func (c *Client) GetTickers() (coinTickers []*CoinTicker, err error) {
	tickersUnconverted, err := c.GetTickersUnconverted()
	if err != nil {
		return nil, err
	}

	for _, ticker := range tickersUnconverted {
		ct, err := ticker.convert()
		if err != nil {
			return coinTickers, err
		}
		coinTickers = append(coinTickers, ct)
	}

	return coinTickers, nil
}

// GetTickerByIDUnconverted gets ticker information for specific coin by id (eg. btc-bitcoin). Returned data is in original string format.
func (c *Client) GetTickerByIDUnconverted(id string) (tickerUnconverted *CoinTickerUnconverted, err error) {
	url := fmt.Sprintf("%s/ticker/%s", baseURL, id)

	body, err := sendGET(c.httpClient, url)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(body, &tickerUnconverted); err != nil {
		return tickerUnconverted, err
	}

	return tickerUnconverted, err
}

// GetTickerByID gets ticker information for specific coin by id (eg. btc-bitcoin). Returned data is automatically parsed.
func (c *Client) GetTickerByID(id string) (coinTickers *CoinTicker, err error) {
	tickerUnconverted, err := c.GetTickerByIDUnconverted(id)
	if err != nil {
		return nil, err
	}

	return tickerUnconverted.convert()

}

// GetCoins returns list of all active coins listed on coinpaprika.
func (c *Client) GetCoins() (coins []*CoinInfo, err error) {
	url := fmt.Sprintf("%s/coins", baseURL)

	body, err := sendGET(c.httpClient, url)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(body, &coins); err != nil {
		return coins, err
	}

	return coins, err
}

func constructSearchURL(query string, options *SearchOptions) string {
	url := fmt.Sprintf("%s/search?q=%s", baseURL, query)

	if options == nil {
		return url
	}

	if options.Categories != "" {
		url = fmt.Sprintf("%s&c=%s", url, options.Categories)
	}
	if options.Limit != 0 {
		url = fmt.Sprintf("%s&limit=%v", url, options.Limit)
	}

	return url
}

// SearchOptions specifies optional parameters for search endpoint.
type SearchOptions struct {
	// Comma separated categories to include in search results.
	// Available options: currencies|exchanges|icos|people|tags.
	// Eg. "currencies,exchanges"
	Categories string

	// The number of results per category.
	Limit int
}

// Search returns a list of currencies, exchanges, icos, people and tags for given query.
func (c *Client) Search(query string, options *SearchOptions) (searchResult *SearchResult, err error) {
	url := constructSearchURL(query, options)

	body, err := sendGET(c.httpClient, url)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(body, &searchResult); err != nil {
		return searchResult, err
	}

	return searchResult, nil
}

func sendGET(client *http.Client, url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", userAgent)

	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code: %v, body: %s", response.StatusCode, string(body))
	}

	return body, nil
}
