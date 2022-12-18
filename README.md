# go-apicaller
get some info from different apis and save to MongoDB

Try develop app with principles of Clean Code.
Full-stack app using go (backend), js (frontend), mongo and redis (databases).

## Description 

Get usefull info from external apis, save it to two different databases and compare with info for previous period.\
Using websockets for updating UI in real-time, so user get fresh info.\
Save all needed stats to NoSQL database, cause data is unrelated.\
Stats are stored in MongoDB and Redis. If one of them is inaccessible - get stats from another.\
All info is pulled at certain appropriate time intervals depending on api type.

## Use

https://github.com/go-co-op/gocron - for starting handling call to external apis in appropriate time.\
https://github.com/gorilla/websocket - for real-time communication.\
https://go.mongodb.org/mongo-driver/mongo  - driver for using MongoDB.\
https://github.com/go-redis/redis - driver for using Redis.\
https://github.com/joho/godotenv - for setting and use of enviroment variables.\
https://github.com/spf13/viper - for work with config.yml.\

## Appearance
![api-caller](https://user-images.githubusercontent.com/116604417/208270781-e32b245d-5fda-4339-afbe-8cf405d20b7f.png)
