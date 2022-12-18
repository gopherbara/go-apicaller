package mongodb

import (
	"context"
	"errors"
	"fmt"
	"github.com/vv-projects/go-apicaller/pkg/models/baseobjects"
	"github.com/vv-projects/go-apicaller/pkg/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type WeatherMongo struct {
	MongoDB *mongo.Database
}

func NewMongoWeather(db *mongo.Database) *WeatherMongo {
	return &WeatherMongo{MongoDB: db}
}

func (w WeatherMongo) SaveWeather(weather baseobjects.WeatherObject) error {
	collection := w.MongoDB.Collection(repository.WeaherTable)
	res, insertErr := collection.InsertOne(context.Background(), weather)
	if insertErr != nil {
		return insertErr
	}
	fmt.Println(res)
	return nil
}

func (w WeatherMongo) GetWeatherByDate(date string) ([]baseobjects.WeatherObject, error) {
	collection := w.MongoDB.Collection(repository.WeaherTable)
	if collection == nil {
		return nil, errors.New(fmt.Sprintf("no such collection in mongo db %s", repository.WeaherTable))
	}
	query := bson.D{
		{"datecall", bson.D{{"$regex", date}}},
	}
	cur, err := collection.Find(context.Background(), query)
	defer cur.Close(context.Background())
	if err != nil {
		return nil, err
	}

	weather := make([]baseobjects.WeatherObject, 0)
	if err := cur.All(context.Background(), &weather); err != nil {
		return nil, err
	}
	return weather, nil
}
