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

type CryptoHandler struct {
	manager  *socket.Manager
	services *service.Service
	settings map[string]interface{}
	Compared ComparedCryptoObject
}

type ComparedCryptoObject struct {
	ApiName      string                  `json:"api_name,omitempty"`
	PrevPeriod   string                  `json:"prev_period,omitempty"`
	LastUpdated  string                  `json:"last_updated"`
	CryptoChange map[string]GrowthObject `json:"crypto_change,omitempty"`
}

func NewCryptoHandler(manager *socket.Manager, services *service.Service, settings map[string]interface{}) *CryptoHandler {
	return &CryptoHandler{
		manager:  manager,
		services: services,
		settings: settings,
		Compared: ComparedCryptoObject{},
	}
}

func (s *CryptoHandler) Handle() {
	sa := apihandler.NewCoinApi("CoinApi", s.settings)

	timeCall := time.Now().Format("2006-01-02 15:04")
	prevDate := time.Now().AddDate(0, 0, sa.Settings.General.PreviousPeriod).Format("2006-01-02")

	prevCrypto, _ := s.services.GetCryptoByDate(prevDate)
	sa.ProcessApiCall()
	nowCrypto := sa.Response.ToBase(sa.Settings.Name, timeCall)
	if sa.Settings.General.IsProd {
		s.services.SaveCrypto(nowCrypto)
	}

	s.Compared = CompareCrypto(prevCrypto, nowCrypto)
	s.Send()
}

func (s *CryptoHandler) Send() {
	str, _ := json.Marshal(s.Compared)
	msg := socket.Message{
		ApiType: "crypto",
		Body:    string(str),
	}
	s.manager.UpdateMessage <- msg
}

func CompareCrypto(prev baseobjects.CryptoObject, now baseobjects.CryptoObject) ComparedCryptoObject {

	cryptoChange := make(map[string]GrowthObject)
	for cryptoNow, priceNow := range now.CryptoPrice {
		crypto := GrowthObject{
			ObjectChange: cryptoNow,
			TodayValue:   utils.RoundFloat(priceNow),
		}
		if pricePrev, ok := prev.CryptoPrice[cryptoNow]; ok {
			crypto.PreviousValue = utils.RoundFloat(pricePrev)
		} else {
			crypto.PreviousValue = 0
		}
		crypto.Growth = utils.RoundFloat((crypto.TodayValue - crypto.PreviousValue) / crypto.TodayValue * 100)
		cryptoChange[cryptoNow] = crypto
	}

	return ComparedCryptoObject{
		ApiName:      now.ApiName,
		PrevPeriod:   prev.LastUpdated,
		LastUpdated:  now.LastUpdated,
		CryptoChange: cryptoChange,
	}
}
