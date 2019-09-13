package zreg

import "strings"

//检测是否含有http开头的数据
func IsHttpUrl(m_url string) bool {
	is_success := false
	m_url = strings.TrimSpace(m_url)
	if strings.HasPrefix(m_url, "http") {
		is_success = true
	}
	return is_success
}
