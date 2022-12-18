package handler

import (
	"encoding/json"
	"github.com/vv-projects/go-apicaller/pkg/apihandler"
	"github.com/vv-projects/go-apicaller/pkg/models/baseobjects"
	"github.com/vv-projects/go-apicaller/pkg/service"
	"github.com/vv-projects/go-apicaller/pkg/socket"
	"time"
)

type GeoLocationHandler struct {
	manager  *socket.Manager
	services *service.Service
	settings map[string]interface{}
	Geo      baseobjects.GeoLocationObject
}

func NewGeoLocationHandler(manager *socket.Manager, services *service.Service, settings map[string]interface{}) *GeoLocationHandler {
	return &GeoLocationHandler{
		manager:  manager,
		services: services,
		settings: settings,
		Geo:      baseobjects.GeoLocationObject{},
	}
}

func (q *GeoLocationHandler) Handle() {
	qa := apihandler.NewIpApi("IpApi", q.settings)

	timeCall := time.Now().Format("2006-01-02 15:04")

	qa.ProcessApiCall()
	q.Geo = qa.Response.ToBase(qa.Settings.Name, timeCall)
	if qa.Settings.General.IsProd {
		q.services.SaveGeoLocation(q.Geo)
	}

	q.Send()
}

func (q *GeoLocationHandler) Send() {
	str, _ := json.Marshal(q.Geo)
	msg := socket.Message{
		ApiType: "geo",
		Body:    string(str),
	}
	q.manager.UpdateMessage <- msg
}
