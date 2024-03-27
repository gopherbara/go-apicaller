package redisdb

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/gopherbara/go-apicaller/pkg/repository"
	"github.com/spf13/viper"
	"os"
)

type RedisConfig struct {
	Addr     string
	Password string
	DB       int
}

func GetRedisConfig() RedisConfig {
	return RedisConfig{
		Addr:     fmt.Sprintf("%s:%s", viper.GetString("redisdb.host"), viper.GetString("redisdb.port")),
		Password: os.Getenv("REDISPASSWORD"),
		DB:       viper.GetInt("redisdb.db"),
	}
}

func NewRedisRepos(redisdb *redis.Client) *repository.Repository {
	return &repository.Repository{
		Quote:        NewRedisQuote(redisdb),
		GeoLocation:  NewRedisGeoLocation(redisdb),
		Weather:      NewRedisWeather(redisdb),
		CurrencyRate: NewRedisCurrency(redisdb),
		Crypto:       NewRedisCrypto(redisdb),
		Stocks:       NewRedisStocks(redisdb),
	}
}

func getRedisClient(addr string, password string, db int) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db, // use default DB
	})
}

func NewRedisDB() *redis.Client {
	config := GetRedisConfig()
	return redis.NewClient(&redis.Options{
		Addr:     config.Addr,
		Password: config.Password,
		DB:       config.DB, // use default DB
	})
}
