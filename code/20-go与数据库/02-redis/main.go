package main

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

// redis连接池  初始化一些连接  可以节省连接redis时间提高效率
var pool *redis.Pool

// 初始化连接池
func init() {
	pool = &redis.Pool{
		MaxIdle:     8,   //最大空间链接数
		MaxActive:   0,   //最大和数据库的最大链接数,0代表没有限制
		IdleTimeout: 300, //最大空闲时间
		Dial: func() (redis.Conn, error) { //初始化连接
			return redis.Dial("tcp", "127.0.0.1:6379")
		},
	}

}

func main() {
	// 连接池使用
	conn := pool.Get()
	defer conn.Close()

	// go get -u github.com/go-redis/redis
	// go get -u github.com/gomodule/redigo/redis //更接近cli操作
	// conn, err := redis.Dial("tcp", "localhost:6379")
	// if err != nil {
	// 	fmt.Println("conn redis failed", err)
	// 	return
	// }
	// defer conn.Close()

	// 1.set string
	_, err := conn.Do("Set", "name", "张三")
	if err != nil {
		fmt.Println("set err :", err)
	}
	_, _ = conn.Do("Mset", "age", "18", "gender", "男")

	// 2.get string
	res, _ := redis.String(conn.Do("Get", "name"))
	fmt.Println(res)
	ress, _ := redis.Strings(conn.Do("Mget", "name", "age", "gender"))
	fmt.Println(ress)

	// 3. set hash
	_, _ = conn.Do("Hset", "user", "name", "soul")
	_, _ = conn.Do("Hmset", "user", "age", "20", "gender", "男")

	// 4.get hash
	uname, _ := redis.String(conn.Do("Hget", "user", "name"))
	uage, _ := redis.Int(conn.Do("Hget", "user", "age"))
	keys, _ := redis.Strings(conn.Do("Hkeys", "user")) //stings   string的切片
	vals, _ := redis.Strings(conn.Do("Hvals", "user"))
	all, _ := redis.StringMap(conn.Do("Hgetall", "user"))
	fmt.Println(uname, uage)
	fmt.Printf("keys:%v values:%v all:%v", keys, vals, all)
}
