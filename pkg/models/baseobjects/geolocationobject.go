package baseobjects

type GeoLocationObject struct {
	ApiName     string  `json:"apiname,omitempty"`
	DateCall    string  `json:"datecall,omitempty"`
	Ip          string  `json:"ip,omitempty"`
	City        string  `json:"city,omitempty"`
	Country     string  `json:"country,omitempty"`
	CountryCode string  `json:"countrycode,omitempty"`
	Lat         float64 `json:"lat,omitempty"`
	Lon         float64 `json:"lon,omitempty"`
}
