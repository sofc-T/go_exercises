package services

import (
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"context"

	"errors"

	"github.com/sofc-t/task_manager/models"
)
var db *mongo.Client



func SetUpDataBase(uri string)  (*mongo.Client, error){
	clientOption := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil{
		log.Fatal("Couldn't connect to database")
		return nil, err
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil{
		log.Fatal("Couldn't connect to database")
		return nil,  err
	}
	db = client
	return client, nil
}


func FetchTasks() ([]models.Task, error ){
	client := db.Database("task_manager").Collection("tasks")
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
	filter := bson.D{{Key: "id" , Value: id}}
	client := db.Database("task_manager").Collection("tasks")
	var task models.Task
	err := client.FindOne(context.TODO(), filter).Decode(&task)

	if err != nil{
		return task, errors.New("failed to load Data")
	}

	return task ,nil
}


func UpdateTask(id int, title string) (models.Task, error){
	client := db.Database("task_manager").Collection("tasks")
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
	client := db.Database("task_manager").Collection("tasks")
	filter := bson.D{{Key: "id" , Value: id}}

	_, err := client.DeleteOne(context.TODO(), filter)
	if err != nil{
		return  errors.New("failed to load Data")
	}
	return nil

}


func InsertTask(id int, title string) (models.Task , error){
	client := db.Database("task_manager").Collection("tasks")
	task := models.Task{Id: id, Title: title}

	_, err := client.InsertOne(context.TODO(), task)
	if err != nil{
		return task, errors.New("failed to load Data")
	}

	return task, nil
	

}

