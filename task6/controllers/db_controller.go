package controllers

import  "github.com/sofc-t/task_manager/task6/services"


func SetUpDataBase() {
	services.CreateDatabaseInstance()
}