//https://github.com/polaris1119/times/blob/master/times.go
package ztime

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

//设置时区
/*var countryTz = map[string]string{
    "Local":   "Local",
    "Hungary": "Europe/Budapest",
    "Egypt":   "Africa/Cairo",
    "China":   "Asia/Shanghai",
    "America": "America/Los_Angeles",
}
*/
const DefaultTimeZone = "Asia/Shanghai" //"Asia/Chongqing" ,"Asia/Shanghai" "Local"

var days = [...]string{
	"日",
	"一",
	"二",
	"三",
	"四",
	"五",
	"六",
}
var daysMon = [...]string{
	"一",
	"二",
	"三",
	"四",
	"五",
	"六",
	"日",
}

//数字补0
func FormatNum(num string) string {
	b, err := strconv.Atoi(num)
	if err != nil {
		return err.Error()
	}
	var snum string
	if b < 10 && b >= 0 {
		snum = "0" + strconv.Itoa(b)
	} else {
		snum = num
	}

	return snum
}

//数字补0
func FormatNumInt(n int) string {
	var snum string
	if n < 10 && n >= 0 {
		snum = "0" + strconv.Itoa(n)
	} else {
		snum = strconv.Itoa(n)
	}

	return snum
}

//日期补0,"2012-11-12 23:32:1" 或 "2012-6-12"变成2012-11-12 23:32:01 和2012-06-12
func FormatDateTime(value string) string {
	value = strings.Replace(value, " ", ",", -1) //将空格变成,
	s := Explode(",", value)
	var bstr string
	for _, v := range s {
		if v == "" { //有空值
			continue
		}
		s1 := Explode("-", v)
		if len(s1) > 1 { //年月日
			for k, v1 := range s1 {
				if k == 0 {
					bstr += FormatNum(v1)
				} else {
					bstr += "-" + FormatNum(v1)
				}
			}
		} else { //时分秒
			s2 := Explode(":", v)
			if len(s2) > 1 {
				for k1, v2 := range s2 {
					if k1 == 0 {
						bstr += " " + FormatNum(v2)
					} else {
						bstr += ":" + FormatNum(v2)
					}
				}

			} else {
				bstr += FormatNum(v)
			}
		}
	}

	return bstr
}

//$str = "Hello world. I love Shanghai!";
//print_r (explode(" ",$str));
func Explode(separator string, str string) []string {
	stringSlice := strings.Split(str, separator)
	return stringSlice
}

/*func formatDate(date string) string {
	dates := strings.Split(strings.TrimSpace(date), "/")
	return dates[0] + "/" + dates[1] + "/20" + dates[2]
}*/

//得到x年x月x日是星期几
func GetWeekday(y, m, d string) string {
	t := Strtotime(y + "-" + FormatNum(m) + "-" + FormatNum(d)) //time.Time
	//t := time.Now() //时间戳//time.Time
	wday := int(t.Weekday())
	s1 := days[wday]
	//s1 := t.Weekday().String() //英文
	return s1
}

func GetWeekdayNum(y, m, d string) int {
	t := Strtotime(y + "-" + FormatNum(m) + "-" + FormatNum(d))
	wday := int(t.Weekday())
	return wday
}

// Date 跟 PHP 中 date 类似的使用方式，如果 ts 没传递，则使用当前时间
func DateInt64(format string, ts int64) string { //int64时间戳格式化 int64->string
	ts1 := time.Unix(ts, 0)
	return Date(format, ts1)
}
func Date(format string, ts ...time.Time) string {
	patterns := []string{
		// 年
		"Y", "2006", // 4 位数字完整表示的年份
		"y", "06", // 2 位数字表示的年份

		// 月
		"m", "01", // 数字表示的月份，有前导零
		"n", "1", // 数字表示的月份，没有前导零
		"M", "Jan", // 三个字母缩写表示的月份
		"F", "January", // 月份，完整的文本格式，例如 January 或者 March

		// 日
		"d", "02", // 月份中的第几天，有前导零的 2 位数字
		"j", "2", // 月份中的第几天，没有前导零

		"D", "Mon", // 星期几，文本表示，3 个字母
		"l", "Monday", // 星期几，完整的文本格式;L的小写字母

		// 时间
		"g", "3", // 小时，12 小时格式，没有前导零
		"G", "15", // 小时，24 小时格式，没有前导零
		"h", "03", // 小时，12 小时格式，有前导零
		"H", "15", // 小时，24 小时格式，有前导零

		"a", "pm", // 小写的上午和下午值
		"A", "PM", // 小写的上午和下午值

		"i", "04", // 有前导零的分钟数
		"s", "05", // 秒数，有前导零
	}
	replacer := strings.NewReplacer(patterns...)
	format = replacer.Replace(format)

	t := time.Now()
	if len(ts) > 0 {
		t = ts[0]
	}

	loc, err := time.LoadLocation(DefaultTimeZone)
	if err != nil {
		//panic(err)
	}

	return t.In(loc).Format(format)
}

////value:"2012-11-12 23:32:01" 或 "2012-11-12"
func Strtotime(value string) time.Time { //给当前日期加上后面的+0800 CST
	if value == "" {
		return time.Time{}
	}

	//格式化日期,将2017-1-1和2017-01-01都统一变成2017-01-01
	value = FormatDateTime(value)

	loc, err := time.LoadLocation(DefaultTimeZone)
	if err != nil {
	}

	zoneName, offset := time.Now().In(loc).Zone()
	zoneValue := offset / 3600 * 100
	if zoneValue > 0 {
		value += fmt.Sprintf(" +%04d", zoneValue)
	} else {
		value += fmt.Sprintf(" -%04d", zoneValue)
	}

	if zoneName != "" {
		value += " " + zoneName
	}
	return StrToTime(value)
}

