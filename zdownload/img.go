/*
下载文件工具包
*/
package zdownload

import (
	//"fmt"
	"github.com/zouhuigang/package/zhttp"
)

func Img(url string) string {
	//防止程序出错
	defer func() {
		if r := recover(); r != nil {
			//log.Println("[E]", r)
		}
	}()

	respone := zhttp.HttpGet(url)
	return respone
}
