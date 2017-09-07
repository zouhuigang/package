//https://github.com/kinwyb/golang/blob/0935b112572b08ecfefaf1f128418bf8f7cef188/payment/wxpay/wxpay.go
package zweixin

import (
	//"crypto/sha1"
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"sort"
	"strings"
	"time"
)

//NonceStr:   必须, 32个字符以内, 商户生成的随机字符串
//随机32位字符串
func NonceStr() string {
	chars := []byte("abcdefghijklmnopqrstuvwxyz0123456789")
	rand.Seed(time.Now().UnixNano())
	result := [32]byte{}
	for i := 0; i < 32; i++ {
		result[i] = chars[rand.Int31n(35)]
	}
	return string(result[:])
}

//过滤
func paraFilter(params map[string]string) []string {
	keys := make([]string, 0)
	for k, v := range params {
		if k == "sign" || strings.TrimSpace(v) == "" {
			delete(params, k)
		} else {
			keys = append(keys, k)
		}
	}
	sort.Strings(keys)
	return keys
}

//拼接字符串 按照“参数=参数值”的模式用“&”字符拼接成字符串
func createLinkString(keys []string, args map[string]string) string {
	buf := bytes.NewBufferString("")
	for _, k := range keys {
		buf.WriteString(k)
		buf.WriteString("=")
		buf.WriteString(args[k])
		buf.WriteString("&")
	}
	buf.Truncate(buf.Len() - 1)
	return buf.String()
}

//签名
func Signature(args map[string]string) {
	keys := paraFilter(args)
	signStr := createLinkString(keys, args)
	sign := md5.Sum([]byte(signStr))
	args["sign"] = strings.ToUpper(hex.EncodeToString(sign[:]))
}

/*
// 微信公众号 明文模式/URL认证 签名
func Sign(token, timestamp, nonce string) (signature string) {
	strs := sort.StringSlice{token, timestamp, nonce}
	strs.Sort()

	buf := make([]byte, 0, len(token)+len(timestamp)+len(nonce))

	buf = append(buf, strs[0]...)
	buf = append(buf, strs[1]...)
	buf = append(buf, strs[2]...)

	hashsum := sha1.Sum(buf)
	return hex.EncodeToString(hashsum[:])
}

// 微信公众号/企业号 密文模式消息签名
func MsgSign(token, timestamp, nonce, encryptedMsg string) (signature string) {
	strs := sort.StringSlice{token, timestamp, nonce, encryptedMsg}
	strs.Sort()

	buf := make([]byte, 0, len(token)+len(timestamp)+len(nonce)+len(encryptedMsg))

	buf = append(buf, strs[0]...)
	buf = append(buf, strs[1]...)
	buf = append(buf, strs[2]...)
	buf = append(buf, strs[3]...)

	hashsum := sha1.Sum(buf)
	return hex.EncodeToString(hashsum[:])
}*/
