### golang时间操作包，日历包

说明：日历最多6行，所以可以固定6行数目。

### 1.时间日期补零

	fmt.Printf("%v\n%v\n%v\n%v\n", ztime.FormatDateTime("2012-11-12 23:32:1"), 
	ztime.FormatDateTime("2012-6-12"), ztime.FormatDateTime("2012-6-12  12:11:1"), 
	ztime.FormatDateTime("2012-6-12  0:0:1"))

输出：

	
	2012-11-12 23:32:01
	2012-06-12
	2012-06-12 12:11:01
	2012-06-12 0:0:01

### 2.得到年月日属于星期几

	w2 := ztime.GetWeekday("2017", "07", "9")
	fmt.Println("星期"+w2)
	w3 := ztime.GetWeekdayNum("2017", "07", "9")
	fmt.Println(w3)

输出：

	星期日
	0


### 3.判断闰年和平年

	b:=ztime.Is_LeapYear(2017)

输出：

	false //不是闰年


### 4.得到一个月有多少天

	num:=ztime.GetMonthDayNum(2017,2)

输出：

	28


### 5.日历功能

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
	
	}

输出:

	
	2017年7月
	日历第1行,数据为:[0 0 0 0 0 0 1]
	日历第2行,数据为:[2 3 4 5 6 7 8]
	日历第3行,数据为:[9 10 11 12 13 14 15]
	日历第4行,数据为:[16 17 18 19 20 21 22]
	日历第5行,数据为:[23 24 25 26 27 28 29]
	日历第6行,数据为:[30 31 0 0 0 0 0]
	2019年2月
	日历第1行,数据为:[0 0 0 0 0 1 2]
	日历第2行,数据为:[3 4 5 6 7 8 9]
	日历第3行,数据为:[10 11 12 13 14 15 16]
	日历第4行,数据为:[17 18 19 20 21 22 23]
	日历第5行,数据为:[24 25 26 27 28 0 0]

0代表：空。

[0,0,0,0,0,0,0] 代表: [日,一，二，三，四，五，六]



### 6.今天的时间，当地时区

	y, m, d, ymd, timestamp := ztime.GetTodayYMD()

输出：

	今天的时间是:2017年7月18 [2017-07-18] 时间戳:[1500307200]

### 7.判断时间是否是当天时间

	timestamp1 := ztime.Is_Today("2017-7-18")
	fmt.Printf("%v %v\n", timestamp1, ztime.Is_Today("2019-7-18"))

输出：

	true,false