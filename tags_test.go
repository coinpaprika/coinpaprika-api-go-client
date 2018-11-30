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
	paprikaClient, err := NewClient()
	suite.NoError(err)
	suite.NotNil(paprikaClient)

	suite.paprikaClient = paprikaClient
}

func (suite *TagsTestSuite) TestGetTags() {
	tags, err := suite.paprikaClient.GetTags(nil)
	suite.NoError(err)
	suite.NotEmpty(tags)

	suite.Nil(tags[0].Coins)
}

func (suite *TagsTestSuite) TestGetTagsWithCoins() {
	tags, err := suite.paprikaClient.GetTags(&TagsOptions{AdditionalFields: "coins"})
	suite.NoError(err)
	suite.NotEmpty(tags)

	suite.NotNil(tags[0].Coins)
}

func (suite *TagsTestSuite) TestGetTagByID() {
	tag, err := suite.paprikaClient.GetTagByID("cryptocurrency", nil)
	suite.NoError(err)
	suite.NotNil(tag)
	suite.Nil(tag.Coins)
}

func (suite *TagsTestSuite) TestGetTagByIDWithCoins() {
	tag, err := suite.paprikaClient.GetTagByID("cryptocurrency", &TagsOptions{AdditionalFields: "coins"})
	suite.NoError(err)
	suite.NotNil(tag)
	suite.NotNil(tag.Coins)
}

func TestTagsTestSuite(t *testing.T) {
	suite.Run(t, new(TagsTestSuite))
}
