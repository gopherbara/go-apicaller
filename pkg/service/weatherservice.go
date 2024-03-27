package service

import (
	"github.com/gopherbara/go-apicaller/pkg/models/baseobjects"
	"github.com/gopherbara/go-apicaller/pkg/repository"
	"log"
	"math/rand"
)

type WeatherService struct {
	mongo repository.Weather
	redis repository.Weather
}

func NewWeatherService(mongorepos repository.Weather, redisrepos repository.Weather) *WeatherService {
	return &WeatherService{
		mongo: mongorepos,
		redis: redisrepos,
	}
}

func (w WeatherService) SaveWeather(weather baseobjects.WeatherObject) bool {
	errRedis := w.redis.SaveWeather(weather)
	if errRedis != nil {
		log.Printf("can`t save weather for redis: %s", errRedis)
		// bad that can`t save to redisdb, but not critical
	}

	err := w.mongo.SaveWeather(weather)
	if err != nil {
		log.Printf("can`t save weather for mongo: %s", err)
		return false
	}
	return true
}

func (w WeatherService) GetWeatherByDate(date string) (baseobjects.WeatherObject, bool) {
	quotesRedis, err := w.redis.GetWeatherByDate(date)
	if err == nil && len(quotesRedis) > 0 {
		return quotesRedis[0], true
	}
	weatherMongo, err := w.mongo.GetWeatherByDate(date)
	if err != nil || len(weatherMongo) == 0 {
		log.Printf("can`t get weather from db, cause : %s", err)
		return baseobjects.WeatherObject{}, false
	}
	rand := rand.Intn(len(weatherMongo))
	return weatherMongo[rand], true
}
