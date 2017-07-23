### 加密解密包

	AES：更快，兼容设备，安全级别高；
	
	SHA1：公钥后处理回传
	
	DES：本地数据，安全级别低
	
	RSA：非对称加密，有公钥和私钥
	
	MD5：防篡改



### base64 encode/decode:

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


输出:

	base64 encode:aHR0cDovL21kdGVzdC5lbWF0b25nLmNvbS9taWNyby1pbWZhbnMvd2FwL2FkY2FsbGJhY2svcmVkcGFja2V0

	base64 decode:http://mdtest.ematong.com/micro-imfans/wap/adcallback/redpacket


### aes


My name is Astaxie=>5072eadc20720cdb321b7c62947982d8227d
5072eadc20720cdb321b7c62947982d8227d=>My name is Astaxie




http://www.philo.top/2015/03/18/golang-js-des/
