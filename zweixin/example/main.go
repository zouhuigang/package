package main

import (
	"fmt"
	//"github.com/zouhuigang/package/ztime"
	"github.com/zouhuigang/package/zweixin"
)

func main() {
	Wx := map[string]string{}
	zwx, err := zweixin.NewZweixin(`wxfbc9af730c9f8168`, `4bb0a866be2838c20be7cadda0d29968`, `192.168.99.100`, 3306, `anooc`, ``)
	if err == nil {
		token := zwx.GetAccessToken()
		token1 := zwx.GetJsApiTicket()
		fmt.Printf("token:%v\njstoke:%v\n", token, token1)

		//验证前面是否正确
		Wx["jsapi_ticket"] = `kgt8ON7yVITDhtdwci0qedb8EuKl7VzW2NoBNJA819yQXNy4bd6IlLzxolhEatYfgOdvteSiqGXQlbmgsCusDQ`
		Wx["noncestr"] = `spybo2yt3ohu4jr8yaw6ik6vl3k6vhpg`
		Wx["timestamp"] = `1505096462`
		Wx["url"] = `https://www.anooc.com/edu/teacher/scan`
		zweixin.Signature(Wx)
		fmt.Println(Wx) //6d61cdd2d481ecb8a8b04e842aac91f613091043
	}
}
