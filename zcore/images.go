package zcore

import (
	"strings"
)

//判断https或http前缀
func IsHttpOrHttps(value string, prevalue string) (bool, string) {
	if !strings.HasPrefix(value, "http") {
		value = prevalue + value
		return false, value
	}

	return true, value
}
