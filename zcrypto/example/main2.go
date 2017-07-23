package main

import (
	"fmt"
	"github.com/zouhuigang/package/zcrypto"
)

func main() {
	encodeStr := "zouhuigang952750120"
	passKey := "bztaxyz12798akljzmknm.ahkjkljl;k"
	encode := zcrypto.AesEncodeCFB(encodeStr, passKey)
	fmt.Printf("aes cfb mode encode:%x\n", encode)

	//encode = "5072eadc20720cdb321b7c62947982d8227d"
	decode := zcrypto.AesDecodeCFB(encode, passKey)
	fmt.Printf("aes  cfb mode decode:%s\n", decode)
}
