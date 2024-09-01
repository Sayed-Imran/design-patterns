package main

import (
	"github.com/sayed-imran/go-design-pattern/config"
	"github.com/sayed-imran/go-design-pattern/db"
)

func main() {
	appConfig := config.NewConfig()
	mongoRepo := db.CreateNewMongoDbRepo(appConfig)
}
