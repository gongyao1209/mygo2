package db

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"os"
)

var c redis.Conn

func init()  {
	var err error
	c, err = redis.Dial("tcp", "localhost:6379")
	errCheck(err)
	fmt.Println("redis conn")
}

func Get(key interface{}) (r string) {
	var err error
	r, err = redis.String(c.Do("GET", key))
	errCheck(err)

	return
}

func SET(key, value interface{}) int {
	r, err := redis.String(c.Do("SET", key, value))
	errCheck(err)

	if r == "OK" {
		return 1
	}

	return 0
}

func SETNX(key, value interface{}) int64 {
	r, err := redis.Int64(c.Do("SETNX", key, value))
	errCheck(err)

	return r
}

func EXPIRE(key interface{}, secend int) int64 {
	r, err := redis.Int64(c.Do("EXPIRE", key, secend))

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



func Test()  {

	//var c redis.Conn
	//c = RedisConn()


	//对本次连接进行set操作
	res,setErr := c.Do("SETNX","url", "xxbandy.github.io1")
	errCheck(setErr)
	fmt.Println(res)

	_,setErr = c.Do("EXPIRE", "url", 10)
	errCheck(setErr)

	//使用redis的string类型获取set的k/v信息
	//r,getErr := redis.StringMap((*c).Do("HGETALL","url"))
	//errCheck(getErr)
	//fmt.Println(r["1"])
}

func errCheck(err error) {
	if err != nil {
		fmt.Println("sorry,has some error:",err)
		os.Exit(-1)
	}
}