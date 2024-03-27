package apiobjects

import (
	"github.com/gopherbara/go-apicaller/pkg/models/baseobjects"
)

type ChuckNorrisQuotesSettings struct {
	General     GeneralSettings
	Name        string
	Description string
	Url         string
}

type ChuckNorrisQuotesResponse struct {
	Categories []string `json:"categories"`
	CreatedAt  string   `json:"created_at"`
	Id         string   `json:"id"`
	IconUrl    string   `json:"icon_url"` // mostly not found
	UpdatedAt  string   `json:"updated_at"`
	Url        string   `json:"url"`   // full url to joke
	Value      string   `json:"value"` // actual quote
}

func (cnq ChuckNorrisQuotesResponse) ToBase(name string, timeApi string) baseobjects.QuotesObject {
	return baseobjects.QuotesObject{
		ApiName:     name,
		DateCall:    timeApi,
		Quote:       cnq.Value,
		DateCreated: cnq.CreatedAt,
	}
}
