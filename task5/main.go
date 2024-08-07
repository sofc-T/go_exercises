package main

import (
	"github.com/sofc-t/task_manager/router"
	"github.com/sofc-t/task_manager/services"
)

func main(){	
	services.SetUpDataBase("mongodb://localhost:27017")
	router := router.SetUpRouoter()
	router.Run(":8080")
		
}
