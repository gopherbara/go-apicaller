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

type StocksMongo struct {
	MongoDB *mongo.Database
}

func NewMongoStocks(db *mongo.Database) *StocksMongo {
	return &StocksMongo{MongoDB: db}
}

func (s StocksMongo) SaveStocks(stocks baseobjects.StocksObject) error {
	collection := s.MongoDB.Collection(repository.StocksTable)
	res, insertErr := collection.InsertOne(context.Background(), stocks)
	if insertErr != nil {
		return insertErr
	}
	fmt.Println(res)
	return nil
}

func (s StocksMongo) GetStocksByDate(date string) ([]baseobjects.StocksObject, error) {
	collection := s.MongoDB.Collection(repository.StocksTable)
	if collection == nil {
		return nil, errors.New(fmt.Sprintf("no such collection in mongo db %s", repository.StocksTable))
	}
	query := bson.D{
		{"datecall", bson.D{{"$regex", date}}},
	}
	cur, err := collection.Find(context.Background(), query)
	defer cur.Close(context.Background())
	if err != nil {
		return nil, err
	}

	stocks := make([]baseobjects.StocksObject, 0)
	if err := cur.All(context.Background(), &stocks); err != nil {
		return nil, err
	}
	return stocks, nil
}

func (s StocksMongo) GetStocksByDateApi(date string, apiName string) ([]baseobjects.StocksObject, error) {
	collection := s.MongoDB.Collection(repository.StocksTable)
	if collection == nil {
		return nil, errors.New(fmt.Sprintf("no such collection in mongo db %s", repository.StocksTable))
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

	stocks := make([]baseobjects.StocksObject, 0)
	if err := cur.All(context.Background(), &stocks); err != nil {
		return nil, err
	}
	return stocks, nil
}
