package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"ljtTest/myProject/db"
	"ljtTest/myProject/model"
	"log"
	"net/http"
	"os"
	"strconv"
)

//有问题的
func UpLoads(c *gin.Context) {
	//从参数中获取
	//name := c.Param("username")
	form, err := c.MultipartForm()
	if err != nil{
		fmt.Println("UpLoads call MultipartForm() err=",err)
		return
	}
	files := form.File["file"]

	for index, file := range files {
		log.Println(file.Filename)
		dst := fmt.Sprintf("./%s_%d",file.Filename,index)
		err = c.SaveUploadedFile(file,dst)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"err message":err.Error(),
			})
			fmt.Println("UpLoads call SaveUploadedFile()err=",err)
			return
		}
	}
	c.JSON(http.StatusOK,gin.H{
		"message":fmt.Sprintf("%dfiles uploaded!",len(files)),
	})
}

func UpLoad(c *gin.Context) {
	//从参数中获取
	name := c.Param("username")

	//要先判断这个username是否在数据库中存在
	var demo model.DemoOrder
	err := db.FindByName(name,&demo)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"message":"use no exist",
		})
		return
	}
	//金主文件夹，没有就创建
	upLoadDir := "/home/loading/"+name
	_,err = os.Stat(upLoadDir)
	if os.IsNotExist(err) {
		os.Mkdir(upLoadDir,os.ModePerm)
	}

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{
			"err message":err.Error(),
		})
		fmt.Println("UpLoad call FormFile() err=",err)
		return
	}
	log.Println(file.Filename)
	upLoadDir = upLoadDir+"/"+"file"
	//修改金主的fileurl
	err = db.UpdateByFileUrl(demo.ID,upLoadDir)
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{
			"errMessage":"update fileurl failed",
		})
		return
	}
	//保存文件
	err = c.SaveUploadedFile(file,upLoadDir)
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{
			"errMessage":err.Error(),
		})
	}else{
		c.JSON(200,gin.H{
			"message":"upload sucess",
		})
	}
}

//form表单传参数创建金主
func CreateDemo(c *gin.Context) {
	name := c.PostForm("name")
	no := c.PostForm("no")
	amount := c.PostForm("amount")
	status := c.PostForm("status")

	f64, err :=	strconv.ParseFloat(amount,64)
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{
			"message":err.Error(),
		})
		return
	}
	demo := model.DemoOrder{
		OrderNo:  no,
		UserName: name,
		Amount:   f64,
		Status:   status,
	}
	 err = db.Create(&demo)
	 if err != nil {
		 c.JSON(http.StatusInternalServerError,gin.H{
			 "message":err.Error(),
		 })
		 return
	 }

}


func UpdateAount(c *gin.Context) {

}