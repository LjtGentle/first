/*
逻辑测试部分
*/
package db

import (
	"github.com/jinzhu/gorm"
	"ljtTest/myProject/model"
	"reflect"
	"testing"
	_"time"
)

//pass 测试了创建和按订单号查询
func Test_createandFindbyOrderNo(t *testing.T) {
	want := model.DemoOrder{
		OrderNo:"811",
		UserName:"Gentle",
		Amount:771.88,
		Status:"0",
	}
	err :=  JzAdo.Create(&want)
	if err != nil {
		t.Errorf("call Create err=%#v",err)
	}

	got := model.DemoOrder{
		Model:    gorm.Model{},
		OrderNo:  "",
		UserName: "",
		Amount:   0,
		Status:   "",
		FileUrl:  "",
	}
	err = JzAdo.FindByOrderNo("811",&got)

	if err != nil {
		t.Errorf("call FindByOrderNo err=%#v\n",err)
	}
	if !reflect.DeepEqual(want.OrderNo,got.OrderNo) {
		t.Errorf("excepted:%#v, got :%v\n",want.OrderNo, got.OrderNo)
	}
	if !reflect.DeepEqual(want.Amount,got.Amount) {
		t.Errorf("excepted:%#v, got :%#v\n",want.Amount,got.OrderNo)
	}
	if !reflect.DeepEqual(want.UserName,got.UserName) {
		t.Errorf("excepted:%#v, got :%#v\n",want.UserName,got.UserName)
	}
	if !reflect.DeepEqual(want.Status,got.Status) {
		t.Errorf("excepted:%#v, got :%#v\n",want.Status,got.Status)
	}
}

//测试批量创建 pass 创建多个和用名字查询
func Test_creates(t *testing.T){
	wants := []model.DemoOrder{
		{
			Status:"3",
			UserName:"Gentle1",
			OrderNo:"1111",
			Amount:111.11,
		},
		{
			Status:"2",
			UserName:"Gentle2",
			OrderNo:"2222",
			Amount:222.22,
		},

		{
			Status:"1",
			UserName:"Gentle3",
			OrderNo:"3333",
			Amount:333.33,
		},
	}
	sum,err := JzAdo.Creates(&wants)
	if err != nil {
		t.Errorf("err=%#v\n",err)
	}
	if sum != 3 {
		t.Errorf("要插入的个数为3，实际插入的个数为：%#v\n",sum)
	}
	if err != nil {
		t.Errorf("call Creates err = %#v\n",err)
	}

	//第一个
	for _,v := range wants {
		got:= model.DemoOrder{}
		err = JzAdo.FindByName(v.UserName,&got)
		if err != nil {
			t.Error("call FindByName err=",err)
		}
		if got.Status != v.Status || got.Amount != v.Amount ||got.OrderNo != v.OrderNo {
			t.Errorf("excepted status :%#v, got status:%#v\n",v.Status,got.Status)
			t.Errorf("excepted Amount :%#v, got Amount:%#v\n",v.Amount,got.Amount)
			t.Errorf("excepted OrderNo :%#v, got OrderNo:%#v\n",v.OrderNo,got.OrderNo)
		}
	}

}

//测试 更新金额和按id查询 pass
func Test_DateByAmountAndFindById(t *testing.T) {
	//更新金额
	err := JzAdo.UpdateByAmout(9,999.99)
	if err != nil {
		t.Error("call UpdateByAmout err=",err)
	}

	//按id查询
	demo := model.DemoOrder{}
	demo.ID = 9
	err = JzAdo.FindByID(&demo)
	if err != nil {
		t.Error("call FindByID err= ",err)
	}

	if demo.Amount != 999.99 {
		t.Errorf("excepted:999.99, got:%#v\n",demo.Amount)
	}
}

//测试更新状态 pass
func Test_updateStatus(t *testing.T) {
	//更状态
	err := JzAdo.UpdateByStatus(9,"99")
	if err != nil {
		t.Error("call UpdateByStatus err=",err)
	}
	//根据Id查询
	demo := model.DemoOrder{}
	demo.ID = 9
	err = JzAdo.FindByID(&demo)
	if err != nil {
		t.Error("call FindByID err=",err)
	}
	if demo.Status != "99" {
		t.Errorf("excepted:99, got:%#v\n",demo.Status)
	}

}


//测试更新fileurl pass
func Test_updateFileUrl(t *testing.T) {
	err := JzAdo.UpdateByFileUrl(1,"/home")
	if err != nil {
		t.Error("call UpdateByFileUrl err=",err)
	}
	demo := model.DemoOrder{}
	demo.ID = 1

	err = JzAdo.FindByID(&demo)
	if  err != nil {
		t.Error("call FindByID err= ",err)
	}
	if demo.FileUrl != "/home" {
		t.Errorf("excepted:/home, got:%#v\n",demo.FileUrl)
	}
}

