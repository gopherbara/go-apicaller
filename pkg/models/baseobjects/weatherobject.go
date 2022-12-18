package baseobjects

type WeatherObject struct {
	ApiName     string  `json:"apiname,omitempty"`
	DateCall    string  `json:"datecall,omitempty"`
	LastUpdated string  `json:"lastupdated,omitempty"`
	TempC       float64 `json:"tempc,omitempty"`
	TempF       float64 `json:"tempf,omitempty"`
	WindMph     float64 `json:"windmph,omitempty"`
	WindKph     float64 `json:"windkph,omitempty"`
	WindDir     string  `json:"winddir,omitempty"`
	Cloud       int     `json:"cloud,omitempty"`
	Sky         string  `json:"sky"`
	City        string  `json:"city"`
}
