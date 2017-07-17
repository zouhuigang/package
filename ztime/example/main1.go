package main

import (
	"fmt"
	"github.com/zouhuigang/package/ztime"
)

func main() {
	//日期补零
	fmt.Printf("%v\n%v\n%v\n%v\n", ztime.FormatDateTime("2012-11-12 23:32:1"), ztime.FormatDateTime("2012-6-12"), ztime.FormatDateTime("2012-6-12  12:11:1"), ztime.FormatDateTime("2012-6-12  0:0:1"))
}
