package model

import (
	"fmt"

	"github.com/go-redis/redis"
)

// 定义一个UserDao结构体
// 完成对User结构体的各种操作

type UserDao struct {
	pool *redis.Client
}