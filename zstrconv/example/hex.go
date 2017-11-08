package main

import (
	"fmt"
	"github.com/zouhuigang/package/zstrconv"
)

func main() {
	fmt.Println(zstrconv.Hex10ToAny(20013, 16))
	fmt.Println(zstrconv.AnyToHex10("4e2d", 16))

	fmt.Println(zstrconv.AnyToHex10("0061", 16))

}
