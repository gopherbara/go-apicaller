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

type GeoLocationMongo struct {
	MongoDB *mongo.Database
}

func NewMongoGeoLocation(db *mongo.Database) *GeoLocationMongo {
	return &GeoLocationMongo{MongoDB: db}
}

func (g GeoLocationMongo) SaveGeoLocation(geo baseobjects.GeoLocationObject) error {
	collection := g.MongoDB.Collection(repository.GeoLocationTable)
	res, insertErr := collection.InsertOne(context.Background(), geo)
	if insertErr != nil {
		return insertErr
	}
	fmt.Println(res)
	return nil
}

func (g GeoLocationMongo) GetGeoLocationByDate(date string) ([]baseobjects.GeoLocationObject, error) {
	collection := g.MongoDB.Collection(repository.GeoLocationTable)
	if collection == nil {
		return nil, errors.New(fmt.Sprintf("no such collection in mongo db %s", repository.GeoLocationTable))
	}
	query := bson.D{
		{"datecall", bson.D{{"$regex", date}}},
	}
	cur, err := collection.Find(context.Background(), query)
	defer cur.Close(context.Background())
	if err != nil {
		return nil, err
	}

	geo := make([]baseobjects.GeoLocationObject, 0)
	if err := cur.All(context.Background(), &geo); err != nil {
		return nil, err
	}
	return geo, nil
}
