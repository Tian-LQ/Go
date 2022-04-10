package main

import (
	"fmt"
	"github.com/go-kratos/kratos/pkg/cache/redis"
	"testing"
	"time"
)

func TestRedisConnection(t *testing.T) {
	c, _ := redis.Dial("tcp", "127.0.0.1:3389",
		redis.DialDatabase(0),
		redis.DialPassword("hello"),
		redis.DialReadTimeout(time.Second*10))
	fmt.Println(c)
}
