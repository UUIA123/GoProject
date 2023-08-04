package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
)

type ArticleModel struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type Response struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

func _bindJson1(c *gin.Context, obj any) (err error) {
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

func _getList(c *gin.Context) {

	articleList := []ArticleModel{
		{"GO语言标题", "Go语言内容"},
		{"Python标题", "Python内容"},
		{"java语言标题", "Java内容"},
	}
	c.JSON(200, Response{0, articleList, "成功"})
}

func _getDetail(c *gin.Context) {
	// 获取params中的id
	fmt.Println(c.Param("id"))
	article := ArticleModel{
		"title", "Content",
	}
	c.JSON(200, Response{200, article, "查询成功"})
}

func _create(c *gin.Context) {
	// 接收前端传入的JSON数据
	var article ArticleModel
	err := _bindJson1(c, &article)
	if err != nil {
		fmt.Println(err)
		return
	}
	c.JSON(200, Response{200, article, "添加成功"})
}

func _update(c *gin.Context) {
	fmt.Println(c.Param("id"))
	// 接收前端传入的JSON数据
	var article ArticleModel
	err := _bindJson1(c, &article)
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(200, Response{200, article, "修改成功"})
}

func _delete(c *gin.Context) {
	fmt.Println(c.Param("id"))
	c.JSON(200, Response{200, map[string]string{}, "删除成功"})
}

func main() {
	r := gin.Default()
	r.GET("articles", _getList)       // 文章列表
	r.GET("articles/:id", _getDetail) // 文章详情
	r.POST("articles", _create)       // 添加文章
	r.PUT("articles/:id", _update)    //修改文章
	r.DELETE("articles/:id", _delete) //删除文章

	r.Run(":8080")
}
