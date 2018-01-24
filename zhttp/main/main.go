package main

import (
	"fmt"
	"github.com/zouhuigang/package/zhttp"
	"log"
)

type ticket_callback struct {
	Errcode int64  `json:"errcode"`
	Errmsg  string `json:"errmsg"`
	Ticket  string `json:"ticket"`
}

func main() {
	appid := `wxfbc9af730c9f8168`
	appsec := `4bb0a866be2838c20be7cadda0d29968`
	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s", appid, appsec)
	zhttp.GET(url)

}

func main2() {
	callback := &ticket_callback{}

	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/ticket/getticket?access_token=%s&type=jsapi", "ss")
	err := zhttp.GETWithUnmarshal(url, callback)
	if err != nil {
		log.Print(err.Error())
		///return "", err
	}

	log.Printf("=======%v\n", callback)

}
