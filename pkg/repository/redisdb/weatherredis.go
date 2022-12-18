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

type WeatherRedis struct {
	RedisDB *redis.Client
}

func NewRedisWeather(db *redis.Client) *WeatherRedis {
	return &WeatherRedis{RedisDB: db}
}

func (w WeatherRedis) SaveWeather(weather baseobjects.WeatherObject) error {

	jsonWeather, err := json.Marshal(weather)
	if err != nil {
		return err
	}
	date := weather.DateCall
	if i := strings.Index(weather.DateCall, " "); i >= 0 {
		date = weather.DateCall[:i]
	}
	w.RedisDB.Set(fmt.Sprintf("%s_%s", repository.WeaherTable, date), jsonWeather, 0)

	return nil
}

func (w WeatherRedis) GetWeatherByDate(date string) ([]baseobjects.WeatherObject, error) {
	weathers := make([]baseobjects.WeatherObject, 0)

	if i := strings.Index(date, " "); i >= 0 {
		date = date[:i]
	}
	info, err := w.RedisDB.Get(fmt.Sprintf("%s_%s", repository.WeaherTable, date)).Result()
	if err == redis.Nil {
		return weathers, errors.New("no such key")
	} else if err != nil {
		return weathers, err
	}
	var weather baseobjects.WeatherObject
	err = json.Unmarshal([]byte(info), &weather)
	if err != nil {
		return nil, err
	}
	weathers = append(weathers, weather)
	return weathers, nil
}
