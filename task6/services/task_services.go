package services

import (
	"context"
	"log"
	"go.mongodb.org/mongo-driver/bson"
	"github.com/sofc-t/task_manager/task6/models"
	"errors"
)



func FetchTasks() ([]models.Task, error ){
	if Client == nil || Db_mongo_name == ""{
		log.Fatal("Trying to access a collection before connecting to database")
	}
	Client := Client.Database("task_manager").Collection("tasks")
	cursor, err := Client.Find(context.TODO(), bson.D{})

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
		log.Fatal("Trying to access a collection before connecting to database")
	}
	filter := bson.D{{Key: "id" , Value: id}}
	Client := Client.Database("task_manager").Collection("tasks")
	var task models.Task
	err := Client.FindOne(context.TODO(), filter).Decode(&task)

	if err != nil{
		return task, errors.New("failed to load Data")
	}

	return task ,nil
}


func UpdateTask(id int, title string) (models.Task, error){
	if Client == nil || Db_mongo_name == ""{
		log.Fatal("Trying to access a collection before connecting to database")
	}
	Client := Client.Database("task_manager").Collection("tasks")
	filter := bson.D{{Key: "id" , Value: id}}
	update := bson.D{{ Key: "$set", Value: bson.D{{Key: "title", Value: title}}}}

	_, err := Client.UpdateOne(context.TODO(), filter, update)
	task := models.Task{Id: id, Title: title}

	if err != nil{
		return task, errors.New("failed to load Data")
	}
	return task, nil
}

func DeleteTask(id int) (error){
	if Client == nil || Db_mongo_name == ""{
		log.Fatal("Trying to access a collection before connecting to database")
	}
	Client := Client.Database("task_manager").Collection("tasks")
	filter := bson.D{{Key: "id" , Value: id}}

	_, err := Client.DeleteOne(context.TODO(), filter)
	if err != nil{
		return  errors.New("failed to load Data")
	}
	return nil

}


func InsertTask(id int, title string) (models.Task , error){
	if Client == nil || Db_mongo_name == ""{
		log.Fatal("Trying to access a collection before connecting to database")
	}
	Client := Client.Database("task_manager").Collection("tasks")
	task := models.Task{Id: id, Title: title}

	_, err := Client.InsertOne(context.TODO(), task)
	if err != nil{
		return task, errors.New("failed to load Data")
	}

	return task, nil
	

}

