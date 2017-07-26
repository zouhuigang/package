package main

import (
	"fmt"
	"github.com/zouhuigang/package/ztime"
)

func main() {
	timestamp := ztime.NowTimeStamp()
	fmt.Println(timestamp)

	sw1, w1 := ztime.SWeek(2017, 7, 26)
	fmt.Printf("%v,%v\n", sw1, w1)

	sw2, w2 := ztime.SWeek(2017, 7, 23)
	fmt.Printf("%v,%v\n", sw2, w2)

	sw3, w3 := ztime.SWeek(2017, 7, 24)
	fmt.Printf("%v,%v\n", sw3, w3)

	sw5, w5 := ztime.SWeek(2017, 7, 29)
	fmt.Printf("%v,%v\n", sw5, w5)

	sw4, w4 := ztime.SWeek(2017, 7, 30)
	fmt.Printf("%v,%v\n", sw4, w4)

	sw6, w6 := ztime.SWeek(2017, 8, 3)
	fmt.Printf("%v,%v\n", sw6, w6)

	y1, m1, d1 := ztime.PreWeekMon(2017, 8, 3)
	y2, m2, d2 := ztime.NextWeekMon(2017, 8, 3)
	fmt.Printf("[%v-%v-%v],[%v-%v-%v]\n", y1, m1, d1, y2, m2, d2)
}
