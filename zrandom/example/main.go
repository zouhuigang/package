package main

import (
	"fmt"

	"github.com/zouhuigang/package/zrandom"
)

func main() {
	s := zrandom.Digits(6)
	s1 := zrandom.DigitsToString(6)
	fmt.Println(s, s1)
}
