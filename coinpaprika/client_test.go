package coinpaprika

import (
	"io"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

// recordingRoundTripper is a stub http.RoundTripper that records every
// request it receives instead of sending it anywhere, so tests in this
// package can assert what the SDK built without any network access.
type recordingRoundTripper struct {
	requests  []*http.Request
	responses []*http.Response
}

func (rt *recordingRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	rt.requests = append(rt.requests, req)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(strings.NewReader("{}")),
		Header:     make(http.Header),
		Request:    req,
	}
	rt.responses = append(rt.responses, resp)
	return resp, nil
}

type ClientTestSuite struct {
	suite.Suite
	paprikaClient *Client
}

func (suite *ClientTestSuite) SetupTest() {
	paprikaClient := NewClient(nil)
	suite.NotNil(paprikaClient)

	suite.paprikaClient = paprikaClient
}

func (suite *ClientTestSuite) TestNewClientDefault() {
	client := NewClient(nil)
	suite.NotNil(client)
	suite.Equal(http.DefaultClient, client.httpClient)
}

func (suite *ClientTestSuite) TestNewClientSetHTTPClient() {
	customHTTPClient := &http.Client{Timeout: time.Second * 10}

	client := NewClient(customHTTPClient)
	suite.NotNil(client)
	suite.Equal(customHTTPClient, client.httpClient)
}

// TestWithAPIKeyDoesNotMutateDefaultClient proves WithAPIKey(nil-client case)
// no longer installs authTransport onto http.DefaultClient, which used to
// leak the API key to every host any other code in the process reaches
// through http.Get, http.Post or http.DefaultClient.
func (suite *ClientTestSuite) TestWithAPIKeyDoesNotMutateDefaultClient() {
	originalTransport := http.DefaultClient.Transport
	originalBaseURL := baseURL
	defer func() {
		http.DefaultClient.Transport = originalTransport
		baseURL = originalBaseURL
	}()

	client := NewClient(nil, WithAPIKey("dummy-verifier-key"))
	suite.NotNil(client)

	suite.Equal(originalTransport, http.DefaultClient.Transport)
	_, isAuthTransport := http.DefaultClient.Transport.(*authTransport)
	suite.False(isAuthTransport, "http.DefaultClient.Transport must not become *authTransport")
}

// TestAuthTransportOnlyAddsAuthorizationForConfiguredHost proves the key is
// attached only to https requests aimed at the configured Pro API host, not
// to requests the same transport sees for an unrelated host or over plain
// http (the shape of a cross-host redirect).
func (suite *ClientTestSuite) TestAuthTransportOnlyAddsAuthorizationForConfiguredHost() {
	originalBaseURL := baseURL
	defer func() { baseURL = originalBaseURL }()

	rt := &recordingRoundTripper{}
	client := NewClient(&http.Client{Transport: rt}, WithAPIKey("dummy-verifier-key"))

	transport := client.httpClient.Transport
	suite.NotNil(transport)

	proReq, err := http.NewRequest("GET", "https://api-pro.coinpaprika.com/v1/coins/btc-bitcoin", nil)
	suite.NoError(err)
	proResp, err := transport.RoundTrip(proReq)
	suite.NoError(err)
	defer proResp.Body.Close()

	unrelatedReq, err := http.NewRequest("GET", "https://unrelated-host.example/unrelated", nil)
	suite.NoError(err)
	unrelatedResp, err := transport.RoundTrip(unrelatedReq)
	suite.NoError(err)
	defer unrelatedResp.Body.Close()

	plainHTTPReq, err := http.NewRequest("GET", "http://api-pro.coinpaprika.com/v1/coins/btc-bitcoin", nil)
	suite.NoError(err)
	plainHTTPResp, err := transport.RoundTrip(plainHTTPReq)
	suite.NoError(err)
	defer plainHTTPResp.Body.Close()

	suite.Require().Len(rt.requests, 3)
	suite.NotEmpty(rt.requests[0].Header.Get("Authorization"), "the configured https API host must receive the key")
	suite.Empty(rt.requests[1].Header.Get("Authorization"), "an unrelated host must never receive the key")
	suite.Empty(rt.requests[2].Header.Get("Authorization"), "a plain http request must never receive the key")
}

func (suite *ClientTestSuite) TestEncodePathSegmentRejectsUnsafeIDs() {
	badIDs := []string{
		"",
		".",
		"..",
		"a/b",
		"a\\b",
		"btc?x=1",
		"btc#f",
		"a b",
		"../key/info",
		"%2e%2e",
		"%2f",
		"a%b",
	}
	for _, id := range badIDs {
		_, err := encodePathSegment(id)
		suite.Error(err, "expected id %q to be rejected", id)
	}
}

func (suite *ClientTestSuite) TestEncodePathSegmentEncodesSafeIDs() {
	cases := map[string]string{
		"btc-bitcoin":     "btc-bitcoin",
		"vitalik-buterin": "vitalik-buterin",
		"eth-ethereum":    "eth-ethereum",
	}
	for id, want := range cases {
		got, err := encodePathSegment(id)
		suite.NoError(err)
		suite.Equal(want, got)
	}
}

func TestClientTestSuite(t *testing.T) {
	suite.Run(t, new(ClientTestSuite))
}
