package config

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

type Config struct {
	Mongo  *mongo.Client
	InProd bool
}


func connectDB(mongoURI string) *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to MongoDB!")
	return client
}

func NewConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}
	mongoURI := os.Getenv("MONGODB_URI")
	inProd := false
	if os.Getenv("IN_PROD") == "true" {
		inProd = true
	}
	if mongoURI == "" {
		log.Fatal("MONGODB_URI environment variable is required")
	}

	return &Config{
		Mongo:  connectDB(mongoURI),
		InProd: inProd,
	}
}
