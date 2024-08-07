package controllers

import (
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/sofc-t/task_manager/models"
	"github.com/sofc-t/task_manager/services"
)


func GetAllTasksHandler( t *services.TaskManager, ctx *gin.Context ){
	ctx.IndentedJSON(http.StatusAccepted, t.GetAllTasks())
	
}


func GetTaskHandler(t * services.TaskManager, ctx *gin.Context){	
	idStr := ctx.Param("id")
	if idStr == ""{
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message" : "Invalid ID"})
		return
	}

	id, err := strconv.Atoi(idStr)
        if err != nil {            
            ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid ID"})
            return
        }
	
	task := t.GetTask(id)
	if task == nil{
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message":"Task Doesn't Exist"})
		return 
	}

	ctx.IndentedJSON( http.StatusOK, task)
}

func UpdateTaskHandler(t *services.TaskManager, ctx *gin.Context ){
	var task models.Task 
	if err := ctx.BindJSON(&task); err != nil{
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invaid Task Information"})
		return 
	}
	idStr := ctx.Param("id")
	if idStr == ""{
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message" : "Invalid ID"})
		return
	}

	id, err := strconv.Atoi(idStr)
        if err != nil {            
            ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid ID"})
            return
        }

	result := t.UpdateTask(id, task.Title)
	if result == nil{
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Couldn't Update Task"})
		return 
	}

	ctx.IndentedJSON(http.StatusAccepted,  task)
	

}



func DeleteTaskHandler(t *services.TaskManager, ctx *gin.Context){
	idStr := ctx.Param("id")
	if idStr == ""{
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": ctx.Request.Body})
		return 
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {            
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"id": id, "err":err})
		return
	}

	result := t.DeleteTask(id)
	if result == nil{
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"message":"Could not Delete Task"})
		return 
	}

	ctx.Status(http.StatusNoContent)
	return

}


func CreateTaskHandler(t *services.TaskManager,ctx *gin.Context){
	var task models.Task 
	err := ctx.BindJSON(&task)
	if err != nil{
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message":"Invalid Input"})
		return 
	}

	result := t.CreateTask(task.Id, task.Title)
	ctx.IndentedJSON(http.StatusCreated, result)
	
	
}