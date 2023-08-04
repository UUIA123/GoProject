package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

// 只有Content-Type为form-data才能上传文件
func main() {
	r := gin.Default()
	r.POST("upload", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		fmt.Println(file.Size / 1024)
		fmt.Println(file.Filename)
		readerFile, _ := file.Open()
		writleFile, _ := os.Create("./upload/13.png")
		defer writleFile.Close()
		n, _ := io.Copy(writleFile, readerFile)
		fmt.Println(n)
		// 读取上传文件

		//data, _ := io.ReadAll(readerFile)
		//fmt.Println(string(data))

		// 保存到指定目录
		// c.SaveUploadedFile(file, "./upload/12.png")
		c.JSON(200, gin.H{"msg": "上传成功"})
	})
	// 上传多文件
	r.POST("uploads", func(c *gin.Context) {
		form, _ := c.MultipartForm()
		files := form.File["upload[]"]
		for _, file := range files {
			c.SaveUploadedFile(file, "./upload/"+file.Filename)
		}
		c.JSON(200, gin.H{"msg": "上传成功"})
	})
	r.Run(":8080")
}
