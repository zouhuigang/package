/*
时间的转换类

*/
package ztime

import (
	"math"
)

//得到月份属于第几季度
func GetQuarterByMonth(month int) int {
	m := float64(month)
	q := math.Ceil(m / 3.0)
	return int(q)
}
