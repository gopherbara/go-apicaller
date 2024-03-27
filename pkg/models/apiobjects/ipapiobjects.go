package apiobjects

import (
	"github.com/gopherbara/go-apicaller/pkg/models/baseobjects"
)

type IpApiSettings struct {
	General     GeneralSettings
	Name        string
	Description string
	Url         string
	IP          string
}

type IpApiResponse struct {
	Status      string  `json:"status"`
	Country     string  `json:"country"`
	CountryCode string  `json:"countryCode"`
	Region      string  `json:"region"`
	RegionName  string  `json:"regionName"`
	City        string  `json:"city"`
	Zip         string  `json:"zip"`
	Lat         float64 `json:"lat"`
	Lon         float64 `json:"lon"`
	Timezone    string  `json:"timezone"`
	Isp         string  `json:"isp"`
	Org         string  `json:"org"`
	As          string  `json:"as"`
	Query       string  `json:"query"`
}

func (ia IpApiResponse) ToBase(name string, date string) baseobjects.GeoLocationObject {
	return baseobjects.GeoLocationObject{
		ApiName:     name,
		DateCall:    date,
		Ip:          ia.Query,
		City:        ia.City,
		Country:     ia.Country,
		CountryCode: ia.CountryCode,
		Lat:         ia.Lat,
		Lon:         ia.Lon,
	}
}
