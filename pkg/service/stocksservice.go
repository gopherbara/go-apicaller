package service

import (
	"github.com/gopherbara/go-apicaller/pkg/models/baseobjects"
	"github.com/gopherbara/go-apicaller/pkg/repository"
	"log"
	"math/rand"
)

type StocksService struct {
	mongo repository.Stocks
	redis repository.Stocks
}

func NewStocksService(mongorepos repository.Stocks, redisrepos repository.Stocks) *StocksService {
	return &StocksService{
		mongo: mongorepos,
		redis: redisrepos,
	}
}

func (s StocksService) SaveStocks(stocks baseobjects.StocksObject) bool {
	errRedis := s.redis.SaveStocks(stocks)
	if errRedis != nil {
		log.Printf("can`t save stocks for redis: %s", errRedis)
		// bad that can`t save to redisdb, but not critical
	}

	err := s.mongo.SaveStocks(stocks)
	if err != nil {
		log.Printf("can`t save stocks for mongo: %s", err)
		return false
	}
	return true
}

func (s StocksService) GetStocksByDate(date string) (baseobjects.StocksObject, bool) {
	stocksRedis, err := s.redis.GetStocksByDate(date)
	if err == nil && len(stocksRedis) > 0 {
		return stocksRedis[0], true
	}
	stocksMongo, err := s.mongo.GetStocksByDate(date)
	if err != nil || len(stocksMongo) == 0 {
		log.Printf("can`t get stocks from db, cause : %s", err)
		return baseobjects.StocksObject{}, false
	}
	rand := rand.Intn(len(stocksMongo))
	return stocksMongo[rand], true
}

func (s StocksService) GetStocksByDateApi(date string, apiName string) (baseobjects.StocksObject, bool) {
	//stocksRedis, err := s.redis.GetStocksByDate(date)
	//if err == nil && len(stocksRedis) > 0 {
	//	return stocksRedis[0], true
	//}
	stocksMongo, err := s.mongo.GetStocksByDateApi(date, apiName)
	if err != nil || len(stocksMongo) == 0 {
		log.Printf("can`t get stocks from db, cause : %s", err)
		return baseobjects.StocksObject{}, false
	}
	rand := rand.Intn(len(stocksMongo))
	return stocksMongo[rand], true
}
