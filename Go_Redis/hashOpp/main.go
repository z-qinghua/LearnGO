package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

var rdb *redis.Client

func main() {
	// 初始化连接
	rdb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
	_, err := rdb.Ping().Result()
	if err != nil {
		fmt.Println("conn failed...")
		return
	}
	fmt.Println("conn succ...")

	// 根据key和field字段设置，field字段的值
	// user_1 是hash key，username 是字段名, tizi365是字段值
	err = rdb.HSet("usr1", "username", "zeng").Err()
	if err != nil {
		fmt.Println("Hset failed...")
		panic(err)
	}

	// 根据key和field字段，查询field字段的值
	// user_1 是hash key，username是字段名
	username, err := rdb.HGet("usr1", "username").Result()
	if err != nil {
		fmt.Println("HGet failed...")
		return
	}
	fmt.Println("username =", username)

	// 根据key和field字段设置，field字段的值
	// user_1 是hash key，username 是字段名, tizi365是字段值
	err = rdb.HSet("usr1", "email", "hua@henu.com").Err()
	if err != nil {
		fmt.Println("Hset failed...")
		panic(err)
	}

	// 根据key查询所有字段和值
	// 一次性返回key=usr1的所有hash字段和值
	data, err := rdb.HGetAll("usr1").Result()
	if err != nil {
		fmt.Println("HGetAll failed...")
		panic(err)
	}
	//  data是一个map类型，这里使用使用循环迭代输出

	fmt.Println("-------HGetAll--------")
	for field, val := range data {
		fmt.Println(field, val)
	}

	// 根据key返回所有字段名
	// keys是一个string数组
	fmt.Println("-------HKeys--------")

	keys, err := rdb.HKeys("usr1").Result()
	if err != nil {
		fmt.Println("HKeys failed...")
		panic(err)
	}
	fmt.Println(keys)

	// 根据key，查询hash的字段数量
	fmt.Println("-------HLen--------")

	size, err := rdb.HLen("usr1").Result()
	if err != nil {
		fmt.Println("HLen failed...")
		panic(err)
	}
	fmt.Println(size)

	// 根据key和多个字段名，批量查询多个hash字段值
	// HMGet支持多个field字段名，意思是一次返回多个字段值
	vals, err := rdb.HMGet("usr1", "username", "email").Result()
	if err != nil {
		fmt.Println("HMGet failed..")
		panic(err)
	}
	// vals一个数组
	fmt.Println(vals)

	// 根据key和多个字段名和字段值，批量设置hash字段值
	// 初始化hash数据的多个字段值

	datas := make(map[string]interface{})
	datas["id"] = "1"
	datas["username"] = "aaa"
	// 一次性保存多个hash字段值
	err = rdb.HMSet("key", datas).Err()
	if err != nil {
		fmt.Println("HMSet failed...")
		panic(err)
	}

	// 根据key和多个字段名，批量查询多个hash字段值
	// HMGet支持多个field字段名，意思是一次返回多个字段值

	vals, err = rdb.HMGet("key", "id", "username").Result()
	if err != nil {
		fmt.Println("HMGet failed..")
		panic(err)
	}
	// vals一个数组
	fmt.Println("-----HMSet-----")
	fmt.Println(vals)

	// 如果field字段不存在，则设置hash字段值,存在则不做改变
	err = rdb.HSetNX("key", "ids", 100).Err()
	if err != nil {
		panic(err)
	}

	// 根据key和字段名，删除hash字段，支持批量删除hash字段
	// 删除一个字段id
	rdb.HDel("key", "id")

	// 删除多个字段
	rdb.HDel("key", "ids", "username")


	// 监测hash字段是否存在
	err = rdb.HExists("key", "id").Err()
	if err != nil {
		panic(err)
	}

}
