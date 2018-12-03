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
	paprikaClient, err := NewClient()
	suite.NoError(err)
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
}

func TestGlobalTestSuite(t *testing.T) {
	suite.Run(t, new(GlobalTestSuite))
}
