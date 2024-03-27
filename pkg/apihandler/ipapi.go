package apihandler

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gopherbara/go-apicaller/pkg/models/apiobjects"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type IpApi struct {
	Settings apiobjects.IpApiSettings
	Response apiobjects.IpApiResponse
}

func NewIpApi(name string, settings map[string]interface{}) *IpApi {
	gla := &IpApi{
		Settings: apiobjects.IpApiSettings{
			Name:    name,
			General: apiobjects.FillGeneralSettings(settings),
		},
	}
	if err := gla.FillSettings(settings); err != nil {
		log.Printf("error from %s on fill setiings: %s", name, err)
		return nil
	}
	return gla
}

func (g *IpApi) ProcessApiCall() {
	var resp string
	if g.Settings.General.IsProd {
		rq := g.BuildApiRequest()
		resp, _ = g.GetResponse(rq)
	} else {
		resp, _ = g.GetFromFile()
	}
	g.ParseResponse(resp)
	log.Printf("apihandler %s end call to external api", g.Settings.Name)

	//g.Show()
}

func (g *IpApi) FillSettings(settings map[string]interface{}) error {
	if v, ok := settings[g.Settings.Name]; ok {
		set := v.(map[string]interface{})
		for key, val := range set {
			switch key {
			case "url":
				g.Settings.Url = val.(string)
				break
			case "ip":
				g.Settings.IP = val.(string)
				break
			default:
				break
			}
		}
	} else {
		return errors.New("no settings for api")

	}
	return nil

}

func (g *IpApi) BuildApiRequest() *http.Request {
	url := g.Settings.Url
	if g.Settings.IP != "" {
		url = fmt.Sprintf("%s/%s", g.Settings.Url, g.Settings.IP)
	}
	rq, _ := http.NewRequest("GET", url, nil)
	return rq
}

func (g *IpApi) GetFromFile() (string, error) {
	fileName := fmt.Sprintf("responses/%sResponseExample.txt", g.Settings.Name)
	file, err := os.Open(fileName)
	defer file.Close()

	if err == nil {
		byteValue, _ := ioutil.ReadAll(file)
		return string(byteValue), nil
	}
	return "", err
}

func (g *IpApi) GetResponse(request *http.Request) (string, error) {
	res, _ := http.DefaultClient.Do(request)

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	return string(body), err
}

func (g *IpApi) ParseResponse(response string) error {
	err := json.Unmarshal([]byte(response), &g.Response)
	return err
}

func (g *IpApi) Show() {
	fmt.Printf("Country:%s \n City:%s", g.Response.Country, g.Response.City)
}
