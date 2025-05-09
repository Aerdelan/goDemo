package router

import (
	"my-gin-app/controller"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/hello", controller.HelloHandler)
	r.GET("/users", controller.GetUsers)
}
