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

func (suite *CoinsTestSuite) TestGetTwitterTimelineByID() {
	timeline, err := suite.paprikaClient.Coins.GetTwitterTimelineByID("btc-bitcoin")
	suite.NoError(err)
	suite.NotEmpty(timeline)
}

func TestCoinsTestSuite(t *testing.T) {
	suite.Run(t, new(CoinsTestSuite))
}
