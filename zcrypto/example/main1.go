package main

import (
	"fmt"
	"github.com/zouhuigang/package/zcrypto"
)

func main() {
	encodeStr := "http://mdtest.ematong.com/micro-imfans/wap/adcallback/redpacket"
	encode := zcrypto.Base64Encode(encodeStr)
	fmt.Printf("base64 encode:%v\n", encode)

	decodeStr := "aHR0cDovL21kdGVzdC5lbWF0b25nLmNvbS9taWNyby1pbWZhbnMvd2FwL2FkY2FsbGJhY2svcmVkcGFja2V0"
	decode := zcrypto.Base64Decode(decodeStr)
	fmt.Printf("base64 decode:%v\n", decode)
}
