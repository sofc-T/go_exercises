package services

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/sofc-t/task_manager/task6/models"
	"go.mongodb.org/mongo-driver/bson"
)

type TaskManagerInterface interface{
	GetAllTasks() []models.Task 
	GetTask(id int) models.Task
	UpdateTask(id int, task models.Task) models.Task 
	DeleteTask(id int)
	CreateTask(id int, Title string)
}



func FetchTasks() ([]models.Task, error ){
	if Client == nil || Db_mongo_name == ""{
		log.Fatal("Trying to access a collection before connecting to databasee")
	}
	taskCollection := Client.Database(Db_mongo_name).Collection("tasks")
	cursor, err := taskCollection.Find(context.TODO(), bson.D{})

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
	if Client == nil || Db_mongo_name == ""{
		log.Fatal("Trying to access a collection before connecting to databasee")
	}
	filter := bson.D{{Key: "id" , Value: id}}
	taskCollection := Client.Database(Db_mongo_name).Collection("tasks")
	var task models.Task
	err := taskCollection.FindOne(context.TODO(), filter).Decode(&task)

	if err != nil{
		return task, errors.New("failed to load Data")
	}

	return task ,nil
}


func UpdateTask(id int, title string) (models.Task, error){
	if Client == nil || Db_mongo_name == ""{
		log.Fatal("Trying to access a collection before connecting to databasee")
	}
	taskCollection := Client.Database(Db_mongo_name).Collection("tasks")
	filter := bson.D{{Key: "id" , Value: id}}
	update := bson.D{{ Key: "$set", Value: bson.D{{Key: "title", Value: title}}}}

	_, err := taskCollection.UpdateOne(context.TODO(), filter, update)
	task := models.Task{Id: id, Title: title}

	if err != nil{
		return task, errors.New("failed to load Data")
	}
	return task, nil
}

func DeleteTask(id int) (error){
	if Client == nil || Db_mongo_name == ""{
		log.Fatal("Trying to access a collection before connecting to databasee")
	}
	taskCollection := Client.Database(Db_mongo_name).Collection("tasks")
	filter := bson.D{{Key: "id" , Value: id}}

	_, err := taskCollection.DeleteOne(context.TODO(), filter)
	if err != nil{
		return  errors.New("failed to load Data")
	}
	return nil

}


func InsertTask(tid int, id int, title string) (models.Task , error){
	fmt.Println(Db_mongo_name, Client)
	if Client == nil || Db_mongo_name == ""{
		log.Fatal("Trying to access a collection before connecting to databasee")
	}
	taskCollection := Client.Database(Db_mongo_name).Collection("tasks")
	task := models.Task{Id: tid, TaskId: id, Title: title}

	_, err := taskCollection.InsertOne(context.TODO(), task)
	if err != nil{
		log.Printf("task not creatd")
		return task, errors.New("failed to load Data")
	}

	return task, nil
	

}

