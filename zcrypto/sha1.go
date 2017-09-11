package zcrypto

import (
	"crypto/sha1"
	"fmt"
	"io"
)

func PhpSha1(txt string) string {
	h := sha1.New()
	io.WriteString(h, txt)
	sha1 := fmt.Sprintf("%x\n", h.Sum(nil))
	return sha1
}
