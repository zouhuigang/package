package main

import (
	"crypto/md5"
	"fmt"
	"github.com/zouhuigang/package/zcrypto"
)

func main() {

	data := zcrypto.Md5("zouhuigang 952750120@qq.com http://s") //32 md5

	fmt.Printf("32位md5 hash %s\n", data)

	data = zcrypto.Md5To16("zouhuigang 952750120@qq.com http://s")
	fmt.Printf("16位md5 hash %s\n", data)

	data1 := zcrypto.Md5HexString(md5.Sum([]byte(data))) //p.key.Marshal()

	fmt.Printf("md5  %s %x\n", data1, data1)

	data3 := []byte("zouhuigang 952750120@qq.com http://s")
	fmt.Printf("%x", md5.Sum(data3))

}
