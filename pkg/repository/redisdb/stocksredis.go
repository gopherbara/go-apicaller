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

type StocksRedis struct {
	RedisDB *redis.Client
}

func NewRedisStocks(db *redis.Client) *StocksRedis {
	return &StocksRedis{RedisDB: db}
}

func (s StocksRedis) SaveStocks(stocks baseobjects.StocksObject) error {
	jsonObj, err := json.Marshal(stocks)
	if err != nil {
		return err
	}
	date := stocks.DateCall
	if i := strings.Index(stocks.DateCall, " "); i >= 0 {
		date = stocks.DateCall[:i]
	}
	s.RedisDB.Set(fmt.Sprintf("%s_%s", repository.StocksTable, date), jsonObj, 0)

	return nil
}

func (s StocksRedis) GetStocksByDate(date string) ([]baseobjects.StocksObject, error) {
	stocks := make([]baseobjects.StocksObject, 0)

	if i := strings.Index(date, " "); i >= 0 {
		date = date[:i]
	}
	info, err := s.RedisDB.Get(fmt.Sprintf("%s_%s", repository.WeaherTable, date)).Result()
	if err == redis.Nil {
		return stocks, errors.New("no such key")
	} else if err != nil {
		return stocks, err
	}
	var stock baseobjects.StocksObject
	err = json.Unmarshal([]byte(info), &stock)
	if err != nil {
		return nil, err
	}
	stocks = append(stocks, stock)
	return stocks, nil
}

func (s StocksRedis) GetStocksByDateApi(date string, apiName string) ([]baseobjects.StocksObject, error) {
	return nil, errors.New("not implemented yet and really not sure if needed")
}
