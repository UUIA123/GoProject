package main

import "github.com/gin-gonic/gin"

func main() {

	r := gin.Default()

	r.GET("/download", func(c *gin.Context) {
		// 表示文件流，唤醒浏览器下载，一般设置了这个就要设置文件名
		c.Header("Content-Type", "application/octet-stream")
		// 设置文件名
		c.Header("Content-Disposition", "attachment;filename="+"a.png")
		c.File("./upload/13.png")
	})

	r.Run(":8080")
}
