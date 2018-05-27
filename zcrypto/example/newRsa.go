package main

import (
	"github.com/zouhuigang/package/zcrypto"
	"log"
)

func main() {
	// 生成 2048 位密钥对文件 指定名称
	err := zcrypto.NewRSAFile("id_rsa.pub", "id_rsa", 2048)
	if err != nil {
		log.Fatalln(err)
	}
}
