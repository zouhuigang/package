package ztime

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

//最简单的计算年龄
func GetAge(year int) (age int) {
	if year <= 0 {
		age = -1
	}
	nowyear := time.Now().Year()
	age = nowyear - year
	return
}

//根据1992-01-01分割出生日
func ParseBirthday(birthday string, spec string) (string, string, string, error) {

	if spec != "" {
		birthday = strings.Replace(birthday, spec, "", -1)
	}
	idLength := len(birthday)
	if idLength != 8 {
		return "", "", "", errors.New("生日非法")
	}
	year := string(birthday[0:4])
	month := string(birthday[4:6])
	day := string(birthday[6:8])
	return year, month, day, nil
}

func ParseBirthdayInt(birthday string, spec string) (int, int, int, error) {

	ys, ms, ds, err := ParseBirthday(birthday, spec)
	if err != nil {
		return 0, 0, 0, err
	}
	var year, month, day int
	year, err = strconv.Atoi(ys)
	if err != nil {
		return 0, 0, 0, err
	}
	month, err = strconv.Atoi(ms)
	if err != nil {
		return 0, 0, 0, err
	}
	day, err = strconv.Atoi(ds)
	if err != nil {
		return 0, 0, 0, err
	}
	return year, month, day, nil
}
