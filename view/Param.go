package main

import "github.com/gin-gonic/gin"

/**
// 常用的校验器
// 不能为空，并且不能没有这个字段
required:必填字段，如: binding:"required"
min最小长度，如: binding:"min=5"
max 最大长度，如: binding:"max=10"
len 长度，如: binding:"len=6"
eq等于，如: binding:"eq=3"
ne 不等于，如: binding:"ne=12"
gt 大于，如: binding:"gt=10"
gte 大于等于，如: binding:"gte=10"
lt 小于，如: binding:"lt=10"
lte 小于等于，如: binding:"te=10"

// 字符串的校验
contains=a 包含a的字符
excludes=a 不包含的字符
startswith 字符串前缀
endswith	字符串后缀
eqfield 等于其他字段的值，如: PassWord string binding:"eqfield=ConfirmPassword"
nefield 不等于其他字段的值

// 数组
dive
LikeList string `json:"list" binding:"required,dive,startswith=like"`

// 网络验证
ip
url
uri  // uri的范围比URL的范围要广

// 日期验证
Date string `json:"date" binding:"datetime=2006-01-02 15:04:06"`

忽略字段，如: binding:"-"
*/

// SignUserInfo /
type SignUserInfo struct {
	Name string `json:"name" binding:"min=3,max=6,"`
	Age  int    `json:"age" binding:"lt=30,gt=18"`
	Sex  string `json:"sex" binding:"oneof=man women"`
}

func main() {
	r := gin.Default()

	var user SignUserInfo
	r.POST("/", func(c *gin.Context) {
		err := c.ShouldBind(&user)
		if err != nil {
			c.JSON(100, gin.H{"data": err})
			return
		}
		c.JSON(100, gin.H{"data": user})
	})

	r.Run(":8080")

}
