package coinpaprika

import (
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
	suite.NoError(err)
	suite.NotEmpty(tickers)

	suite.Len(tickers, 10)
}

func TestTickerTestSuite(t *testing.T) {
	suite.Run(t, new(TickerTestSuite))
}
