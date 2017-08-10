/*
string 转换成int
作者:邹慧刚
邮箱：952750120@qq.com
github.com:github.com/zouhuigang
*/
package zcore

import (
	"fmt"
	"log"
	"runtime"
	"strconv"
	"strings"
)

// 字符串转int
func StringToInt(s string, defaultVal ...int) int {
	getDefault := func() int {
		if len(defaultVal) > 0 {
			return defaultVal[0]
		}
		return 0
	}

	if s == "" {
		return getDefault()
	}

	i, err := strconv.Atoi(strings.TrimSpace(s))
	if err != nil {
		msg := "zcore StringToInt strconv.Atoi error:" + err.Error()
		// 加上文件调用和行号
		_, callerFile, line, ok := runtime.Caller(1)
		if ok {
			msg += fmt.Sprintf("file:%s,line:%d", callerFile, line)
		}
		log.Println(msg)
		return getDefault()
	}

	return i
}

//字符串转int64
func StringToInt64(s string, defaultVal ...int64) int64 {
	getDefault := func() int64 {
		if len(defaultVal) > 0 {
			return defaultVal[0]
		}
		return 0
	}

	if s == "" {
		return getDefault()
	}

	i, err := strconv.ParseInt(strings.TrimSpace(s), 10, 64)
	if err != nil {
		msg := "zcore StringToInt64 strconv.ParseInt error:" + err.Error()
		// 加上文件调用和行号
		_, callerFile, line, ok := runtime.Caller(1)
		if ok {
			msg += fmt.Sprintf("file:%s,line:%d", callerFile, line)
		}
		log.Println(msg)
		log.Println("zcore StringToInt64 strconv.ParseInt error:", err)
		return getDefault()
	}

	return i
}
