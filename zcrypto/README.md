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



### php sha1算法

	package main

	import (
		"fmt"
		"github.com/zouhuigang/package/zcrypto"
	)

	func main() {
		str := `jsapi_ticket=kgt8ON7yVITDhtdwci0qedb8EuKl7VzW2NoBNJA819yQXNy4bd6IlLzxolhEatYfgOdvteSiqGXQlbmgsCusDQ&noncestr=spybo2yt3ohu4jr8yaw6ik6vl3k6vhpg&timestamp=1505096462&url=https://www.anooc.com/edu/teacher/scan`
	
		sha1 := zcrypto.PhpSha1(str)
		fmt.Println(sha1)
	}

输出：6d61cdd2d481ecb8a8b04e842aac91f613091043,跟微信接口的一样https://mp.weixin.qq.com/debug/cgi-bin/sandbox?t=jsapisign



### 金额（money）加密与解密

> 为了防止存入mysql中的金额数据，被其他程序员更改或知道mysql密码的用户知道，乱更改金额，导致金额不一致，所以有了这个加密与解密函数，在程序服务端控制加解密。



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
	
		s := zcrypto.Encode_money(699.34, I_key, E_key, Iv)//加密
		s1 := zcrypto.Decprice(s, I_key, E_key)//解密
		fmt.Println(s, s1)
	}


输出：


	em91aHVpZ2FuZzEyMzQ1NsaLy7RJrXzKNpJF9g  699.34






