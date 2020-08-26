package myredis

import "github.com/go-redis/redis"

// 声明一个全局的rdb变量
var rdb *redis.Client

// InitClient 初始化连接
func InitClient() (rdb *redis.Client, err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err = rdb.Ping().Result()
	if err != nil {
		return rdb, err
	}
	return rdb, nil
}
