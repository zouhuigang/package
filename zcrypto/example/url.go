package main

import (
	"fmt"
	"github.com/zouhuigang/package/zcrypto"
)

func main() {

	url1 := zcrypto.Urlencode("http://192.168.99.100:8089/edu/teacher/tlogin?backurl=/edu/teacher/sign?qrcode=B934D55BD7120E2")

	fmt.Printf("url encode  %s\n", url1)

	data := zcrypto.Urldecode("http%3a%2f%2f192.168.99.100%3a8089%2fedu%2fteacher%2ftlogin%3fbackurl%3d%2fedu%2fteacher%2fsign%3fqrcode%3dB934D55BD7120E2")
	fmt.Printf("url decode %s\n", data)

}
