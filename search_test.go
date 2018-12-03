package coinpaprika

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type SearchTestSuite struct {
	suite.Suite
	paprikaClient *Client
}

func (suite *SearchTestSuite) SetupTest() {
	paprikaClient, err := NewClient()
	suite.NoError(err)
	suite.NotNil(paprikaClient)

	suite.paprikaClient = paprikaClient
}

func (suite *SearchTestSuite) TestSearch() {
	searchResult, err := suite.paprikaClient.Search.Search("a", nil)
	suite.NoError(err)

	suite.NotNil(searchResult.Currencies)
	suite.NotNil(searchResult.Exchanges)
	suite.NotNil(searchResult.ICOS)
	suite.NotNil(searchResult.People)
	suite.NotNil(searchResult.Tags)
}

func (suite *SearchTestSuite) TestSearchLimit() {
	searchResult, err := suite.paprikaClient.Search.Search("a", &SearchOptions{Limit: 1})
	suite.NoError(err)

	suite.Len(searchResult.Currencies, 1)
	suite.Len(searchResult.Exchanges, 1)
	suite.Len(searchResult.ICOS, 1)
	suite.Len(searchResult.People, 1)
	suite.Len(searchResult.Tags, 1)
}

func (suite *SearchTestSuite) TestSearchCategories() {
	searchResult, err := suite.paprikaClient.Search.Search("a", &SearchOptions{Categories: "currencies,exchanges"})
	suite.NoError(err)

	suite.NotNil(searchResult.Currencies)
	suite.NotNil(searchResult.Exchanges)
	suite.Nil(searchResult.ICOS)
	suite.Nil(searchResult.People)
	suite.Nil(searchResult.Tags)
}

func TestSearchTestSuite(t *testing.T) {
	suite.Run(t, new(SearchTestSuite))
}
