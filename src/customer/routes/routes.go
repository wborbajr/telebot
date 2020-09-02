package routes

import (
	"strconv"
	"os"
	"log"
	"net/http"
	user "github.com/wborbajr/telebot/src/customer/controllers/user"

	"github.com/gin-gonic/gin"
)

//StartGin function
func StartGin() {

	PORT, err := strconv.Atoi(os.Getenv("APP_PORT"))
    if err != nil {
        PORT = 3030
    }

	router := gin.Default()
	api := router.Group("/api/v1")
	{
		api.GET("/users", user.GetAllUser)
		api.GET("/users/:id", user.GetUser)
		api.PUT("/users/:id", user.UpdateUser)
		api.POST("/users", user.CreateUser)
		api.DELETE("/users/:id", user.DeleteUser)
	}
	
	router.NoRoute(func(c *gin.Context) {
		c.AbortWithStatus(http.StatusNotFound)
	})

	log.Fatal(router.Run(":"+strconv.Itoa(PORT)))
}