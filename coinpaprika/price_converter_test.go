package coinpaprika

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type PriceConverterTestSuite struct {
	suite.Suite
	paprikaClient *Client
}

func (suite *PriceConverterTestSuite) SetupTest() {
	paprikaClient := NewClient(nil)
	suite.NotNil(paprikaClient)

	suite.paprikaClient = paprikaClient
}

func (suite *PriceConverterTestSuite) TestPriceConverter() {
	ret, err := suite.paprikaClient.PriceConverter.PriceConverter(
		&PriceConverterOptions{
			BaseCurrencyID: "btc-bitcoin", QuoteCurrencyID: "usd-us-dollars", Amount: 1.5,
		},
	)
	suite.NoError(err)

	suite.Equal("btc-bitcoin", *ret.BaseCurrencyID)
	suite.Equal("Bitcoin", *ret.BaseCurrencyName)
	suite.NotZero(ret.BasePriceLastUpdated)
	suite.Equal("usd-us-dollars", *ret.QuoteCurrencyID)
	suite.Equal("US Dollars", *ret.QuoteCurrencyName)
	suite.NotZero(ret.QuotePriceLastUpdated)
	suite.Equal(1.5, *ret.Amount)
	suite.NotZero(ret.Price)
}

func TestPriceConverterTestSuite(t *testing.T) {
	suite.Run(t, new(PriceConverterTestSuite))
}
