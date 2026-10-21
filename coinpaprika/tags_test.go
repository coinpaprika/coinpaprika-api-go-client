package coinpaprika

import (
	"net/http"
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

func (suite *TagsTestSuite) TestGetWithICOs() {
	tag, err := suite.paprikaClient.Tags.GetByID("cryptocurrency", &TagsOptions{AdditionalFields: "icos"})
	suite.NoError(err)
	suite.NotNil(tag)
	suite.NotEmpty(tag.ICOs)
}

// TestGetByIDRejectsPathEscapingIDs proves a tagID that could retarget the
// endpoint or inject a query is rejected before any request is built.
func (suite *TagsTestSuite) TestGetByIDRejectsPathEscapingIDs() {
	rt := &recordingRoundTripper{}
	client := NewClient(&http.Client{Transport: rt})

	badIDs := []string{"a/b", "..", "../key/info", "cryptocurrency?x=1", "cryptocurrency#f"}
	for _, id := range badIDs {
		_, err := client.Tags.GetByID(id, nil)
		suite.Error(err, "id %q should be rejected before a request is built", id)
	}
	suite.Empty(rt.requests, "no request should reach the transport for a rejected id")
}

// TestGetByIDEncodesSafeID proves a normal tagID still produces exactly one
// encoded path segment under /tags, unchanged from before this fix.
func (suite *TagsTestSuite) TestGetByIDEncodesSafeID() {
	rt := &recordingRoundTripper{}
	client := NewClient(&http.Client{Transport: rt})

	_, _ = client.Tags.GetByID("cryptocurrency", nil)
	suite.Require().Len(rt.requests, 1)
	suite.Equal("/v1/tags/cryptocurrency", rt.requests[0].URL.Path)
}

func TestTagsTestSuite(t *testing.T) {
	suite.Run(t, new(TagsTestSuite))
}
