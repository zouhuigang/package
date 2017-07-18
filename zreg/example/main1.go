package main

import (
	"fmt"
	"github.com/zouhuigang/package/zreg"
)

func main() {

	//为空
	s := []string{"", "zou4009@qq.com", " 邹慧刚"}
	for _, v := range s {
		fmt.Printf("是否为空%v : %v\n", v, zreg.IsNull(v))
	}

}
