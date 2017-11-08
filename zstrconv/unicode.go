/*
中文转unicode
*/
package zstrconv

import (
	"strconv"
)

//仅仅将中文->unicode
func OnlyChineseToUnicode(txt string) string {
	textQuoted := strconv.QuoteToASCII(txt)
	return textQuoted
}

//将unicode转换成中文
func UnicodeToChinese(strVal string) (string, error) {
	txt, err := strconv.Unquote(strVal)
	if err != nil {
		return ``, err
	}
	return txt, nil
}

//将中文字母符号等->unicode
/*
unicode实际上是ascii，Hex(16进制)
*/
func AllToUnicode(strVal string) string {
	//num2 := []byte(s)
	return ``
}
