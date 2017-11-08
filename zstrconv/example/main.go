package main

import (
	"fmt"
	"github.com/zouhuigang/package/zstrconv"
)

func main() {
	s := "abc中"
	slen := len(s)
	fmt.Printf("%s,%d", s, slen)
}

func main3() {
	s := []byte("abc中")
	t := zstrconv.Tool_DecimalByteSlice2HexString(s) //16进制
	fmt.Printf("%s\n%x", t, s)
}

func main2() {
	s := zstrconv.OnlyChineseToUnicode(`中过asdasdfrwer1312asf撒打算`)
	fmt.Println(s)

	//转换为中文,`"\u` + strVal + `"`
	s2, err := zstrconv.UnicodeToChinese(`"\u` + `0061` + `"`) //\u0061\u0062\u0063\u0064
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%s\n", s2)
}
