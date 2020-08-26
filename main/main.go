package main

import (
	// 系统包
	"database/sql"
	"fmt"

	// 自定义包
	"github.com/qiulc/gotest/conn/myredis"
	"github.com/qiulc/gotest/conn/mysql"

	// 外包
	"github.com/go-redis/redis"
)

var db *sql.DB
var rdb *redis.Client

func main() {
	fmt.Println("你好啊 ")
	var err error
	db, err = mysql.InitDB()
	if err != nil {
		fmt.Println("连接Mysql数据库失败===>", err)
		return
	}
	fmt.Println("Mysql已连接...")
	rdb, err = myredis.InitClient()

	if err != nil {
		fmt.Println("连接Redis数据库失败===>", err)
		return
	}
	fmt.Println("Redis已连接...")
}
