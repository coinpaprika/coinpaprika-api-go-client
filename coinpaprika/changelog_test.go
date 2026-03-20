package coinpaprika

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
)

type ChangelogTestSuite struct {
	suite.Suite
	paprikaClient *Client
}

func (suite *ChangelogTestSuite) SetupTest() {
	paprikaClient := NewClient(nil)
	suite.NotNil(paprikaClient)
	suite.paprikaClient = paprikaClient
}

func (suite *ChangelogTestSuite) TestGetIDChanges() {
	_, err := suite.paprikaClient.Changelog.GetIDChanges(nil)
	if err != nil {
		// Expected to fail without a Starter+ API key
		suite.True(strings.Contains(err.Error(), "status code: 4"))
	}
}

func (suite *ChangelogTestSuite) TestGetIDChangesWithOptions() {
	options := &ChangelogOptions{Page: 1, Limit: 10}
	_, err := suite.paprikaClient.Changelog.GetIDChanges(options)
	if err != nil {
		// Expected to fail without a Starter+ API key
		suite.True(strings.Contains(err.Error(), "status code: 4"))
	}
}

func TestChangelogTestSuite(t *testing.T) {
	suite.Run(t, new(ChangelogTestSuite))
}
