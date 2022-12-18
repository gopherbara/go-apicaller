package redisdb

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/vv-projects/go-apicaller/pkg/models/baseobjects"
	"github.com/vv-projects/go-apicaller/pkg/repository"
	"strings"
)

type CurrencyRedis struct {
	RedisDB *redis.Client
}

func NewRedisCurrency(db *redis.Client) *CurrencyRedis {
	return &CurrencyRedis{RedisDB: db}
}

func (c CurrencyRedis) SaveCurrencyRate(currency baseobjects.CurrencyObject) error {
	json, err := json.Marshal(currency)
	if err != nil {
		return err
	}
	date := currency.DateCall
	if i := strings.Index(currency.DateCall, " "); i >= 0 {
		date = currency.DateCall[:i]
	}
	c.RedisDB.Set(fmt.Sprintf("%s_%s", repository.CurrencyTable, date), json, 0)

	return nil
}

func (c CurrencyRedis) GetCurrencyRateByDate(date string) ([]baseobjects.CurrencyObject, error) {
	currencys := make([]baseobjects.CurrencyObject, 0)

	if i := strings.Index(date, " "); i >= 0 {
		date = date[:i]
	}
	info, err := c.RedisDB.Get(fmt.Sprintf("%s_%s", repository.CurrencyTable, date)).Result()
	if err == redis.Nil {
		return currencys, errors.New("no such key")
	} else if err != nil {
		return currencys, err
	}

	var currency baseobjects.CurrencyObject
	err = json.Unmarshal([]byte(info), &currency)
	if err != nil {
		return nil, err
	}
	currencys = append(currencys, currency)
	return currencys, nil
}

func (c CurrencyRedis) GetCurrencyRateByDateApi(date string, apiName string) ([]baseobjects.CurrencyObject, error) {
	return nil, errors.New("not implemented yet and really not sure if needed")

}
