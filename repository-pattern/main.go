package main

import (
	"github.com/sayed-imran/go-design-pattern/config"
	"github.com/sayed-imran/go-design-pattern/db"
	"github.com/sayed-imran/go-design-pattern/handlers"
)

func main() {
	appConfig := config.NewConfig()
	mongoRepo := db.CreateNewMongoDbRepo(appConfig)
	repo := handlers.NewRepository(mongoRepo)
	handlers.NewHandler(repo)
}
