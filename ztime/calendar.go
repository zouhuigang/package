/*
返回日历
*/
package ztime

import (
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

//下一个月
func NextMonth(y, m int) (next_year int, next_month int) {
	_, y = LimitYear(y)
	_, m = LimitYear(m)

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
	_, m = LimitYear(m)

	pre_month = m - 1
	if pre_month < 1 {
		pre_year = y - 1
		pre_month = 12
	} else {
		pre_year = y
	}
	return
}
