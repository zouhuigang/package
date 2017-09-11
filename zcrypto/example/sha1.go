package main

import (
	"fmt"
	"github.com/zouhuigang/package/zcrypto"
)

func main() {
	str := `jsapi_ticket=kgt8ON7yVITDhtdwci0qedb8EuKl7VzW2NoBNJA819yQXNy4bd6IlLzxolhEatYfgOdvteSiqGXQlbmgsCusDQ&noncestr=spybo2yt3ohu4jr8yaw6ik6vl3k6vhpg&timestamp=1505096462&url=https://www.anooc.com/edu/teacher/scan`

	sha1 := zcrypto.PhpSha1(str)
	fmt.Println(sha1)
}
