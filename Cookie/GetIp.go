package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/ip", func(c *gin.Context) {
		fmt.Println(c.ClientIP())
		c.JSON(200, gin.H{"msg": "请求成功", "ip": c.ClientIP()})
	})
	r.Run(":8080")
}
