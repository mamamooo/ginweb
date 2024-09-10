package main

import "github.com/gin-gonic/gin"

func main() {
	// 创建一个Gin实例
	r := gin.Default()

	// 定义一个简单的路由
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// 运行Gin服务器
	r.Run() // 默认监听在 0.0.0.0:8080
}
