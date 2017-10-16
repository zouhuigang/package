package main

import (
	"fmt"
	"github.com/zouhuigang/package/zphone/txt2dat"
)

func main() {
	zphoneWrite := txt2dat.LoadTxtFile(`D:\mnt\anooc_go\src\github.com\zouhuigang\package\zphone\txt2dat\example\output.txt`)

	err := zphoneWrite.WriteDat("phone2.dat", "1705") //输出的文件名和版本号，版本号为4位
	if err != nil {
		fmt.Printf("write error: %s", err)
	}

	fmt.Printf("write success..")

}
