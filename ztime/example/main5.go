package main

import (
	"fmt"
	"github.com/zouhuigang/package/ztime"
)

func main() {
	//得到2017-9-1这天所在的周
	weekInfo, w1 := ztime.SWeek(2017, 9, 1)
	fmt.Printf("%v,%v\n", weekInfo, w1)

	getWeekInfo := ztime.GetWeekInfo(weekInfo, ".")
	fmt.Printf("周详细信息:\n%v\n", getWeekInfo)

}
