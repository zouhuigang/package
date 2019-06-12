package zfile

import (
	"os"
	"path/filepath"
	"strings"
)

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
