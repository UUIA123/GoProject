package main

import "github.com/gin-gonic/gin"

type _UserInfo struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
type Res struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

func _UserList(c *gin.Context) {
	var userList []_UserInfo = []_UserInfo{
		{"张三", 12},
		{"赵四", 23},
		{"王五", 24},
	}
	c.JSON(200, Res{0, userList, "请求成功"})
}

// Middleware 中间件可以使用  闭包写法
func Middleware(msg string) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token == "1234" {
			c.Next()
			return
		}
		c.JSON(200, Res{1001, nil, msg})
		c.Abort()
	}
}

// UserGroup 将路由分组打包成函数
func _UserGroup(r *gin.RouterGroup) {
	user := r.Group("123").Use(Middleware("用户校验失败"))
	{
		user.GET("/user", _UserList)
	}
}

func main() {

	r := gin.Default()
	// 路由分组  访问必须加"/api"
	api := r.Group("api")
	{
		api.GET("/", _UserList)
		api.GET("login", Middleware("登录失败"), func(c *gin.Context) {

		})

	}
	_UserGroup(api)

	// api还可以嵌套分组
	// abc := api.Group("abc")

	r.Run(":8080")
}
