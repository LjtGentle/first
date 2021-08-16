package router

import (
	"bytes"
	"io"
	"os"
	"path/filepath"

	"encoding/json"

	"github.com/stretchr/testify/assert"
	"ljtTest/myProject/model"
	"mime/multipart"
	"net/http"
	"net/http/httptest"

	"testing"
)

//pass
func Test_CreateDemo(t *testing.T) {
	router := Router01()
	post :=  make(map[string]string)
	post["name"] = "tom22"
	post["no"] = "222456"
	post["status"] = "22"
	post["amount"] = "22.22"
	body := new(bytes.Buffer)
	w1 := multipart.NewWriter(body)
	for k, v := range post {
		w1.WriteField(k,v)
	}
	w1.Close()
	postbyte, err := json.Marshal(post)
	req ,err := http.NewRequest(http.MethodPost,"http://127.0.0.1:8080/create",body)
	if err != nil {
		t.Error("call http.NewRequest err=",err)
		return
	}
	req.Header.Set("Content-Type",w1.FormDataContentType())

	w := httptest.NewRecorder()
	w.Write(postbyte)
	router.ServeHTTP(w,req)
	assert.Equal(t,http.StatusOK,w.Code)
	assert.Equal(t,"{\"amount\":\"22.22\"," +
		"\"name\":\"tom22\",\"no\":\"222456\",\"status\":\"22\"}{\"message\":\"create jinzhu success\"}",w.Body.String())
}

//pass
func Test_UpdateAount(t *testing.T) {
	router := Router01()
	post :=  make(map[string]string)
	post["id"] = "1"
	post["amount"] ="16"

	body := new(bytes.Buffer)
	w1 := multipart.NewWriter(body)
	for k, v := range post {
		w1.WriteField(k,v)
	}
	w1.Close()

	req ,err := http.NewRequest(http.MethodPut,"http://127.0.0.1:8080/update/amount",body)
	if err != nil {
		t.Error("call http.NewRequest err=",err)
		return
	}
	req.Header.Set("Content-Type",w1.FormDataContentType())

	w := httptest.NewRecorder()
	//postByte ,_:= json.Marshal(&post)
	//w.Write(postByte)
	router.ServeHTTP(w,req)
	assert.Equal(t,http.StatusOK,w.Code)

	var response map[string]string
	err = json.Unmarshal([]byte(w.Body.String()), &response)
	if err != nil {
		t.Error("call json.Unmarshal err=",err)
		t.Error("w.Body.String = ",w.Body.String())
		return
	}
	//t.Errorf("response=%#v\n",response)
	value, exists := response["message"]
	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Equal(t, "update aount success", value)

}

//测试更新状态的路由 pass
func Test_UpdateStatus(t *testing.T) {
	router := Router01()
	//准备map --form表单
	post := make(map[string]string)
	post["id"] = "1"
	post["status"] = "111"

	//准备body
	body := new(bytes.Buffer)
	w1 := multipart.NewWriter(body)
	for k, v := range post {
		w1.WriteField(k,v)
	}
	w1.Close()

	//发出请求
	req ,err := http.NewRequest(http.MethodPut,"http://127.0.0.1:8080/update/status",body)
	if err != nil {
		t.Error("call http.NewRequest err=",err)
		return
	}
   //指定Content-Type
	req.Header.Set("Content-Type",w1.FormDataContentType())

	w := httptest.NewRecorder()
	router.ServeHTTP(w,req)
	//判断返回的错误码
	assert.Equal(t,http.StatusOK,w.Code)
	var response map[string]string
	err = json.Unmarshal([]byte(w.Body.String()), &response)
	if err != nil {
		t.Error("call json.Unmarshal err=",err)
		t.Error("w.Body.String = ",w.Body.String())
		return
	}
	value, exists := response["Message"]
	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Equal(t, "update status success", value)
}

