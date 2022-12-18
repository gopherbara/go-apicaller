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
	"strings"
)

type CoinApi struct {
	Settings apiobjects.CoinSettings
	Response apiobjects.CoinApiResponses
}

func NewCoinApi(name string, settings map[string]any) *CoinApi {
	ca := &CoinApi{
		Settings: apiobjects.CoinSettings{
			Name:    name,
			General: apiobjects.FillGeneralSettings(settings),
		},
	}
	if err := ca.FillSettings(settings); err != nil {
		log.Printf("error from %s on fill setiings: %s", name, err)
		return nil
	}
	return ca
}

func (c *CoinApi) ProcessApiCall() {
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

func (c *CoinApi) FillSettings(settings map[string]interface{}) error {
	if v, ok := settings[c.Settings.Name]; ok {
		set := v.(map[string]interface{})
		for key, val := range set {
			switch key {
			case "url":
				c.Settings.Url = val.(string)
				break
			case "cryptos":
				c.Settings.Cryptos = make([]string, len(val.([]interface{})))
				for i, v := range val.([]interface{}) {
					c.Settings.Cryptos[i] = fmt.Sprint(v)
				}
				break
			default:
				break
			}
		}
	} else {
		return errors.New("no settings for api")
	}
	c.Settings.APIKey = os.Getenv("COINAPI_KEY")
	return nil
}

func (c *CoinApi) BuildApiRequest() *http.Request {
	var url string
	if len(c.Settings.Cryptos) > 0 {
		url = fmt.Sprintf("%s?filter_asset_id=%s", c.Settings.Url, strings.Join(c.Settings.Cryptos, "%2C"))
	} else {
		url = c.Settings.Url
	}

	rq, _ := http.NewRequest("GET", url, nil)

	rq.Header.Add("X-CoinAPI-Key", c.Settings.APIKey)
	rq.Header.Add("Accept", "application/json")
	//rq.Header.Add("Accept-Encoding", "deflate, gzip")
	return rq
}

func (c *CoinApi) GetFromFile() (string, error) {
	fileName := fmt.Sprintf("responses/%sResponseExample.txt", c.Settings.Name)
	file, err := os.Open(fileName)
	defer file.Close()

	if err == nil {
		byteValue, _ := ioutil.ReadAll(file)
		return string(byteValue), nil
	}
	return "", err
}

func (c *CoinApi) GetResponse(request *http.Request) (string, error) {
	res, _ := http.DefaultClient.Do(request)

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	return string(body), err
}

func (c *CoinApi) ParseResponse(response string) error {
	err := json.Unmarshal([]byte(response), &c.Response)
	return err
}

func (c *CoinApi) Show() {
	for _, val := range c.Response {
		fmt.Println(val.AssetId, val.PriceUsd)
	}
}
