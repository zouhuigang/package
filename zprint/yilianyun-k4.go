/*
对接易联云K4云打印机
*/

package zprint

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	. "github.com/bitly/go-simplejson"
	"github.com/satori/go.uuid"
	"github.com/zouhuigang/package/zcrypto"
	"github.com/zouhuigang/package/ztime"
)

/*
使用说明

第一次运行请先实例化对象（）

*/

type OperationInterface struct {

	//所有的接口操作都在这个结构体下
	Client_id     string
	Client_secret string
	Access_token  string
	Refresh_token string
}

type OurApplication struct {
	/*
		自有应用
		通过继承OperationInterface获得所有的操作方法，自身只有获取token的方法
	*/
	OperationInterface
}

type OpenApplication struct {
	/*
		开放应用
		通过继承OperationInterface获得所有的操作方法，自身只有获取token的方法
		code 需要手工获取
	*/
	OperationInterface
	Code string
}

func (self OurApplication) AddPrinter(machine_code string, msign string) {
	url := "https://open-api.10ss.net/printer/addprinter"
	uid := uuid.Must(uuid.NewV4())
	timestamp := time.Now().Unix()
	fmtStr := fmt.Sprintf("client_id=%s&sign=%s&timestamp=%d&id=%s&access_token=%s&machine_code=%s&msign=%s", self.Client_id, self.Sign(timestamp), timestamp, uid, self.Access_token, machine_code, msign)
	resp := httpPost(fmtStr, url)
	if self.CheckResponseStatus(resp) {
		fmt.Println(resp)
	} else {
		fmt.Println(resp)
	}
}

func (self OperationInterface) Sign(push_time int64) string {
	s := fmt.Sprintf("%s%d%s", self.Client_id, push_time, self.Client_secret)
	sign := zcrypto.Md5(s)
	return sign
}

func httpPost(data string, url string) (response string) {
	resp, err := http.Post(url, "application/x-www-form-urlencoded", strings.NewReader(data))
	if err != nil {
		fmt.Errorf("网络错误: %s", err)
		return "网络错误"
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Errorf("数据解析错误：%s", err)
		return "数据解析错误"
	}
	stats := string(body)
	return stats
}

func (self *OurApplication) GetToken() string {
	uid := uuid.Must(uuid.NewV4())
	timestamp := time.Now().Unix()
	url := "https://open-api.10ss.net/oauth/oauth"
	fmtStr := fmt.Sprintf("client_id=%s&grant_type=client_credentials&scope=all&sign=%s&timestamp=%d&id=%s", self.Client_id, self.Sign(timestamp), timestamp, uid)
	resp := httpPost(fmtStr, url)
	return resp
	/*js, err := NewJson([]byte(resp))
	if err != nil {
		//fmt.Println(err.Error)
		return js, err
	}

	return js, nil*/

	//status := js.Get("error").MustString()
	//if status != "0" {
	//	fmt.Println("error : ", resp)
	//	}

	//fmt.Println("Access_token : ", js.Get("body").Get("Access_token").MustString())
	//fmt.Println("refresh_token : ", js.Get("body").Get("refresh_token").MustString())
}

func (self *OpenApplication) GetToken() {
	uid := uuid.Must(uuid.NewV4())
	timestamp := time.Now().Unix()
	url := "https://open-api.10ss.net/oauth/oauth"
	fmtStr := fmt.Sprintf("client_id=%s&grant_type=authorization_code&scope=all&sign=%s&timestamp=%d&id=%s&code=%s", self.Client_id, self.Sign(timestamp), timestamp, uid, self.Code)
	fmt.Println(fmtStr)
	resp := httpPost(fmtStr, url)
	js, err := NewJson([]byte(resp))

	if err != nil {
		fmt.Println(err.Error)
	}

	status := js.Get("error").MustString()
	if status != "0" {
		fmt.Println("error : ", resp)
	}

	fmt.Println("Access_token : ", js.Get("body").Get("Access_token").MustString())
	fmt.Println("refresh_token : ", js.Get("body").Get("refresh_token").MustString())
}

func (self OperationInterface) CheckResponseStatus(resp string) (stats bool) {
	js, err := NewJson([]byte(resp))
	if err != nil {
		fmt.Errorf("服务器返回非预期的数据格式：", err)
		return false
	}
	status := js.Get("error").MustString()
	if status != "0" {
		fmt.Errorf("状态码异常", err)
		return false
	}
	return true
}

