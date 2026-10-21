package coinpaprika

import (
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type TickerTestSuite struct {
	suite.Suite
	paprikaClient *Client
}

func (suite *TickerTestSuite) SetupTest() {
	paprikaClient := NewClient(nil)
	suite.NotNil(paprikaClient)

	suite.paprikaClient = paprikaClient
}

func (suite *TickerTestSuite) TestList() {
	tickers, err := suite.paprikaClient.Tickers.List(nil)
	suite.NoError(err)
	suite.NotEmpty(tickers)
}

func (suite *TickerTestSuite) TestGetByID() {
	ticker, err := suite.paprikaClient.Tickers.GetByID("btc-bitcoin", nil)
	suite.NoError(err)
	suite.NotEmpty(ticker)

	suite.Equal("Bitcoin", *ticker.Name)
	suite.Equal("BTC", *ticker.Symbol)
	suite.Equal("btc-bitcoin", *ticker.ID)

	suite.NotZero(ticker.Quotes["USD"].Price)
}

func (suite *TickerTestSuite) TestGetByIDWithQuotes() {
	options := &TickersOptions{Quotes: "BTC,ETH"}
	ticker, err := suite.paprikaClient.Tickers.GetByID("btc-bitcoin", options)
	suite.NoError(err)
	suite.NotEmpty(ticker)

	suite.Equal("Bitcoin", *ticker.Name)
	suite.Equal("BTC", *ticker.Symbol)
	suite.Equal("btc-bitcoin", *ticker.ID)

	suite.Nil(ticker.Quotes["USD"].Price)
	suite.NotNil(ticker.Quotes["BTC"].Price)
	suite.NotNil(ticker.Quotes["ETH"].Price)
}

func (suite *TickerTestSuite) TestGetHistoricalTickersByID() {
	start := time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2018, 2, 1, 0, 0, 0, 0, time.UTC)
	options := &TickersHistoricalOptions{Start: start, End: end, Limit: 10, Interval: "10m"}

	tickers, err := suite.paprikaClient.Tickers.GetHistoricalTickersByID("btc-bitcoin", options)
	if err != nil {
		// Historical minute-level data may require a paid plan
		suite.Contains(err.Error(), "status code: 4")
		return
	}
	suite.NotEmpty(tickers)
	suite.Len(tickers, 10)
}

// TestGetByIDRejectsPathEscapingIDs proves a coinID that could retarget the
// endpoint or inject a query is rejected before any request is built.
func (suite *TickerTestSuite) TestGetByIDRejectsPathEscapingIDs() {
	rt := &recordingRoundTripper{}
	client := NewClient(&http.Client{Transport: rt})

	badIDs := []string{"a/b", "..", "../key/info", "btc?x=1", "btc#f"}
	for _, id := range badIDs {
		_, err := client.Tickers.GetByID(id, nil)
		suite.Error(err, "id %q should be rejected before a request is built", id)
	}
	suite.Empty(rt.requests, "no request should reach the transport for a rejected id")
}

// TestGetByIDEncodesSafeID proves a normal coinID still produces exactly one
// encoded path segment under /tickers, unchanged from before this fix.
func (suite *TickerTestSuite) TestGetByIDEncodesSafeID() {
	rt := &recordingRoundTripper{}
	client := NewClient(&http.Client{Transport: rt})

	_, _ = client.Tickers.GetByID("btc-bitcoin", nil)
	suite.Require().Len(rt.requests, 1)
	suite.Equal("/v1/tickers/btc-bitcoin", rt.requests[0].URL.Path)
}

func TestTickerTestSuite(t *testing.T) {
	suite.Run(t, new(TickerTestSuite))
}
