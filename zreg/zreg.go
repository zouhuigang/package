package zreg

import (
	"net/http"
	"regexp"
	"strings"
)

//去除空字符串
func Trim(str string) string {
	// 去除空格
	str = strings.Replace(str, " ", "", -1)
	//fmt.Println(strings.TrimSpace(" \t\n anooc.com \n\t\r\n"))
	return str
}

//判断是否为空,为空true
func IsNull(s string) bool {
	s = Trim(s)
	if s == "" {
		return true
	}
	//新加
	if len(s) == 0 {
		//为空的处理
		return true
	}
	return false
}

//是手机号
func Is_phone(phone string) bool {
	//reg := `^1([38][0-9]|14[57]|5[^4])\d{8}$`
	reg := `^(1[3|4|5|7|8][0-9]\d{4,8})$`
	rgx := regexp.MustCompile(reg)
	return rgx.MatchString(phone)
}

//是移动端 req := this.Request(ctx)
func IsMobile(req *http.Request) bool {
	userAgent := req.Header.Get("User-Agent")
	mobileRe, _ := regexp.Compile("(?i:Mobile|iPod|iPhone|Android|Opera Mini|BlackBerry|webOS|UCWEB|Blazer|PSP)")
	flagM := mobileRe.FindString(userAgent)

	if !IsNull(flagM) {
		return true
	} else {
		return false
	}

}

//是邮箱2-10位@xxxx.xx
func Is_email(email string) bool {
	if m, _ := regexp.MatchString(`^([\w\.\_]{2,10})@(\w{1,}).([a-z]{2,4})$`, email); !m {
		return false
	}
	return true
}

//身份证
func Is_usercard(usercard string) bool {
	//验证15位身份证，15位的是全部数字
	if m, _ := regexp.MatchString(`^(\d{15})$`, usercard); !m {
		return false
	}

	//验证18位身份证，18位前17位为数字，最后一位是校验位，可能为数字或字符X。
	if m, _ := regexp.MatchString(`^(\d{17})([0-9]|X)$`, usercard); !m {
		return false
	}

	return true

}

//验证全部中文
func Is_chinese(str string) bool {
	if m, _ := regexp.MatchString("^[\\x{4e00}-\\x{9fa5}]+$", str); !m {
		return false
	}
	return true
}

//验证全为英文
func Is_english(engname string) bool {
	if m, _ := regexp.MatchString("^[a-zA-Z]+$", engname); !m {
		return false
	}
	return true
}

//验证select是否是之前设定的值，防止更改 	//slice := []string{"apple", "pear", "banane"}
func Is_select(slice []string, selects string) bool {
	for _, v := range slice {
		if v == selects {
			return true
		}
	}
	return false
}

//复选框//slice := []string{"football", "basketball", "tennis"}
func Is_checkbox(slice []interface{}, selects []interface{}) bool {
	a := Slice_diff(selects, slice)
	if a == nil {
		return true
	}
	return false
}

//https://github.com/astaxie/beeku
func Slice_diff(slice1, slice2 []interface{}) (diffslice []interface{}) {
	for _, v := range slice1 {
		if !In_slice(v, slice2) {
			diffslice = append(diffslice, v)
		}
	}
	return
}
func In_slice(val interface{}, slice []interface{}) bool {
	for _, v := range slice {
		if v == val {
			return true
		}
	}
	return false
}

//单选 	slice := []int{1, 2}
func Is_radio(slice []int, selects int) bool {
	for _, v := range slice {
		if v == selects {
			return true
		}
	}
	return false
}

//判断是正整数positive
func IsPosInt(n int) bool {
	if n <= 0 {
		return false
	}
	return true
}