//应用管理平台
//https://dev.10ss.net/admin/listinfo?id=1096322761
//接口文档-http://doc2.10ss.net/372519
func (self OperationInterface) Print(machine_code string, content string) string {
	url := "https://open-api.10ss.net/print/index"
	uid := uuid.Must(uuid.NewV4())
	timestamp := ztime.NowTimeStamp()
	origin_id := fmt.Sprintf("%d", uid)
	origin_id = zcrypto.Md5(origin_id)
	fmtStr := fmt.Sprintf("client_id=%s&access_token=%s&machine_code=%s&content=%s&origin_id=%s&sign=%s&id=%s&timestamp=%d",
		self.Client_id, self.Access_token, machine_code, content,
		origin_id, self.Sign(timestamp), uid, timestamp)
	fmt.Println(fmtStr)

	resp := httpPost(fmtStr, url)
	return resp

	/*if self.CheckResponseStatus(resp) {
		fmt.Println(resp)
	} else {
		fmt.Println(resp)
	}*/
}

func (self OperationInterface) DelPrint(machine_code string) {
	url := "https://open-api.10ss.net/printer/deleteprinter"
	uid := uuid.Must(uuid.NewV4())
	timestamp := time.Now().Unix()
	fmtStr := fmt.Sprintf("client_id=%s&sign=%s&timestamp=%d&id=%s&access_token=%s&machine_code=%s", self.Client_id, self.Sign(timestamp), timestamp, uid, self.Access_token, machine_code)
	resp := httpPost(fmtStr, url)
	if self.CheckResponseStatus(resp) {
		fmt.Println(resp)
	} else {
		fmt.Println(resp)
	}
}

func (self OperationInterface) AddPrintMenu(machine_code string, content string) {
	url := "https://open-api.10ss.net/printmenu/addprintmenu"
	uid := uuid.Must(uuid.NewV4())
	timestamp := time.Now().Unix()
	fmtStr := fmt.Sprintf("client_id=%s&sign=%s&timestamp=%d&id=%s&access_token=%s&machine_code=%s&content=%s", self.Client_id, self.Sign(timestamp), timestamp, uid, self.Access_token, machine_code, content)
	fmt.Println(fmtStr)
	resp := httpPost(fmtStr, url)
	if self.CheckResponseStatus(resp) {
		fmt.Println(resp)
	} else {
		fmt.Println(resp)
	}
}

func (self OperationInterface) ShutdownRestart(machine_code string, response_type string) {
	url := "https://open-api.10ss.net/printer/shutdownrestart"
	uid := uuid.Must(uuid.NewV4())
	timestamp := time.Now().Unix()
	fmtStr := fmt.Sprintf("client_id=%s&sign=%s&timestamp=%d&id=%s&access_token=%s&machine_code=%s&response_type=%s", self.Client_id, self.Sign(timestamp), timestamp, uid, self.Access_token, machine_code, response_type)
	fmt.Println(fmtStr)
	resp := httpPost(fmtStr, url)
	if self.CheckResponseStatus(resp) {
		fmt.Println(resp)
	} else {
		fmt.Println(resp)
	}
}

func (self OperationInterface) SetSound(machine_code string, response_type string, voice string) {
	url := "https://open-api.10ss.net/printer/setsound"
	uid := uuid.Must(uuid.NewV4())
	timestamp := time.Now().Unix()
	fmtStr := fmt.Sprintf("client_id=%s&sign=%s&timestamp=%d&id=%s&access_token=%s&machine_code=%s&response_type=%s&voice=%s", self.Client_id, self.Sign(timestamp), timestamp, uid, self.Access_token, machine_code, response_type, voice)
	fmt.Println(fmtStr)
	resp := httpPost(fmtStr, url)
	if self.CheckResponseStatus(resp) {
		fmt.Println(resp)
	} else {
		fmt.Println(resp)
	}
}

func (self OperationInterface) PrintInfo(machine_code string) {
	url := "https://open-api.10ss.net/printer/printinfo"
	uid := uuid.Must(uuid.NewV4())
	timestamp := time.Now().Unix()
	fmtStr := fmt.Sprintf("client_id=%s&sign=%s&timestamp=%d&id=%s&access_token=%s&machine_code=%s", self.Client_id, self.Sign(timestamp), timestamp, uid, self.Access_token, machine_code)
	fmt.Println(fmtStr)
	resp := httpPost(fmtStr, url)
	if self.CheckResponseStatus(resp) {
		fmt.Println(resp)
	} else {
		fmt.Println(resp)
	}
}

func (self OperationInterface) GetVersion(machine_code string) {
	url := "https://open-api.10ss.net/printer/getversion"
	uid := uuid.Must(uuid.NewV4())
	timestamp := time.Now().Unix()
	fmtStr := fmt.Sprintf("client_id=%s&sign=%s&timestamp=%d&id=%s&access_token=%s&machine_code=%s", self.Client_id, self.Sign(timestamp), timestamp, uid, self.Access_token, machine_code)
	fmt.Println(fmtStr)
	resp := httpPost(fmtStr, url)
	if self.CheckResponseStatus(resp) {
		fmt.Println(resp)
	} else {
		fmt.Println(resp)
	}
}

