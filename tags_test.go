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

func (suite *TagsTestSuite) TestListTags() {
	tags, err := suite.paprikaClient.ListTags(nil)
	suite.NoError(err)
	suite.NotEmpty(tags)

	suite.Nil(tags[0].Coins)
}

func (suite *TagsTestSuite) TestListTagsWithCoins() {
	tags, err := suite.paprikaClient.ListTags(&TagsOptions{AdditionalFields: "coins"})
	suite.NoError(err)
	suite.NotEmpty(tags)

	suite.NotNil(tags[0].Coins)
}

func (suite *TagsTestSuite) TestGetTag() {
	tag, err := suite.paprikaClient.GetTag("cryptocurrency", nil)
	suite.NoError(err)
	suite.NotNil(tag)
	suite.Nil(tag.Coins)
}

func (suite *TagsTestSuite) TestGetTagWithCoins() {
	tag, err := suite.paprikaClient.GetTag("cryptocurrency", &TagsOptions{AdditionalFields: "coins"})
	suite.NoError(err)
	suite.NotNil(tag)
	suite.NotNil(tag.Coins)
}

func TestTagsTestSuite(t *testing.T) {
	suite.Run(t, new(TagsTestSuite))
}
