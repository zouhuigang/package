package main

import (
	"fmt"
	"github.com/zouhuigang/package/zcore"
)

func main() {
	int1 := zcore.StringToInt("1000", 9)
	int2 := zcore.StringToInt("", 2)
	fmt.Println(int1, int2)
}
