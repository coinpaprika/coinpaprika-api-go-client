package coinpaprika

import (
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type ClientTestSuite struct {
	suite.Suite
	paprikaClient *Client
}

func (suite *ClientTestSuite) SetupTest() {
	paprikaClient, err := NewClient()
	suite.NoError(err)
	suite.NotNil(paprikaClient)

	suite.paprikaClient = paprikaClient
}

func (suite *ClientTestSuite) TestNewClientDefault() {
	client, err := NewClient()
	suite.NoError(err)
	suite.NotNil(client)
	suite.Equal(http.DefaultClient, client.httpClient)
}

func (suite *ClientTestSuite) TestNewClientSetHTTPClient() {
	customHTTPClient := &http.Client{Timeout: time.Second * 10}

	client, err := NewClient(SetHTTPClient(customHTTPClient))
	suite.NoError(err)
	suite.NotNil(client)
	suite.Equal(customHTTPClient, client.httpClient)
}

func (suite *ClientTestSuite) TestGetGlobalStats() {
	globalStats, err := suite.paprikaClient.GetGlobalStats()
	suite.NoError(err)

	suite.NotZero(globalStats.MarketCapUSD)
	suite.NotZero(globalStats.Volume24hUSD)
	suite.NotZero(globalStats.LastUpdated)
	suite.NotZero(globalStats.BitcoinDominancePercentage)
	suite.NotZero(globalStats.CryptocurrenciesNumber)
}

func (suite *ClientTestSuite) TestGetTickers() {
	tickers, err := suite.paprikaClient.GetTickers()
	suite.NoError(err)
	suite.NotEmpty(tickers)
}

func (suite *ClientTestSuite) TestGetTickersUnconverted() {
	tickers, err := suite.paprikaClient.GetTickersUnconverted()
	suite.NoError(err)
	suite.NotEmpty(tickers)
}

func (suite *ClientTestSuite) TestGetTickerByIDUnconverted() {
	ticker, err := suite.paprikaClient.GetTickerByIDUnconverted("btc-bitcoin")
	suite.NoError(err)
	suite.NotEmpty(ticker)

	suite.Equal("Bitcoin", ticker.Name)
	suite.Equal("BTC", ticker.Symbol)
	suite.Equal("btc-bitcoin", ticker.ID)
}

func (suite *ClientTestSuite) TestGetTickerByID() {
	ticker, err := suite.paprikaClient.GetTickerByID("btc-bitcoin")
	suite.NoError(err)
	suite.NotEmpty(ticker)

	suite.Equal("Bitcoin", ticker.Name)
	suite.Equal("BTC", ticker.Symbol)
	suite.Equal("btc-bitcoin", ticker.ID)

	suite.NotNil(ticker.PriceUSD)
	suite.NotZero(ticker.PriceUSD)
}

func (suite *ClientTestSuite) TestGetCoins() {
	coins, err := suite.paprikaClient.GetCoins()
	suite.NoError(err)
	suite.NotEmpty(coins)
}

func TestClientTestSuite(t *testing.T) {
	suite.Run(t, new(ClientTestSuite))
}
