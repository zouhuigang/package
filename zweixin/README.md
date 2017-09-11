### 微信开发专用包


调用js：

	data := map[string]interface{}{}
	appid := `wxfbc9af730c9f8168`
	appsecret := `4bb0a866be2838c20be7cadda0d29968`
	Wx := map[string]string{}
	zwx, err := zweixin.NewZweixin(appid, appsecret, `192.168.99.100`, 3306, `anooc`, ``)
	if err == nil {
		jsapi_ticket := zwx.GetJsApiTicket()
		fmt.Printf("jstoke:%v\n", jsapi_ticket)

		//验证前面是否正确
		Wx["jsapi_ticket"] = jsapi_ticket
		Wx["noncestr"] = zweixin.NonceStr()
		timestamp := fmt.Sprintf("%d", ztime.NowTimeStamp())
		Wx["timestamp"] = timestamp
		Wx["url"] = `https://www.anooc.com/edu/teacher/scan`
		zweixin.Signature(Wx)
		Wx["appid"] = appid
		fmt.Println(Wx) //6d61cdd2d481ecb8a8b04e842aac91f613091043
	}

	data["wx"] = Wx

