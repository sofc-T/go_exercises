package main

import (
	"github.com/sofc-t/task_manager/task6/router"
	"github.com/sofc-t/task_manager/task6/controllers"
)

func main(){	
	controllers.SetUpDataBase()
	router := router.SetUpRouoter()
	router.Run(":8080")
		
}
