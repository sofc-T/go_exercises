package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sofc-t/task_manager/services"
	"github.com/sofc-t/task_manager/models"
	
)


func GetAllTasksHandler( ctx *gin.Context ){
	tasks, err := services.FetchTasks()
	if err != nil{
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch services"})
	}
	ctx.IndentedJSON(http.StatusAccepted, tasks)
	
}


func GetTaskHandler(ctx *gin.Context){	
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
	
	task, err := services.FindTask(id)

	if err != nil{
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch services"})
	}


	ctx.IndentedJSON( http.StatusOK, task)
}

func UpdateTaskHandler(ctx *gin.Context ){
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

	result, err := services.UpdateTask(id, task.Title)
	if err != nil{
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch services"})
	}


	ctx.IndentedJSON(http.StatusAccepted,  result)
	

}



func DeleteTaskHandler(ctx *gin.Context){
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

	err = services.DeleteTask(id)
	if err != nil{
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"message":"Could not Delete Task"})
		return 
	}

	ctx.Status(http.StatusNoContent)
	return

}


func CreateTaskHandler(ctx *gin.Context){
	var task models.Task 
	err := ctx.BindJSON(&task)
	if err != nil{
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message":"Invalid Input"})
		return 
	}

	result, err := services.InsertTask(task.Id, task.Title)
	if err != nil{
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"message":"Could not Delete Task"})
		return 
	}
	ctx.IndentedJSON(http.StatusCreated, result)
	
	
}