package coinpaprika

import (
	"net/http"
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

// TestGetByPlatformRejectsPathEscapingIDs proves a platformID that could
// retarget the endpoint or inject a query is rejected before any request is
// built.
func (suite *ContractsTestSuite) TestGetByPlatformRejectsPathEscapingIDs() {
	rt := &recordingRoundTripper{}
	client := NewClient(&http.Client{Transport: rt})

	badIDs := []string{"a/b", "..", "../key/info", "eth-ethereum?x=1", "eth-ethereum#f"}
	for _, id := range badIDs {
		_, err := client.Contracts.GetByPlatform(id)
		suite.Error(err, "id %q should be rejected before a request is built", id)
	}
	suite.Empty(rt.requests, "no request should reach the transport for a rejected id")
}

// TestGetTickerByContractAddressRejectsPathEscapingIDs proves both the
// platformID and the contractAddress are validated, since either one is
// formatted unescaped into the path.
func (suite *ContractsTestSuite) TestGetTickerByContractAddressRejectsPathEscapingIDs() {
	rt := &recordingRoundTripper{}
	client := NewClient(&http.Client{Transport: rt})

	_, err := client.Contracts.GetTickerByContractAddress("a/b", "0xdac17f958d2ee523a2206206994597c13d831ec7", nil)
	suite.Error(err, "a bad platformID should be rejected before a request is built")

	_, err = client.Contracts.GetTickerByContractAddress("eth-ethereum", "..", nil)
	suite.Error(err, "a bad contractAddress should be rejected before a request is built")

	suite.Empty(rt.requests, "no request should reach the transport for a rejected id")
}

// TestGetByPlatformEncodesSafeID proves a normal platformID still produces
// exactly one encoded path segment under /contracts, unchanged from before
// this fix.
func (suite *ContractsTestSuite) TestGetByPlatformEncodesSafeID() {
	rt := &recordingRoundTripper{}
	client := NewClient(&http.Client{Transport: rt})

	_, _ = client.Contracts.GetByPlatform("eth-ethereum")
	suite.Require().Len(rt.requests, 1)
	suite.Equal("/v1/contracts/eth-ethereum", rt.requests[0].URL.Path)
}

func TestContractsTestSuite(t *testing.T) {
	suite.Run(t, new(ContractsTestSuite))
}
