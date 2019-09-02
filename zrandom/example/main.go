package main

import (
	"fmt"

	"github.com/zouhuigang/package/zrandom"
)

func main() {
	s := zrandom.Digits(6)
	fmt.Println(s)
}
