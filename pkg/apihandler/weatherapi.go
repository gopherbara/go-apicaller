package apihandler

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/vv-projects/go-apicaller/pkg/models/apiobjects"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type WeatherApi struct {
	Settings apiobjects.WeatherSettings
	Response apiobjects.OpenWeatherResponse
}

func NewWeatherApi(name string, settings map[string]interface{}) *WeatherApi {
	wa := &WeatherApi{
		Settings: apiobjects.WeatherSettings{
			Name:    name,
			General: apiobjects.FillGeneralSettings(settings),
		},
		Response: apiobjects.OpenWeatherResponse{},
	}
	if err := wa.FillSettings(settings); err != nil {
		log.Printf("error from %s on fill setiings: %s", name, err)
		return nil
	}
	return wa
}

func (w *WeatherApi) ProcessApiCall() {
	if w.Settings.City == "" {
		return
	}
	var resp string
	if w.Settings.General.IsProd {
		rq := w.BuildApiRequest()
		resp, _ = w.GetResponse(rq)
	} else {
		resp, _ = w.GetFromFile()
	}
	w.ParseResponse(resp)
	log.Printf("apihandler %s end call to external api", w.Settings.Name)

	//w.Show()
}

func (w *WeatherApi) FillSettings(settings map[string]interface{}) error {
	//w.Settings = models.WeatherSettings{}
	if v, ok := settings[w.Settings.Name]; ok {
		set := v.(map[string]interface{})
		for key, val := range set {
			switch key {
			case "apihost":
				w.Settings.APIHost = val.(string)
				break
			case "url":
				w.Settings.Url = val.(string)
				break
			case "city":
				w.Settings.City = val.(string)
				break
			default:
				break
			}
		}
	} else {
		return errors.New("no settings for api")
	}
	w.Settings.APIKey = os.Getenv("WEATHERAPI_KEY")
	return nil
}

func (w *WeatherApi) BuildApiRequest() *http.Request {
	rq, _ := http.NewRequest("GET", w.Settings.Url+w.Settings.City, nil)

	rq.Header.Add("X-RapidAPI-Key", w.Settings.APIKey)
	rq.Header.Add("X-RapidAPI-Host", w.Settings.APIHost)
	return rq
}

func (w *WeatherApi) GetResponse(rq *http.Request) (string, error) {
	res, _ := http.DefaultClient.Do(rq)

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	return string(body), err
}

func (w WeatherApi) GetFromFile() (string, error) {
	fileName := fmt.Sprintf("responses/%sResponseExample.txt", w.Settings.Name)
	file, err := os.Open(fileName)
	defer file.Close()

	if err == nil {
		byteValue, _ := ioutil.ReadAll(file)
		return string(byteValue), nil
	}
	return "", err
}

func (w *WeatherApi) ParseResponse(response string) error {
	err := json.Unmarshal([]byte(response), &w.Response)
	return err
}

// temp show
func (w *WeatherApi) Show() {
	fmt.Println(w.Response)
}
