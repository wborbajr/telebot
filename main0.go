package main

import (
	"net/http"
	"fmt"
	"log"
	"github.com/gin-gonic/gin"
)

func main()  {
	fmt.Println("TeleBOT v1.0.0")

	router := gin.Default()

	v1 := router.Group("/api/v1/telebot") 
	{
		v1.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "TeleBOT it is working!!!",
			})
		})
	}
	
	// listen and server the router with the port
	log.Fatal(router.Run(":3030"))

}