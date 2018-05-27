package zcrypto

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"io/ioutil"
)

const (

	// RSAAlgorithmSign RSA签名算法
	RSAAlgorithmSign = crypto.SHA256
)

// ReadRSAKeyPair 读取RSA密钥对
// pubKeyFilename: 公钥文件名称   priKeyFilename: 私钥文件名
func ReadRSAKeyPair(pubKeyFilename, priKeyFilename string) ([]byte, []byte, error) {
	pub, err := ioutil.ReadFile(pubKeyFilename)
	if err != nil {
		return nil, nil, err
	}

	pri, err := ioutil.ReadFile(priKeyFilename)
	if err != nil {
		return nil, nil, err
	}
	return pub, pri, nil
}

// GoRSA RSA加密解密
type GoRSA struct {
	PublicKey  *rsa.PublicKey
	PrivateKey *rsa.PrivateKey
}

// NewGoRSA 初始化 GoRSA对象
func NewGoRSA(pubKeyFilename, priKeyFilename string) (*GoRSA, error) {

	publicKey, privateKey, err := ReadRSAKeyPair(pubKeyFilename, priKeyFilename)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(publicKey)
	if block == nil {
		return nil, errors.New("public key error")
	}

	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	pub := pubInterface.(*rsa.PublicKey)
	block, _ = pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("private key error")
	}

	priv, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	pri, ok := priv.(*rsa.PrivateKey)
	if ok {
		return &GoRSA{
			PublicKey:  pub,
			PrivateKey: pri,
		}, nil
	}
	return nil, errors.New("private key not supported")
}

// PublicEncrypt 公钥加密
func (r *GoRSA) PublicEncrypt(data []byte) ([]byte, error) {
	partLen := r.PublicKey.N.BitLen()/8 - 11
	chunks := split(data, partLen)
	buffer := bytes.NewBufferString("")

	for _, chunk := range chunks {
		bytes, err := rsa.EncryptPKCS1v15(rand.Reader, r.PublicKey, chunk)
		if err != nil {
			return nil, err
		}
		buffer.Write(bytes)
	}

	return buffer.Bytes(), nil
}

// PrivateDecrypt 私钥解密
func (r *GoRSA) PrivateDecrypt(encrypted []byte) ([]byte, error) {

	partLen := r.PublicKey.N.BitLen() / 8
	chunks := split(encrypted, partLen)
	buffer := bytes.NewBufferString("")

	for _, chunk := range chunks {
		decrypted, err := rsa.DecryptPKCS1v15(rand.Reader, r.PrivateKey, chunk)
		if err != nil {
			return nil, err
		}
		buffer.Write(decrypted)
	}
	return buffer.Bytes(), nil
}

// Sign 数据进行签名
func (r *GoRSA) Sign(data string) (string, error) {

	h := RSAAlgorithmSign.New()
	h.Write([]byte(data))
	hashed := h.Sum(nil)
	sign, err := rsa.SignPKCS1v15(rand.Reader, r.PrivateKey, RSAAlgorithmSign, hashed)
	if err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(sign), err
}

// Verify 数据验证签名
func (r *GoRSA) Verify(data string, sign string) error {

	h := RSAAlgorithmSign.New()
	h.Write([]byte(data))
	hashed := h.Sum(nil)
	decodedSign, err := base64.RawURLEncoding.DecodeString(sign)
	if err != nil {
		return err
	}
	return rsa.VerifyPKCS1v15(r.PublicKey, RSAAlgorithmSign, hashed, decodedSign)
}

func split(buf []byte, lim int) [][]byte {
	var chunk []byte
	chunks := make([][]byte, 0, len(buf)/lim+1)
	for len(buf) >= lim {
		chunk, buf = buf[:lim], buf[lim:]
		chunks = append(chunks, chunk)
	}
	if len(buf) > 0 {
		chunks = append(chunks, buf[:len(buf)])
	}
	return chunks
}
