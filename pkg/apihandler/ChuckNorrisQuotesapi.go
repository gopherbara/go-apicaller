package apihandler

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/vv-projects/go-apicaller/pkg/models/apiobjects"
	"github.com/vv-projects/go-apicaller/pkg/service"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type ChuckNorrisQuotesApi struct {
	Settings apiobjects.ChuckNorrisQuotesSettings
	Response apiobjects.ChuckNorrisQuotesResponse
	Service  *service.Service
}

func NewChuckNorrisQuotesApi(name string, settings map[string]interface{}) *ChuckNorrisQuotesApi {
	cnq := &ChuckNorrisQuotesApi{
		Settings: apiobjects.ChuckNorrisQuotesSettings{
			Name:    name,
			General: apiobjects.FillGeneralSettings(settings),
		},
	}
	if err := cnq.FillSettings(settings); err != nil {
		log.Printf("error from %s on fill setiings: %s", name, err)
		return nil
	}
	return cnq
}

func (cnq *ChuckNorrisQuotesApi) ProcessApiCall() {
	var resp string
	if cnq.Settings.General.IsProd {
		rq := cnq.BuildApiRequest()
		resp, _ = cnq.GetResponse(rq)
	} else {
		resp, _ = cnq.GetFromFile()
	}
	cnq.ParseResponse(resp)
	log.Printf("apihandler %s end call to external api", cnq.Settings.Name)
	//cnq.Show()
}

func (cnq *ChuckNorrisQuotesApi) FillSettings(settings map[string]interface{}) error {
	if v, ok := settings[cnq.Settings.Name]; ok {
		set := v.(map[string]interface{})
		for key, val := range set {
			switch key {
			case "url":
				cnq.Settings.Url = val.(string)
				break
			default:
				break
			}
		}
		return nil
	}
	return errors.New("no settings for api")
}

func (cnq *ChuckNorrisQuotesApi) BuildApiRequest() *http.Request {
	rq, _ := http.NewRequest("GET", cnq.Settings.Url, nil)
	return rq
}

func (cnq *ChuckNorrisQuotesApi) GetFromFile() (string, error) {
	fileName := fmt.Sprintf("responses/%sResponseExample.txt", cnq.Settings.Name)
	file, err := os.Open(fileName)
	defer file.Close()

	if err == nil {
		byteValue, _ := ioutil.ReadAll(file)
		return string(byteValue), nil
	}
	return "", err
}

func (cnq *ChuckNorrisQuotesApi) GetResponse(rq *http.Request) (string, error) {
	res, _ := http.DefaultClient.Do(rq)

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	return string(body), err
}

func (cnq *ChuckNorrisQuotesApi) ParseResponse(response string) error {
	err := json.Unmarshal([]byte(response), &cnq.Response)
	return err
}

func (cnq *ChuckNorrisQuotesApi) Show() {
	fmt.Println(cnq.Response.Value)
}
