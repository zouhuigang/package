package main

import (
	"fmt"
	"github.com/zouhuigang/package/zcrypto"
	"log"
)

func main() {
	// 指定 公钥文件名 和 私钥文件名
	gorsa, err := zcrypto.NewGoRSA("id_rsa.pub", "id_rsa")
	if err != nil {
		log.Fatalln(err)
	}

	// 明文字符
	rawStr := "O8Hp8WQbFPT7b5AUsEMVLtIU3MVYOrt8 20180527zouhuigang 952750120@qq.com http://s"

	// 使用公钥加密
	encrypt, err := gorsa.PublicEncrypt([]byte(rawStr))
	if err != nil {
		log.Fatalln(err)
	}

	// 使用私钥解密
	decrypt, err := gorsa.PrivateDecrypt(encrypt)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(decrypt))
}

func main2() {

	data, err := zcrypto.RsaEncode([]byte("20180527zouhuigang 952750120@qq.com http://s")) //RSA加密
	if err != nil {
		panic(err)
	}
	fmt.Printf("RSA加密 %x\n", string(data))
	origData, err := zcrypto.RsaDecode(data) //RSA解密
	if err != nil {
		panic(err)
	}
	fmt.Println("RSA解密", string(origData))
}
