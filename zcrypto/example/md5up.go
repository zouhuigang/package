package main

import (
	"fmt"
	"github.com/zouhuigang/package/zcrypto"
)

func main() {

	data := zcrypto.Md5UP("zouhuigang 952750120@qq.com http://s") //32 md5

	fmt.Printf("32位大写md5 hash %s\n", data)

	data = zcrypto.Md5To16UP("zouhuigang 952750120@qq.com http://s")
	fmt.Printf("16位大写md5 hash %s\n", data)

}
