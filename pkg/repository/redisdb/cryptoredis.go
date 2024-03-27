package redisdb

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/gopherbara/go-apicaller/pkg/models/baseobjects"
	"github.com/gopherbara/go-apicaller/pkg/repository"
	"strings"
)

type CryptoRedis struct {
	RedisDB *redis.Client
}

func NewRedisCrypto(db *redis.Client) *CryptoRedis {
	return &CryptoRedis{RedisDB: db}
}

func (c CryptoRedis) SaveCrypto(crypto baseobjects.CryptoObject) error {
	json, err := json.Marshal(crypto)
	if err != nil {
		return err
	}
	date := crypto.DateCall
	if i := strings.Index(crypto.DateCall, " "); i >= 0 {
		date = crypto.DateCall[:i]
	}
	c.RedisDB.Set(fmt.Sprintf("%s_%s", repository.CryptoTable, date), json, 0)

	return nil
}

func (c CryptoRedis) GetCryptoByDate(date string) ([]baseobjects.CryptoObject, error) {
	cryptos := make([]baseobjects.CryptoObject, 0)

	if i := strings.Index(date, " "); i >= 0 {
		date = date[:i]
	}
	info, err := c.RedisDB.Get(fmt.Sprintf("%s_%s", repository.CryptoTable, date)).Result()
	if err == redis.Nil {
		return cryptos, errors.New("no such key")
	} else if err != nil {
		return cryptos, err
	}

	var crypto baseobjects.CryptoObject
	err = json.Unmarshal([]byte(info), &crypto)
	if err != nil {
		return nil, err
	}
	cryptos = append(cryptos, crypto)
	return cryptos, nil
}

func (c CryptoRedis) GetCryptoByDateApi(date string, apiName string) ([]baseobjects.CryptoObject, error) {
	return nil, errors.New("not implemented yet and really not sure if needed")

}
