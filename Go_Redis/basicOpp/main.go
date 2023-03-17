package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

//redis

var rdb *redis.Client
//初始化连接
func initRedis() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
		Password: "",
		DB: 0,
	})
	_, err = rdb.Ping().Result()
	
	return
}
// func Set(key, value string)  {
	
// }

func main()  {
	err := initRedis()
	if err != nil {
		fmt.Println("connect redis faild")
		return
	}
	fmt.Println("连接redis成功")

	// 设置一个key值
	// 第三个参数代表key的过期时间，0代表不会过期。
	err = rdb.Set("key", "value", 0).Err()
	if err!= nil {
		fmt.Println("set failed..")
		panic(err)
	}

	// 查询key的值
	// Result函数返回两个值，第一个是key的值，第二个是错误信息
	val, err := rdb.Get("key").Result()
	// 查询是否出错
	if err!=nil {
		fmt.Println("Ger failed...")
		panic(err)
	}
	fmt.Println("key =", val)

	// 如果key不存在，则设置这个key的值
	// 第三个参数代表key的过期时间，0代表不会过期。
	err = rdb.SetNX("name","zeng", 0).Err()
	if err != nil {
		fmt.Println("SetNX failed...")
	}

	// 查询key的值
	// Result函数返回两个值，第一个是key的值，第二个是错误信息
	val, err = rdb.Get("name").Result()
	// 查询是否出错
	if err!=nil {
		fmt.Println("Get failed...")
		panic(err)
	}
	fmt.Println("name =", val)	

	// 批量查询key的值
	// MGet函数可以传入任意个key，一次性返回多个值。
	// 这里Result返回两个值，第一个值是一个数组，第二个值是错误信息
	vals, err := rdb.MGet("key", "name").Result()
	if err != nil {
		fmt.Println("MGet failed..")
		panic(err)
	}
	fmt.Println(vals)

	// 批量设置key的值
	err = rdb.MSet("key1","aaa", "key2", "bbb", "key3", "ccc").Err()
	if err != nil {
		fmt.Println("MSet failed...")
		panic(err)
	}

	// 删除key操作，支持批量删除
	// 删除
	rdb.Del("key")
	// 删除多个key, Del函数支持删除多个key
	err = rdb.Del("key1", "key2").Err()
	if err != nil {
		fmt.Println("Del failed..")
		panic(err)
	}
	// 设置key的过期时间,单位秒
	rdb.Expire("name", 3)

	// 设置一个key的值，并返回这个key的旧值
	// Result函数返回两个值，第一个是key的值，第二个是错误信息
	oldVal, err := rdb.GetSet("key3", "ddd").Result()
	if err != nil {
		fmt.Println("GetSet failed...")
		panic(err)
	}
	// 打印key的旧值
	fmt.Println("key =", oldVal)
}