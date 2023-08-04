package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// 查询参数
func _query(c *gin.Context) {
	println(c.Query("user"))      // 直接查URL后面的参数
	println(c.GetQuery("user"))   // 判读是否传参
	println(c.QueryArray("user")) // 返回多个相同的查询参数
	c.DefaultQuery("add", "gd")
	// url?user=zhansan&user=wangwu
}

// 动态参数
func _param(c *gin.Context) {
	println(c.Param("user_id"))
	println(c.Param("book_id"))
}

// 表单参数
func _form(c *gin.Context) {
	println(c.PostForm("name"))
	println(c.PostFormArray("name"))
	println(c.DefaultPostForm("addr", "广东")) // 如果没有参数，则使用默认参数
	form, err := c.MultipartForm()           // 能够接收所有string类型的参数
	fmt.Println(form, err)

}

func _bindJson(c *gin.Context, obj any) (err error) {
	body, _ := c.GetRawData()
	contentType := c.GetHeader("Content-Type")
	switch contentType {
	case "application/json":
		err := json.Unmarshal(body, &obj)
		if err != nil {
			fmt.Println(err.Error())
			return err
		}
	}
	return nil
}

// 原始参数
func _raw(c *gin.Context) {
	data, _ := c.GetRawData()
	var user User
	err := _bindJson(c, &user)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(data)
}
func main() {
	r := gin.Default()
	r.GET("/query", _query)
	r.GET("/param/:user_id", _param) // 自动选择匹配的
	r.GET("/param/:user_id/:book_id", _param)
	r.POST("/form", _form)
	r.POST("/raw", _raw)
	r.Run(":8080")
}
