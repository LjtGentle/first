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
	//更新
	r.PUT("/update/amount",handler.UpdateAount)
	r.PUT("/update/status",handler.UpdateStatus)

	return r
}



//用于上传or下载
func Router02() http.Handler {
	r := gin.Default()
	r.POST("/upload/:username",handler.UpLoad)
	r.GET("/download/:id",handler.DownLoad)

	r.GET("/downloads",handler.DownLoadAll)
	return r
}