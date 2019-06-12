package zcrypto

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"strings"
)

// //MD5加密
// //@param			str			需要加密的字符串
// //@param			salt		盐值
// //@return			CryptStr	加密后返回的字符串
// func Md5crypt(str string, salt ...interface{}) (CryptStr string) {
// 	if l := len(salt); l > 0 {
// 		slice := make([]string, l+1)
// 		str = fmt.Sprintf(str+strings.Join(slice, "%v"), salt...)
// 	}
// 	return fmt.Sprintf("%x", md5.Sum([]byte(str)))
// }

//md5 32位 hash值
func Md5(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

//32位大写
func Md5UP(text string) string {
	st := Md5(text)
	return strings.ToUpper(st)
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

//16位大写
func Md5To16UP(text string) string {
	st := Md5To16(text)
	return strings.ToUpper(st)
}

func Md5HexString(md5 [16]byte) (s string) {
	s = fmt.Sprintf("% x", md5)
	s = strings.Replace(s, " ", ":", -1)
	return s
}

/*文件md5
f, err := os.Open(*path)
	if err != nil {
		fmt.Println("Open", err)
		return
	}

	defer f.Close()*/
func FileToMd5(f io.Reader) (error, string) {
	hasher := md5.New()
	if _, err := io.Copy(hasher, f); err != nil {
		return err, ""
	}
	return nil, hex.EncodeToString(hasher.Sum(nil))
}
