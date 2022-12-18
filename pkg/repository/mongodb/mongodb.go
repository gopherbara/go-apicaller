package mongodb

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"github.com/vv-projects/go-apicaller/pkg/repository"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

type MongoConfig struct {
	Url    string
	DBName string
}

func GetMongoConfig() MongoConfig {
	return MongoConfig{
		Url:    viper.GetString("mongodb.uri"),
		DBName: viper.GetString("mongodb.dbname"),
	}
}

func NewMongoDB() (*mongo.Client, *mongo.Database, error) {
	cfg := GetMongoConfig()

	client, err := mongo.NewClient(options.Client().ApplyURI(fmt.Sprintf("%s", cfg.Url)))
	if err != nil {
		log.Fatalf("error on creating mongo client: %s", err.Error())
		return client, nil, err
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	//defer client.Disconnect(ctx)
	if err != nil {
		log.Fatalf("error on connecting to mongo db: %s", err.Error())
		return client, nil, err
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatalf("error on ping mongo db: %s", err.Error())
		return client, nil, err
	}

	db := client.Database(cfg.DBName)

	return client, db, nil
}

func NewMongoRepos(db *mongo.Database) *repository.Repository {
	return &repository.Repository{
		Quote:        NewMongoQuote(db),
		GeoLocation:  NewMongoGeoLocation(db),
		Weather:      NewMongoWeather(db),
		CurrencyRate: NewMongoCurrency(db),
		Crypto:       NewMongoCrypto(db),
		Stocks:       NewMongoStocks(db),
	}
}
