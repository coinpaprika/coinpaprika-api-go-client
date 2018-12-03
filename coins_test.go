package coinpaprika

import (
	"testing"

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

func (suite *ClientTestSuite) TestListCoins() {
	coins, err := suite.paprikaClient.Coins.List()
	suite.NoError(err)
	suite.NotEmpty(coins)
}

func TestCoinsTestSuite(t *testing.T) {
	suite.Run(t, new(CoinsTestSuite))
}
