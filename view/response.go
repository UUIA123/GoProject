package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func _string(context *gin.Context) {
	context.String(http.StatusOK, "你好！！")
}
func _json(context *gin.Context) {

	// json 响应结构体
	//type UserInfo struct {
	//	UserName string `json:"user_name"`
	//	Age      int    `json:"age"`
	//	PassWord string `json:"-"` // 不进行JSON序列化
	//}
	//user := UserInfo{"张三", 2, "123122"}
	//context.JSON(http.StatusOK, user)

	// JSON响应Map
	//userMap := map[string]string{
	//	"user_name": "hzangsan",
	//	"age":       "123",
	//}
	//context.JSON(http.StatusOK, userMap)

	// JSON响应gin.H{}
	context.JSON(http.StatusOK, gin.H{"user_name": "张三", "age": "12"})
}
func _xml(c *gin.Context) {
	c.XML(http.StatusOK, gin.H{"user": "User标签text", "message": "heu", "http_status": http.StatusOK})
}
func _yaml(c *gin.Context) {
	c.YAML(http.StatusOK, gin.H{"user": "User标签text", "message": "heu", "http_status": http.StatusOK})
}
func _html(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{"msg": "go后端传的数据"})
}

// 重定向
func _redirect(c *gin.Context) {
	// 301 永久重定向
	// 302 临时重定向
	c.Redirect(301, "https://www.baidu.com")
	//c.Redirect(301, "/html")
}

func main() {
	router := gin.Default()
	// 加载templates下的所有文件
	router.LoadHTMLGlob("templates/*")
	// 在goland中没有相对路径而言，只有相对于项目路径
	//router.StaticFile("/static/img/image", "./static/img/image.png") //指定允许访问文件 访问路径 + 文件路径
	router.StaticFS("/static/img", http.Dir("./static/img")) //指定访问文件夹 访问前缀 + 文件夹路径
	//
	router.GET("/hello", _string)
	router.GET("/json", _json)
	router.GET("/xml", _xml)
	router.GET("/yaml", _yaml)
	router.GET("/html", _html)
	router.GET("/baidu", _redirect)
	router.Run(":8080")
}
