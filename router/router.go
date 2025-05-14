package router

import (
	"my-gin-app/controller"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*") // 允许所有来源
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true") // 允许携带 Cookie

		// 如果是 OPTIONS 请求，直接返回状态码 204
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func SetupRoutes(r *gin.Engine) {
	// 使用跨域中间件
	r.Use(CORSMiddleware())

	// 定义路由
	r.GET("/hello", controller.HelloHandler)
	r.GET("/users", controller.GetUsers)
	r.GET("/getUser", controller.GetUserByID)
}
