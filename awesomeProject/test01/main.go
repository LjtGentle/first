package main

import (
	_ "database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "log"
	_"time"
)

type Demo_order struct {
	gorm.Model        //将ID 、CreatedAt、UpdataAt、DeleteAt注入
	Order_no   string `gorm:"unique;not null"`
	User_name  string
	Amount     float64
	Status     string
	File_url   string
}

var mysqlArgs = "root:123456@(localhost)/TEST?charset=utf8mb4&parseTime=True&loc=Local"

// func (do *Demo_order) Create() (err error) {
// 	db.Create(do).Update("CreateAt", time.Now())
// }

func main() {
	//打开数据库
	db, err := gorm.Open("mysql", mysqlArgs)
	if err != nil {
		panic(err)
		return
	}
	user := Demo_order{}
	users := make([]Demo_order,10)
	db.First(&user)//查询第一条数据
	fmt.Println(user)
	db.Last(&user)
	fmt.Println(user)
	result2 := db.Find(&user).Where("id=?",1)
	if result2.Error != nil {
		fmt.Println("2:",result2.Error)
		return
	}
	fmt.Println("user=",user)
	result := db.Find(&users)
	if result.Error != nil {
		panic(result.Error)
		return
	}
	fmt.Println("aaa")
	fmt.Println("users=",users)
	for _, value := range users {
		fmt.Println(value)
	}
	fmt.Println("bbb")

}
