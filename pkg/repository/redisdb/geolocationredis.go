package redisdb

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/gopherbara/go-apicaller/pkg/models/baseobjects"
	"github.com/gopherbara/go-apicaller/pkg/repository"
	"strings"
)

type GeoLocationRedis struct {
	RedisDB *redis.Client
}

func NewRedisGeoLocation(db *redis.Client) *GeoLocationRedis {
	return &GeoLocationRedis{RedisDB: db}
}

func (g GeoLocationRedis) SaveGeoLocation(geo baseobjects.GeoLocationObject) error {
	json, err := json.Marshal(geo)
	if err != nil {
		return err
	}
	date := geo.DateCall
	if i := strings.Index(geo.DateCall, " "); i >= 0 {
		date = geo.DateCall[:i]
	}
	g.RedisDB.Set(fmt.Sprintf("%s_%s", repository.GeoLocationTable, date), json, 0)

	return nil
}

func (g GeoLocationRedis) GetGeoLocationByDate(date string) ([]baseobjects.GeoLocationObject, error) {
	geos := make([]baseobjects.GeoLocationObject, 0)

	if i := strings.Index(date, " "); i >= 0 {
		date = date[:i]
	}
	info, err := g.RedisDB.Get(fmt.Sprintf("%s_%s", repository.GeoLocationTable, date)).Result()
	if err == redis.Nil {
		return geos, errors.New("no such key")
	} else if err != nil {
		return geos, err
	}

	var geo baseobjects.GeoLocationObject
	err = json.Unmarshal([]byte(info), &geos)
	if err != nil {
		return nil, err
	}
	geos = append(geos, geo)
	return geos, nil
}
