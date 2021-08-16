package db

import (
	"fmt"
	_"fmt"
	_"fmt"
	_"github.com/jinzhu/gorm/dialects/mysql" //mysql驱动
	"github.com/jinzhu/gorm"
	"ljtTest/myProject/model"
)


type JinZhuAdo struct {
	ConStr string
}

var JzAdo JinZhuAdo

func init() {
	JzAdo.ConStr= "root:123456@(localhost)/TEST?charset=utf8mb4&parseTime=True&loc=Local"
}

//自动迁移
func (jinZhuAdo *JinZhuAdo) InitTable(value interface{}) (err error) {
	db, err := gorm.Open("mysql",jinZhuAdo.ConStr)
	if err != nil{
		panic(err)
		return
	}
	defer db.Close()

	result := db.Debug().AutoMigrate(value)
	if result.Error != nil {
		err = result.Error
	}
	return
}

//新建一个记录 --test
func (jinZhuAdo *JinZhuAdo)Create(value interface{})(err error){
	db, err := gorm.Open("mysql",jinZhuAdo.ConStr)
	if err != nil{
		panic(err)
		return
	}
	defer db.Close()

	result := db.Debug().Create(value)
 	return result.Error
}

//批量新建 --test
func (jinZhuAdo *JinZhuAdo)Creates(slice *[]model.DemoOrder)(sum int , err error) {
	db, err := gorm.Open("mysql",jinZhuAdo.ConStr)
	if err != nil{
		panic(err)
		return
	}
	defer db.Close()
	sum = 0
	for index, value := range *slice {
		result := db.Debug().Create(&value)
		if result.Error  != nil {
			err = result.Error
			sum = index  //出错的时候返回插入出错的slice下标
			return
		}
		sum += int(result.RowsAffected) //正常的情况记录着插入记录的条数
	}
	return
}
//更新amount --test
func (jinZhuAdo *JinZhuAdo)UpdateByAmout(id uint, amount float64)(err error) {
	db, err := gorm.Open("mysql",jinZhuAdo.ConStr)
	if err != nil{
		panic(err)
		return
	}

	result := db.Debug().Model(&model.DemoOrder{}).
		Where("id=?",id).Update("amount",amount)
	if result.Error != nil {
		err = result.Error
	}
	return
}

//更新status --test
func (jinZhuAdo *JinZhuAdo)UpdateByStatus(id uint, status string) (err error) {
	db, err := gorm.Open("mysql",jinZhuAdo.ConStr)
	if err != nil{
		panic(err)
		return
	}

	result := db.Debug().Model(&model.DemoOrder{}).
		Where("id=?",id).Update("status",status)
	if result.Error != nil {
		err = result.Error
	}
	return
}


//更新file_url --test
func(jinZhuAdo *JinZhuAdo) UpdateByFileUrl(id uint, fileUrl string) (err error) {
	// fmt.Println("come in UpdateByFileUrl")
	// defer fmt.Println("out of UpdateByFileUrl")
	db, err := gorm.Open("mysql",jinZhuAdo.ConStr)
	if err != nil{
		panic(err)
		return
	}

	result := db.Debug().Model(&model.DemoOrder{}).
		Where("id=?",id).Update("file_url",fileUrl)
	if result.Error != nil {
		err = result.Error
	}
	return
}


//根据id查询 --test
func (jinZhuAdo *JinZhuAdo)FindByID(value interface{})(err error) {
	db, err := gorm.Open("mysql",jinZhuAdo.ConStr)
	if err != nil{
		panic(err)
		return
	}
	defer db.Close()

	result := db.Debug().Limit(1).Find(value)//根据id查数据
	if result.Error != nil{
		err = result.Error
	}
	return
}

//根据UserName查询 --test
func (jinZhuAdo *JinZhuAdo)FindByName(name string,value interface{})(err error) {
	db, err := gorm.Open("mysql",jinZhuAdo.ConStr)
	if err != nil{
		panic(err)
		return
	}
	defer db.Close()

	result := db.Debug().Where("user_name=?",name).Find(value)
	if result.Error != nil{
		err = result.Error
	}
	return
}


//根据订单号查询 --test
func (jinZhuAdo *JinZhuAdo)FindByOrderNo(orderNo string,value interface{}) (err error){
	db, err := gorm.Open("mysql",jinZhuAdo.ConStr)
	if err != nil{
		panic(err)
		return
	}
	defer db.Close()

	result := db.Debug().Where("order_no=?",orderNo).Find(value)
	if result.Error != nil{
		err = result.Error
	}
	return
}

//批量查询
//查询全部 --test
func (jinZhuAdo *JinZhuAdo)FindAll(values *[]model.DemoOrder)(sum int, err error) {
	db, err := gorm.Open("mysql",jinZhuAdo.ConStr)
	if err != nil{
		panic(err)
		return
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

//模糊查询
//按大概的创建时间查询 --test
func (jinZhuAdo *JinZhuAdo)FindAboutCreateTime(demos *[]model.DemoOrder,time string)(err error) {
	db, err := gorm.Open("mysql",jinZhuAdo.ConStr)
	if err != nil{
		panic(err)
		return
	}
	defer db.Close()

	str := "%"+time+"%"
	res := db.Debug().Where("created_at LIKE?",str).Find(&demos)
	err = res.Error
	return
}

//条件查询
//根据创建的时间排序查询 --test
func (jinZhuAdo *JinZhuAdo)OrderCreateTime(demos *[]model.DemoOrder,isDesc bool)(err error ) {
	db, err := gorm.Open("mysql",jinZhuAdo.ConStr)
	if err != nil{
		panic(err)
		return
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
//金额排序查询 ---test
func (jinZhuAdo *JinZhuAdo)OrderAmount(demos *[]model.DemoOrder,isDesc bool)(err error ) {
	db, err := gorm.Open("mysql",jinZhuAdo.ConStr)
	if err != nil{
		panic(err)
		return
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

//金币排名前几或是后几名 --test
func(jinZhuAdo *JinZhuAdo) OrderAmountRank(demos *[]model.DemoOrder,limit int, isDesc bool)(err error) {
	db, err := gorm.Open("mysql",jinZhuAdo.ConStr)
	if err != nil{
		panic(err)
		return
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