//测试ShowJinZhuByTime pass
func Test_ShowJinZhuByTime(t *testing.T) {
	router := Router01()
	//准备map --form表单
	post := make(map[string]string)
	post["time"] = "2021-08-13"

	//准备body
	body := new(bytes.Buffer)
	w1 := multipart.NewWriter(body)
	for k, v := range post {
		w1.WriteField(k,v)
	}
	w1.Close()

	//发出请求
	req ,err := http.NewRequest(http.MethodGet,"http://127.0.0.1:8080/show/times",body)
	if err != nil {
		t.Error("call http.NewRequest err=",err)
		return
	}
	//指定Content-Type
	req.Header.Set("Content-Type",w1.FormDataContentType())

	w := httptest.NewRecorder()
	router.ServeHTTP(w,req)
	//判断返回的错误码
	assert.Equal(t,http.StatusOK,w.Code)
	var response map[string]model.DemoOrder
	err = json.Unmarshal([]byte(w.Body.String()), &response)
	if err != nil {
		t.Error("call json.Unmarshal err=",err)
		t.Error("w.Body.String = ",w.Body.String())
		return
	}

	// 处理返回值
	for _,v := range response {
		//time 转string
		assert.Contains(t,v.CreatedAt.Format("2006-01-02 15:04:05"),"2021-08-13")

	}

}

//路由根据金钱排序  pass
func Test_ShowJinZhuByMoney(t *testing.T) {
	router := Router01()
	//准备map --form表单
	post := make(map[string]string)
	post["flag"] = "desc" //降序

	//准备body
	body := new(bytes.Buffer)
	w1 := multipart.NewWriter(body)
	for k, v := range post {
		w1.WriteField(k,v)
	}
	w1.Close()
	//发出请求
	req ,err := http.NewRequest(http.MethodGet,"http://127.0.0.1:8080/show/moneny",body)
	if err != nil {
		t.Error("call http.NewRequest err=",err)
		return
	}

	//指定Content-Type
	req.Header.Set("Content-Type",w1.FormDataContentType())

	w := httptest.NewRecorder()
	router.ServeHTTP(w,req)

	//判断返回的错误码
	assert.Equal(t,http.StatusOK,w.Code)
	var response map[int]model.DemoOrder
	err = json.Unmarshal([]byte(w.Body.String()), &response)
	if err != nil {
		t.Error("call json.Unmarshal err=",err)
		t.Error("w.Body.String = ",w.Body.String())
		return
	}

	//处理返回值
	for i,_ := range response {
		if i > 1 {
			if response[i].Amount > response[i-1].Amount {
				t.Errorf("上一个的Amount为：%#v, 当前的Amount为：%#v\n",response[i-1].Amount,response[i].Amount)
			}
		}

	}
}

//路由根据金额排名 pass
func Test_ShowJinZhuByMoneyRank(t *testing.T) {
	router := Router01()
	//准备map --form表单
	post := make(map[string]string)
	post["flag"] = "desc" //降序
	post["limit"] = "5"

	//准备body
	body := new(bytes.Buffer)
	w1 := multipart.NewWriter(body)
	for k, v := range post {
		w1.WriteField(k,v)
	}
	w1.Close()
	//发出请求
	req ,err := http.NewRequest(http.MethodGet,"http://127.0.0.1:8080/show/rankmoney",body)
	if err != nil {
		t.Error("call http.NewRequest err=",err)
		return
	}

	//指定Content-Type
	req.Header.Set("Content-Type",w1.FormDataContentType())

	w := httptest.NewRecorder()
	router.ServeHTTP(w,req)

	//判断返回的错误码
	assert.Equal(t,http.StatusOK,w.Code)
	var response map[int]model.DemoOrder
	err = json.Unmarshal([]byte(w.Body.String()), &response)
	if err != nil {
		t.Error("call json.Unmarshal err=",err)
		t.Error("w.Body.String = ",w.Body.String())
		return
	}

	//处理返回值
	j := 0
	for i,_ := range response {
		j++
		if i > 1 {
			if response[i].Amount > response[i-1].Amount {
				t.Errorf("上一个的Amount为：%#v, 当前的Amount为：%#v\n",response[i-1].Amount,response[i].Amount)
			}
		}

	}
	if j != 5 {
		t.Errorf("excepted:5, got:%#v\n",j)
	}

}

