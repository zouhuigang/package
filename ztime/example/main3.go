package main

import (
	"fmt"
	"github.com/zouhuigang/package/ztime"
)

func main() {
	y, m, d, ymd, timestamp := ztime.GetTodayYMD()
	fmt.Printf("今天的时间是:%v年%v月%v [%v] 时间戳:[%v]\n", y, m, d, ymd, timestamp)

	timestamp1 := ztime.Is_Today("2017-7-18")
	fmt.Printf("%v %v\n", timestamp1, ztime.Is_Today("2019-7-18"))

	s1, e1 := ztime.MonthStartEndOfTimeStamp(2017, 8)
	fmt.Printf("%v %v\n", s1, e1)

}
