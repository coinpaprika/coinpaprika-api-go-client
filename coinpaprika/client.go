package coinpaprika

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"unicode"

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
		Contracts      ContractsService
		Key            KeyService
		Changelog      ChangelogService
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
		host          string
	}
)

// RoundTrip adds the Authorization header only to https requests aimed at the
// configured API host, so the key is never sent to any other host this process
// talks to, and never survives a redirect to a different host or to plain http.
func (t *authTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	if req.URL.Scheme == "https" && req.URL.Host == t.host {
		req = req.Clone(req.Context())
		req.Header.Set("Authorization", t.apiKey)
	}
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
	c.Contracts.httpClient = c.httpClient
	c.Key.httpClient = c.httpClient
	c.Changelog.httpClient = c.httpClient

	return c
}

// WithAPIKey sets API key enabling access to Coinpaprika Pro API.
// https://api-pro.coinpaprika.com is used.
func WithAPIKey(apiKey string) ClientOptions {
	return func(a *Client) {
		baseURL = baseProURL

		proHost := baseProURL
		if parsed, err := url.Parse(baseProURL); err == nil {
			proHost = parsed.Host
		}

		baseTransport := http.DefaultTransport
		if a.httpClient != nil && a.httpClient.Transport != nil {
			baseTransport = a.httpClient.Transport
		}

		// Never mutate a client the SDK does not own: copy the caller's
		// client (or http.DefaultClient, if none was given) instead of
		// replacing its Transport in place.
		copied := http.Client{}
		if a.httpClient != nil {
			copied = *a.httpClient
		}
		copied.Transport = &authTransport{
			apiKey:        apiKey,
			baseTransport: baseTransport,
			host:          proHost,
		}
		a.httpClient = &copied
	}
}

// encodePathSegment validates a caller-supplied id and returns it encoded as a
// single, safe path segment. It rejects ids that could change which endpoint,
// path or query the request actually hits: empty ids, ids that are (or, after
// one percent-decode, decode to) "." or "..", and ids containing "/", "\",
// "?", "#", "%", whitespace, or control characters. The check runs on the raw
// id, not only the encoded form, because url.PathEscape alone leaves "." and
// ".." untouched.
func encodePathSegment(id string) (string, error) {
	if id == "" {
		return "", fmt.Errorf("id must not be empty")
	}
	if id == "." || id == ".." {
		return "", fmt.Errorf("id must not be %q", id)
	}

	decoded, decodeErr := url.QueryUnescape(id)
	if decodeErr != nil {
		decoded = id
	}
	if decoded == "." || decoded == ".." {
		return "", fmt.Errorf("id must not decode to %q", decoded)
	}

	for _, r := range decoded {
		switch {
		case r == '/' || r == '\\' || r == '?' || r == '#' || r == '%':
			return "", fmt.Errorf("id must not contain %q", string(r))
		case unicode.IsSpace(r) || unicode.IsControl(r):
			return "", fmt.Errorf("id must not contain whitespace or control characters")
		}
	}

	escaped := url.PathEscape(id)
	if strings.ContainsAny(escaped, "/\\") {
		// Defensive: escaped must still be a single path segment.
		return "", fmt.Errorf("id did not encode to a single path segment")
	}

	return escaped, nil
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
