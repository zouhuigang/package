package main

import (
	"fmt"
	"github.com/zouhuigang/package/ztime"
)

func main() {
	year, month, monthData := ztime.ShowMonth(2017, 7)
	fmt.Printf("%v年%v月\n", year, month)
	for k, v := range monthData {
		fmt.Printf("日历第%v行,数据为:%v\n", k+1, v)
	}

	year, month, monthData = ztime.ShowMonth(2019, 2)
	fmt.Printf("%v年%v月\n", year, month)
	for k, v := range monthData {
		fmt.Printf("日历第%v行,数据为:%v\n", k+1, v)
	}

	year, month, monthData = ztime.ShowMonth(2017, 10)
	fmt.Printf("%v年%v月\n", year, month)
	for k, v := range monthData {
		fmt.Printf("日历第%v行,数据为:%v\n", k+1, v)
	}

}
