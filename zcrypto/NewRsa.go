/*
生成rsa文件
*/

package zcrypto

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/asn1"
	"encoding/pem"
	"io"
	"os"
)

// NewRSAFile 生成密钥对文件
// pubKeyFilename: 公钥文件名 priKeyFilename: 私钥文件名 kekeyLength: 密钥长度
func NewRSAFile(pubKeyFilename, priKeyFilename string, keyLength int) error {
	if pubKeyFilename == "" {
		pubKeyFilename = "id_rsa.pub"
	}
	if priKeyFilename == "" {
		priKeyFilename = "id_rsa"
	}

	if keyLength == 0 || keyLength < 1024 {
		keyLength = 1024
	}

	// 创建公钥文件
	pubWriter, err := os.Create(pubKeyFilename)
	if err != nil {
		return err
	}
	defer pubWriter.Close()

	// 创建私钥文件
	priWriter, err := os.Create(priKeyFilename)
	if err != nil {
		return err
	}
	defer priWriter.Close()

	// 生成密钥对
	err = WriteRSAKeyPair(pubWriter, priWriter, keyLength)
	if err != nil {
		return err
	}

	return nil
}

// WriteRSAKeyPair 生成RSA密钥对
func WriteRSAKeyPair(publicKeyWriter, privateKeyWriter io.Writer, keyLength int) error {

	// 生成私钥文件
	privateKey, err := rsa.GenerateKey(rand.Reader, keyLength)
	if err != nil {
		return err
	}

	derStream := MarshalPKCS8PrivateKey(privateKey)
	block := &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: derStream,
	}
	err = pem.Encode(privateKeyWriter, block)
	if err != nil {
		return err
	}

	// 生成公钥文件
	publicKey := &privateKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)

	if err != nil {
		return err
	}

	block = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPkix,
	}

	err = pem.Encode(publicKeyWriter, block)
	if err != nil {
		return err
	}

	return nil

}

// MarshalPKCS8PrivateKey 私钥解析
func MarshalPKCS8PrivateKey(key *rsa.PrivateKey) []byte {

	info := struct {
		Version             int
		PrivateKeyAlgorithm []asn1.ObjectIdentifier
		PrivateKey          []byte
	}{}

	info.Version = 0
	info.PrivateKeyAlgorithm = make([]asn1.ObjectIdentifier, 1)
	info.PrivateKeyAlgorithm[0] = asn1.ObjectIdentifier{1, 2, 840, 113549, 1, 1, 1}
	info.PrivateKey = x509.MarshalPKCS1PrivateKey(key)
	k, _ := asn1.Marshal(info)
	return k

}
