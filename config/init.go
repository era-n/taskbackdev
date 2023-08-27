package config

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Config struct {
	MongoUrl string
	Secret   string
}

func LoadConfig() (Config, error) {
	err := godotenv.Load(".env")
	cfg := Config{
		MongoUrl: os.Getenv("MONGO_URL"),
		Secret:   os.Getenv("SECRET"),
	}
	if err != nil {
		return cfg, err
	}

	return cfg, err
}

func InitMongoConn() *mongo.Client {
	cfg, _ := LoadConfig()

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(cfg.MongoUrl))
	if err != nil {
		panic(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	return client
}
