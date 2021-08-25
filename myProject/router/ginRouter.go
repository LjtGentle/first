package router

import (
	"github.com/gin-gonic/gin"
	"myProject/handler"
	"net/http"
)

/*
GET 	查询
POST 	创建
PUT 	更新
DELETE 	删除
 */

// Router01 服务1
func Router01()http.Handler {
	r := gin.Default()
	//创建金主，form表单
	r.POST("/create",handler.CreateDemo)
	//更新
	r.PUT("/update/amount",handler.UpdateAount)
	r.PUT("/update/status",handler.UpdateStatus)

	showGroup := r.Group("/show")
	{
		showGroup.GET("/times",handler.ShowJinZhuByTime) //模糊时间
		showGroup.GET("/moneny",handler.ShowJinZhuByMoneny)
		showGroup.GET("/rankmoney",handler.ShowJinZhuByMonenyRank)
		showGroup.GET("/orderno",handler.ShowJinZhuByOrderNo)
		showGroup.GET("/time",handler.ShowJinZhuByOrderTime) //创建时间排序
	}
	return r
}



// Router02 服务2 用于上传or下载
func Router02() http.Handler {
	r := gin.Default()
	r.POST("/upload/:username",handler.UpLoad)
	r.GET("/download/:id",handler.DownLoad)

	r.GET("/downloads",handler.DownLoadAll)
	return r
}

