package main

import (
	"github.com/gin-gonic/gin"
)

// gin中的 bind 方便将前端传递的数据与结构体进行参数绑定，以及参数校验

type UserInfo struct {
	Name string `json:"name" form:"name" uri:"name"`
	Age  int    `json:"age" form:"age" uri:"age"`
	Sex  string `json:"sex" form:"sex" uri:"sex"`
}

func main() {

	r := gin.Default()
	// 校验JSON参数
	r.POST("/", func(c *gin.Context) {

		var userInfo UserInfo
		// 校验JSON参数
		err := c.ShouldBindJSON(&userInfo)
		if err != nil {
			c.JSON(200, gin.H{"msg": err})
			return
		}
		c.JSON(200, userInfo)
	})
	// 校验param参数
	r.POST("/query", func(c *gin.Context) {

		var userInfo UserInfo
		//
		err := c.ShouldBindQuery(&userInfo)
		if err != nil {
			c.JSON(200, gin.H{"msg": err})
			return
		}
		c.JSON(200, userInfo)
	})

	// 校验param参数
	r.POST("/uri/:name/:age/:sex", func(c *gin.Context) {

		var userInfo UserInfo
		//
		err := c.ShouldBindUri(&userInfo)
		if err != nil {
			c.JSON(200, gin.H{"msg": err})
			return
		}
		c.JSON(200, userInfo)
	})

	// 绑定form-data
	r.POST("/form", func(c *gin.Context) {
		var userInfo UserInfo
		//
		err := c.ShouldBind(&userInfo)
		if err != nil {
			c.JSON(200, gin.H{"msg": err})
			return
		}
		c.JSON(200, userInfo)
	})

	r.Run(":8080")
}
