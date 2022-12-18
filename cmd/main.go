package main

import (
	"context"
	"encoding/json"
	"github.com/go-co-op/gocron"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"github.com/vv-projects/go-apicaller/pkg/handler"
	"github.com/vv-projects/go-apicaller/pkg/repository/mongodb"
	"github.com/vv-projects/go-apicaller/pkg/repository/redisdb"
	"github.com/vv-projects/go-apicaller/pkg/service"
	"github.com/vv-projects/go-apicaller/pkg/socket"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Printf("error loading env variables: %s", err.Error())
	}
	if err := getDBConfig(); err != nil {
		log.Printf("error on initializing db config: %s", err.Error())
	}

}

func main() {
	quit := make(chan os.Signal, 1)

	settings, err := getApiSettings()
	if err != nil {
		log.Printf("error on api settings: %s", err.Error())
	}

	manager := socket.NewManager()
	go manager.Manage()
	setupRoutes(manager)

	redisDB := redisdb.NewRedisDB()
	defer redisDB.Close()
	reposRedis := redisdb.NewRedisRepos(redisDB)

	client, mongoDB, err := mongodb.NewMongoDB()
	defer client.Disconnect(context.Background())
	if err != nil {
		log.Printf("failed to initialize mongo db: %s", err.Error())
	}
	reposMongo := mongodb.NewMongoRepos(mongoDB)

	services := service.NewService(reposMongo, reposRedis)

	quoteHandler := handler.NewQuotesHandler(manager, services, settings)
	stocksHandler := handler.NewStocksHandler(manager, services, settings)
	geoHandler := handler.NewGeoLocationHandler(manager, services, settings)
	cryptoHandler := handler.NewCryptoHandler(manager, services, settings)
	currencyHandler := handler.NewCurrencyHandler(manager, services, settings)
	weatherHandler := handler.NewWeatherHandler(manager, services, settings)

	scheduler := gocron.NewScheduler(time.UTC)
	scheduler.Every(2).Hours().Do(quoteHandler.Handle)
	scheduler.Every(1).Days().Do(stocksHandler.Handle)
	scheduler.Every(1).Days().Do(geoHandler.Handle)
	scheduler.Every(1).Days().Do(cryptoHandler.Handle)
	scheduler.Every(1).Days().Do(currencyHandler.Handle)
	scheduler.Every(1).Days().Do(weatherHandler.Handle)
	scheduler.StartAsync()

	go http.ListenAndServe(":4000", nil)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Println("api caller stopped")
}

func setupRoutes(manager *socket.Manager) {
	go manager.Manage()
	http.HandleFunc("/ws", func(writer http.ResponseWriter, request *http.Request) {
		wsEndpoind(manager, writer, request)
	})
}

func wsEndpoind(manager *socket.Manager, w http.ResponseWriter, r *http.Request) {
	conn, err := socket.Upgrade(w, r)
	if err != nil {
		log.Println(err)
	}
	client := socket.NewClient(manager, conn)
	manager.Register <- client
	go client.Read()

}

func getApiSettings() (map[string]interface{}, error) {
	jsonFile, err := os.Open("configs/apisettings.json")
	byteValue, _ := ioutil.ReadAll(jsonFile)
	apisettings := make(map[string]interface{})
	json.Unmarshal(byteValue, &apisettings)
	return apisettings, err
}

func getDBConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("databaseconfig") // set name of file
	return viper.ReadInConfig()
}
