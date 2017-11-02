package main

import (
	"fmt"
	"github.com/zouhuigang/package/zbucket"
	"time"
)

func main() {
	var bucketFillDuring time.Duration = time.Millisecond * 1000 * (60 * 60 * 10) //单位秒,每隔10个小时存入一个令牌
	var bucketMax int64 = 1                                                       //存入的令牌

	stuBucket := zbucket.ZbucketOperationInterface{bucketFillDuring, bucketMax}

	s1 := stuBucket.JuBuckets("1")
	if s1 {
		fmt.Println("1 pass\n")
	} else {
		fmt.Println("1 not pass\n")
	}

	s1 = stuBucket.JuBuckets("1")
	if s1 {
		fmt.Println("1 pass\n")
	} else {
		fmt.Println("1 not pass\n")
	}

	s2 := stuBucket.JuBuckets("2")
	if s2 {
		fmt.Println("2 pass\n")
	} else {
		fmt.Println("2 not pass\n")
	}

	s2 = stuBucket.JuBuckets("2")
	if s2 {
		fmt.Println("2 pass\n")
	} else {
		fmt.Println("2 not pass\n")
	}

}
