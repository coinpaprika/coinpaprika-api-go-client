package coinpaprika

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
)

type KeyTestSuite struct {
	suite.Suite
	paprikaClient *Client
}

func (suite *KeyTestSuite) SetupTest() {
	paprikaClient := NewClient(nil)
	suite.NotNil(paprikaClient)
	suite.paprikaClient = paprikaClient
}

func (suite *KeyTestSuite) TestGetInfo() {
	_, err := suite.paprikaClient.Key.GetInfo()
	if err != nil {
		// Expected to fail without a Pro+ API key
		suite.True(strings.Contains(err.Error(), "status code: 4"))
	}
}

func TestKeyTestSuite(t *testing.T) {
	suite.Run(t, new(KeyTestSuite))
}
