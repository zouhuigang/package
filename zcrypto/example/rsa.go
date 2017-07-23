package main

import (
	"fmt"
	"github.com/zouhuigang/package/zcrypto"
)

func main() {

	data, err := zcrypto.RsaEncode([]byte("zouhuigang 952750120@qq.com http://s")) //RSA加密
	if err != nil {
		panic(err)
	}
	fmt.Printf("RSA加密 %x\n", string(data))
	origData, err := zcrypto.RsaDecode(data) //RSA解密
	if err != nil {
		panic(err)
	}
	fmt.Println("RSA解密", string(origData))
}
