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

type AlphaVantageApi struct {
	Settings apiobjects.AlphaVantageSettings
	Response apiobjects.AlphaVantageResponse
}

func NewAlphaVantageApi(name string, settings map[string]interface{}) *AlphaVantageApi {
	ava := &AlphaVantageApi{
		Settings: apiobjects.AlphaVantageSettings{
			Name:    name,
			General: apiobjects.FillGeneralSettings(settings),
		},
		Response: apiobjects.AlphaVantageResponse{},
	}
	if err := ava.FillSettings(settings); err != nil {
		log.Printf("error from %s on fill setiings: %s", name, err)
		return nil
	}
	return ava
}

func (a *AlphaVantageApi) ProcessApiCall() {
	var resp string
	if a.Settings.General.IsProd {
		rq := a.BuildApiRequest()
		resp, _ = a.GetResponse(rq)
	} else {
		resp, _ = a.GetFromFile()
	}
	a.ParseResponse(resp)
	log.Printf("apihandler %s end call to external api", a.Settings.Name)
	//a.Show()
}

func (a *AlphaVantageApi) FillSettings(settings map[string]interface{}) error {
	if v, ok := settings[a.Settings.Name]; ok {
		set := v.(map[string]interface{})
		for key, val := range set {
			switch key {
			case "url":
				a.Settings.Url = val.(string)
				break
			case "stock":
				a.Settings.Stock = val.(string)
				break

			case "interval":
				a.Settings.Interval = val.(string)
				break
			default:
				break
			}
		}
	} else {
		return errors.New("no settings for api")
	}
	a.Settings.APIKey = os.Getenv("ALPHAVANTAGEAPI_KEY")
	return nil
}

func (a *AlphaVantageApi) BuildApiRequest() *http.Request {

	url := fmt.Sprintf("%ssymbol=%s&interval=%s&apikey=%s", a.Settings.Url, a.Settings.Stock, a.Settings.Interval, a.Settings.APIKey)
	rq, _ := http.NewRequest("GET", url, nil)

	return rq
}

func (a *AlphaVantageApi) GetFromFile() (string, error) {
	fileName := fmt.Sprintf("responses/%sResponseExample.txt", a.Settings.Name)
	file, err := os.Open(fileName)
	defer file.Close()

	if err == nil {
		byteValue, _ := ioutil.ReadAll(file)
		return string(byteValue), nil
	}
	return "", err
}

func (a *AlphaVantageApi) GetResponse(request *http.Request) (string, error) {
	res, _ := http.DefaultClient.Do(request)

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	return string(body), err
}

func (a *AlphaVantageApi) ParseResponse(response string) error {
	err := json.Unmarshal([]byte(response), &a.Response)
	return err
}

func (a *AlphaVantageApi) Show() {
	fmt.Println(a.Response.TimeSeriesFiels[a.Response.MetaData.LastRefreshed])
}
