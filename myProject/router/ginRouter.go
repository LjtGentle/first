package router

import (
	"github.com/gin-gonic/gin"
	"ljtTest/myProject/handler"
	"net/http"
)

/*
GET 	查询
POST 	创建
PUT 	更新
DELETE 	删除
 */
func Router01()http.Handler {
	r := gin.Default()
	//创建金主，form表单
	r.POST("/create",handler.CreateDemo)
	r.PUT("/update/amount",handler.UpdateAount)
	return r
}



//用于上传or下载
func Router02() http.Handler {
	r := gin.Default()
	r.GET("/upload/:username",handler.UpLoad)

	return r
}