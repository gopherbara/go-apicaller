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

type StocksHandler struct {
	manager  *socket.Manager
	services *service.Service
	settings map[string]interface{}
	Compared ComparedStocksObject
}

type ComparedStocksObject struct {
	ApiName      string                  `json:"api_name,omitempty"`
	PrevPeriod   string                  `json:"prev_period,omitempty"`
	LastUpdated  string                  `json:"last_updated"`
	StocksChange map[string]GrowthObject `json:"stocks_change,omitempty"`
}

func NewStocksHandler(manager *socket.Manager, services *service.Service, settings map[string]interface{}) *StocksHandler {
	return &StocksHandler{
		manager:  manager,
		services: services,
		settings: settings,
		Compared: ComparedStocksObject{},
	}
}

func (s *StocksHandler) Handle() {
	// another stock apis can be chosed
	sa := apihandler.NewAlphaVantageApi("AlphaVantage", s.settings)

	timeCall := time.Now().Format("2006-01-02 15:04")
	prevDate := time.Now().AddDate(0, 0, sa.Settings.General.PreviousPeriod).Format("2006-01-02")

	prevStocks, _ := s.services.GetStocksByDate(prevDate)
	sa.ProcessApiCall()
	nowStocks := sa.Response.ToBase(sa.Settings.Name, timeCall)
	if sa.Settings.General.IsProd {
		s.services.SaveStocks(nowStocks)
	}

	s.Compared = CompareStocks(prevStocks, nowStocks)
	s.Send()
}

func (s *StocksHandler) Send() {
	str, _ := json.Marshal(s.Compared)
	msg := socket.Message{
		ApiType: "stocks",
		Body:    string(str),
	}
	s.manager.UpdateMessage <- msg
}

func CompareStocks(prev baseobjects.StocksObject, now baseobjects.StocksObject) ComparedStocksObject {

	stocksChange := make(map[string]GrowthObject)
	for stockNow, priceNow := range now.StockPrice {
		stock := GrowthObject{
			ObjectChange: stockNow,
			TodayValue:   utils.RoundFloat(priceNow),
		}
		if pricePrev, ok := prev.StockPrice[stockNow]; ok {
			stock.PreviousValue = utils.RoundFloat(pricePrev)
		} else {
			stock.PreviousValue = 0
		}
		stock.Growth = utils.RoundFloat((stock.TodayValue - stock.PreviousValue) / stock.TodayValue * 100)
		stocksChange[stockNow] = stock
	}

	return ComparedStocksObject{
		ApiName:      now.ApiName,
		PrevPeriod:   prev.LastUpdated,
		LastUpdated:  now.LastUpdated,
		StocksChange: stocksChange,
	}
}
