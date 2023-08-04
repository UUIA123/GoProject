package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/index", func(context *gin.Context) {
		context.String(200, "Hello world")
	})

	// 内网0.0.0.0:8080
	router.Run(":8000")

	// 原生HTTP
}
