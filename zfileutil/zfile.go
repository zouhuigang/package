package zfileutil

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/zouhuigang/package/zfile"
)

//返回的目录扫描结果
type FileList struct {
	IsDir   bool   //是否是目录
	Path    string //文件路径
	Ext     string //文件扩展名
	Name    string //文件名
	Size    int64  //文件大小
	ModTime int64  //文件修改时间戳
}

//目录扫描
//@param			dir			需要扫描的目录
//@return			fl			文件列表
//@return			err			错误
func ScanFiles(dir string) (fl []FileList, err error) {
	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err == nil {
			path = strings.Replace(path, "\\", "/", -1) //文件路径处理

			//如果含有特殊字符等，则重命名
			var new_file_name string = info.Name()
			var check_file_name string = zfile.CheckFileName(info.Name())
			var new_path string = path
			if check_file_name != new_file_name {
				new_file_name = check_file_name
				new_path = filepath.Join(filepath.Dir(path), new_file_name) //+filepath.Ext(path)
				os.Rename(path, new_path)
			}

			fl = append(fl, FileList{
				IsDir:   info.IsDir(),
				Path:    new_path,
				Ext:     strings.ToLower(filepath.Ext(path)),
				Name:    new_file_name,
				Size:    info.Size(),
				ModTime: info.ModTime().Unix(),
			})
		}
		return err
	})
	return
}
