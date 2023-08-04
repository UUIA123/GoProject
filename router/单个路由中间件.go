package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func m1(c *gin.Context) {

	fmt.Println("m1----")
	// c.Abort()// 终止中间件
	c.Next() // 表示后面的内容为响应中中间件
}
func m2(c *gin.Context) {

}
func m3(c *gin.Context) {

}

func main() {

	r := gin.Default()

	r.GET("/", m1, m2, m3)

	r.Run(":8080")
}
