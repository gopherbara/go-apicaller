package repository

import (
	"github.com/vv-projects/go-apicaller/pkg/models/baseobjects"
)

const (
	WeaherTable      = "weather"
	StocksTable      = "stocks"
	QuotesTable      = "quotes"
	CryptoTable      = "crypto"
	CurrencyTable    = "currency"
	GeoLocationTable = "geolocation"
)

// interfaces for all kind of apis which must implement dbs
// all dbs
type Quote interface {
	SaveQuote(quote baseobjects.QuotesObject) error
	GetQuoteByDate(date string) ([]baseobjects.QuotesObject, error)
}

type GeoLocation interface {
	SaveGeoLocation(geo baseobjects.GeoLocationObject) error
	GetGeoLocationByDate(date string) ([]baseobjects.GeoLocationObject, error)
}

type Weather interface {
	SaveWeather(weather baseobjects.WeatherObject) error
	GetWeatherByDate(date string) ([]baseobjects.WeatherObject, error)
}

type CurrencyRate interface {
	SaveCurrencyRate(currency baseobjects.CurrencyObject) error
	GetCurrencyRateByDate(date string) ([]baseobjects.CurrencyObject, error)
	GetCurrencyRateByDateApi(date string, apiName string) ([]baseobjects.CurrencyObject, error)
}

type Stocks interface {
	SaveStocks(stocks baseobjects.StocksObject) error
	GetStocksByDate(date string) ([]baseobjects.StocksObject, error)
	GetStocksByDateApi(date string, apiName string) ([]baseobjects.StocksObject, error)
}

type Crypto interface {
	SaveCrypto(crypto baseobjects.CryptoObject) error
	GetCryptoByDate(date string) ([]baseobjects.CryptoObject, error)
	GetCryptoByDateApi(date string, apiName string) ([]baseobjects.CryptoObject, error)
}

type Repository struct {
	Quote
	GeoLocation
	Weather
	CurrencyRate
	Stocks
	Crypto
}
