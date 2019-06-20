package db

import (
	"flag"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"os"
	"sync"
	"time"
)

var (
	pool        *redis.Pool
	redisServer = flag.String("redisServer", RedisUrl, "")
)

const (
	RedisUrl            = "localhost:6379"
	redisMaxIdle        = 3
	redisIdleTimeoutSec = 1
	redisPassword       = ""
)

//获取redis连接池
func newPool() *redis.Pool {
	return &redis.Pool{
		// 最大空闲连接数，如果没有连接连接数据库 也得保存这么大的连接，不至于为空。
		MaxIdle: redisMaxIdle,
		// 连接池支持的最大连接数，这是运行时候的连接数。设置为0的时候 不限
		MaxActive: 0,
		//在此期间保持空闲关闭连接
		IdleTimeout: redisIdleTimeoutSec * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", RedisUrl)
			if err != nil {
				return nil, err
			}
			//if _, err := c.Do("AUTH", redisPassword); err != nil {
			//  c.Close()
			//  return nil, err
			//}
			if _, err := c.Do("SELECT", 1); err != nil {
				c.Close()
				return nil, err
			}
			return c, nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}
}

func init() {
	pool = newPool()
}

func Get(key interface{}) (r string) {
	conn := pool.Get()
	defer conn.Close()

	var err error
	r, err = redis.String(conn.Do("GET", key))
	errCheck(err)

	return
}

func SET(key, value interface{}) int {
	conn := pool.Get()
	defer conn.Close()

	r, err := redis.String(conn.Do("SET", key, value))
	errCheck(err)

	if r == "OK" {
		return 1
	}

	return 0
}

func SETNX(key, value interface{}) int64 {
	conn := pool.Get()
	defer conn.Close()

	r, err := redis.Int64(conn.Do("SETNX", key, value))
	errCheck(err)

	return r
}

func EXPIRE(key interface{}, secend int) int64 {
	conn := pool.Get()
	defer conn.Close()

	r, err := redis.Int64(conn.Do("EXPIRE", key, secend))

	errCheck(err)

	return r
}

//func RedisConn() redis.Conn {
//	c, err := redis.Dial("tcp", "localhost:6379")
//	errCheck(err)
//
//	fmt.Println("redis conn")
//
//	return c
//}

func Test() {
	var wg sync.WaitGroup

	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(Get("name2"))
		}()
	}

	for i := 0; i < 10; i++ {
		if i > 5 {
			time.Sleep(1 * time.Second)
		}
		fmt.Println("i = ", i, ", ActiveCount: ", pool.ActiveCount())
	}

	wg.Wait()
}

func errCheck(err error) {
	if err != nil {
		fmt.Println("sorry,has some error:", err)
		os.Exit(-1)
	}
}
