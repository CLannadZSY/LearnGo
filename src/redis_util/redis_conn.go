package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

func main() {
	//connMethOne()
	connMethTwo()
}

// 使用参数配置方式连接
func connMethOne() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       1,
	})
	ctx := context.Background()
	val, err := rdb.Get(ctx, "test_key").Result()
	if err != nil {
		if err == redis.Nil {
			fmt.Println("key does not exists")
			return
		}
		panic(err)
	}
	fmt.Println(val)
}

// 使用字符串url连接, 推荐使用这种
func connMethTwo() {
	redisUrlDBOne := "redis://username:password@127.0.0.1:6379/1"
	opt, err := redis.ParseURL(redisUrlDBOne)
	if err != nil {
		panic(err)
	}
	rdb := redis.NewClient(opt)
	ctx := context.Background()
	val, err := rdb.Get(ctx, "test_key").Result()
	if err != nil {
		if err == redis.Nil {
			fmt.Println("key does not exists")
			return
		}
		panic(err)
	}
	fmt.Println(val)

	//	要执行任意/自定义命令
	get := rdb.Do(ctx, "get", "key")
	if err := get.Err(); err != nil {
		if err == redis.Nil {
			fmt.Println("key does not exists")
			return
		}
		panic(err)
	}
	fmt.Println(get.Val().(string))
}
