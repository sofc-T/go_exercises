package main

import (
	"github.com/sofc-t/task_manager/services"
	"github.com/sofc-t/task_manager/controllers"
	"github.com/sofc-t/task_manager/models"
	"github.com/gin-gonic/gin"
	
)

func main(){
	task_manager := &services.TaskManager{
		Tasks: make([]*models.Task, 0),
		
	}

	router := gin.Default()

	router.GET("/tasks", func (ctx *gin.Context){
		controllers.GetAllTasksHandler(task_manager, ctx)
	})

	router.GET("/task/:id", func(ctx *gin.Context){
		controllers.GetTaskHandler(task_manager, ctx)
	})

	router.POST("/task/", func(ctx *gin.Context){
		controllers.CreateTaskHandler(task_manager, ctx)
	})

	router.DELETE("/task/:id", func(ctx *gin.Context){
		controllers.DeleteTaskHandler(task_manager, ctx)
	})

	router.PUT("/task/:id", func(ctx *gin.Context){
		controllers.UpdateTaskHandler(task_manager, ctx)
	})


	router.Run(":8080")

	



		
}