//路由根据订单查询 pass
func Test_ShowJinZhuByOrderNo(t *testing.T) {
	router := Router01()
	//准备map --form表单
	post := make(map[string]string)
	post["no"] = "01234"

	//准备body
	body := new(bytes.Buffer)
	w1 := multipart.NewWriter(body)
	for k, v := range post {
		w1.WriteField(k,v)
	}
	w1.Close()
	//发出请求
	req ,err := http.NewRequest(http.MethodGet,"http://127.0.0.1:8080/show/orderno",body)
	if err != nil {
		t.Error("call http.NewRequest err=",err)
		return
	}

	//指定Content-Type
	req.Header.Set("Content-Type",w1.FormDataContentType())

	w := httptest.NewRecorder()
	router.ServeHTTP(w,req)

	//判断返回的错误码
	assert.Equal(t,http.StatusOK,w.Code)
	var response map[string]model.DemoOrder
	err = json.Unmarshal([]byte(w.Body.String()), &response)
	if err != nil {
		t.Error("call json.Unmarshal err=",err)
		t.Error("w.Body.String = ",w.Body.String())
		return
	}

	for _, value := range response {
		if value.OrderNo !=  "01234"||value.UserName != "jack2"|| value.Status != "5" || value.Amount !=8888 ||value.ID != 5 {
			t.Errorf("got:%#v\n",value)

		}
	}
}

//路由根据时间排序查询 pass
func Test_ShowJinZhuByOrderTime(t *testing.T) {
	router := Router01()
	//准备map --form表单
	post := make(map[string]string)
	post["flag"] = "desc" //降序

	//准备body
	body := new(bytes.Buffer)
	w1 := multipart.NewWriter(body)
	for k, v := range post {
		w1.WriteField(k,v)
	}
	w1.Close()
	//发出请求
	req ,err := http.NewRequest(http.MethodGet,"http://127.0.0.1:8080/show/time",body)
	if err != nil {
		t.Error("call http.NewRequest err=",err)
		return
	}

	//指定Content-Type
	req.Header.Set("Content-Type",w1.FormDataContentType())

	w := httptest.NewRecorder()
	router.ServeHTTP(w,req)

	//判断返回的错误码
	assert.Equal(t,http.StatusOK,w.Code)
	var response map[int]model.DemoOrder
	err = json.Unmarshal([]byte(w.Body.String()), &response)
	if err != nil {
		t.Error("call json.Unmarshal err=",err)
		t.Error("w.Body.String = ",w.Body.String())
		return
	}

	//处理返回值
	j := 0
	for i,_ := range response {
		j++
		if i > 1 {
			strTime1 :=response[i].CreatedAt.Format("2006-01-02 15:04:05")
			strTime2 := response[i-1].CreatedAt.Format("2006-01-02 15:04:05")
			if  strTime1 > strTime2 {
				t.Errorf("上一个的Amount为：%#v, 当前的Amount为：%#v\n",strTime2,strTime1)
			}
		}

	}
	if j != 12 {
		t.Errorf("excepted:5, got:%#v\n",j)
	}

}

//
func Test_UpLoad(t *testing.T) {
	router := Router02()
	//要上传的文件路径
	path := "/home/weilijie/loading/test.txt"
	file, err := os.Open(path)
	if err != nil {
		t.Error("call os.Open err=",err)
		return
	}
	defer file.Close()
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("Gentle_file",filepath.Base(path))

	if err != nil {
		t.Error("call writer.CreateFormFile err=",err)
	}
	defer writer.Close()
	io.Copy(part,file)



	//发出请求
	req ,err := http.NewRequest(http.MethodPost,"http://127.0.0.1:8080/upload/Gentle",body)
	if err != nil {
		t.Error("call http.NewRequest err=",err)
		return
	}

	w := httptest.NewRecorder()
	router.ServeHTTP(w,req)
	//判断返回的错误码
	assert.Equal(t,http.StatusOK,w.Code)


}
func Test_DownLoad(t *testing.T) {

}
func Test_DownLoadAll(t *testing.T) {

}
