package services

import (
	"context"
	"log"
	"os"
	"time"
	"go.mongodb.org/mongo-driver/bson"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/sofc-t/task_manager/task6/models"
	"errors"
)

var (
	mongo_uri, db_mongo_name string 
	client *mongo.Client
)


func CreateDatabaseInstance() *mongo.Client{
	err := godotenv.Load(".env")
	if err != nil{
		log.Fatal("Couldn't load .env")

	}

	mongo_uri = os.Getenv("db_mongo_uri")
	db_mongo_name = os.Getenv("db_mongo_name")
	if mongo_uri == "" || db_mongo_name == ""{
		log.Fatal("Couldn't find mongo db uri or db_mongo_name in .env ")
	}

	clientOptions := options.Client().ApplyURI(mongo_uri)
	ctx, cancel := context.WithTimeout(context.Background(), 15 * time.Second)


	client, err := mongo.Connect(ctx, clientOptions)
	defer cancel()

	if err != nil{
		log.Fatal("Couldn't create ")
	}

	return client
}

func GetDatabaseCollection(colleciton_name string) *mongo.Collection{
	if client == nil || db_mongo_name == ""{
		log.Fatal("Trying to access a collection before connecting to database")
	}

	var collection = client.Database(db_mongo_name).Collection(colleciton_name)
	return collection
}


func FetchTasks() ([]models.Task, error ){
	if client == nil || db_mongo_name == ""{
		log.Fatal("Trying to access a collection before connecting to database")
	}
	client := client.Database("task_manager").Collection("tasks")
	cursor, err := client.Find(context.TODO(), bson.D{})

	if err != nil{
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var tasks []models.Task

	if err:= cursor.All(context.TODO(), &tasks); err != nil{
		return nil, err
	}

	return tasks, nil
}


func FindTask(id int) (models.Task , error){
	if client == nil || db_mongo_name == ""{
		log.Fatal("Trying to access a collection before connecting to database")
	}
	filter := bson.D{{Key: "id" , Value: id}}
	client := client.Database("task_manager").Collection("tasks")
	var task models.Task
	err := client.FindOne(context.TODO(), filter).Decode(&task)

	if err != nil{
		return task, errors.New("failed to load Data")
	}

	return task ,nil
}


func UpdateTask(id int, title string) (models.Task, error){
	if client == nil || db_mongo_name == ""{
		log.Fatal("Trying to access a collection before connecting to database")
	}
	client := client.Database("task_manager").Collection("tasks")
	filter := bson.D{{Key: "id" , Value: id}}
	update := bson.D{{ Key: "$set", Value: bson.D{{Key: "title", Value: title}}}}

	_, err := client.UpdateOne(context.TODO(), filter, update)
	task := models.Task{Id: id, Title: title}

	if err != nil{
		return task, errors.New("failed to load Data")
	}
	return task, nil
}

func DeleteTask(id int) (error){
	if client == nil || db_mongo_name == ""{
		log.Fatal("Trying to access a collection before connecting to database")
	}
	client := client.Database("task_manager").Collection("tasks")
	filter := bson.D{{Key: "id" , Value: id}}

	_, err := client.DeleteOne(context.TODO(), filter)
	if err != nil{
		return  errors.New("failed to load Data")
	}
	return nil

}


func InsertTask(id int, title string) (models.Task , error){
	if client == nil || db_mongo_name == ""{
		log.Fatal("Trying to access a collection before connecting to database")
	}
	client := client.Database("task_manager").Collection("tasks")
	task := models.Task{Id: id, Title: title}

	_, err := client.InsertOne(context.TODO(), task)
	if err != nil{
		return task, errors.New("failed to load Data")
	}

	return task, nil
	

}

