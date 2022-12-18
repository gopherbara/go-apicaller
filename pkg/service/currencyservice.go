package service

import (
	"github.com/vv-projects/go-apicaller/pkg/models/baseobjects"
	"github.com/vv-projects/go-apicaller/pkg/repository"
	"log"
	"math/rand"
)

type CurrencyService struct {
	mongo repository.CurrencyRate
	redis repository.CurrencyRate
}

func NewCurrencyService(mongorepos repository.CurrencyRate, redisrepos repository.CurrencyRate) *CurrencyService {
	return &CurrencyService{
		mongo: mongorepos,
		redis: redisrepos,
	}
}

func (c CurrencyService) SaveCurrencyRate(currency baseobjects.CurrencyObject) bool {
	errRedis := c.redis.SaveCurrencyRate(currency)
	if errRedis != nil {
		log.Printf("can`t save currency for redis: %s", errRedis)
		// bad that can`t save to redisdb, but not critical
	}

	err := c.mongo.SaveCurrencyRate(currency)
	if err != nil {
		log.Printf("can`t save currency for mongo: %s", err)
		return false
	}
	return true
}

func (c CurrencyService) GetCurrencyRateByDate(date string) (baseobjects.CurrencyObject, bool) {
	currencyRedis, err := c.redis.GetCurrencyRateByDate(date)
	if err == nil && len(currencyRedis) > 0 {
		return currencyRedis[0], true
	}
	currencyMongo, err := c.mongo.GetCurrencyRateByDate(date)
	if err != nil || len(currencyMongo) == 0 {
		log.Printf("can`t get currency from db, cause : %s", err)
		return baseobjects.CurrencyObject{}, false
	}
	rand := rand.Intn(len(currencyMongo))
	return currencyMongo[rand], true
}

func (c CurrencyService) GetCurrencyRateByDateApi(date string, apiName string) (baseobjects.CurrencyObject, bool) {
	currencyMongo, err := c.mongo.GetCurrencyRateByDateApi(date, apiName)
	if err != nil || len(currencyMongo) == 0 {
		log.Printf("can`t get currency from db, cause : %s", err)
		return baseobjects.CurrencyObject{}, false
	}
	rand := rand.Intn(len(currencyMongo))
	return currencyMongo[rand], true
}
