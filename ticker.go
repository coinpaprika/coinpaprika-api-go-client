package coinpaprika

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/pkg/errors"
)

type TickerService service

const conversionErrFMT = "conversion error in coin: %s, field: %s"

// CoinTickerUnconverted represents ticker information in original format returned by coinpaprika API.
type CoinTickerUnconverted struct {
	ID                string `json:"id"`
	Name              string `json:"name"`
	Symbol            string `json:"symbol"`
	Rank              string `json:"rank"`
	PriceUSD          string `json:"price_usd"`
	PriceBTC          string `json:"price_btc"`
	Volume24hUSD      string `json:"volume_24h_usd"`
	MarketCapUSD      string `json:"market_cap_usd"`
	CirculatingSupply string `json:"circulating_supply"`
	TotalSupply       string `json:"total_supply"`
	MaxSupply         string `json:"max_supply"`
	PercentChange1h   string `json:"percent_change_1h"`
	PercentChange24h  string `json:"percent_change_24h"`
	PercentChange7d   string `json:"percent_change_7d"`
	LastUpdated       string `json:"last_updated"`
}

func convertStrToFloatPtr(value string) (converted *float64, err error) {
	if value != "" {
		f, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return nil, err
		}
		converted = &f
	}
	return converted, nil
}

func convertStrToIntPtr(value string) (converted *int64, err error) {
	if value != "" {
		f, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return nil, err
		}
		converted = &f
	}
	return converted, nil
}

func (u *CoinTickerUnconverted) convert() (*CoinTicker, error) {
	rank, err := strconv.ParseInt(u.Rank, 10, 64)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(conversionErrFMT, u.Name, "rank"))
	}

	priceUSD, err := convertStrToFloatPtr(u.PriceUSD)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(conversionErrFMT, u.Name, "price_usd"))
	}

	priceBTC, err := convertStrToFloatPtr(u.PriceBTC)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(conversionErrFMT, u.Name, "price_btc"))
	}

	volume24hUSD, err := convertStrToIntPtr(u.Volume24hUSD)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(conversionErrFMT, u.Name, "volume_24h_usd"))
	}

	marketCapUSD, err := convertStrToIntPtr(u.MarketCapUSD)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(conversionErrFMT, u.Name, "market_cap_usd"))
	}

	circulatingSupply, err := convertStrToIntPtr(u.CirculatingSupply)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(conversionErrFMT, u.Name, "circulating_supply"))
	}

	totalSupply, err := convertStrToIntPtr(u.TotalSupply)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(conversionErrFMT, u.Name, "total_supply"))
	}

	maxSupply, err := convertStrToIntPtr(u.MaxSupply)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(conversionErrFMT, u.Name, "max_supply"))
	}

	percentChange1h, err := convertStrToFloatPtr(u.PercentChange1h)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(conversionErrFMT, u.Name, "percent_change_1h"))
	}

	percentChange24h, err := convertStrToFloatPtr(u.PercentChange24h)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(conversionErrFMT, u.Name, "percent_change_24h"))
	}

	percentChange7d, err := convertStrToFloatPtr(u.PercentChange7d)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(conversionErrFMT, u.Name, "percent_change_7d"))
	}

	lastUpdated, err := strconv.ParseInt(u.LastUpdated, 10, 64)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(conversionErrFMT, u.Name, "last_updated"))
	}

	ct := CoinTicker{
		ID:                u.ID,
		Name:              u.Name,
		Symbol:            u.Symbol,
		Rank:              rank,
		PriceUSD:          priceUSD,
		PriceBTC:          priceBTC,
		Volume24hUSD:      volume24hUSD,
		MarketCapUSD:      marketCapUSD,
		CirculatingSupply: circulatingSupply,
		TotalSupply:       totalSupply,
		MaxSupply:         maxSupply,
		PercentChange1h:   percentChange1h,
		PercentChange24h:  percentChange24h,
		PercentChange7d:   percentChange7d,
		LastUpdated:       time.Unix(lastUpdated, 0),
	}

	return &ct, nil
}

// CoinTicker represents ticker information parsed to appropriate types.
type CoinTicker struct {
	ID                string    `json:"id"`
	Name              string    `json:"name"`
	Symbol            string    `json:"symbol"`
	Rank              int64     `json:"rank"`
	PriceUSD          *float64  `json:"price_usd"`
	PriceBTC          *float64  `json:"price_btc"`
	Volume24hUSD      *int64    `json:"volume_24h_usd"`
	MarketCapUSD      *int64    `json:"market_cap_usd"`
	CirculatingSupply *int64    `json:"circulating_supply"`
	TotalSupply       *int64    `json:"total_supply"`
	MaxSupply         *int64    `json:"max_supply"`
	PercentChange1h   *float64  `json:"percent_change_1h"`
	PercentChange24h  *float64  `json:"percent_change_24h"`
	PercentChange7d   *float64  `json:"percent_change_7d"`
	LastUpdated       time.Time `json:"last_updated"`
}

// GetTickersUnconverted gets ticker information for all coins listed on coinpaprika. Returned data is in original string format.
func (s *TickerService) GetTickersUnconverted() (tickersUnconverted []*CoinTickerUnconverted, err error) {
	url := fmt.Sprintf("%s/ticker", baseURL)

	body, err := sendGET(s.httpClient, url)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(body, &tickersUnconverted); err != nil {
		return tickersUnconverted, err
	}

	return tickersUnconverted, err
}

// GetTickers gets ticker information for all coins listed on coinpaprika. Returned data is automatically parsed.
func (s *TickerService) GetTickers() (coinTickers []*CoinTicker, err error) {
	tickersUnconverted, err := s.GetTickersUnconverted()
	if err != nil {
		return nil, err
	}

	for _, ticker := range tickersUnconverted {
		ct, err := ticker.convert()
		if err != nil {
			return coinTickers, err
		}
		coinTickers = append(coinTickers, ct)
	}

	return coinTickers, nil
}

// GetTickerByIDUnconverted gets ticker information for specific coin by id (eg. btc-bitcoin). Returned data is in original string format.
func (s *TickerService) GetTickerByIDUnconverted(id string) (tickerUnconverted *CoinTickerUnconverted, err error) {
	url := fmt.Sprintf("%s/ticker/%s", baseURL, id)

	body, err := sendGET(s.httpClient, url)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(body, &tickerUnconverted); err != nil {
		return tickerUnconverted, err
	}

	return tickerUnconverted, err
}

// GetTickerByID gets ticker information for specific coin by id (eg. btc-bitcoin). Returned data is automatically parsed.
func (s *TickerService) GetTickerByID(id string) (coinTickers *CoinTicker, err error) {
	tickerUnconverted, err := s.GetTickerByIDUnconverted(id)
	if err != nil {
		return nil, err
	}

	return tickerUnconverted.convert()

}
