package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {

	// r := gin.Default()
	// r.GET("/", func(ctx *gin.Context) {
	// 	// 注册一个路由
	// 	// 以 JSON 格式响应
	// 	ctx.JSON(http.StatusOK, gin.H{"Hello": "World!"})
	// })

	// r.Run()
	// new 一个 Gin Engine 实例
	r := gin.New()

	// 注册中间件
	r.Use(gin.Logger(), gin.Recovery())
	// 注册一个路由
	r.GET("/", func(c *gin.Context) {

		// 以 JSON 格式响应
		c.JSON(http.StatusOK, gin.H{
			"Hello": "World!",
		})
	})

	r.NoRoute(func(c *gin.Context) {
		// 获取标头信息的 Accept 信息
		acceptString := c.Request.Header.Get("Accept")
		if strings.Contains(acceptString, "text/html") {
			c.String(http.StatusNotFound, "页面返回 404")
		} else {
			c.JSON(http.StatusNotFound, gin.H{
				"error_code":    404,
				"error_message": "路由未定义，请确认 url 和请求方法是否正确。",
			})
		}
	})

	// 运行服务，默认为 8080，我们指定端口为 8000
	r.Run(":8000")

}
