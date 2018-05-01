/*
连接redis的处理类
*/
package zredis

import (
	"github.com/garyburd/redigo/redis"
	"log"
	"time"
)

var (
	// RedisPool RedisPool连接池实例
	RedisPool *redis.Pool
	config    *RedisConfig
)

/*
;redis.host = 127.0.0.1:6379   ; redis连接地址
;redis.db = 1                  ; 建议选择未被使用的数据库或新建redis实例，防止key冲突
;redis.password =              ; redis密码, 无需密码留空
;redis.max_idle = 10           ; redis连接池最大空闲连接数
;redis.max_active = 0          ; redis连接池最大激活连接数, 0为不限制
;redis.connect_timeout = 5000  ; redis连接超时时间, 单位毫秒
;redis.read_timeout = 180000   ; redis读取超时时间, 单位毫秒
;redis.write_timeout = 3000    ; redis写入超时时间, 单位毫秒
*/

type RedisConfig struct {
	Host           string
	Db             int
	Password       string
	MaxIdle        int // 连接池最大空闲连接i
	MaxActive      int // 连接池最大激活连接数
	ConnectTimeout int // 连接超时, 单位毫秒
	ReadTimeout    int // 读取超时, 单位毫秒
	WriteTimeout   int // 写入超时, 单位毫秒
}

// 初始化连接池
func InitRedisPool(redisconfig *RedisConfig) *redis.Pool {
	config = redisconfig //赋予全局值
	pool := &redis.Pool{
		MaxIdle:      config.MaxIdle,   // 连接池最大空闲连接
		MaxActive:    config.MaxActive, // 连接池最大激活连接数
		IdleTimeout:  300 * time.Second,
		Dial:         redisDial,
		TestOnBorrow: redisTestOnBorrow,
		Wait:         true,
	}

	return pool
}

// 连接redis
func redisDial() (redis.Conn, error) {
	conn, err := redis.Dial(
		"tcp",
		config.Host,
		redis.DialConnectTimeout(time.Duration(config.ConnectTimeout)*time.Millisecond),
		redis.DialReadTimeout(time.Duration(config.ReadTimeout)*time.Millisecond),
		redis.DialWriteTimeout(time.Duration(config.WriteTimeout)*time.Millisecond),
	)
	if err != nil {
		return nil, err
	}
	//密码认证
	if config.Password != "" {
		if _, err := conn.Do("AUTH", config.Password); err != nil {
			conn.Close()
			return nil, err
		}
	}
	// 选择db
	if _, err := conn.Do("SELECT", config.Db); err != nil {
		conn.Close()
		return nil, err
	}
	log.Println("redis连接成功")
	return conn, nil

}

// 从池中取出连接后，判断连接是否有效
func redisTestOnBorrow(conn redis.Conn, t time.Time) error {
	_, err := conn.Do("PING")
	if err != nil {
		log.Printf("从redis连接池取出的连接无效#%s", err.Error())
	}

	return err
}

// 执行redis命令, 执行完成后连接自动放回连接池
func ExecRedisCommand(command string, args ...interface{}) (interface{}, error) {
	redis := RedisPool.Get()
	defer redis.Close()

	return redis.Do(command, args...)
}
