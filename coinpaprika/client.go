package coinpaprika

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"

	"github.com/google/go-querystring/query"
)

const (
	userAgent   = "Coinpaprika API Client - Go"
	baseFreeURL = "https://api.coinpaprika.com/v1"
	baseProURL  = "https://api-pro.coinpaprika.com/v1"
)

var baseURL = baseFreeURL

// Client can be used to get data from coinpaprika API.
type (
	Client struct {
		httpClient     *http.Client
		Tickers        TickersService
		Search         SearchService
		PriceConverter PriceConverterService
		Coins          CoinsService
		Global         GlobalService
		Tags           TagsService
		People         PeopleService
		Exchanges      ExchangesService
	}
)

type (
	service struct {
		httpClient *http.Client
	}
)

type (
	authTransport struct {
		baseTransport http.RoundTripper
		apiKey        string
	}
)

func (t *authTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Add("Authorization", t.apiKey)
	return t.baseTransport.RoundTrip(req)
}

type ClientOptions func(a *Client)

// NewClient creates a new client to work with coinpaprika API.
func NewClient(httpClient *http.Client, opts ...ClientOptions) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	c := &Client{
		httpClient: httpClient,
	}

	for _, opt := range opts {
		opt(c)
	}

	c.Tickers.httpClient = c.httpClient
	c.Search.httpClient = c.httpClient
	c.PriceConverter.httpClient = c.httpClient
	c.Coins.httpClient = c.httpClient
	c.Global.httpClient = c.httpClient
	c.Tags.httpClient = c.httpClient
	c.People.httpClient = c.httpClient
	c.Exchanges.httpClient = c.httpClient

	return c
}

// WithAPIKey sets API key enabling access to Coinpaprika Pro API.
// https://api-pro.coinpaprika.com is used.
func WithAPIKey(apiKey string) ClientOptions {
	return func(a *Client) {
		baseURL = baseProURL

		baseTransport := http.DefaultTransport
		if a.httpClient.Transport != nil {
			baseTransport = a.httpClient.Transport
		}

		a.httpClient.Transport = &authTransport{
			apiKey:        apiKey,
			baseTransport: baseTransport,
		}
	}
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

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code: %v, body: %s", response.StatusCode, string(body))
	}

	return body, nil
}
