package mongodb

import (
	"context"
	"errors"
	"fmt"
	"github.com/gopherbara/go-apicaller/pkg/models/baseobjects"
	"github.com/gopherbara/go-apicaller/pkg/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type CurrencyMongo struct {
	MongoDB *mongo.Database
}

func NewMongoCurrency(db *mongo.Database) *CurrencyMongo {
	return &CurrencyMongo{MongoDB: db}
}

func (c CurrencyMongo) SaveCurrencyRate(currency baseobjects.CurrencyObject) error {
	collection := c.MongoDB.Collection(repository.CurrencyTable)
	res, insertErr := collection.InsertOne(context.Background(), currency)
	if insertErr != nil {
		return insertErr
	}
	fmt.Println(res)
	return nil
}

func (c CurrencyMongo) GetCurrencyRateByDate(date string) ([]baseobjects.CurrencyObject, error) {
	collection := c.MongoDB.Collection(repository.CurrencyTable)
	if collection == nil {
		return nil, errors.New(fmt.Sprintf("no such collection in mongo db %s", repository.CurrencyTable))
	}
	query := bson.D{
		{"datecall", bson.D{{"$regex", date}}},
	}
	cur, err := collection.Find(context.Background(), query)
	defer cur.Close(context.Background())
	if err != nil {
		return nil, err
	}

	currency := make([]baseobjects.CurrencyObject, 0)
	if err := cur.All(context.Background(), &currency); err != nil {
		return nil, err
	}
	return currency, nil
}

func (c CurrencyMongo) GetCurrencyRateByDateApi(date string, apiName string) ([]baseobjects.CurrencyObject, error) {
	collection := c.MongoDB.Collection(repository.CurrencyTable)
	if collection == nil {
		return nil, errors.New(fmt.Sprintf("no such collection in mongo db %s", repository.CurrencyTable))
	}
	query := bson.D{
		{"datecall", bson.D{{"$regex", date}}},
		{"apiname", apiName},
	}
	cur, err := collection.Find(context.Background(), query)
	defer cur.Close(context.Background())
	if err != nil {
		return nil, err
	}

	currency := make([]baseobjects.CurrencyObject, 0)
	if err := cur.All(context.Background(), &currency); err != nil {
		return nil, err
	}
	return currency, nil
}
