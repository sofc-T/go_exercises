package services

import (
	"context"
	"log"
	"os"
	"time"	
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"	
)

var (
	Mongo_uri, Db_mongo_name string 
	Client *mongo.Client
)

var (
	admin  = "admin"
	guest = "user"
	Admin = &admin 
	Guest = &guest
)


func CreateDatabaseInstance() *mongo.Client{
	err := godotenv.Load(".env")
	if err != nil{
		log.Fatal("Couldn't load .env")

	}

	Mongo_uri = os.Getenv("db_mongo_uri")
	Db_mongo_name = os.Getenv("db_mongo_name")
	if Mongo_uri == "" || Db_mongo_name == ""{
		log.Fatal("Couldn't find mongo db uri or Db_mongo_name in .env ")
	}

	clientOptions := options.Client().ApplyURI(Mongo_uri)
	ctx, cancel := context.WithTimeout(context.Background(), 15 * time.Second)


	Client, err := mongo.Connect(ctx, clientOptions)
	defer cancel()

	if err != nil{
		log.Fatal("Couldn't create ")
	}

	return Client
}

func GetDatabaseCollection(colleciton_name string) *mongo.Collection{
	if Client == nil || Db_mongo_name == ""{
		log.Fatal("Trying to access a collection before connecting to database")
	}

	var collection = Client.Database(Db_mongo_name).Collection(colleciton_name)
	return collection
}


