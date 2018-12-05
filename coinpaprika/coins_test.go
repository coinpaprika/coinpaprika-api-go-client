package coinpaprika

import (
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
	suite.NoError(err)
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

func (suite *CoinsTestSuite) TestGetHistoricalOHLCVByCoinID() {
	options := &HistoricalOHLCVOptions{
		Start: time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC),
		End:   time.Date(2018, 1, 11, 0, 0, 0, 0, time.UTC),
		Quote: "btc",
	}
	entries, err := suite.paprikaClient.Coins.GetHistoricalOHLCVByCoinID("btc-bitcoin", options)
	suite.NoError(err)
	suite.NotEmpty(entries)
	suite.Len(entries, 10)
}

func TestCoinsTestSuite(t *testing.T) {
	suite.Run(t, new(CoinsTestSuite))
}
