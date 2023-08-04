package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func m10(c *gin.Context) {
	fmt.Println("--全局中间件")
	c.Set("name", "张三")
	c.Set("user", User{
		"张三",
		12,
	})
}

func main() {

	r := gin.Default()
	r.Use(m10)
	r.GET("/m10", func(c *gin.Context) {
		name, _ := c.Get("name")
		fmt.Println(name)

		_user, _ := c.Get("user")
		user, ok := _user.(User) // 断言是不是同一个类型
		fmt.Println(user, ok)

		c.JSON(200, gin.H{"msg": "响应成功"})
	})
	r.Run(":8080")
}
