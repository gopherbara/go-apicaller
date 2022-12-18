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

type WeatherHandler struct {
	manager  *socket.Manager
	services *service.Service
	settings map[string]interface{}
	Compared ComparedWeatherObject
}

type ComparedWeatherObject struct {
	ApiName           string         `json:"api_name,omitempty"`
	City              string         `json:"city,omitempty"`
	PrevPeriod        string         `json:"prev_period,omitempty"`
	LastUpdated       string         `json:"last_updated"`
	WeatherDifference []GrowthObject `json:"weather_difference,omitempty"`
	WindDir           string         `json:"wind_dir"`
	Sky               string         `json:"sky"`
}

func NewWeatherHandler(manager *socket.Manager, services *service.Service, settings map[string]interface{}) *WeatherHandler {
	return &WeatherHandler{
		manager:  manager,
		services: services,
		settings: settings,
		Compared: ComparedWeatherObject{},
	}
}

func (w *WeatherHandler) Handle() {
	wa := apihandler.NewWeatherApi("Weather", w.settings)

	timeCall := time.Now().Format("2006-01-02 15:04")
	prevDate := time.Now().AddDate(0, 0, wa.Settings.General.PreviousPeriod).Format("2006-01-02")

	prevWeather, _ := w.services.GetWeatherByDate(prevDate)
	wa.ProcessApiCall()
	nowWeather := wa.Response.ToBase(wa.Settings.Name, timeCall, wa.Settings.City)
	if wa.Settings.General.IsProd {
		w.services.SaveWeather(nowWeather)
	}

	w.Compared = CompareWeather(prevWeather, nowWeather)
	w.Send()
}

func (w *WeatherHandler) Send() {
	str, _ := json.Marshal(w.Compared)
	msg := socket.Message{
		ApiType: "weather",
		Body:    string(str),
	}
	w.manager.UpdateMessage <- msg
}

func CompareWeather(prev baseobjects.WeatherObject, now baseobjects.WeatherObject) ComparedWeatherObject {

	weatherDifference := make([]GrowthObject, 0, 4)
	tempC := GrowthObject{
		ObjectChange:  "Temp, C",
		TodayValue:    utils.RoundFloat(now.TempC),
		PreviousValue: utils.RoundFloat(prev.TempC),
		Growth:        utils.RoundFloat(now.TempC - prev.TempC), // just -, no need in %
	}
	tempF := GrowthObject{
		ObjectChange:  "Temp,F",
		TodayValue:    utils.RoundFloat(now.TempF),
		PreviousValue: utils.RoundFloat(prev.TempF),
		Growth:        utils.RoundFloat(now.TempF - prev.TempF), // just -, no need in %
	}

	windKph := GrowthObject{
		ObjectChange:  "Wind(kph)",
		TodayValue:    utils.RoundFloat(now.WindKph),
		PreviousValue: utils.RoundFloat(prev.WindKph),
		Growth:        utils.RoundFloat(now.WindKph - prev.WindKph), // just -, no need in %
	}
	windMph := GrowthObject{
		ObjectChange:  "Wind(mph)",
		TodayValue:    utils.RoundFloat(now.WindMph),
		PreviousValue: utils.RoundFloat(prev.WindMph),
		Growth:        utils.RoundFloat(now.WindMph - prev.WindMph), // just -, no need in %
	}

	weatherDifference = append(weatherDifference, tempC)
	weatherDifference = append(weatherDifference, tempF)
	weatherDifference = append(weatherDifference, windKph)
	weatherDifference = append(weatherDifference, windMph)

	return ComparedWeatherObject{
		ApiName:           now.ApiName,
		PrevPeriod:        prev.LastUpdated,
		LastUpdated:       now.LastUpdated,
		WeatherDifference: weatherDifference,
		WindDir:           now.WindDir,
		Sky:               now.Sky,
		City:              now.City,
	}
}
