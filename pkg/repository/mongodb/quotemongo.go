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

type QuoteMongo struct {
	MongoDB *mongo.Database
}

func NewMongoQuote(db *mongo.Database) *QuoteMongo {
	return &QuoteMongo{MongoDB: db}
}

func (q QuoteMongo) SaveQuote(quote baseobjects.QuotesObject) error {
	collection := q.MongoDB.Collection(repository.QuotesTable)
	res, insertErr := collection.InsertOne(context.Background(), quote)
	if insertErr != nil {
		return insertErr
	}
	fmt.Println(res)
	return nil
}

func (q QuoteMongo) GetQuoteByDate(date string) ([]baseobjects.QuotesObject, error) {
	collection := q.MongoDB.Collection(repository.QuotesTable)
	if collection == nil {
		return nil, errors.New(fmt.Sprintf("no such collection in mongo db %s", repository.QuotesTable))
	}
	query := bson.D{
		{"datecall", bson.D{{"$regex", date}}},
	}
	cur, err := collection.Find(context.Background(), query)
	defer cur.Close(context.Background())
	if err != nil {
		return nil, err
	}

	quotes := make([]baseobjects.QuotesObject, 0)
	if err := cur.All(context.Background(), &quotes); err != nil {
		return nil, err
	}
	return quotes, nil
}
