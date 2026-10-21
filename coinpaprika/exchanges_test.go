package coinpaprika

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/suite"
)

type ExchangesTestSuite struct {
	suite.Suite
	paprikaClient *Client
}

func (suite *ExchangesTestSuite) SetupTest() {
	paprikaClient := NewClient(nil)
	suite.NotNil(paprikaClient)

	suite.paprikaClient = paprikaClient
}

func (suite *ExchangesTestSuite) TestList() {
	exchanges, err := suite.paprikaClient.Exchanges.List(nil)
	suite.NoError(err)
	suite.NotEmpty(exchanges)
}

func (suite *ExchangesTestSuite) TestGetByID() {
	exchange, err := suite.paprikaClient.Exchanges.GetByID("binance", nil)
	suite.NoError(err)
	suite.NotEmpty(exchange)
}

func (suite *ExchangesTestSuite) TestGetMarketsByExchangeID() {
	markets, err := suite.paprikaClient.Exchanges.GetMarketsByExchangeID("binance", nil)
	suite.NoError(err)
	suite.NotEmpty(markets)
}

// TestGetByIDRejectsPathEscapingIDs proves an exchangeID that could retarget
// the endpoint or inject a query is rejected before any request is built.
func (suite *ExchangesTestSuite) TestGetByIDRejectsPathEscapingIDs() {
	rt := &recordingRoundTripper{}
	client := NewClient(&http.Client{Transport: rt})

	badIDs := []string{"a/b", "..", "../key/info", "binance?x=1", "binance#f"}
	for _, id := range badIDs {
		_, err := client.Exchanges.GetByID(id, nil)
		suite.Error(err, "id %q should be rejected before a request is built", id)
	}
	suite.Empty(rt.requests, "no request should reach the transport for a rejected id")
}

// TestGetByIDEncodesSafeID proves a normal exchangeID still produces exactly
// one encoded path segment under /exchanges, unchanged from before this fix.
func (suite *ExchangesTestSuite) TestGetByIDEncodesSafeID() {
	rt := &recordingRoundTripper{}
	client := NewClient(&http.Client{Transport: rt})

	_, _ = client.Exchanges.GetByID("binance", nil)
	suite.Require().Len(rt.requests, 1)
	suite.Equal("/v1/exchanges/binance", rt.requests[0].URL.Path)
}

func TestExchangesTestSuite(t *testing.T) {
	suite.Run(t, new(ExchangesTestSuite))
}
