package coinpaprika

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"

	"github.com/google/go-querystring/query"
)

const (
	userAgent = "Coinpaprika API Client - Go"
	baseURL   = "https://api.coinpaprika.com/v1"
)

// Client can be used to get data from coinpaprika API.
type Client struct {
	httpClient *http.Client
	Tickers    TickersService
	Search     SearchService
	Coins      CoinsService
	Global     GlobalService
	Tags       TagsService
	People     PeopleService
}

type service struct {
	httpClient *http.Client
}

// NewClient creates a new client to work with coinpaprika API.
func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	c := &Client{
		httpClient: httpClient,
	}

	c.Tickers.httpClient = c.httpClient
	c.Search.httpClient = c.httpClient
	c.Coins.httpClient = c.httpClient
	c.Global.httpClient = c.httpClient
	c.Tags.httpClient = c.httpClient
	c.People.httpClient = c.httpClient

	return c
}

func constructURL(rawURL string, options interface{}) (string, error) {
	if v := reflect.ValueOf(options); v.Kind() == reflect.Ptr && v.IsNil() {
		return rawURL, nil
	}

	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return rawURL, err
	}

	values, err := query.Values(options)
	if err != nil {
		return rawURL, err
	}

	parsedURL.RawQuery = values.Encode()
	return parsedURL.String(), nil
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
