package coinpaprika

import (
	"testing"

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
	tickers, err := suite.paprikaClient.Tickers.List()
	suite.NoError(err)
	suite.NotEmpty(tickers)
}

func (suite *TickerTestSuite) TestGetByID() {
	ticker, err := suite.paprikaClient.Tickers.GetByID("btc-bitcoin")
	suite.NoError(err)
	suite.NotEmpty(ticker)

	suite.Equal("Bitcoin", *ticker.Name)
	suite.Equal("BTC", *ticker.Symbol)
	suite.Equal("btc-bitcoin", *ticker.ID)

	suite.NotZero(ticker.Quotes["USD"].Price)
}

func TestTickerTestSuite(t *testing.T) {
	suite.Run(t, new(TickerTestSuite))
}
