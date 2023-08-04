package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// 自定义go中间件  拦截器
func myHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Set("userSession", "userid-1")
		context.Next() // 放行

		// context.Abort() // 阻止
	}
}

func main() {

	// 创建一个服务
	ginServer := gin.Default()
	//ginServer.Use(favicon.New("./icon.png"))

	// 加载静态页面
	ginServer.LoadHTMLGlob("templates/*")
	// 加载静态资源
	ginServer.Static("/static", "./static")

	// 响应一个页面给前端
	ginServer.GET("/index", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", gin.H{
			"msg": "这是go后端传递的参数",
		})
	})

	// 访问地址
	ginServer.GET("/hello", func(context *gin.Context) {
		context.JSON(200, gin.H{"msg": "hello world"})
	})

	ginServer.POST("/user", func(context *gin.Context) {
		context.JSON(200, gin.H{"msg": "POST hello world"})
	})

	// 后端接收前端参数   restful
	ginServer.GET("/user/info/:userid/:username", myHandler(), func(context *gin.Context) {

		// 取出中间件
		usersession := context.MustGet("userSession").(string)
		log.Println("========>", usersession)
		userid := context.Param("userid")
		username := context.Param("username")
		context.JSON(http.StatusOK, gin.H{
			"userid":   userid,
			"username": username,
		})
	})

	// 前端给后端传递JSON
	ginServer.POST("/json", func(context *gin.Context) {
		data, _ := context.GetRawData()
		var m map[string]interface{}
		_ = json.Unmarshal(data, &m)
		context.JSON(http.StatusOK, m)
	})

	// 支持函数式编程  接收from表单数据
	ginServer.POST("/user/add", func(context *gin.Context) {
		username := context.PostForm("username")
		password := context.PostForm("password")
		context.JSON(http.StatusOK, gin.H{
			"msg":      "ok",
			"username": username,
			"password": password,
		})
	})

	// 路由
	ginServer.GET("/test", func(context *gin.Context) {
		// 重定向 301
		context.Redirect(http.StatusMovedPermanently, "https://www.baidu.com")
	})

	// 404
	ginServer.NoRoute(func(context *gin.Context) {
		context.HTML(http.StatusNotFound, "404.html", nil)
	})

	// 路由组 /user/add
	userGroup := ginServer.Group("/user")
	{
		userGroup.GET("/add")
		userGroup.POST("login")
		userGroup.POST("logout")
	}

	// 服务器端口
	ginServer.Run(":8080")

}
