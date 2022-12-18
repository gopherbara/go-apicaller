package handler

import (
	"encoding/json"
	"github.com/vv-projects/go-apicaller/pkg/apihandler"
	"github.com/vv-projects/go-apicaller/pkg/models/baseobjects"
	"github.com/vv-projects/go-apicaller/pkg/service"
	"github.com/vv-projects/go-apicaller/pkg/socket"
	"time"
)

type QuotesHandler struct {
	manager  *socket.Manager
	services *service.Service
	settings map[string]interface{}
	Quote    baseobjects.QuotesObject
}

func NewQuotesHandler(manager *socket.Manager, services *service.Service, settings map[string]interface{}) *QuotesHandler {
	return &QuotesHandler{
		manager:  manager,
		services: services,
		settings: settings,
		Quote:    baseobjects.QuotesObject{},
	}
}

func (q *QuotesHandler) Handle() {
	qa := apihandler.NewChuckNorrisQuotesApi("ChuckNorrisQuotes", q.settings)

	timeCall := time.Now().Format("2006-01-02 15:04")

	qa.ProcessApiCall()
	q.Quote = qa.Response.ToBase(qa.Settings.Name, timeCall)
	if qa.Settings.General.IsProd {
		q.services.SaveQuote(q.Quote)

	}

	q.Send()
}

func (q *QuotesHandler) Send() {
	str, _ := json.Marshal(q.Quote)
	msg := socket.Message{
		ApiType: "quote",
		Body:    string(str),
	}
	q.manager.UpdateMessage <- msg
}
