package main

import (
	"fmt"
	"github.com/zouhuigang/package/zcrypto"
)

func main() {
	var I_key string = "123456"   //原始值
	var E_key string = "78910111" //加密过的
	//Hex 16位数 zouhuigang123456
	var Iv []byte = []byte("zouhuigang123456")

	s := zcrypto.Encode_money(699.34, I_key, E_key, Iv)
	s1 := zcrypto.Decprice(s, I_key, E_key)
	fmt.Println(s, s1)
}
