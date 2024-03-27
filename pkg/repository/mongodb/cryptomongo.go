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

type CryptoMongo struct {
	MongoDB *mongo.Database
}

func NewMongoCrypto(db *mongo.Database) *CryptoMongo {
	return &CryptoMongo{MongoDB: db}
}

func (c CryptoMongo) SaveCrypto(crypto baseobjects.CryptoObject) error {
	collection := c.MongoDB.Collection(repository.CryptoTable)
	res, insertErr := collection.InsertOne(context.Background(), crypto)
	if insertErr != nil {
		return insertErr
	}
	fmt.Println(res)
	return nil
}

func (c CryptoMongo) GetCryptoByDate(date string) ([]baseobjects.CryptoObject, error) {
	collection := c.MongoDB.Collection(repository.CryptoTable)
	if collection == nil {
		return nil, errors.New(fmt.Sprintf("no such collection in mongo db %s", repository.CryptoTable))
	}
	query := bson.D{
		{"datecall", bson.D{{"$regex", date}}},
	}
	cur, err := collection.Find(context.Background(), query)
	defer cur.Close(context.Background())
	if err != nil {
		return nil, err
	}

	crypto := make([]baseobjects.CryptoObject, 0)
	if err := cur.All(context.Background(), &crypto); err != nil {
		return nil, err
	}
	return crypto, nil
}

func (c CryptoMongo) GetCryptoByDateApi(date string, apiName string) ([]baseobjects.CryptoObject, error) {
	collection := c.MongoDB.Collection(repository.CryptoTable)
	if collection == nil {
		return nil, errors.New(fmt.Sprintf("no such collection in mongo db %s", repository.CryptoTable))
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

	crypto := make([]baseobjects.CryptoObject, 0)
	if err := cur.All(context.Background(), &crypto); err != nil {
		return nil, err
	}
	return crypto, nil
}
