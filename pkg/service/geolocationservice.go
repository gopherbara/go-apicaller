package service

import (
	"github.com/vv-projects/go-apicaller/pkg/models/baseobjects"
	"github.com/vv-projects/go-apicaller/pkg/repository"
	"log"
	"math/rand"
)

type GeoLocationService struct {
	mongo repository.GeoLocation
	redis repository.GeoLocation
}

func NewGeoLocationService(mongorepos repository.GeoLocation, redisrepos repository.GeoLocation) *GeoLocationService {
	return &GeoLocationService{
		mongo: mongorepos,
		redis: redisrepos,
	}
}

func (g GeoLocationService) SaveGeoLocation(geo baseobjects.GeoLocationObject) bool {
	errRedis := g.redis.SaveGeoLocation(geo)
	if errRedis != nil {
		log.Printf("can`t create geo for redis: %s", errRedis)
		// bad that can`t save to redisdb, but not critical
	}

	err := g.mongo.SaveGeoLocation(geo)
	if err != nil {
		log.Printf("can`t create geo for mongo: %s", err)
		return false
	}
	return true
}

func (g GeoLocationService) GetGeoLocationByDate(date string) (baseobjects.GeoLocationObject, bool) {
	geosRedis, err := g.redis.GetGeoLocationByDate(date)
	if err == nil && len(geosRedis) > 0 {
		return geosRedis[0], true
	}
	geosMongo, err := g.mongo.GetGeoLocationByDate(date)
	if err != nil || len(geosMongo) == 0 {
		log.Printf("can`t get geos from db, cause : %s", err)
		return baseobjects.GeoLocationObject{}, false
	}
	geoRand := rand.Intn(len(geosMongo))

	return geosMongo[geoRand], true
}
