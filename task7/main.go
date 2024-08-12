package main

import (
	"context"
	"log"
	"os"
	"time"


	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/sofc-t/task_manager/task7/router"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Couldn't load .env file")
	}

	MongoURI := os.Getenv("db_mongo_uri")
	DbMongoName := os.Getenv("db_mongo_name")
	if MongoURI == "" || DbMongoName == "" {
		log.Fatal("Couldn't find MongoDB URI or DbMongoName in .env")
	}

	clientOptions := options.Client().ApplyURI(MongoURI)
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("Couldn't create MongoDB client")
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}

	r := routers.SetUpRouter(1000*time.Second, *client.Database(DbMongoName))

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("Port not Specified")
	}
	r.Run(":" + port)
}
