/*
返回日历
*/
package ztime

import (
	"fmt"
	"strconv"
	"time"
)

//展示日历,返回年月及当月信息情况
func ShowMonth(year, month int) (int, int, [][]int) {
	if month > 12 { //处理出现月份大于12的情况
		month = 1
		year++
	}
	if month < 1 { //处理出现月份小于1的情况
		month = 12
		year--
	}

	days := GetMonthDayNum(year, month)                                      //得到一个月有多少天
	dayofweek := GetWeekdayNum(strconv.Itoa(year), strconv.Itoa(month), "1") //得到给定的月份的 1号 是星期几
	monthData := CShowDate(days, dayofweek)
	return year, month, monthData
}

//删除函数
//func remove(s []string, i int) []string {
//	return append(s[:i], s[i+1:]...)
//}

//返回日历的月份和前一个月及后一个月的链接
func CShowDate(days int, dayofweek int) [][]int {
	var month_data [][]int
	var line_data []int        //存储每一行数据
	firstnums := dayofweek + 1 //1号位于第几个

	line_data = make([]int, 7)  //存储每一行数据,星期一到星期日，共7天
	for i := 0; i < days; i++ { //输出天数信息
		mi := firstnums % 7
		if mi == 1 { //每行等于1的时候，保存之前的值，并创建存储
			month_data = append(month_data, line_data)
			line_data = make([]int, 7)
		}

		if mi == 0 {
			line_data[6] = i + 1
		} else {
			line_data[mi-1] = i + 1
		}

		firstnums++
	}
	month_data = append(month_data, line_data) //最后一行数据

	//排除第一行全是0的值,即第一行没有当前月份的日期
	var first_del int = 0
	for _, v := range month_data[0] {
		if v == 0 {
			first_del++
		}

	}
	if first_del == len(month_data[0]) {
		index := 0
		month_data = append(month_data[:index], month_data[index+1:]...)
	}

	return month_data

}

//判断是否是闰年
//1、非整百年：能被4整除的为闰年。（如2004年就是闰年,2100年不是闰年）
//2、整百年：能被400整除的是闰年。(如2000年是闰年，1900年不是闰年)
func Is_LeapYear(year int) bool {
	//isLeapYear = (year % 4 == 0 && year % 100 != 0) || (year % 400 == 0) //判断是否是闰年
	if (year%4 == 0 && year%100 != 0) || (year%400 == 0) {
		return true
	}

	return false
}

//得到某年某月有多少天
//1,3,5,7,8,10,12有31天
//平年二月有28天,闰年的2月有29天
//4,6,9,11有30天
func GetMonthDayNum(year, month int) int {
	var numberOfDays int
	if month == 4 || month == 6 || month == 9 || month == 11 {
		numberOfDays = 30
	} else if month == 2 {
		leap := Is_LeapYear(year)
		if leap {
			numberOfDays = 29
		} else {
			numberOfDays = 28
		}

	} else {
		numberOfDays = 31
	}

	return numberOfDays

}

//windows 1900年-2099年,限制年份范围,如果不满足，则设置为当前年月
func LimitYear(y int) (bool, int) {
	if y >= 1900 && y <= 2099 {
		return true, y
	}
	return false, int(time.Now().Year())
}

//限制月份1-12
func LimitMonth(m int) (bool, int) {
	if m >= 1 && m <= 12 {
		return true, m
	}
	return false, int(time.Now().Month())
}

//限制天
func LimitDay(y, m, d int) (bool, int) {
	mday := GetMonthDayNum(y, m) //当月有多少天
	if d >= 1 && m <= mday {
		return true, d
	}
	return false, int(time.Now().Day())
}

//下一个月
func NextMonth(y, m int) (next_year int, next_month int) {
	_, y = LimitYear(y)
	_, m = LimitMonth(m)

	next_month = m + 1
	if next_month > 12 {
		next_year = y + 1
		next_month = 1
	} else {
		next_year = y
	}
	return
}

//上一个月
func PreMonth(y, m int) (pre_year int, pre_month int) { //上一个月
	_, y = LimitYear(y)
	_, m = LimitMonth(m)

	pre_month = m - 1
	if pre_month < 1 {
		pre_year = y - 1
		pre_month = 12
	} else {
		pre_year = y
	}
	return
}

//当月时间段,月初-月末时间戳
func MonthStartEndOfTimeStamp(y, m int) (month_start, month_end int64) {
	_, y = LimitYear(y)
	_, m = LimitMonth(m)
	m_days := GetMonthDayNum(y, m)
	m_start := strconv.Itoa(y) + "-" + strconv.Itoa(m) + "-" + "01"
	m_end := strconv.Itoa(y) + "-" + strconv.Itoa(m) + "-" + strconv.Itoa(m_days) + " " + "23:59:59"
	//转化为时间戳
	month_start = StrToTimestamp(m_start)
	month_end = StrToTimestamp(m_end)
	//month_start = 1
	//month_end = 1
	fmt.Println(m_start, m_end)
	return
}

