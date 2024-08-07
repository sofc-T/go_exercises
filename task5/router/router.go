package router

import (
	"github.com/sofc-t/task_manager/controllers"
	"github.com/gin-gonic/gin"
	
)

func SetUpRouoter() *gin.Engine{	

	router := gin.Default()

	router.GET("/tasks", controllers.GetAllTasksHandler)

	router.GET("/task/:id", controllers.GetTaskHandler)

	router.POST("/task/", controllers.CreateTaskHandler)

	router.DELETE("/task/:id",controllers.DeleteTaskHandler)

	router.PUT("/task/:id", controllers.UpdateTaskHandler)

	
	return router
		
}
