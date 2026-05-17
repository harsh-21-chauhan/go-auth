package main

import (

    "github.com/harsh-21-chauhan/go-auth/config"
	"github.com/gin-gonic/gin"
	"github.com/harsh-21-chauhan/go-auth/routes"
)

func main(){

	 config.ConnectDB()

	 
	r := gin.Default()



	routes.SetupRoutes(r)

	r.Run(":8080")
	

}