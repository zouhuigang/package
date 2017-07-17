package main

import (
	"fmt"
	"github.com/zouhuigang/package/ztime"
)

func main() {
	daynum := ztime.GetMonthDayNum(2017, 2)
	fmt.Println(daynum)

	daynum = ztime.GetMonthDayNum(2000, 2)
	fmt.Println(daynum)

	w1 := ztime.GetWeekday("2017", "07", "01")
	fmt.Println(w1)

	w2 := ztime.GetWeekday("2017", "07", "9")
	fmt.Println(w2)
	w3 := ztime.GetWeekdayNum("2017", "07", "9")
	fmt.Println(w3)

	//php strtotime 将字符串转时间戳
	fmt.Printf("%v\n%v\n", ztime.Strtotime("2017-07-12"), ztime.Strtotime("2017-7-2"))

	//将1,2,3,4,5,6,01,02,03等统一格式化为:01,02,03,04
	fmt.Println(ztime.FormatNum("7"), ztime.FormatNum("10"), ztime.FormatNum("0"), ztime.FormatNum("02"))

}
