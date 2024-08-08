package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sofc-t/task_manager/task6/utils"
)


func AuthenticationMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")	
		if token == ""{
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Authorization Required"})
			ctx.Abort()
			return 
		}
		claims, err := Utils.ValidateToken(token)

		if err != nil{
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": "iNVALID tOKEN"})
			ctx.Abort()
			return
		}

		ctx.Set("email", claims.Email)
		ctx.Set("first_name", claims.Name)
		ctx.Set("uid",claims.Uid)
		ctx.Set("role", claims.Role)
		ctx.Next()
	}

}


func AdminMiddleware() gin.HandlerFunc{
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token == ""{
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Authorization Required"})
			ctx.Abort()
			return
		}

		claims, err := Utils.ValidateToken(token)
		if err != nil{
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": "iNVALID tOKEN"})
			ctx.Abort()
			return 
		}

		if claims.Role != "admin"{
			ctx.JSON(http.StatusForbidden, gin.H{"message":"Forbbiden Content "})
			ctx.Abort()
			return
		}
		

		ctx.Next()
	}
}


func AuthenticationandAuthorizeMiddleware() gin.HandlerFunc{
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token == ""{
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Authorization Required"})
			ctx.Abort()
			return
		}

		claims, err := Utils.ValidateToken(token)
		if err != nil{
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": "iNVALID tOKEN"})
			ctx.Abort()
			return 
		}
		if claims.Role != "admin"{
			ctx.JSON(http.StatusForbidden, gin.H{"message":"Forbbiden Content "})
			ctx.Abort()
			return
		}
		ctx.Set("email", claims.Email)
		ctx.Set("first_name", claims.Name)
		ctx.Set("uid",claims.Uid)
		ctx.Set("role", claims.Role)

		ctx.Next()
	}
}