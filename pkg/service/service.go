package service

import (
	"github.com/vv-projects/go-apicaller/pkg/models/baseobjects"
	"github.com/vv-projects/go-apicaller/pkg/repository"
)

//type Weather interface {
//	CreateWeatherItem(weather baseobjects.WeatherObject) error
//	GetWeatherByStrDate(date string) (baseobjects.WeatherObject, error)
//	//GetWeatherByDate(date time.Time) (baseobjects.WeatherObject, error)
//}

type ObjectType interface {
	baseobjects.WeatherObject | baseobjects.QuotesObject
}

type Quotes interface {
	SaveQuote(quote baseobjects.QuotesObject) bool
	GetQuoteByDate(date string) (baseobjects.QuotesObject, bool)
}

type GeoLocation interface {
	SaveGeoLocation(geo baseobjects.GeoLocationObject) bool
	GetGeoLocationByDate(date string) (baseobjects.GeoLocationObject, bool)
}

type Weather interface {
	SaveWeather(weather baseobjects.WeatherObject) bool
	GetWeatherByDate(date string) (baseobjects.WeatherObject, bool)
}

type CurrencyRate interface {
	SaveCurrencyRate(currency baseobjects.CurrencyObject) bool
	GetCurrencyRateByDate(date string) (baseobjects.CurrencyObject, bool)
	GetCurrencyRateByDateApi(date string, apiName string) (baseobjects.CurrencyObject, bool)
}

type Stocks interface {
	SaveStocks(stocks baseobjects.StocksObject) bool
	GetStocksByDate(date string) (baseobjects.StocksObject, bool)
	GetStocksByDateApi(date string, apiName string) (baseobjects.StocksObject, bool)
}

type Crypto interface {
	SaveCrypto(crypto baseobjects.CryptoObject) bool
	GetCryptoByDate(date string) (baseobjects.CryptoObject, bool)
	GetCryptoByDateApi(date string, apiName string) (baseobjects.CryptoObject, bool)
}

type Service struct {
	Quotes
	GeoLocation
	Weather
	CurrencyRate
	Stocks
	Crypto
}

func NewService(mongoDB *repository.Repository, redisDB *repository.Repository) *Service {
	return &Service{
		Quotes:       NewQuoteService(mongoDB.Quote, redisDB.Quote),
		GeoLocation:  NewGeoLocationService(mongoDB.GeoLocation, redisDB.GeoLocation),
		Weather:      NewWeatherService(mongoDB.Weather, redisDB.Weather),
		CurrencyRate: NewCurrencyService(mongoDB.CurrencyRate, redisDB.CurrencyRate),
		Stocks:       NewStocksService(mongoDB.Stocks, redisDB.Stocks),
		Crypto:       NewCryptoService(mongoDB.Crypto, redisDB.Crypto),
	}

}
