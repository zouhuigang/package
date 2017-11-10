package main

import (
	"fmt"
	"github.com/zouhuigang/package/ztime"
)

func main() {
	utc_time := `1510329455` //2017/11/10 23:57:35.168935
	y, m, d, ymd, q := ztime.GetTodayYMD()
	fmt.Printf("y%v, m%v, d%v, ymd%v, q%v\n", y, m, d, ymd, q)
	nowtime := ztime.NowTimeStamp()
	fmt.Printf("utc:%v,now:%v,nowtime:%v\n", utc_time, q, nowtime)

}
func main2() {

	for i := 1; i <= 12; i++ {
		q := ztime.GetQuarterByMonth(i)
		fmt.Printf("%v月,第%d季度\n", i, q)

	}

}
