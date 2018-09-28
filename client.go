package coinpaprika

import (
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
	TickerService
	SearchService
	CoinsService
	GlobalService
}

type service struct {
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

	c.TickerService.httpClient = c.httpClient
	c.SearchService.httpClient = c.httpClient
	c.CoinsService.httpClient = c.httpClient
	c.GlobalService.httpClient = c.httpClient

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
