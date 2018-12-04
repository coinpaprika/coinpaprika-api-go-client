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
	suite.NoError(err)
	suite.NotEmpty(person)
}

func TestPeopleTestSuite(t *testing.T) {
	suite.Run(t, new(PeopleTestSuite))
}
