package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

func main() {
	r := gin.Default()
	// 请求头的各种获取方式
	r.GET("/", func(c *gin.Context) {
		// 获取一个请求头 返回切片中的第一个数据
		fmt.Println(c.GetHeader("User-Agent"))
		// 如果使用Get方法 或是 .GetHeader,那么可以不用区分大小写，并且返回第一个value
		fmt.Println("c.Request.Header", c.Request.Header)

		fmt.Println("", c.Request.Header.Get("User-Agent"))
		// 自定义请求头 可以覆盖原有请求头 不区分大小写
	})
	// 识别爬虫
	r.GET("/index", func(c *gin.Context) {
		userAgent := c.GetHeader("User-Agent")
		// 正则匹配
		// 字符串包含匹配
		if strings.Contains(userAgent, "python") {
			// 是爬虫
			c.JSON(0, gin.H{"data": "这是响应给爬虫的数据"})
			return
		}
		c.JSON(0, gin.H{"data": "这是响应给用户的数据"})
	})

	r.GET("res", func(c *gin.Context) {
		c.Header("Token", "这是后端设置的请求头")
		// 设置响应头
		c.JSON(0, gin.H{"data": "看看响应头"})
	})
	r.Run(":8080")
}
