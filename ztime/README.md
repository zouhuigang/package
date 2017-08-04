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


### 8.时间戳转换

	从字符串转为时间戳，第一个参数是格式，第二个是要转换的时间字符串
	tm2, _ := time.Parse("01/02/2006", "02/08/2015")
	fmt.Println(tm2.Unix())
	fmt.Println(time.Now().Unix()) //获取当前秒
	fmt.Println(time.Now().UnixNano())//获取当前纳秒
	fmt.Println(time.Now().UnixNano()/1e6)//将纳秒转换为毫秒
	fmt.Println(time.Now().UnixNano()/1e9)//将纳秒转换为秒
	c := time.Unix(time.Now().UnixNano()/1e9,0) //将毫秒转换为 time 类型
	fmt.Println(c.String()) //输出当前英文时间戳格式  


### 9.当月时间段,月初-月末时间戳,方面mysql查询数据

	s1, e1 := ztime.MonthStartEndOfTimeStamp(2017, 07)
	fmt.Printf("%v %v\n", s1, e1)

输出:

	1498838400 1501516799

分别对应：

	2017/7/1 0:0:0 2017/7/31 23:59:59


### 10.int64->string或int64转time.Time

	s2 := ztime.Int64ToTime(1501603200)//中国时间，东八区时间
	fmt.Printf("%v %v\n", s2, ztime.Date("Y-m-d H:i:s", s2)) 

	s3 := ztime.DateInt64("Y-m-d H:i:s", 1501603200)//中国时间，东八区时间
	fmt.Printf("%v\n", s3)


### 11.得到当前时间戳（东八区）

	timestamp := ztime.NowTimeStamp()


### 12.得到指定日期的周，上一周，下一周信息

	sw6, w6 := ztime.SWeek(2017, 8, 3)
	fmt.Printf("%v,%v\n", sw6, w6)

输出：

	[2017-07-31 2017-08-01 2017-08-02 2017-08-03 2017-08-04 2017-08-05 2017-08-06],4
	格式说明：[星期一，星期二，星期三，星期四，星期五，星期六，星期日],当前日期星期四


### 13.输入指定日期，得到上一周周一信息和下一周周一信息

	y1, m1, d1 := ztime.PreWeekMon(2017, 8, 3)
	y2, m2, d2 := ztime.NextWeekMon(2017, 8, 3)
	fmt.Printf("[%v-%v-%v],[%v-%v-%v]\n", y1, m1, d1, y2, m2, d2)

输出：[2017-7-24],[2017-8-7]


### 14.得到一周具体的信息


	package main

	import (
		"fmt"
		"github.com/zouhuigang/package/ztime"
	)
	
	func main() {
		//得到2017-07-26这天所在的周
		weekInfo, w1 := ztime.SWeek(2017, 7, 26)
		fmt.Printf("%v,%v\n", weekInfo, w1)
	
		getWeekInfo := ztime.GetWeekInfo(weekInfo, "/")
		fmt.Printf("周详细信息:\n%v\n", getWeekInfo)
	
	}


输出：


	周详细信息:
	[{2017-07-24 07.24 一} {2017-07-25 07.25 二} {2017-07-26 07.26 三} {2017-07-27 0
	7.27 四} {2017-07-28 07.28 五} {2017-07-29 07.29 六} {2017-07-30 07.30 日}]