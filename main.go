package main

import (
	"fmt"
	"gohub/bootstrap"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	// new 一个 Gin Engine 实例
	r := gin.New()

	bootstrap.SetupRoute(r)

	r.GET("/", func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{
			"HELLO": "WORLD",
		})
	})

	// 运行服务
	err := r.Run(":3000")
	if err != nil {
		fmt.Println(err.Error())
	}
}
