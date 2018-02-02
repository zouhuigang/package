package main

import (
	"fmt"
	"github.com/zouhuigang/package/zcore"
)

func main() {
	a := []string{"hello", "", "  ", "world", "yes", "hello", "nihao", "shijie", "hello", "yes", "nihao", "good"}
	a = zcore.RemoveDuplicatesAndEmpty(a)
	fmt.Printf("%v", a)
}

func main2() {
	domain := "http://www.anooc.com/"
	_, s1 := zcore.IsHttpOrHttps("images/qqzengLogo.png", domain)
	_, s2 := zcore.IsHttpOrHttps("../image/qqzeng_ip_weixin.png", domain)
	_, s3 := zcore.IsHttpOrHttps("https://www.qqzeng.com/ip/images/qqzengLogo.png", domain)

	fmt.Printf("%s\n%s\n%s\n", s1, s2, s3)
}
