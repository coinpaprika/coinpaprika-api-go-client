package coinpaprika

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/suite"
)

type PeopleTestSuite struct {
	suite.Suite
	paprikaClient *Client
}

func (suite *PeopleTestSuite) SetupTest() {
	paprikaClient := NewClient(nil)
	suite.NotNil(paprikaClient)

	suite.paprikaClient = paprikaClient
}

func (suite *PeopleTestSuite) TestGetByID() {
	person, err := suite.paprikaClient.People.GetByID("vitalik-buterin")
	if err != nil {
		// The people endpoint may return 404 for some IDs
		suite.Contains(err.Error(), "status code: 404")
		return
	}
	suite.NotEmpty(person)
}

// TestGetByIDRejectsPathEscapingIDs proves a personID that could retarget the
// endpoint or inject a query is rejected before any request is built.
func (suite *PeopleTestSuite) TestGetByIDRejectsPathEscapingIDs() {
	rt := &recordingRoundTripper{}
	client := NewClient(&http.Client{Transport: rt})

	badIDs := []string{"a/b", "..", "../key/info", "vitalik?x=1", "vitalik#f"}
	for _, id := range badIDs {
		_, err := client.People.GetByID(id)
		suite.Error(err, "id %q should be rejected before a request is built", id)
	}
	suite.Empty(rt.requests, "no request should reach the transport for a rejected id")
}

// TestGetByIDEncodesSafeID proves a normal personID still produces exactly
// one encoded path segment under /people, unchanged from before this fix.
func (suite *PeopleTestSuite) TestGetByIDEncodesSafeID() {
	rt := &recordingRoundTripper{}
	client := NewClient(&http.Client{Transport: rt})

	_, _ = client.People.GetByID("vitalik-buterin")
	suite.Require().Len(rt.requests, 1)
	suite.Equal("/v1/people/vitalik-buterin", rt.requests[0].URL.Path)
}

func TestPeopleTestSuite(t *testing.T) {
	suite.Run(t, new(PeopleTestSuite))
}
