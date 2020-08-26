package routes

import (
	"os"
	"log"
	"net/http"
	user "github.com/wborbajr/telebot/controllers/user"

	"github.com/gin-gonic/gin"
)

//StartGin function
func StartGin() {

	PORT := os.Getenv("APP_PORT")
    if PORT == "" {
        PORT = "3030"
    }

	router := gin.Default()
	api := router.Group("/api/v1")
	{
		api.GET("/users", user.GetAllUser)
		api.POST("/users", user.CreateUser)
		api.GET("/users/:id", user.GetUser)
		api.PUT("/users/:id", user.UpdateUser)
		api.DELETE("/users/:id", user.DeleteUser)
	}
	
	router.NoRoute(func(c *gin.Context) {
		c.AbortWithStatus(http.StatusNotFound)
	})

	log.Fatal(router.Run(":"+PORT))
}