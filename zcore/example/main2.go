package main

import (
	"fmt"
	"github.com/zouhuigang/package/zcore"
)

func main() {
	domain := "http://www.anooc.com/"
	_, s1 := zcore.IsHttpOrHttps("images/qqzengLogo.png", domain)
	_, s2 := zcore.IsHttpOrHttps("../image/qqzeng_ip_weixin.png", domain)
	_, s3 := zcore.IsHttpOrHttps("https://www.qqzeng.com/ip/images/qqzengLogo.png", domain)

	fmt.Printf("%s\n%s\n%s\n", s1, s2, s3)
}
