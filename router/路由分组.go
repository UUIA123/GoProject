package main

import "github.com/gin-gonic/gin"

type UserInfo struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
type Response struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

func UserList(c *gin.Context) {
	var userList []UserInfo = []UserInfo{
		{"张三", 12},
		{"赵四", 23},
		{"王五", 24},
	}
	c.JSON(200, Response{0, userList, "请求成功"})
}

// UserGroup 将路由分组打包成函数
func UserGroup(r *gin.RouterGroup) {
	user := r.Group("123")
	{
		user.GET("/user", UserList)
	}
}

func main() {

	r := gin.Default()
	// 路由分组  访问必须加"/api"
	api := r.Group("api")
	{
		api.GET("/", UserList)
	}
	UserGroup(api)

	// api还可以嵌套分组
	// abc := api.Group("abc")

	r.Run(":8080")
}
