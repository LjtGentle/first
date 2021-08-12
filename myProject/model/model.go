package model

import (
	"github.com/jinzhu/gorm"
)

type DemoOrder struct {
	gorm.Model
	OrderNo string  `gorm:"unique;not null"`//订单号 唯一且不为空
	UserName string `gorm:"unique;not null"`//用户名 唯一且不为空
	Amount float64  `gorm:"default:0.0"`//金额 默认值为0.0
	Status string   //状态
	FileUrl string  //文件存放的路径
}