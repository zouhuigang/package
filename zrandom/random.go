package zrandom

//复制的是https://github.com/dchest/captcha/blob/master/random.go
//需要改版，先用着
import (
	"crypto/rand"
	"io"
)

//随机字符串
var randStr = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")

const idLen = 20

//随机数字
func Digits(length int) []byte {
	return randomBytesMod(length, 10)
}

func DigitsToString(length int) string {
	return RandomDigitsToString(randomBytesMod(length, 10))
}

func randomId() string {
	b := randomBytesMod(idLen, byte(len(randStr)))
	for i, c := range b {
		b[i] = randStr[c]
	}
	return string(b)
}

func randomBytesMod(length int, mod byte) (b []byte) {
	if length == 0 {
		return nil
	}
	if mod == 0 {
		panic("captcha: bad mod argument for randomBytesMod")
	}
	maxrb := 255 - byte(256%int(mod))
	b = make([]byte, length)
	i := 0
	for {
		r := randomBytes(length + (length / 4))
		for _, c := range r {
			if c > maxrb {
				// Skip this number to avoid modulo bias.
				continue
			}
			b[i] = c % mod
			i++
			if i == length {
				return
			}
		}
	}

}

func randomBytes(length int) (b []byte) {
	b = make([]byte, length)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		panic("captcha: error reading random source: " + err.Error())
	}
	return
}

//将处理过的[]byte转成string
func RandomDigitsToString(b []byte) string {
	var numList = []byte("0123456789")
	var new_byte = make([]byte, len(b))
	for i, c := range b {
		new_byte[i] = numList[c]
	}
	return string(new_byte)
}
