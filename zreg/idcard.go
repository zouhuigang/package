package zreg

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/zouhuigang/package/ztime"
)

/*
362202199608096171
前1、2位数字表示：所在省份的代码；
第3、4位数字表示：所在城市的代码；
第5、6位数字表示：所在区县的代码；
第7~14位数字表示：出生年、月、日；
第15、16位数字表示：所在地的派出所的代码；
第17位数字表示性别：奇数表示男性，偶数表示女性；
第18位数字是校检码：也有的说是个人信息码，一般是随计算机的随机产生， 用来检验身份证的正确性。校检码可以是0~9的数字，有时也用x表示
*/

const m_GENDER_MALE string = "M"   //1
const m_GENDER_FEMALE string = "F" //0

//解析身份证中的信息
type IdCard struct {
	Length   int
	Year     string
	Month    string
	Day      string
	Sex      string
	Birthday time.Time
}

//解析出性别
func getGender(idLength int, idNumber string) string {
	var genderStr string
	var gender string
	if idLength == 18 {
		gender = string(idNumber[16]) //16
	} else {
		gender = string(idNumber[14]) //14
	}

	genderInt, _ := strconv.Atoi(gender)

	if genderInt%2 == 0 {
		genderStr = m_GENDER_FEMALE
	} else {
		genderStr = m_GENDER_MALE
	}
	return genderStr
}

//解析出生日
func getBirthday(idLength int, idNumber string) string {
	var birthday string
	if idLength == 18 {
		birthday = string(idNumber[6:14])
	} else {
		birthday = "19" + string(idNumber[6:12])
	}
	return birthday
}

//解析出年月日
func getBirthdayYMD(idLength int, idNumber string) (string, string, string) {
	birthday := getBirthday(idLength, idNumber)
	year := string(birthday[0:4])
	month := string(birthday[4:6])
	day := string(birthday[6:8])
	return year, month, day
}

//解析出身份证,身份证最少为15位
func PraseIdCard(idCard string) (*IdCard, error) {
	// 去除空格
	idNumber := strings.Replace(idCard, " ", "", -1)
	// 去除换行符
	idNumber = strings.Replace(idCard, "\n", "", -1)
	idLength := len(idNumber)
	if idLength < 15 {
		return nil, errors.New("身份证错误")
	}

	myidcard := new(IdCard)
	myidcard.Length = idLength
	myidcard.Sex = getGender(idLength, idNumber)
	myidcard.Year, myidcard.Month, myidcard.Day = getBirthdayYMD(idLength, idNumber)
	birthday := fmt.Sprintf("%s-%s-%s", myidcard.Year, myidcard.Month, myidcard.Day)
	myidcard.Birthday = ztime.Strtotime(birthday)

	return myidcard, nil
}
