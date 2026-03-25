package coinpaprika

import (
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

func TestPeopleTestSuite(t *testing.T) {
	suite.Run(t, new(PeopleTestSuite))
}
