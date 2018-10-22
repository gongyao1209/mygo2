package db

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"os"
)

func RedisConn() redis.Conn {
	c, err := redis.Dial("tcp", "localhost:6379")
	errCheck(err)

	return c
}



func Test()  {

	var c redis.Conn
	c = RedisConn()


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