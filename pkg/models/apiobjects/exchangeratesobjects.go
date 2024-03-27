package apiobjects

import "github.com/gopherbara/go-apicaller/pkg/models/baseobjects"

type ExchangeRatesSettings struct {
	General      GeneralSettings
	Name         string
	Description  string
	Url          string
	FromCurrency string
	Currencies   []string
	APIKey       string
}
type ExchangeRatesResponse struct {
	FromCurrency string             `json:"base"`
	Date         string             `json:"date"`
	Rates        map[string]float64 `json:"rates"`
	Success      bool               `json:"success"`
	Timestamp    int                `json:"timestamp"`
}

func (er ExchangeRatesResponse) ToBase(name string, date string) baseobjects.CurrencyObject {
	return baseobjects.CurrencyObject{
		ApiName:      name,
		DateCall:     date,
		FromCurrency: er.FromCurrency,
		LastUpdated:  er.Date,
		Rates:        er.Rates,
	}
}
