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

type QuoteRedis struct {
	RedisDB *redis.Client
}

func NewRedisQuote(db *redis.Client) *QuoteRedis {
	return &QuoteRedis{RedisDB: db}
}

func (q QuoteRedis) SaveQuote(quote baseobjects.QuotesObject) error {
	jsonQuote, err := json.Marshal(quote)
	if err != nil {
		return err
	}
	date := quote.DateCall
	if i := strings.Index(quote.DateCall, " "); i >= 0 {
		date = quote.DateCall[:i]
	}
	q.RedisDB.Set(fmt.Sprintf("%s_%s", repository.QuotesTable, date), jsonQuote, 0)

	return nil
}

// get last quote in redis for today
func (q QuoteRedis) GetQuoteByDate(date string) ([]baseobjects.QuotesObject, error) {
	quotes := make([]baseobjects.QuotesObject, 0)

	if i := strings.Index(date, " "); i >= 0 {
		date = date[:i]
	}
	info, err := q.RedisDB.Get(fmt.Sprintf("%s_%s", repository.QuotesTable, date)).Result()
	if err == redis.Nil {
		return quotes, errors.New("no such key")
	} else if err != nil {
		return quotes, err
	}
	var quote baseobjects.QuotesObject
	err = json.Unmarshal([]byte(info), &quote)
	if err != nil {
		return nil, err
	}
	quotes = append(quotes, quote)
	return quotes, nil
}
