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
	"github.com/tealeg/xlsx/v3"
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
	fmt.Println("name=",name)
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
	upLoadDir := "/home/weilijie/loading/"+name
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
	upLoadDir = upLoadDir+"/"+file.Filename
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
			"errMessage save file":err.Error(),
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
			"errMessage":err.Error(),
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

	 }else{
	 	c.JSON(200,gin.H{
	 		"message":"create jinzhu success",
		})
	 }

}

//form表单传入数据aount id
func UpdateAount(c *gin.Context) {
	aount := c.PostForm("aount")
	id := c.PostForm("id")
	aountf64 ,err:= strconv.ParseFloat(aount,64)
	idu64 ,err := strconv.ParseUint(id,10,64)
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{
			"errMessage":err.Error(),
		})
		return
	}

	err = db.UpdateByAmout(uint(idu64),aountf64)
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{
			"errMessage":err.Error(),
		})
	}else{
		c.JSON(http.StatusOK,gin.H{
			"message":"update aount success",
		})
	}
}

//form表单传入数据 id status
func UpdateStatus(c *gin.Context) {
	id := c.PostForm("id")
	status := c.PostForm("status")

	id64 ,err := strconv.ParseUint(id,10,64)
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{
			"errMessage":err.Error(),
		})
		return
	}
	err = db.UpdateByStatus(uint(id64),status)
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{
			"errMessage":err.Error(),
		})
	}else{
		c.JSON(http.StatusOK,gin.H{
			"Message":"update status success",
		})
	}

}

//下载
func DownLoad(c *gin.Context) {
	id := c.Param("id")

	id64 ,err := strconv.ParseUint(id,10,64)
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{
			"errMessage":err.Error(),
		})
		return
	}
	//找到file_url
	demo := model.DemoOrder{}
	demo.ID = uint(id64)
	err = db.FindByID(&demo)
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{
			"errMessage":err.Error(),
		})
	}

	if demo.FileUrl =="" {
		c.JSON(http.StatusBadRequest, gin.H{
			"errMessage":"你没有可下载的文件",
		})
		return
	}
	c.Header("Content-Type", "application/octet-stream")

	c.Header("Content-Disposition", "attachment; filename="+demo.UserName)
	c.Header("Content-Transfer-Encoding", "binary")
	c.File(demo.FileUrl)
}


//下载全部金主的信息
func DownLoadAll(c *gin.Context) {
	filex := xlsx.NewFile()
	sheet,err := filex.AddSheet("金主们的全部信息")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"errMessage": err.Error(),
		})
		return
	}

	//把金主们全部读出来
	//demos := make([]*model.DemoOrder,100)
	var demos []model.DemoOrder
	_, err= db.FindAll(&demos)
	fmt.Println("DownLoadAll->",demos)
	if err !=  nil {
		c.JSON(http.StatusInternalServerError,gin.H{
			"errMessage":err.Error(),
		})
		return
	}
	//创建表头
	row := sheet.AddRow()
	row.SetHeightCM(1)//设置行的高度
	cell := row.AddCell()
	cell.Value = "id"
	cell = row.AddCell()
	cell.Value = "created_at"
	cell = row.AddCell()
	cell.Value = "updated_at"
	cell = row.AddCell()
	cell.Value = "deleted_at"
	cell = row.AddCell()
	cell.Value = "order_no"
	cell = row.AddCell()
	cell.Value = "user_name"
	cell = row.AddCell()
	cell.Value = "amount"
	cell = row.AddCell()
	cell.Value = "status"
	cell = row.AddCell()
	cell.Value = "file_url"

	for _, value := range demos {
		row := sheet.AddRow()
		row.SetHeightCM(1)//设置行的高度
		cell = row.AddCell()
		cell.Value = fmt.Sprintf("%#v",value.ID)
		cell := row.AddCell()
		cell.Value =value.CreatedAt.Format("2006-01-02 15:04:05")
		cell = row.AddCell()
		cell.Value =value.UpdatedAt.Format("2006-01-02 15:04:05")
		cell = row.AddCell()
		if value.DeletedAt != nil {
			cell.Value = value.DeletedAt.Format("2006-01-02 15:04:05")
		}else{
			cell.Value = "NULL"
		}
		cell = row.AddCell()
		cell.Value = value.OrderNo
		cell = row.AddCell()
		cell.Value = value.UserName
		cell = row.AddCell()
		cell.Value =  fmt.Sprintf("%#v",value.Amount)
		cell = row.AddCell()
		cell.Value =  value.Status
		cell = row.AddCell()
		cell.Value = value.FileUrl
	}

	fmt.Println("DownLoadAll=--------------33333")
	//保存到服务器本地
	path := "jinzhus.xlsx"
	err = filex.Save(path)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"errMessage":err.Error(),
		})
		return
	}

	//给用户下载
	c.Header("Content-Type", "application/octet-stream")

	c.Header("Content-Disposition", "attachment; filename="+"message.xlsx")
	c.Header("Content-Transfer-Encoding", "binary")
	c.File(path)
}