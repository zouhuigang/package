//https://blog.tanteng.me/2017/03/php-urldecode-in-go/
package zcrypto

import (
	"net/url"
)

func Urlencode(urlStr string) string {
	return url.QueryEscape(urlStr)
}

func Urldecode(urlStr string) string {
	url, err := url.QueryUnescape(urlStr)
	if err != nil {

	}
	return url
}
