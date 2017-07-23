/*
AES加密解密
*/
package zcrypto

//http://studygolang.com/articles/7302
//https://www.oschina.net/code/snippet_197499_25891
import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
)

var commonIV = []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}

func passKeyEncode(passKey string) cipher.Block {
	// 创建加密算法aes
	c, err := aes.NewCipher([]byte(passKey))
	if err != nil {
		fmt.Printf("Error: NewCipher(%d bytes) = %s", len(passKey), err)
		//os.Exit(-1)
	}

	return c
}

////(参数key必须是16、24或者32位的[]byte，分别对应AES-128, AES-192或AES-256算法)
//fmt.Printf("%s=>%x\n", strByte, ciphertext) ,fmt.Println(len(passKey))
//passKey:aes的加密字符串
func AesEncodeCFB(str string, passKey string) string {
	strByte := []byte(str)

	c := passKeyEncode(passKey)

	//加密字符串
	cfb := cipher.NewCFBEncrypter(c, commonIV)
	ciphertext := make([]byte, len(strByte))
	cfb.XORKeyStream(ciphertext, strByte)

	return string(ciphertext)
}

func AesDecodeCFB(str string, passKey string) string {
	strByte := []byte(str)
	c := passKeyEncode(passKey)
	// 解密字符串
	cfbdec := cipher.NewCFBDecrypter(c, commonIV)
	decode := make([]byte, len(strByte))
	cfbdec.XORKeyStream(decode, strByte)
	//fmt.Printf("%x=>%s\n", decode, strByte)
	return string(decode)
}
