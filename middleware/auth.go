package middleware

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/harsh-21-chauhan/go-auth/helpers"
)

func Authenciated()gin.HandlerFunc{

	return func (c *gin.Context){
		
		authHeader := c.GetHeader("Authorization")

		if authHeader == ""{
			c.JSON(http.StatusUnauthorized,gin.H{"error":"invalid authorization header request"})
			c.Abort()
			return
		}

		// Remove the word Bearer from - Bearer <token>

		 tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

		 claims, err := helpers.ValidateToken(tokenStr)
   


		if err!=nil{
			log.Printf("Token validation error: %v", err)
			c.JSON(http.StatusUnauthorized,gin.H{"error":"invalid token"})
			c.Abort()
			return
		}
	
         c.Set("userID", claims["user_id"])
         c.Set("role", claims["role"])

		c.Next()

	
	}
}