//测试查找全部 pass
func Test_findAll(t *testing.T) {
	demos := make([]model.DemoOrder,5)
	sum, err := JzAdo.FindAll(&demos)
	if err != nil {
		t.Error("call FindAll err=",err)
	}
	if sum != 9 || len(demos) != 9 {
		t.Errorf("excepted:9,got sum=%#v, len=%#v\n",sum,len(demos))
	}
	for _,v := range demos {
		demo := model.DemoOrder{}
		demo.ID = v.ID
		err = JzAdo.FindByID(&demo)
		if err != nil {
			t.Error("call FindByID err=",err)
		}
		if demo.FileUrl != v.FileUrl || demo.Status != v.Status || demo.Amount != v.Amount ||demo.UserName != v.UserName {
			t.Errorf("expeted fileurl:%#v,got fileurl:%#v\n",demo.FileUrl,v.FileUrl)
			t.Errorf("expeted Status:%#v,got Status:%#v\n",demo.Status,v.Status)
			t.Errorf("expeted Amount:%#v,got Amount:%#v\n",demo.Amount,v.Amount)
			t.Errorf("expeted UserName:%#v,got UserName:%#v\n",demo.UserName,v.UserName)
		}
	}
}

//测试创建的时间模糊查询 pass
func Test_FindAboutCreateTime(t *testing.T) {
	demos := make([]model.DemoOrder,5)
	err := JzAdo.FindAboutCreateTime(&demos,"2021-08-13")
	if err != nil {
		t.Error("call FindAboutCreateTime err =",err)
	}

	if len(demos) != 3 {
		t.Errorf("excepted:3,got:%#v\n",len(demos))
	}
	for i, v := range demos {
		demo := model.DemoOrder{}
		demo.ID = v.ID
		err := JzAdo.FindByID(&demo)
		if err != nil {
			t.Error("call FindByID err =",err)
		}
		if demo.UpdatedAt != v.UpdatedAt {
			t.Errorf("[%#v] excepted :%#v, got: %#v\n",i,demo.UpdatedAt,v.UpdatedAt)
		}
	}
}

//测试OrderCreateTime  pass
func Test_OrderCreateTime(t *testing.T) {
	demos := make([]model.DemoOrder,10)
	//降序
	err:= JzAdo.OrderCreateTime(&demos,true)
	if err != nil {
		t.Error("call OrderCreateTime desc err=",err)
	}
	//遍历比较时间
	for i,_ := range demos {
		if i>1 {
			if  demos[i].CreatedAt.After(demos[i-1].CreatedAt){
				t.Error("创建时间不是降序")
				t.Errorf("--前一个是：%#v,后一个是:%#v\n",demos[i-1].CreatedAt,demos[i].CreatedAt)
			}
		}
	}

	//升序
	err = JzAdo.OrderCreateTime(&demos,false)
	if err != nil {
		t.Error("call OrderCreateTime err=",err)
	}
	for j:=range demos {
		if j>1 {
			if demos[j].CreatedAt.Before(demos[j-1].CreatedAt) {
				t.Error("创建时间不是升序")
				t.Errorf("--前一个是：%#v,后一个是:%#v\n",demos[j-1].CreatedAt,demos[j].CreatedAt)
			}
		}
	}

}

//测试OrderAmount pass
func Test_OrderAmount(t *testing.T) {
	demos := make([]model.DemoOrder,10)
	err := JzAdo.OrderAmount(&demos,true)
	if err != nil {
		t.Error("call OrderAmount err=",err)
	}
	//降序 遍历
	for i,_:= range demos {
		if i>1 {
			if demos[i].Amount>demos[i-1].Amount {
				t.Error("金额不是降序")
				t.Errorf("前一个的金额：%#v,当前的金额为：%#v\n",demos[i-1].Amount,demos[i].Amount)
			}
		}
	}

	//升序
	err = JzAdo.OrderAmount(&demos,false)
	if err != nil {
		t.Error("call OrderAmount err=",err)
	}
	for j,_ := range demos {
		if j > 1 {
			if demos[j].Amount<demos[j-1].Amount {
				t.Error("金额不是降序")
				t.Errorf("前一个的金额：%#v,当前的金额为：%#v\n",demos[j-1].Amount,demos[j].Amount)
			}
		}
	}

}

//测试OrderAmountRank --pass
func Test_OrderAmountRank( t *testing.T) {
	demos := make([]model.DemoOrder,5)
	//降序
	err := JzAdo.OrderAmountRank(&demos,5,true)
	if err != nil {
		t.Error("call OrderAmountRank desc err=",err)
	}
	if len(demos) != 5 {
		t.Errorf("excepted:5,got:%#v\n",len(demos))
	}
	for i,_:=range demos {
		if i >1 {
			if demos[i].Amount>demos[i-1].Amount {
				t.Errorf("前一个为：%#v,当前为:%#v\n",demos[i-1].Amount,demos[i].Amount)
			}
		}
	}
	//升序
	err = JzAdo.OrderAmountRank(&demos,5,false)
	if err != nil {
		t.Error("call OrderAmountRank err=",err)
	}
	if len(demos) != 5 {
		t.Errorf("excepted:5,got:%#v\n",len(demos))
	}
	for i,_:=range demos {
		if i >1 {
			if demos[i].Amount<demos[i-1].Amount {
				t.Errorf("前一个为：%#v,当前为:%#v\n",demos[i-1].Amount,demos[i].Amount)
			}
		}
	}

}

