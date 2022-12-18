package apiobjects

import (
	"github.com/vv-projects/go-apicaller/pkg/models/baseobjects"
	"strconv"
)

type AlphaVantageSettings struct {
	General     GeneralSettings
	Name        string
	Description string
	Url         string
	APIKey      string
	Interval    string
	Stock       string
}

type AlphaVantageResponse struct {
	MetaData struct {
		Information   string `json:"1. Information"`
		Symbol        string `json:"2. Symbol"`
		LastRefreshed string `json:"3. Last Refreshed"`
		Interval      string `json:"4. Interval"`
		OutputSize    string `json:"5. Output Size"`
		TimeZone      string `json:"6. Time Zone"`
	} `json:"Meta Data"`
	TimeSeriesFiels map[string]TimeSeries `json:"Time Series (5min)"`
}

type TimeSeries struct {
	Open   string `json:"1. open"`
	High   string `json:"2. high"`
	Low    string `json:"3. low"`
	Close  string `json:"4. close"`
	Volume string `json:"5. volume"`
}

func (avr AlphaVantageResponse) ToBase(name string, date string) baseobjects.StocksObject {
	price, _ := strconv.ParseFloat(avr.TimeSeriesFiels[avr.MetaData.LastRefreshed].Close, 64)
	return baseobjects.StocksObject{
		ApiName:     name,
		DateCall:    date,
		LastUpdated: avr.MetaData.LastRefreshed,
		Stock:       avr.MetaData.Symbol,
		Price:       price,
		StockPrice:  map[string]float64{avr.MetaData.Symbol: price},
	}
}
