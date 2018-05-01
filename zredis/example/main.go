package main

import (
	"fmt"
	"github.com/zouhuigang/package/zredis"
)

func init() {
	redisConfig := new(zredis.RedisConfig)
	redisConfig.Host = "127.0.0.1:6379"
	redisConfig.Db = 1
	redisConfig.Password = ""
	redisConfig.MaxIdle = 10
	redisConfig.MaxActive = 0
	redisConfig.ConnectTimeout = 5000
	redisConfig.ReadTimeout = 180000
	redisConfig.WriteTimeout = 3000
	zredis.RedisPool = zredis.InitRedisPool(redisConfig)
}

func main() {
	zredis.ExecRedisCommand("SET", "test", "success")
	value, _ := zredis.ExecRedisCommand("GET", "test")
	byteValue := value.([]byte)
	fmt.Printf("test:%v", string(byteValue))
	/*
		输出:
			2018/05/01 10:31:07 redis连接成功
			test:success
	*/
}
