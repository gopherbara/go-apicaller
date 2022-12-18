package apihandler

import (
	"net/http"
)

type ApiHandle interface {
	FillSettings(settings map[string]interface{}) error

	ProcessApiCall() // custom process call
	BuildApiRequest() *http.Request
	GetResponse(request *http.Request) (string, error)
	ParseResponse(response string) error

	GetFromFile() (string, error) // for tests, not to call real api
	Show()
}