func StrToTime(value string) time.Time {
	if value == "" {
		return time.Time{}
	}
	layouts := []string{
		"2006-01-02 15:04:05 -0700 MST",
		"2006-01-02 15:04:05 -0700",
		"2006-01-02 15:04:05",
		"2006/01/02 15:04:05 -0700 MST",
		"2006/01/02 15:04:05 -0700",
		"2006/01/02 15:04:05",
		"2006-01-02 -0700 MST",
		"2006-01-02 -0700",
		"2006-01-02",
		"2006/01/02 -0700 MST",
		"2006/01/02 -0700",
		"2006/01/02",
		"2006-01-02 15:04:05 -0700 -0700",
		"2006/01/02 15:04:05 -0700 -0700",
		"2006-01-02 -0700 -0700",
		"2006/01/02 -0700 -0700",
		time.ANSIC,
		time.UnixDate,
		time.RubyDate,
		time.RFC822,
		time.RFC822Z,
		time.RFC850,
		time.RFC1123,
		time.RFC1123Z,
		time.RFC3339,
		time.RFC3339Nano,
		time.Kitchen,
		time.Stamp,
		time.StampMilli,
		time.StampMicro,
		time.StampNano,
	}

	var t time.Time
	var err error
	for _, layout := range layouts {
		t, err = time.Parse(layout, value)
		if err == nil {
			return t
		}
	}
	panic(err)
}

//得到当天的时间,凌晨时间(2017,7,18,"2017-07-18",1500307200)
func GetTodayYMD() (y int, m int, d int, ymd string, timestamp int64) {
	// 一般为CST
	loc, _ := time.LoadLocation(DefaultTimeZone)
	y, m, d = int(time.Now().In(loc).Year()), int(time.Now().In(loc).Month()), int(time.Now().In(loc).Day())
	ymd = time.Now().In(loc).Format("2006-01-02")
	//t, _ := time.Parse("2006-01-02", ymd) //转化为utc时间,得到的是8点的时间
	t, _ := time.ParseInLocation("2006-01-02", ymd, loc) //得到零点的时间
	timestamp = t.Unix()
	return
}

//判断输入的日期是否为当天时间ymd:2017-07-18或2017-7-18,不能输入2017-7-0
func Is_Today(ymd string) bool {
	_, _, _, _, timestamp := GetTodayYMD() //当天凌晨时间戳
	tm := Int64ToTime(timestamp)           //int64转time.Time
	//fmt.Println("time===", tm, timestamp)
	timestamp1 := Strtotime(ymd)
	//fmt.Printf("%v,%v\n", tm, timestamp1)
	if tm.Equal(timestamp1) { //time.Time的比较不能用==来比较,在docker镜像里面会出错
		//fmt.Printf("%v,%v\n", tm, timestamp1)
		return true
	}

	return false
}

//将日期格式化成时间戳(秒)
func StrToTimestamp(ymdhis string) int64 {
	t := Strtotime(ymdhis)
	return t.Unix()
}

//int64转time.Time时间
//var ymdDate string = time.Unix(v, 0).In(DefaultTimeZone).Format("2006-01-02 15:04:05") //int64转string时间
func Int64ToTime(v int64) time.Time {
	loc, err := time.LoadLocation(DefaultTimeZone)
	if err != nil {
		//panic(err)
	}
	return time.Unix(v, 0).In(loc)
}

//得到当前时间戳
func NowTimeStamp() int64 {
	// 一般为CST  t := time.Now().Unix()
	loc, _ := time.LoadLocation(DefaultTimeZone)
	nowtime := time.Now().In(loc).Format("2006-01-02 15:04:05 PM")
	//t, _ := time.Parse("2006-01-02 15:04:05 PM", nowtime)
	t, _ := time.ParseInLocation("2006-01-02 15:04:05 PM", nowtime, loc)
	return t.Unix()
}

//得到指定日期的时间戳
//http://www.01happy.com/golang-time/
/*func GetTimeStamp(ymdhis string) int64 {
	loc, _ := time.LoadLocation(DefaultTimeZone)
	t, _ := time.ParseInLocation("2006-01-02 15:04:05 PM", ymdhis, loc)
	return t.Unix()
}*/

type wList struct {
	NowDate  string //日期2017-01-30
	MonthDay string //01/30
	Whatday  string //星期几
	IsToday  bool   //当前日期是否是今天
}

//将一周的日历拆分成一周的详情,sp为`/`或`.`
func GetWeekInfo(weekInfo [7]string, sp string) []wList {

	var week []wList
	week = make([]wList, 0)
	for k, v := range weekInfo {
		y1, m1, d1 := SliptDate(v)
		md := fmt.Sprintf("%s%s%s", FormatNumInt(m1), sp, FormatNumInt(d1))
		todate := fmt.Sprintf("%s-%s-%s", FormatNumInt(y1), FormatNumInt(m1), FormatNumInt(d1))
		isToday := Is_Today(todate) //2017-7-18
		w := wList{v, md, daysMon[k], isToday}
		week = append(week, w)
	}
	return week
}
