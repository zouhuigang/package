package zfile

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

//判断文件是否存在
func IsFileExist(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}
	return true
}

// 判断文件夹是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

//判断文件夹是否存在，如果不存在，则创建
func PathExistsAndMkDir(_dir string) error {

	exist, err := PathExists(_dir)
	if err != nil {
		return err
	}
	if exist {
		return nil
	}

	// 创建文件夹
	err = os.Mkdir(_dir, os.ModePerm)
	if err != nil {
		return err
	}

	return nil

}

//获取文件后缀
func Ext(path string) string {
	return strings.ToLower(filepath.Ext(path))
}

/*
处理路径
windows下路径为:"D:\\workspacego\\src\\anooc-hujiao\\pdf\\赠险接口文档.docx" =>D:/workspacego/src/anooc-hujiao/pdf/赠险接口文档.docx
linux下路径为:/root/mnt/
其实建议使用 “/” 作为路径分隔符，windows 和 linux 系统都能正常使用。
// 将 path 中平台相关的路径分隔符转换为 '/'
ToSlash(path string) string
// 将 path 中的 '/' 转换为系统相关的路径分隔符
FromSlash(path string) string
*/
func FomartPath(path string) string {
	sysType := runtime.GOOS

	// if sysType == "linux" {
	// 	// LINUX系统
	// }

	if sysType == "windows" {
		// windows系统
		return filepath.ToSlash(path)
	}

	return path
}

//检测文件名是否含有空格等特殊符号，如果有，则重命名
func CheckFileName(fileName string) string {
	// 去除空格
	fileName = strings.Replace(fileName, " ", "", -1)
	// 去除换行符
	fileName = strings.Replace(fileName, "\n", "", -1)
	return fileName
}