func NextDay(y, m, d int) (next_year int, next_month int, next_day int) { //下一天
	_, y = LimitYear(y)
	_, m = LimitMonth(m)
	_, d = LimitDay(y, m, d)

	mday := GetMonthDayNum(y, m) //当月有多少天
	next_day = d + 1
	if next_day > mday { //进入下一个月
		next_year, next_month = NextMonth(y, m)
		next_day = 1
	} else {
		next_year = y
		next_month = m
	}
	return
}

func PreDay(y, m, d int) (pre_year int, pre_month int, pre_day int) { //上一天
	_, y = LimitYear(y)
	_, m = LimitMonth(m)
	_, d = LimitDay(y, m, d)

	pre_day = d - 1
	if pre_day <= 0 { //进入上一个月
		pre_year, pre_month = PreMonth(y, m)
		pre_day = GetMonthDayNum(pre_year, pre_month) //下一个月有多少天
	} else {
		pre_year = y
		pre_month = m
	}
	return
}

//当前时间周,上一周，下一周等,返回当前周的日期，以及当天的星期
func SWeek(y, m, d int) ([7]string, int) { //[一----六日]
	var dayofweek int = GetWeekdayNum(strconv.Itoa(y), strconv.Itoa(m), strconv.Itoa(d)) //得到给定日期是星期几
	if dayofweek == 0 {                                                                  //1，2，3，4，5，6，7代表星期一，星期二....星期日
		dayofweek = 7
	}
	var weekDay [7]string //存储星期数据[一----六日]
	weekDay[dayofweek-1] = strconv.Itoa(y) + "-" + FormatNumInt(m) + "-" + FormatNumInt(d)

	var ny, nm, nd int = y, m, d
	for i := dayofweek + 1; i <= 7; i++ {
		ny, nm, nd = NextDay(ny, nm, nd)
		weekDay[i-1] = strconv.Itoa(ny) + "-" + FormatNumInt(nm) + "-" + FormatNumInt(nd)
	}

	var py, pm, pd int = y, m, d
	var j int
	for j = dayofweek - 1; j > 0; j-- {
		py, pm, pd = PreDay(py, pm, pd)
		weekDay[j-1] = strconv.Itoa(py) + "-" + FormatNumInt(pm) + "-" + FormatNumInt(pd)
	}

	return weekDay, dayofweek
}
func SWeekSun(y, m, d int) ([7]string, int) { //[日一----六]
	var dayofweek int = GetWeekdayNum(strconv.Itoa(y), strconv.Itoa(m), strconv.Itoa(d)) //得到给定日期是星期几
	var weekDay [7]string                                                                //存储星期数据[日一----六]
	weekDay[dayofweek] = strconv.Itoa(y) + "-" + FormatNumInt(m) + "-" + FormatNumInt(d)

	var ny, nm, nd int = y, m, d
	for i := dayofweek + 1; i < 7; i++ {
		ny, nm, nd = NextDay(ny, nm, nd)
		weekDay[i] = strconv.Itoa(ny) + "-" + FormatNumInt(nm) + "-" + FormatNumInt(nd)
	}

	var py, pm, pd int = y, m, d
	var j int
	for j = dayofweek - 1; j >= 0; j-- {
		py, pm, pd = PreDay(py, pm, pd)
		weekDay[j] = strconv.Itoa(py) + "-" + FormatNumInt(pm) + "-" + FormatNumInt(pd)
	}

	return weekDay, dayofweek
}

//分割年月日,如将2017-01-02分割成int
func SliptDate(date string) (int, int, int) {
	s1 := Explode("-", date)
	if len(s1) < 3 {
		return 0, 0, 0
	}
	y, _ := strconv.Atoi(s1[0])
	m, _ := strconv.Atoi(s1[1])
	d, _ := strconv.Atoi(s1[2])
	return y, m, d
}

//上一周周一
func PreWeekMon(y, m, d int) (y1, m1, d1 int) {
	_, y = LimitYear(y)
	_, m = LimitMonth(m)
	_, d = LimitDay(y, m, d)

	nowweek, _ := SWeek(y, m, d)       //得到给定日期的星期信息
	y1, m1, d1 = SliptDate(nowweek[0]) //星期一
	for i := 0; i < 7; i++ {
		y1, m1, d1 = PreDay(y1, m1, d1)
	}
	return
}

//下一周周一信息
func NextWeekMon(y, m, d int) (y1, m1, d1 int) {
	_, y = LimitYear(y)
	_, m = LimitMonth(m)
	_, d = LimitDay(y, m, d)

	nowweek, _ := SWeek(y, m, d)       //得到给定日期的星期信息
	y1, m1, d1 = SliptDate(nowweek[0]) //星期一
	for i := 0; i < 7; i++ {
		y1, m1, d1 = NextDay(y1, m1, d1)
	}
	return
}
