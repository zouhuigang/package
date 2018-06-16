package main

import (
	"fmt"
	"github.com/zouhuigang/package/zip"
)

func main() {
	p, _ := zip.New(`qqzeng.dat`)
	ip := "210.51.200.123"
	ipstr := p.Get(ip)
	fmt.Println(ipstr)
	ip1, _ := p.FindIp("116.232.99.89")
	fmt.Printf("%v", ip1)

}
