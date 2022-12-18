package service

import (
	"github.com/vv-projects/go-apicaller/pkg/models/baseobjects"
	"github.com/vv-projects/go-apicaller/pkg/repository"
	"log"
	"math/rand"
)

type QuotesService struct {
	mongo repository.Quote
	redis repository.Quote
}

func NewQuoteService(mongorepos repository.Quote, redisrepos repository.Quote) *QuotesService {
	return &QuotesService{
		mongo: mongorepos,
		redis: redisrepos,
	}
}

// save to redisdb and mongo
func (q *QuotesService) SaveQuote(quote baseobjects.QuotesObject) bool {
	errRedis := q.redis.SaveQuote(quote)
	if errRedis != nil {
		log.Printf("can`t create quote for redis: %s", errRedis)
		// bad that can`t save to redisdb, but not critical
	}

	err := q.mongo.SaveQuote(quote)
	if err != nil {
		log.Printf("can`t create quote for mongo: %s", err)
		return false
	}
	return true
}

// get random quote from all quotes apis
// get from redisdb.
//if no quote in redisdb, then from mongo
func (q *QuotesService) GetQuoteByDate(date string) (baseobjects.QuotesObject, bool) {
	quotesRedis, err := q.redis.GetQuoteByDate(date)
	if err == nil && len(quotesRedis) > 0 {
		return quotesRedis[0], true
	}
	quotesMongo, err := q.mongo.GetQuoteByDate(date)
	if err != nil || len(quotesMongo) == 0 {
		log.Printf("can`t get quotes from db, cause : %s", err)
		return baseobjects.QuotesObject{}, false
	}
	quoteRand := rand.Intn(len(quotesMongo))

	return quotesMongo[quoteRand], true
}
