package service

import (
	"github.com/vv-projects/go-apicaller/pkg/models/baseobjects"
	"github.com/vv-projects/go-apicaller/pkg/repository"
	"log"
	"math/rand"
)

type CryptoService struct {
	mongo repository.Crypto
	redis repository.Crypto
}

func NewCryptoService(mongorepos repository.Crypto, redisrepos repository.Crypto) *CryptoService {
	return &CryptoService{
		mongo: mongorepos,
		redis: redisrepos,
	}
}

func (c CryptoService) SaveCrypto(crypto baseobjects.CryptoObject) bool {
	errRedis := c.redis.SaveCrypto(crypto)
	if errRedis != nil {
		log.Fatalf("can`t save crypto for redis: %s", errRedis)
		// bad that can`t save to redisdb, but not critical
	}

	err := c.mongo.SaveCrypto(crypto)
	if err != nil {
		log.Fatalf("can`t save crypto for mongo: %s", err)
		return false
	}
	return true
}

func (c CryptoService) GetCryptoByDate(date string) (baseobjects.CryptoObject, bool) {
	cryptoRedis, err := c.redis.GetCryptoByDate(date)
	if err == nil && len(cryptoRedis) > 0 {
		return cryptoRedis[0], true
	}
	cryptoMongo, err := c.mongo.GetCryptoByDate(date)
	if err != nil || len(cryptoMongo) == 0 {
		log.Printf("can`t get crypto from db, cause : %s", err)
		return baseobjects.CryptoObject{}, false
	}
	rand := rand.Intn(len(cryptoMongo))
	return cryptoMongo[rand], true
}

func (c CryptoService) GetCryptoByDateApi(date string, apiName string) (baseobjects.CryptoObject, bool) {
	cryptoMongo, err := c.mongo.GetCryptoByDateApi(date, apiName)
	if err != nil || len(cryptoMongo) == 0 {
		log.Printf("can`t get crypto from db, cause : %s", err)
		return baseobjects.CryptoObject{}, false
	}
	rand := rand.Intn(len(cryptoMongo))
	return cryptoMongo[rand], true
}
