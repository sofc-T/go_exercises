package models

import (
)


type TaskManagerInterface interface{
	GetAllTasks() []Task 
	GetTask(id int) Task
	UpdateTask(id int, task Task) Task 
	DeleteTask(id int)
	CreateTask(id int, Title string)
}



type Task struct{
	Id int  `Json:"id"`
	Title string  `Json:"title"`
}


