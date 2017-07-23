package zcrypto

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
)

//md5 32位 hash值
func Md5(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

//16位md5,MD5形成后是一个32个字符的字符串，去掉前8个，去掉后8个，即32-8-8=16。这就是16位的MD5。
func Md5To16(text string) string {
	m32 := Md5(text)
	rs := []rune(m32)
	/*for i := 0; i < len(rs); i++ {
		fmt.Println("rs[", i, "]=", rs[i], "string=", string(rs[i]))
	}*/

	return string(rs[8:24])

}

func Md5HexString(md5 [16]byte) (s string) {
	s = fmt.Sprintf("% x", md5)
	s = strings.Replace(s, " ", ":", -1)
	return s
}
