package main

import (
	"github.com/zouhuigang/package/zip"
	"log"
)

func main() {
	q := zip.NewQQwry("qqwry.dat")
	q.Find("116.232.99.89")
	zip.VerifyIp("dsad.123.123.1.3.101asd")
	log.Printf("ip:%v, Country:%v, Area:%v", q.Ip, q.Country, q.Area)
}
