package handler

import (
	"encoding/json"
	"github.com/vv-projects/go-apicaller/pkg/apihandler"
	"github.com/vv-projects/go-apicaller/pkg/models/baseobjects"
	"github.com/vv-projects/go-apicaller/pkg/service"
	"github.com/vv-projects/go-apicaller/pkg/socket"
	"github.com/vv-projects/go-apicaller/pkg/utils"
	"time"
)

type CurrencyHandler struct {
	manager  *socket.Manager
	services *service.Service
	settings map[string]interface{}
	Compared ComparedCurrencyObject
}

type ComparedCurrencyObject struct {
	ApiName       string                  `json:"api_name,omitempty"`
	FromCurrency  string                  `json:"from_currency,omitempty"`
	PrevPeriod    string                  `json:"prev_period,omitempty"`
	LastUpdated   string                  `json:"last_updated"`
	CurrencyRates map[string]GrowthObject `json:"currency_rates,omitempty"`
}

func NewCurrencyHandler(manager *socket.Manager, services *service.Service, settings map[string]interface{}) *CurrencyHandler {
	return &CurrencyHandler{
		manager:  manager,
		services: services,
		settings: settings,
		Compared: ComparedCurrencyObject{},
	}
}

func (c *CurrencyHandler) Handle() {
	era := apihandler.NewCurrencyRateApi("ExchangeRates", c.settings)

	timeCall := time.Now().Format("2006-01-02 15:04")
	prevDate := time.Now().AddDate(0, 0, era.Settings.General.PreviousPeriod).Format("2006-01-02")

	prevCurrency, _ := c.services.GetCurrencyRateByDate(prevDate)
	era.ProcessApiCall()
	nowCurrency := era.Response.ToBase(era.Settings.Name, timeCall)
	if era.Settings.General.IsProd {
		c.services.SaveCurrencyRate(nowCurrency)
	}

	c.Compared = CompareCurrency(prevCurrency, nowCurrency)
	c.Send()
}

func (c *CurrencyHandler) Send() {
	str, _ := json.Marshal(c.Compared)
	msg := socket.Message{
		ApiType: "currency",
		Body:    string(str),
	}
	c.manager.UpdateMessage <- msg
}

func CompareCurrency(prev baseobjects.CurrencyObject, now baseobjects.CurrencyObject) ComparedCurrencyObject {
	currencyRates := make(map[string]GrowthObject)
	for currencyNow, rateNow := range now.Rates {
		if prev.FromCurrency == "UAH" {
			rateNow = utils.RoundFloat(1.0 / rateNow)
		}
		currencyRate := GrowthObject{ObjectChange: currencyNow, TodayValue: rateNow}
		if ratePrev, ok := prev.Rates[currencyNow]; ok {
			if prev.FromCurrency == "UAH" {
				ratePrev = utils.RoundFloat(1.0 / ratePrev)
			}
			currencyRate.PreviousValue = ratePrev
		} else {
			currencyRate.PreviousValue = 0
		}
		currencyRate.Growth = utils.RoundFloat((currencyRate.TodayValue - currencyRate.PreviousValue) / currencyRate.TodayValue * 100)

		currencyRates[currencyNow] = currencyRate
	}
	return ComparedCurrencyObject{
		ApiName:       now.ApiName,
		FromCurrency:  now.FromCurrency,
		PrevPeriod:    prev.LastUpdated,
		LastUpdated:   now.LastUpdated,
		CurrencyRates: currencyRates,
	}
}
