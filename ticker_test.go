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
	paprikaClient, err := NewClient()
	suite.NoError(err)
	suite.NotNil(paprikaClient)

	suite.paprikaClient = paprikaClient
}

func (suite *TickerTestSuite) TestConvert() {
	unconverted := TickerUnconverted{
		Name:         "TestCoin",
		Symbol:       "TC",
		Rank:         "100",
		PriceUSD:     "1.234243",
		PriceBTC:     "0.0",
		MarketCapUSD: "",
		Volume24hUSD: "2000",
		LastUpdated:  "1525088839",
	}

	ticker, err := unconverted.convert()
	suite.NoError(err)
	suite.NotNil(ticker)

	if ticker != nil {
		suite.Equal("TestCoin", ticker.Name)
		suite.Equal("TC", ticker.Symbol)
		suite.Equal(int64(100), ticker.Rank)

		suite.NotNil(ticker.PriceUSD)
		suite.Equal(1.234243, *ticker.PriceUSD)

		suite.NotNil(ticker.PriceBTC)
		suite.Zero(*ticker.PriceBTC)

		suite.Nil(ticker.MarketCapUSD)

		suite.NotNil(ticker.Volume24hUSD)
		suite.Equal(int64(2000), *ticker.Volume24hUSD)

		suite.Equal(time.Date(2018, 4, 30, 11, 47, 19, 0, time.UTC), ticker.LastUpdated.UTC())
	}

}

func (suite *TickerTestSuite) TestGetTickers() {
	tickers, err := suite.paprikaClient.Tickers.List()
	suite.NoError(err)
	suite.NotEmpty(tickers)
}

func (suite *TickerTestSuite) TestGetTickersUnconverted() {
	tickers, err := suite.paprikaClient.Tickers.ListUnconverted()
	suite.NoError(err)
	suite.NotEmpty(tickers)
}

func (suite *TickerTestSuite) TestGetTickerByIDUnconverted() {
	ticker, err := suite.paprikaClient.Tickers.GetByIDUnconverted("btc-bitcoin")
	suite.NoError(err)
	suite.NotEmpty(ticker)

	suite.Equal("Bitcoin", ticker.Name)
	suite.Equal("BTC", ticker.Symbol)
	suite.Equal("btc-bitcoin", ticker.ID)
}

func (suite *TickerTestSuite) TestGetTickerByID() {
	ticker, err := suite.paprikaClient.Tickers.GetByID("btc-bitcoin")
	suite.NoError(err)
	suite.NotEmpty(ticker)

	suite.Equal("Bitcoin", ticker.Name)
	suite.Equal("BTC", ticker.Symbol)
	suite.Equal("btc-bitcoin", ticker.ID)

	suite.NotNil(ticker.PriceUSD)
	suite.NotZero(ticker.PriceUSD)
}

func TestTickerTestSuite(t *testing.T) {
	suite.Run(t, new(TickerTestSuite))
}
