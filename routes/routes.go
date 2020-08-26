package routes

import (
	"log"
	"net/http"
	user "github.com/wborbajr/telebot/controllers/user"

	"github.com/gin-gonic/gin"
)

//StartGin function
func StartGin() {
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
	log.Fatal(router.Run(":9090"))
}