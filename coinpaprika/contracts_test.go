package coinpaprika

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
)

type ContractsTestSuite struct {
	suite.Suite
	paprikaClient *Client
}

func (suite *ContractsTestSuite) SetupTest() {
	paprikaClient := NewClient(nil)
	suite.NotNil(paprikaClient)
	suite.paprikaClient = paprikaClient
}

func (suite *ContractsTestSuite) TestListPlatforms() {
	platforms, err := suite.paprikaClient.Contracts.ListPlatforms()
	suite.NoError(err)
	suite.NotEmpty(platforms)
}

func (suite *ContractsTestSuite) TestGetByPlatform() {
	contracts, err := suite.paprikaClient.Contracts.GetByPlatform("eth-ethereum")
	suite.NoError(err)
	suite.NotEmpty(contracts)
}

func (suite *ContractsTestSuite) TestGetTickerByContractAddress() {
	ticker, err := suite.paprikaClient.Contracts.GetTickerByContractAddress("eth-ethereum", "0xdac17f958d2ee523a2206206994597c13d831ec7", nil)
	if err != nil {
		suite.True(strings.Contains(err.Error(), "status code: 4"))
		return
	}
	suite.NotNil(ticker)
}

func (suite *ContractsTestSuite) TestGetHistoricalTickerByContractAddress() {
	_, err := suite.paprikaClient.Contracts.GetHistoricalTickerByContractAddress("eth-ethereum", "0xdac17f958d2ee523a2206206994597c13d831ec7", nil)
	if err != nil {
		suite.True(strings.Contains(err.Error(), "status code: 4"))
	}
}

func TestContractsTestSuite(t *testing.T) {
	suite.Run(t, new(ContractsTestSuite))
}
