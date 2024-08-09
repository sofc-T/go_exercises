package controllers

import (
	// "net/http"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sofc-t/task_manager/task6/models"
	"github.com/sofc-t/task_manager/task6/services"
	
)



func SignUp(ctx *gin.Context) {
	var user *models.User 
	if err := ctx.BindJSON(&user); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message" :  "Invalid Credentials"})
		return
	}


	
	err := services.CreateUser(user)
	if err  != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"message" : "Internal server Error"})
		return

	}

	ctx.IndentedJSON(http.StatusCreated, gin.H{"message": "Signed Up successfully"})
	

}


func Login(ctx *gin.Context) {
	var user *models.User 
	if err := ctx.BindJSON(&user); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message" :  "Invalid Credentials"})
		return
	}


	token, err := services.Login(user)
	if err  != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"message" : err})
		return

	}

	ctx.IndentedJSON(http.StatusCreated, gin.H{"message": "Signed In successfully", "token" : token}, )

}


func GetUseryID(ctx *gin.Context) {
	id := ctx.Param("id")

    if id == "" {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
        return
    }
	user, err := services.FetchUserByID(id)
	if id == "" {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
        return
    }

	ctx.JSON(http.StatusBadRequest, user)


}


func PromoteUser(ctx *gin.Context){
	var req models.PromoteUserRequest
    
    
    if err := ctx.ShouldBindJSON(&req); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

	id := req.ID
    if id == "" {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
        return
    }

	err := services.PromoteUser(id)
	if err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return 
	}

	ctx.JSON(http.StatusAccepted, gin.H{"message" : "User updated successully"})

}