package router

import (
	"github.com/sofc-t/task_manager/task6/controllers"
	"github.com/sofc-t/task_manager/task6/middleware"
	"github.com/gin-gonic/gin"
	
)

func SetUpRouoter() *gin.Engine{	

	r := gin.Default()

    r.POST("/register", controllers.SignUp)
    r.POST("/login", controllers.Login)

    auth := r.Group("/")
    auth.Use(middleware.AuthenticationMiddleware())

    {
        auth.GET("/tasks", controllers.GetAllTasksHandler)                
        auth.GET("/tasks/:id", controllers.GetTaskHandler)            
    }

    
    admin := r.Group("/")
    admin.Use(middleware.AuthenticationandAuthorizeMiddleware())
    
    {
        admin.POST("/tasks", controllers.CreateTaskHandler)               
        admin.PUT("/tasks/:id", controllers.UpdateTaskHandler)            
        admin.DELETE("/tasks/:id", controllers.DeleteTaskHandler)         
        admin.POST("/promote", controllers.PromoteUser)
    }

	
    return r
}


