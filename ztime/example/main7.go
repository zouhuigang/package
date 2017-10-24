package main

import (
	"fmt"
	"github.com/zouhuigang/package/ztime"
)

func main() {

	for i := 1; i <= 12; i++ {
		q := ztime.GetQuarterByMonth(i)
		fmt.Printf("%v月,第%d季度\n", i, q)

	}

}
