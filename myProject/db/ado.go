package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // mysql驱动
	"myProject/model"
)

// JinZhuAdo 用于存放链接数据库的字串
type jinZhuDao struct {
	ConStr string
	db     *gorm.DB
}

//go:generate mockgen -source=ado.go -package=db -destination ./ado_mock.go

// 备注：-source后面接的是待mockgo文件路径 -package 生成包名 -destination 生成路径名

type JinZhu interface {
	OpenMysql() error
	CloseMysql()
	InitTable(interface{})  error
	Create( interface{})  error
	Creates( *[]model.DemoOrder) (int, error)
	UpdateByAmount(uint, float64)  error
	UpdateByStatus(uint,  string)  error
	UpdateByFileURL(uint,string) error
	FindByID(interface{}) error
	FindByName(string,  interface{}) error
	FindByOrderNo( string, interface{}) error
	FindAll( *[]model.DemoOrder) (int, error)
	FindAboutCreateTime( *[]model.DemoOrder, string)  error
	OrderCreateTime(*[]model.DemoOrder, bool) error
	OrderAmount(*[]model.DemoOrder, bool) error
	OrderAmountRank( *[]model.DemoOrder,  int,  bool) error
}

// NewJinZhuDao  创建结构体
func NewJinZhuDao()( *jinZhuDao, func()) {
	jz := &jinZhuDao{
		ConStr: "root:123456@(localhost)/TEST?charset=utf8mb4&parseTime=True&loc=Local",
	}
	err :=jz.OpenMysql()
	if err != nil {
		panic(err)
	}
	return jz, jz.CloseMysql
}

// OpenMysql 开启对数据库的链接
func (jz *jinZhuDao) OpenMysql() (err error) {
	jz.db, err = gorm.Open("mysql", jz.ConStr)
	return err
}
// CloseMysql 关闭对数据库的链接
func (jz *jinZhuDao) CloseMysql() {
	jz.db.Close()
}

// InitTable 自动迁移字段
func (jz *jinZhuDao) InitTable(value interface{}) (err error) {
	result := jz.db.Debug().AutoMigrate(value)
	return result.Error
}

// Create 在数据库中新建一个金主的记录
func (jz *jinZhuDao) Create(value interface{}) (err error) {
	result := jz.db.Debug().Create(value)
	return result.Error
}

// Creates 在数据库中批量新建金主
func (jz *jinZhuDao) Creates(slice *[]model.DemoOrder) (sum int, err error) {
	sum = 0
	for index, value := range *slice {
		result := jz.db.Debug().Create(&value)
		if result.Error != nil {
			err = result.Error
			sum = index // 出错的时候返回插入出错的slice下标
			return
		}
		sum += int(result.RowsAffected) // 正常的情况记录着插入记录的条数
	}
	return
}

// UpdateByAmount 根据id更新数据库中amount字段
func (jz *jinZhuDao) UpdateByAmount(id uint, amount float64) error {
	return jz.db.Debug().Model(&model.DemoOrder{}).
		Where("id=?", id).Update("amount", amount).Error
}

// UpdateByStatus 根据id更新数据库中status字段
func (jz *jinZhuDao) UpdateByStatus(id uint, status string) (err error) {
	result := jz.db.Debug().Model(&model.DemoOrder{}).
		Where("id=?", id).Update("status", status)
	err = result.Error
	return
}

// UpdateByFileURL 根据id更新数据库中file_url字段
func (jz *jinZhuDao) UpdateByFileURL(id uint, fileURL string) (err error) {
	result := jz.db.Debug().Model(&model.DemoOrder{}).
		Where("id=?", id).Update("file_url", fileURL)
	err = result.Error
	return
}

// FindByID 根据id查询金主的全部信息
func (jz *jinZhuDao) FindByID(value interface{}) (err error) {
	result := jz.db.Debug().Limit(1).Find(value) // 根据id查数据
	err = result.Error
	return
}

// FindByName 根据UserName查询金主的详细信息
func (jz *jinZhuDao) FindByName(name string, value interface{}) (err error) {
	result := jz.db.Debug().Where("user_name=?", name).Find(value)
	err = result.Error
	return
}

// FindByOrderNo 根据订单号查询金主的详细信息
func (jz *jinZhuDao) FindByOrderNo(orderNo string, value interface{}) (err error) {
	result := jz.db.Debug().Where("order_no=?", orderNo).Find(value)
	err = result.Error
	return
}

// FindAll 批量查询 查询全部金主的所有信息
func (jz *jinZhuDao) FindAll(values *[]model.DemoOrder) (sum int, err error) {
	result := jz.db.Debug().Find(&values)
	err = result.Error
	sum = int(result.RowsAffected)
	return
}

// FindAboutCreateTime 模糊查询 按大概的创建时间查询一些金主的全部信息
func (jz *jinZhuDao) FindAboutCreateTime(demos *[]model.DemoOrder, time string) (err error) {
	str := "%" + time + "%"
	res := jz.db.Debug().Where("created_at LIKE?", str).Find(&demos)
	err = res.Error
	return
}

// OrderCreateTime 根据创建的时间排序查询一些金主的全部信息
func (jz *jinZhuDao) OrderCreateTime(demos *[]model.DemoOrder, isDesc bool) (err error) {
	if isDesc {
		Result := jz.db.Debug().Order("created_at desc").Find(&demos) //降序
		err = Result.Error
	} else {
		Result := jz.db.Debug().Order("created_at").Find(&demos) //升序
		err = Result.Error
	}
	return
}

// OrderAmount 根据金额排序查询一些金主的全部信息
func (jz *jinZhuDao) OrderAmount(demos *[]model.DemoOrder, isDesc bool) (err error) {
	if isDesc {
		Result := jz.db.Debug().Order("amount desc").Find(&demos) //降序
		err = Result.Error
	} else {
		Result := jz.db.Debug().Order("amount").Find(&demos) //升序
		err = Result.Error
	}
	return
}

// OrderAmountRank 金币排名前几或是后几名的一些金主的全部信息
func (jz *jinZhuDao) OrderAmountRank(demos *[]model.DemoOrder, limit int, isDesc bool) (err error) {
	if isDesc {
		Result := jz.db.Debug().Limit(limit).Order("amount desc").Find(&demos) //降序
		err = Result.Error
	} else {
		Result := jz.db.Debug().Limit(limit).Order("amount").Find(&demos) //升序
		err = Result.Error
	}
	return
}
