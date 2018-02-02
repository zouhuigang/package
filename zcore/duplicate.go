/*
去重
*/
package zcore

import (
	//"fmt"
	"strings"
)

//去重和去空 (空间换时间)
func RemoveDuplicatesAndEmpty(a []string) (ret []string) {

	tempMap := map[string]byte{} // 存放不重复主键

	for _, e := range a {

		if e == "" || len(strings.TrimSpace(e)) == 0 {
			continue
		}
		//fmt.Printf("%d-%s\n", k, e)
		l := len(tempMap)
		tempMap[e] = 0
		if len(tempMap) != l { // 加入map后，map长度变化，则元素不重复
			ret = append(ret, e)
		}
	}

	return
}
