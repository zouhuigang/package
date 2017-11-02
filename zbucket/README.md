###令牌桶

使用：

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


输出：


1 pass

1 not pass

2 pass

2 not pass





网络文章：
	

	http://blog.imlibo.com/2016/06/20/golang-token-bucket/


### 令牌桶


令牌桶算法是网络流量整形（Traffic Shaping）和速率限制（Rate Limiting）中最常使用的一种算法。典型情况下，令牌桶算法用来控制发送到网络上的数据的数目，并允许突发数据的发送。


	http://blog.imlibo.com/2016/06/20/golang-token-bucket/



Golang令牌桶（Token Bucket）限流的实现Jun 20, 2016
背景：作为一个微信公众号后端的消息处理系统，每天面对最多的流量爆发基本都来自某些公众号做营销活动，例如：抽奖、抢优惠、投票、报名……等等。面对突然到来的上百倍峰值，除了消息队列，预留容量以外，我们开始考虑做峰值限流。因为对于大部分营销类活动，消息限流（对被限流的消息直接丢弃，并直接回复：“系统繁忙，请稍后再试。”）并不会对营销的结果有太大影响。
方案：经过了一些限流方案的调查评估，最终选定了“令牌桶”算法进行限流。因为大规模Docker的使用，侯斯特从今年4月开始全面从PHP转向Go。所以直接找到了一个实现很简单轻量的包：

	https://github.com/juju/ratelimit
令牌桶算法：

	https://en.wikipedia.org/wiki/Token_bucket

### 实现： 最简单的办法是import这个包之后，直接定义一个全局变量：

	var tokenBucket *ratelimit.Bucket = nil

在init()中初始化这个变量：

	bucketFillDuring := time.Millisecond * 200
	bucketMax := 20
	tokenBucket = ratelimit.NewBucket(bucketFillDuring, bucketMax)

这里两个参数分别是：令牌填充的时间间隔、令牌桶的最大容量。用白话解释就是：我们定义了一个最大容量是20个令牌的令牌桶，同时每隔200毫秒向桶中添加一个令牌。也就是说每一个程序实例每秒钟可以处理5条消息，同时因为桶有20个容量，所以对突发请求第一秒可以处理25条消息。那么我们就可以轻易地通过调整开启的实例数量确定当前系统的流量限制了。
下面就是限流了。有了这个桶之后在真正开始处理数据之前，可以很简单实现判断是否要对这次处理进行丢弃（被限流）：

	available := tokenBucket.TakeAvailable(1)
	if available <= 0 {
	    /*** 限流处理 ***/
	}

这里就可以简单理解了：在程序真正处理之前，从桶中拿出1个令牌，如果成功了，就继续，如果没成功（返回0），那么就触发限流，回复：“系统繁忙，请稍候再试。”

### 扩展：刚刚我们是对整个系统进行无差别限流。如果我们要针对每一个公众号做单独限流，那么其实也很简单：
初始化全局变量变成一个

	var tokenBuckets map[string]*ratelimit.Bucket

在真正处理数据之前判断用来区分公众号的key是否存在在这个map中。如果没有那么初始化一个令牌桶并放入map；如果已存在则直接使用即可。 需要注意的是因为对全局map做操作，故而要在读写之前做安全锁：

	var ok bool
	var tokenBucket *ratelimit.Bucket = nil

	lock.Lock()
	tokenBucket, ok = tokenBuckets[tokenBucketName]
	lock.Unlock()
	
	if !ok {
	    tokenBucket = ratelimit.NewBucket(bucketFillDuring, bucketMax)
	
	    lock.Lock()
	    tokenBuckets[tokenBucketName] = tokenBucket
	    lock.Unlock()
	}
