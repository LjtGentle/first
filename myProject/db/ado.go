package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // mysql驱动
	"myProject/model"
)
// JinZhuAdo 用于存放链接数据库的字串
type JinZhuAdo struct {
	ConStr string
}

// JzAdo 定义一个全局的结构体变量
var JzAdo JinZhuAdo

// init 对全局结构体变量的字段赋值
func init() {
	JzAdo.ConStr= "root:123456@(localhost)/TEST?charset=utf8mb4&parseTime=True&loc=Local"
}

// InitTable 自动迁移字段
func (jinZhuAdo *JinZhuAdo) InitTable(value interface{}) (err error) {
	db, err := gorm.Open("mysql",jinZhuAdo.ConStr)
	if err != nil{
		panic(err)
	}
	defer db.Close()

	result := db.Debug().AutoMigrate(value)
	if result.Error != nil {
		err = result.Error
	}
	return
}

// Create 在数据库中新建一个金主的记录
func (jinZhuAdo *JinZhuAdo)Create(value interface{})(err error){
	db, err := gorm.Open("mysql",jinZhuAdo.ConStr)
	if err != nil{
		panic(err)
	}
	defer db.Close()

	result := db.Debug().Create(value)
 	return result.Error
}

// Creates 在数据库中批量新建金主
func (jinZhuAdo *JinZhuAdo)Creates(slice *[]model.DemoOrder)(sum int , err error) {
	db, err := gorm.Open("mysql",jinZhuAdo.ConStr)
	if err != nil{
		panic(err)
	}
	defer db.Close()
	sum = 0
	for index, value := range *slice {
		result := db.Debug().Create(&value)
		if result.Error  != nil {
			err = result.Error
			sum = index  // 出错的时候返回插入出错的slice下标
			return
		}
		sum += int(result.RowsAffected) // 正常的情况记录着插入记录的条数
	}
	return
}

// UpdateByAmout 根据id更新数据库中amount字段
func (jinZhuAdo *JinZhuAdo)UpdateByAmout(id uint, amount float64)(err error) {
	db, err := gorm.Open("mysql",jinZhuAdo.ConStr)
	if err != nil{
		panic(err)
	}

	result := db.Debug().Model(&model.DemoOrder{}).
		Where("id=?",id).Update("amount",amount)
	if result.Error != nil {
		err = result.Error
	}
	return
}

// UpdateByStatus 根据id更新数据库中status字段
func (jinZhuAdo *JinZhuAdo)UpdateByStatus(id uint, status string) (err error) {
	db, err := gorm.Open("mysql",jinZhuAdo.ConStr)
	if err != nil{
		panic(err)
	}

	result := db.Debug().Model(&model.DemoOrder{}).
		Where("id=?",id).Update("status",status)
	if result.Error != nil {
		err = result.Error
	}
	return
}


// UpdateByFileURL 根据id更新数据库中file_url字段
func(jinZhuAdo *JinZhuAdo) UpdateByFileURL(id uint, fileURL string) (err error) {
	db, err := gorm.Open("mysql",jinZhuAdo.ConStr)
	if err != nil{
		panic(err)
	}

	result := db.Debug().Model(&model.DemoOrder{}).
		Where("id=?",id).Update("file_url",fileURL)
	if result.Error != nil {
		err = result.Error
	}
	return
}


// FindByID 根据id查询金主的全部信息
func (jinZhuAdo *JinZhuAdo)FindByID(value interface{})(err error) {
	db, err := gorm.Open("mysql",jinZhuAdo.ConStr)
	if err != nil{
		panic(err)
	}
	defer db.Close()

	result := db.Debug().Limit(1).Find(value) // 根据id查数据
	if result.Error != nil{
		err = result.Error
	}
	return
}

// FindByName 根据UserName查询金主的详细信息
func (jinZhuAdo *JinZhuAdo)FindByName(name string,value interface{})(err error) {
	db, err := gorm.Open("mysql",jinZhuAdo.ConStr)
	if err != nil{
		panic(err)
	}
	defer db.Close()

	result := db.Debug().Where("user_name=?",name).Find(value)
	if result.Error != nil{
		err = result.Error
	}
	return
}

// FindByOrderNo 根据订单号查询金主的详细信息
func (jinZhuAdo *JinZhuAdo)FindByOrderNo(orderNo string,value interface{}) (err error){
	db, err := gorm.Open("mysql",jinZhuAdo.ConStr)
	if err != nil{
		panic(err)
	}
	defer db.Close()

	result := db.Debug().Where("order_no=?",orderNo).Find(value)
	if result.Error != nil{
		err = result.Error
	}
	return
}

// FindAll 批量查询 查询全部金主的所有信息
func (jinZhuAdo *JinZhuAdo)FindAll(values *[]model.DemoOrder)(sum int, err error) {
	db, err := gorm.Open("mysql",jinZhuAdo.ConStr)
	if err != nil{
		panic(err)
	}
	defer db.Close()
	result := db.Debug().Find(&values)
	fmt.Println("FindAll->",values)
	if result.Error != nil{
		err = result.Error
	}else{
		sum = int(result.RowsAffected)
	}
	return
}

// FindAboutCreateTime 模糊查询 按大概的创建时间查询一些金主的全部信息
func (jinZhuAdo *JinZhuAdo)FindAboutCreateTime(demos *[]model.DemoOrder,time string)(err error) {
	db, err := gorm.Open("mysql",jinZhuAdo.ConStr)
	if err != nil{
		panic(err)
	}
	defer db.Close()

	str := "%"+time+"%"
	res := db.Debug().Where("created_at LIKE?",str).Find(&demos)
	err = res.Error
	return
}

// OrderCreateTime 根据创建的时间排序查询一些金主的全部信息
func (jinZhuAdo *JinZhuAdo)OrderCreateTime(demos *[]model.DemoOrder,isDesc bool)(err error ) {
	db, err := gorm.Open("mysql",jinZhuAdo.ConStr)
	if err != nil{
		panic(err)
	}
	defer db.Close()
	if isDesc {
		Result := db.Debug().Order("created_at desc").Find(&demos) //降序
		err = Result.Error
	}else {
		Result := db.Debug().Order("created_at").Find(&demos) //升序
		err = Result.Error
	}
	return
}
// OrderAmount 根据金额排序查询一些金主的全部信息
func (jinZhuAdo *JinZhuAdo)OrderAmount(demos *[]model.DemoOrder,isDesc bool)(err error ) {
	db, err := gorm.Open("mysql",jinZhuAdo.ConStr)
	if err != nil{
		panic(err)
	}
	defer db.Close()
	if isDesc {
		Result := db.Debug().Order("amount desc").Find(&demos) //降序
		err = Result.Error
	}else {
		Result := db.Debug().Order("amount").Find(&demos) //升序
		err = Result.Error
	}
	return
}

// OrderAmountRank 金币排名前几或是后几名的一些金主的全部信息
func(jinZhuAdo *JinZhuAdo) OrderAmountRank(demos *[]model.DemoOrder,limit int, isDesc bool)(err error) {
	db, err := gorm.Open("mysql",jinZhuAdo.ConStr)
	if err != nil{
		panic(err)
	}
	defer db.Close()

	if isDesc {
		Result := db.Debug().Limit(limit).Order("amount desc").Find(&demos) //降序
		err = Result.Error
	}else {
		Result := db.Debug().Limit(limit).Order("amount").Find(&demos) //升序
		err = Result.Error
	}
	return
}




