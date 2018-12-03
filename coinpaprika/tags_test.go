package coinpaprika

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type TagsTestSuite struct {
	suite.Suite
	paprikaClient *Client
}

func (suite *TagsTestSuite) SetupTest() {
	paprikaClient := NewClient(nil)
	suite.NotNil(paprikaClient)

	suite.paprikaClient = paprikaClient
}

func (suite *TagsTestSuite) TestList() {
	tags, err := suite.paprikaClient.Tags.List(nil)
	suite.NoError(err)
	suite.NotEmpty(tags)

	suite.Empty(tags[0].Coins)
}

func (suite *TagsTestSuite) TestListWithCoins() {
	tags, err := suite.paprikaClient.Tags.List(&TagsOptions{AdditionalFields: "coins"})
	suite.NoError(err)
	suite.NotEmpty(tags)

	suite.NotEmpty(tags[0].Coins)
}

func (suite *TagsTestSuite) TestGet() {
	tag, err := suite.paprikaClient.Tags.GetByID("cryptocurrency", nil)
	suite.NoError(err)
	suite.NotNil(tag)
	suite.Empty(tag.Coins)
}

func (suite *TagsTestSuite) TestGetWithCoins() {
	tag, err := suite.paprikaClient.Tags.GetByID("cryptocurrency", &TagsOptions{AdditionalFields: "coins"})
	suite.NoError(err)
	suite.NotNil(tag)
	suite.NotEmpty(tag.Coins)
}

func TestTagsTestSuite(t *testing.T) {
	suite.Run(t, new(TagsTestSuite))
}
