package main

import (
	//"fmt"
	"github.com/zouhuigang/package/zcrypto"
	"github.com/zouhuigang/package/zprint"
)

func main() {

	/*
		初次运行请先调用getToken方法获取access_token及refresh_token，如无意外两个token会输出到命令行，复制两个token放入实例中使用
		未获取到token时初始化空字符串，获取到之后再填进来
	*/
	//应用id,应用密钥
	//test := OurApplication{OperationInterface{"1096322761", "0b98b2a5341a5da3762cd20675bc9e95", "", ""}} //未获取到token时初始化空字符串，获取到之后再填进来
	//open := OpenApplication{OperationInterface{"1096845322", "1c498e2c9214c0b712777b572ca2831f", "71ed2a877be6a820ddf86ddd0af9f528", "27d076a292ce8b0708972395850d7595"}, "d92b3a0eb2c1a33bb625e619e40be917"}
	//test.GetToken()
	client_id := "1096322761"
	client_secret := "0b98b2a5341a5da3762cd20675bc9e95"
	test := zprint.OurApplication{zprint.OperationInterface{client_id, client_secret, "9a9383d873e941cca815300e781c5126", "0e3824be65e3454e817bd4b113d166b7"}} //未获取到token时初始化空字符串，获取到之后再填进来
	test.AddPrinter("4004545322", "srkfnt5ytjr4")

	//学优教育辅导签到单打印===========
	title := `<FS><center>╔学优教育1v1辅导签到单╝</center></FS>`
	date := "\r\n辅导科目:语文\r\n辅导日期:2017-10-28\r\n辅导时段:08:00-10:00\r\n"
	name := "学生:何赵懿\r\n教师:李闻达 2017.10.28 14:21\r\n"
	qrcode := "<center> </center><QR>http://www.anooc.com/edu/teacher/sign?qrcode=29A9617D6D26B608</QR>注:此二维码有效期3天，过期作废!\r\n<center> </center>"
	sign := "学生签名:\r\n\r\n\r\n┅┅┅┅┅┅┅┅┅┅┅┅┅┅┅┅\r\n"
	bottomTips := "SCH143-29A9617D6D26B608\r\n*****请务必确认上课信息是否正确\r\n"
	content := zcrypto.Urlencode(title + date + name + qrcode + sign + bottomTips)
	test.Print("4004545322", content) // 打印接口
	// open.ShutdownRestart("400451758","restart")
}
