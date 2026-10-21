package coinpaprika

import (
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type CoinsTestSuite struct {
	suite.Suite
	paprikaClient *Client
}

func (suite *CoinsTestSuite) SetupTest() {
	paprikaClient := NewClient(nil)
	suite.NotNil(paprikaClient)

	suite.paprikaClient = paprikaClient
}

func (suite *CoinsTestSuite) TestList() {
	coins, err := suite.paprikaClient.Coins.List()
	suite.NoError(err)
	suite.NotEmpty(coins)
}

func (suite *CoinsTestSuite) TestGetByID() {
	coin, err := suite.paprikaClient.Coins.GetByID("btc-bitcoin")
	suite.NoError(err)
	suite.NotEmpty(coin)
}

func (suite *CoinsTestSuite) TestGetTwitterTimelineByCoinID() {
	timeline, err := suite.paprikaClient.Coins.GetTwitterTimelineByCoinID("btc-bitcoin")
	if err != nil {
		// This endpoint is deprecated and may return errors
		return
	}
	suite.NotEmpty(timeline)
}

func (suite *CoinsTestSuite) TestGetEventsByCoinID() {
	events, err := suite.paprikaClient.Coins.GetEventsByCoinID("btc-bitcoin")
	suite.NoError(err)
	suite.NotEmpty(events)
}

func (suite *CoinsTestSuite) TestGetExchangesByCoinID() {
	exchanges, err := suite.paprikaClient.Coins.GetExchangesByCoinID("btc-bitcoin")
	suite.NoError(err)
	suite.NotEmpty(exchanges)
}

func (suite *CoinsTestSuite) TestGetMarketsByCoinID() {
	markets, err := suite.paprikaClient.Coins.GetMarketsByCoinID("btc-bitcoin")
	suite.NoError(err)
	suite.NotEmpty(markets)
}

func (suite *CoinsTestSuite) TestGetLatestOHLCVByCoinID() {
	entries, err := suite.paprikaClient.Coins.GetLatestOHLCVByCoinID("btc-bitcoin", nil)
	suite.NoError(err)
	suite.NotEmpty(entries)
	suite.Len(entries, 1)
}

func (suite *CoinsTestSuite) TestGetLatestOHLCVByCoinIDWithQuote() {
	options := &LatestOHLCVOptions{Quote: "btc"}
	entries, err := suite.paprikaClient.Coins.GetLatestOHLCVByCoinID("eth-ethereum", options)
	suite.NoError(err)
	suite.NotEmpty(entries)
	suite.Len(entries, 1)
}

func (suite *CoinsTestSuite) TestGetTodayOHLCVByCoinID() {
	entries, err := suite.paprikaClient.Coins.GetTodayOHLCVByCoinID("btc-bitcoin", nil)
	suite.NoError(err)
	suite.NotEmpty(entries)
}

func (suite *CoinsTestSuite) TestGetMappings() {
	options := &MappingsOptions{Coinpaprika: "btc-bitcoin"}
	_, err := suite.paprikaClient.Coins.GetMappings(options)
	if err != nil {
		// Expected to fail without a Business+ API key
		suite.Contains(err.Error(), "status code: 4")
	}
}

func (suite *CoinsTestSuite) TestGetHistoricalOHLCVByCoinID() {
	options := &HistoricalOHLCVOptions{
		Start: time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC),
		End:   time.Date(2018, 1, 11, 0, 0, 0, 0, time.UTC),
		Quote: "btc",
	}
	entries, err := suite.paprikaClient.Coins.GetHistoricalOHLCVByCoinID("btc-bitcoin", options)
	if err != nil {
		// Historical OHLCV data before a certain date may require a paid plan
		suite.Contains(err.Error(), "status code: 4")
		return
	}
	suite.NotEmpty(entries)
	suite.Len(entries, 11)
}

// TestGetByIDRejectsPathEscapingIDs proves a coinID that could retarget the
// endpoint or inject a query ("a/b", "..", "btc?x=1", "btc#f") is rejected
// before any request is built, using a stub transport so nothing is sent.
func (suite *CoinsTestSuite) TestGetByIDRejectsPathEscapingIDs() {
	rt := &recordingRoundTripper{}
	client := NewClient(&http.Client{Transport: rt})

	badIDs := []string{"a/b", "..", "../key/info", "btc?x=1", "btc#f"}
	for _, id := range badIDs {
		_, err := client.Coins.GetByID(id)
		suite.Error(err, "id %q should be rejected before a request is built", id)
	}
	suite.Empty(rt.requests, "no request should reach the transport for a rejected id")
}

// TestGetByIDEncodesSafeID proves a normal coinID still produces exactly one
// encoded path segment under /coins, unchanged from before this fix.
func (suite *CoinsTestSuite) TestGetByIDEncodesSafeID() {
	rt := &recordingRoundTripper{}
	client := NewClient(&http.Client{Transport: rt})

	_, _ = client.Coins.GetByID("btc-bitcoin")
	suite.Require().Len(rt.requests, 1)
	suite.Equal("/v1/coins/btc-bitcoin", rt.requests[0].URL.Path)
	suite.Empty(rt.requests[0].URL.RawQuery)
}

func TestCoinsTestSuite(t *testing.T) {
	suite.Run(t, new(CoinsTestSuite))
}