func (self OperationInterface) CancelAll(machine_code string) {
	url := "https://open-api.10ss.net/printer/cancelall"
	uid := uuid.Must(uuid.NewV4())
	timestamp := ztime.NowTimeStamp()
	fmtStr := fmt.Sprintf("client_id=%s&sign=%s&timestamp=%d&id=%s&access_token=%s&machine_code=%s", self.Client_id, self.Sign(timestamp), timestamp, uid, self.Access_token, machine_code)
	fmt.Println(fmtStr)
	resp := httpPost(fmtStr, url)
	if self.CheckResponseStatus(resp) {
		fmt.Println(resp)
	} else {
		fmt.Println(resp)
	}
}

func (self OperationInterface) CancelOne(machine_code string, order_id string) {
	url := "https://open-api.10ss.net/printer/cancelone"
	uid := uuid.Must(uuid.NewV4())
	timestamp := time.Now().Unix()
	fmtStr := fmt.Sprintf("client_id=%s&sign=%s&timestamp=%d&id=%s&access_token=%s&machine_code=%s&order_id=%s", self.Client_id, self.Sign(timestamp), timestamp, uid, self.Access_token, machine_code, order_id)
	fmt.Println(fmtStr)
	resp := httpPost(fmtStr, url)
	if self.CheckResponseStatus(resp) {
		fmt.Println(resp)
	} else {
		fmt.Println(resp)
	}
}

func (self OperationInterface) Seticon(machine_code string, img_url string) {
	url := "https://open-api.10ss.net/printer/seticon"
	uid := uuid.Must(uuid.NewV4())
	timestamp := time.Now().Unix()
	fmtStr := fmt.Sprintf("client_id=%s&sign=%s&timestamp=%d&id=%s&access_token=%s&machine_code=%s&img_url=%s", self.Client_id, self.Sign(timestamp), timestamp, uid, self.Access_token, machine_code, img_url)
	fmt.Println(fmtStr)
	resp := httpPost(fmtStr, url)
	if self.CheckResponseStatus(resp) {
		fmt.Println(resp)
	} else {
		fmt.Println(resp)
	}
}

func (self OperationInterface) DeleteIcon(machine_code string) {
	url := "https://open-api.10ss.net/printer/deleteicon"
	uid := uuid.Must(uuid.NewV4())
	timestamp := time.Now().Unix()
	fmtStr := fmt.Sprintf("client_id=%s&sign=%s&timestamp=%d&id=%s&access_token=%s&machine_code=%s", self.Client_id, self.Sign(timestamp), timestamp, uid, self.Access_token, machine_code)
	fmt.Println(fmtStr)
	resp := httpPost(fmtStr, url)
	if self.CheckResponseStatus(resp) {
		fmt.Println(resp)
	} else {
		fmt.Println(resp)
	}
}

func (self OperationInterface) BtnPrint(machine_code string, response_type string) {
	url := "https://open-api.10ss.net/printer/btnprint"
	uid := uuid.Must(uuid.NewV4())
	timestamp := time.Now().Unix()
	fmtStr := fmt.Sprintf("client_id=%s&sign=%s&timestamp=%d&id=%s&access_token=%s&machine_code=%s&response_type=%s", self.Client_id, self.Sign(timestamp), timestamp, uid, self.Access_token, machine_code, response_type)
	fmt.Println(fmtStr)
	resp := httpPost(fmtStr, url)
	if self.CheckResponseStatus(resp) {
		fmt.Println(resp)
	} else {
		fmt.Println(resp)
	}
}

func (self OperationInterface) GetOrder(machine_code string) {
	url := "https://open-api.10ss.net/printer/getorder"
	uid := uuid.Must(uuid.NewV4())
	timestamp := time.Now().Unix()
	fmtStr := fmt.Sprintf("client_id=%s&access_token=%s&machine_code=%s&response_type=%s&sign=%s&id=%s&timestamp=%d",
		self.Client_id, self.Access_token, machine_code, "close", self.Sign(timestamp), uid, timestamp)
	fmt.Println(fmtStr)
	resp := httpPost(fmtStr, url)
	if self.CheckResponseStatus(resp) {
		fmt.Println(resp)
	} else {
		fmt.Println(resp)
	}
}

//订单完成
func finishOrd() {
	//cmd=oauth_finish&machine_code=1212&order_id=613&state=1&print_time=1426908344&origin_id=15966&push_time=1426908399&sign=1F19C52B0EE3FE0F36FEF7487795F9F7
}

func main2() {
	Client_id := "1096322761"
	client_secret := "0b98b2a5341a5da3762cd20675bc9e95"

	//接单成功
	test := OurApplication{OperationInterface{Client_id, client_secret, "9a9383d873e941cca815300e781c5126", "0e3824be65e3454e817bd4b113d166b7"}} //未获取到token时初始化空字符串，获取到之后再填进来

	test.CancelAll("4004545322")

}
