// filepath: d:\goStudy\my-gin-app-1\main.go
package main

import (
	"my-gin-app/model"
	"my-gin-app/router"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化数据库
	model.InitDB()

	r := gin.Default()

	// 路由定义
	router.SetupRoutes(r)

	// 启动服务
	r.Run(":8060")
}
