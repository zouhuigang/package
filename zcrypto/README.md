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



### Sign包

>验证请求传过来的参数是否被更改过

1.微信签名

	package main

	import (
		"fmt"
		"github.com/zouhuigang/package/zcrypto"
	)
	
	func main() {
		str := `jsapi_ticket=kgt8ON7yVITDhtdwci0qedb8EuKl7VzW2NoBNJA819yQXNy4bd6IlLzxolhEatYfgOdvteSiqGXQlbmgsCusDQ&noncestr=spybo2yt3ohu4jr8yaw6ik6vl3k6vhpg&timestamp=1505096462&url=https://www.anooc.com/edu/teacher/scan`
	
		s1 := zcrypto.Sign(str)
		s2 := zcrypto.SignPhpSha1(str)
		fmt.Printf("%s\n%s\n", s1, s2)
	}


输出：

	
	fcedccd3459120b97bfbadd261d61300
	6d61cdd2d481ecb8a8b04e842aac91f613091043


2.易联云签名


	package main

	import (
		"fmt"
		"github.com/zouhuigang/package/zcrypto"
	)
	
	func main() {
		client_id := `1096828699`
		client_secret := `1f66e825d98bc916afc084c59fe3c883`
		var push_time int64 = 1509523926
		str := fmt.Sprintf("%s%d%s", client_id, push_time, client_secret)
		s1 := zcrypto.Sign(str)
		s2 := zcrypto.SignPhpSha1(str)
		fmt.Printf("%s\n%s\n", s1, s2)
	}



输出：

	
	a174d91e441617db9d4a1c41c59af088
	b631528ba6f7bc0e973c0631491baf87144290b4



签名验证规则：

    应用将接收到易联云推送的POST Body
    得到的POST Body里面的sign
    对上一步得到的sign与md5(client_id+push_time+client_secret)校验是否正确
    通过上一步验证,请返回以下示例(格式json),否则将视为推送失败,会再次推送两次,如果均未有返回将放弃推送

	{"data": "OK"}

	字 段	含 义	实 例	备 注
	data	接收成功	OK	接收成功


总结下来，push_time很重要，所以尽量推送的时候，都带一个时间戳回来。然后是用户id或应用id,这个可以后台得到。



