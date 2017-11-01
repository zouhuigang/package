/*
get/Post参数签名包，验证参数在传输过程中，是否被更改过
*/

package zcrypto

//md5--传参数过来
func Sign(data string) string {
	sign := Md5(data)
	return sign
}

//PhpSha1sha1方式
func SignPhpSha1(data string) string {
	sign := PhpSha1(data)
	return sign
}
