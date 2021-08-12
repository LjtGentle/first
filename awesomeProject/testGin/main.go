package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func get(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hello goland----get",
	})
}

func post(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "hello goland --- post",
	})
}

func put(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "hello goland --- put",
	})
}
func mydelete(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "hello goland --- mydelete",
	})
}

/*
RESTful API风格
GET用来获取资源
POST用来新建资源
PUT用来更新资源
DELETE用来删除资源
*/

func main() {
	//创建一个默认的路由引擎
	router := gin.Default()
	//1.定义模版
	//2.解析模版
	router.LoadHTMLGlob("./*")
	router.GET("/fqp", func(c *gin.Context) {

		var msg struct {
			Name    string
			Age     int
			Message string
		}
		msg.Name = "Gentle"
		msg.Age = 23
		msg.Message = "好靓仔"
		c.HTML(http.StatusOK, "fqb.hmpl", gin.H{
			"Name":    msg.Name,
			"Age":     msg.Age,
			"Message": msg.Message,
		})
	})
	//当客户端以GET方式请求路径时，执行
	router.GET("/get", get)
	router.POST("/post", post)
	router.PUT("/put", put)
	router.DELETE("/delete", mydelete)
	//启动HTTP服务器
	//router.Run()//默认8080
	router.Run(":9000") //指定9000端口

}
