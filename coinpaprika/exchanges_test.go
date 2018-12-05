package coinpaprika

import (
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

func TestExchangesTestSuite(t *testing.T) {
	suite.Run(t, new(ExchangesTestSuite))
}
