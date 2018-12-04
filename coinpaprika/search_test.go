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
	paprikaClient := NewClient(nil)
	suite.NotNil(paprikaClient)

	suite.paprikaClient = paprikaClient
}

func (suite *SearchTestSuite) TestSearch() {
	options := &SearchOptions{Query: "a"}
	searchResult, err := suite.paprikaClient.Search.Search(options)
	suite.NoError(err)

	suite.NotNil(searchResult.Currencies)
	suite.NotNil(searchResult.Exchanges)
	suite.NotNil(searchResult.ICOS)
	suite.NotNil(searchResult.People)
	suite.NotNil(searchResult.Tags)
}

func (suite *SearchTestSuite) TestSearchLimit() {
	options := &SearchOptions{Query: "a", Limit: 1}
	searchResult, err := suite.paprikaClient.Search.Search(options)
	suite.NoError(err)

	suite.Len(searchResult.Currencies, 1)
	suite.Len(searchResult.Exchanges, 1)
	suite.Len(searchResult.ICOS, 1)
	suite.Len(searchResult.People, 1)
	suite.Len(searchResult.Tags, 1)
}

func (suite *SearchTestSuite) TestSearchCategories() {
	options := &SearchOptions{Query: "a", Categories: "currencies,exchanges"}
	searchResult, err := suite.paprikaClient.Search.Search(options)
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
