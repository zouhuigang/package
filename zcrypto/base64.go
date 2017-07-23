package zcrypto

import (
	"encoding/base64"
)

func Base64Encode(str string) string {
	encode := base64Encode([]byte(str))
	return string(encode)
}

func Base64Decode(str string) string {
	decode, err := base64Decode([]byte(str))
	if err != nil {
		return err.Error()
	}
	return string(decode)
}

func base64Encode(src []byte) []byte {
	return []byte(base64.StdEncoding.EncodeToString(src))
}

func base64Decode(src []byte) ([]byte, error) {
	return base64.StdEncoding.DecodeString(string(src))
}
