package coinpaprika

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type GlobalTestSuite struct {
	suite.Suite
	paprikaClient *Client
}

func (suite *GlobalTestSuite) SetupTest() {
	paprikaClient := NewClient(nil)

	suite.NotNil(paprikaClient)

	suite.paprikaClient = paprikaClient
}

func (suite *GlobalTestSuite) TestGet() {
	globalStats, err := suite.paprikaClient.Global.Get()
	suite.NoError(err)

	suite.NotZero(globalStats.MarketCapUSD)
	suite.NotZero(globalStats.Volume24hUSD)
	suite.NotZero(globalStats.LastUpdated)
	suite.NotZero(globalStats.BitcoinDominancePercentage)
	suite.NotZero(globalStats.CryptocurrenciesNumber)
	suite.NotZero(globalStats.MarketCapATHValue)
	suite.NotZero(globalStats.MarketCapATHDate)
	suite.NotZero(globalStats.Volume24hATHValue)
	suite.NotZero(globalStats.Volume24hATHDate)
	suite.NotZero(globalStats.Volume24hChange24h)
	suite.NotZero(globalStats.MarketCapChange24h)
}

func TestGlobalTestSuite(t *testing.T) {
	suite.Run(t, new(GlobalTestSuite))
}
