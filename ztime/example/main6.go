package main

import (
	"fmt"
	"github.com/zouhuigang/package/ztime"
)

func main() {

	timestamp1 := ztime.StrToTimestamp("2016-07-27 08:46:15")
	timestamp2 := ztime.StrToTimestamp("2016-07-27 0:0:0")
	fmt.Printf("%v\n%v\n", timestamp1, timestamp2)
	timestamp3 := ztime.Strtotime("2016-07-27 0:0:0")
	fmt.Printf("%v\n", timestamp3)

}
