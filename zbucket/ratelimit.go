/*
令牌桶
*/

package zbucket

import (
	"github.com/juju/ratelimit"
	"sync"
	"time"
)

type ZbucketOperationInterface struct {

	//所有的接口操作都在这个结构体下
	BucketFillDuring time.Duration //存入新令牌的时间间隔,单位毫秒
	BucketMax        int64         //令牌池中，最多能容纳多少令牌
}

var tokenBuckets map[string]*ratelimit.Bucket //每个用户一个令牌池
func init() {
	tokenBuckets = make(map[string]*ratelimit.Bucket)
}

/*
判断用户流量是否放行
*/
func (self ZbucketOperationInterface) JuBuckets(tokenBucketName string) bool { //判断令牌桶
	var ok bool
	var tokenBucket *ratelimit.Bucket = nil
	var lock sync.Mutex

	lock.Lock()
	tokenBucket, ok = tokenBuckets[tokenBucketName]
	lock.Unlock()

	if !ok {
		tokenBucket = ratelimit.NewBucket(self.BucketFillDuring, self.BucketMax)

		lock.Lock()
		tokenBuckets[tokenBucketName] = tokenBucket
		lock.Unlock()
	}
	available := tokenBucket.TakeAvailable(1)
	if available <= 0 {
		/*** 限流处理 ***/
		return false
	}

	return true
}
