package common

import (
	"github.com/go-redis/redis"
	"echo_test/config"
	"time"
)

var (
	Redis = NewRedisClient()
)

//返回redis，单例
func NewRedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     config.Config.RedisAddr,
		Password: config.Config.RedisPassword, // no password set
		DB: config.Config.RedisDb,   // 因为系统是64位的，所以默认的 int 型是 int64
	})
}

//获取
func GetRedis(key string) *redis.StringCmd {
	return Redis.Get(key)
}

//设置
func SetRedis(key string, value interface{}, expirse time.Duration) {
	Redis.Set(key, value, expirse)
}

//删除
func DelRedis(key string) {
	Redis.Del(key)
}

//清空
func ClearRedis() {
	Redis.FlushAll()
}
