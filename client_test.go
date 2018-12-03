package coinpaprika

import (
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type ClientTestSuite struct {
	suite.Suite
	paprikaClient *Client
}

func (suite *ClientTestSuite) SetupTest() {
	paprikaClient := NewClient(nil)
	suite.NotNil(paprikaClient)

	suite.paprikaClient = paprikaClient
}

func (suite *ClientTestSuite) TestNewClientDefault() {
	client := NewClient(nil)
	suite.NotNil(client)
	suite.Equal(http.DefaultClient, client.httpClient)
}

func (suite *ClientTestSuite) TestNewClientSetHTTPClient() {
	customHTTPClient := &http.Client{Timeout: time.Second * 10}

	client := NewClient(customHTTPClient)
	suite.NotNil(client)
	suite.Equal(customHTTPClient, client.httpClient)
}

func TestClientTestSuite(t *testing.T) {
	suite.Run(t, new(ClientTestSuite))
}
