package apihandler

import "C"
import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gopherbara/go-apicaller/pkg/models/apiobjects"
	"github.com/gopherbara/go-apicaller/pkg/service"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

type CurrencyRateApi struct {
	Settings apiobjects.ExchangeRatesSettings
	Response apiobjects.ExchangeRatesResponse
	Service  *service.Service
}

func NewCurrencyRateApi(name string, settings map[string]interface{}) *CurrencyRateApi {
	cra := &CurrencyRateApi{
		Settings: apiobjects.ExchangeRatesSettings{
			Name:    name,
			General: apiobjects.FillGeneralSettings(settings),
		},
	}
	if err := cra.FillSettings(settings); err != nil {
		log.Printf("error from %s on fill setiings: %s", name, err)
		return nil
	}
	return cra
}

func (c *CurrencyRateApi) ProcessApiCall() {
	var resp string
	if c.Settings.General.IsProd {
		rq := c.BuildApiRequest()
		resp, _ = c.GetResponse(rq)
	} else {
		resp, _ = c.GetFromFile()
	}
	c.ParseResponse(resp)
	log.Printf("apihandler %s end call to external api", c.Settings.Name)

	//c.Show()
}

func (c *CurrencyRateApi) FillSettings(settings map[string]interface{}) error {
	if v, ok := settings[c.Settings.Name]; ok {
		set := v.(map[string]interface{})
		for key, val := range set {
			switch key {
			case "url":
				c.Settings.Url = val.(string)
				break
			case "base":
				c.Settings.FromCurrency = val.(string)
				break
			case "currencies":
				c.Settings.Currencies = make([]string, len(val.([]interface{})))
				for i, v := range val.([]interface{}) {
					c.Settings.Currencies[i] = fmt.Sprint(v)
				}
				break
			default:
				break
			}
		}
	} else {
		return errors.New("no settings for api")
	}
	c.Settings.APIKey = os.Getenv("CURRENCYAPI_KEY")
	return nil
}

func (c *CurrencyRateApi) BuildApiRequest() *http.Request {

	url := fmt.Sprintf("%ssymbols=%s&base=%s", c.Settings.Url, strings.Join(c.Settings.Currencies, "%2C"), c.Settings.FromCurrency)
	rq, _ := http.NewRequest("GET", url, nil)

	rq.Header.Add("apikey", c.Settings.APIKey)
	return rq
}

func (c *CurrencyRateApi) GetFromFile() (string, error) {
	fileName := fmt.Sprintf("responses/%sResponseExample.txt", c.Settings.Name)
	file, err := os.Open(fileName)

	if err == nil {
		byteValue, _ := ioutil.ReadAll(file)
		return string(byteValue), nil
	}
	return "", err
}

func (c *CurrencyRateApi) GetResponse(request *http.Request) (string, error) {
	res, _ := http.DefaultClient.Do(request)

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	return string(body), err
}

func (c *CurrencyRateApi) ParseResponse(response string) error {
	err := json.Unmarshal([]byte(response), &c.Response)
	return err
}

func (c *CurrencyRateApi) Show() {
	for k, v := range c.Response.Rates {
		fmt.Println(k, 1.0/v)
	}
}
