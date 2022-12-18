package apiobjects

import "github.com/vv-projects/go-apicaller/pkg/models/baseobjects"

type WeatherSettings struct {
	General     GeneralSettings
	Name        string
	Description string
	APIKey      string
	APIHost     string
	Url         string
	City        string
}

type OpenWeatherResponse struct {
	Location WeatherLocationResponse `json:"location"`
	Current  WeatherCurrentResponse  `json:"current"`
}

type WeatherLocationResponse struct {
	Name           string  `json:"name"`
	Region         string  `json:"region"`
	Country        string  `json:"country"`
	Lat            float64 `json:"lat"`
	Lon            float64 `json:"lon"`
	TzId           string  `json:"tz_id"`
	LocaltimeEpoch int     `json:"localtime_epoch"`
	Localtime      string  `json:"localtime"`
}

type WeatherCurrentResponse struct {
	LastUpdatedEpoch int                      `json:"last_updated_epoch"`
	LastUpdated      string                   `json:"last_updated"`
	TempC            float64                  `json:"temp_c"`
	TempF            float64                  `json:"temp_f"`
	IsDay            int                      `json:"is_day"`
	Condition        WeatherConditionResponse `json:"condition"`
	WindMph          float64                  `json:"wind_mph"`
	WindKph          float64                  `json:"wind_kph"`
	WindDegree       int                      `json:"wind_degree"`
	WindDir          string                   `json:"wind_dir"`
	PressureMb       float64                  `json:"pressure_mb"`
	PressureIn       float64                  `json:"pressure_in"`
	PrecipMm         float64                  `json:"precip_mm"`
	PrecipIn         float64                  `json:"precip_in"`
	Humidity         int                      `json:"humidity"`
	Cloud            int                      `json:"cloud"`
	FeelslikeC       float64                  `json:"feelslike_c"`
	FeelslikeF       float64                  `json:"feelslike_f"`
	VisKm            float64                  `json:"vis_km"`
	VisMiles         float64                  `json:"vis_miles"`
	Uv               float64                  `json:"uv"`
	GustMph          float64                  `json:"gust_mph"`
	GustKph          float64                  `json:"gust_kph"`
}

type WeatherConditionResponse struct {
	Text string `json:"text"`
	Icon string `json:"icon"`
	Code int    `json:"code"`
}

func (owr OpenWeatherResponse) ToBase(name string, date string, city string) baseobjects.WeatherObject {
	return baseobjects.WeatherObject{
		ApiName:     name,
		DateCall:    date,
		LastUpdated: owr.Current.LastUpdated,
		TempC:       owr.Current.TempC,
		TempF:       owr.Current.TempF,
		WindMph:     owr.Current.WindMph,
		WindKph:     owr.Current.WindKph,
		WindDir:     owr.Current.WindDir,
		Cloud:       owr.Current.Cloud,
		Sky:         owr.Current.Condition.Text,
		City:        city,
	}
}
