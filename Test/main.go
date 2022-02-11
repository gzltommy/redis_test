package main

import (
	"fmt"
	"time"

	"github.com/gomodule/redigo/redis"
)

func main() {
	redisHost := "192.168.24.147"
	redisPort := "6379"
	password := ""

	url := fmt.Sprintf("%s:%s", redisHost, redisPort)
	// 连接池
	redisPool := &redis.Pool{
		TestOnBorrow:    nil,
		MaxIdle:         20,
		MaxActive:       100,
		IdleTimeout:     20 * time.Second,
		Wait:            false,
		MaxConnLifetime: 0,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", url)
			if err != nil {
				return nil, err
			}
			if password != "" {
				_, err = c.Do("AUTH", password)
				if err != nil {
					return nil, err
				}
			}
			return c, nil
		},
	}

	// 关闭连接池
	defer redisPool.Close()

	//获取连接
	conn := redisPool.Get()
	//关闭连接
	defer conn.Close()

	//写入数据
	_, err := conn.Do("SET", "go_key", "redigo")
	if err != nil {
		fmt.Println("err while setting:", err)
	}

	//判断key是否存在
	is_key_exit, err := redis.Bool(conn.Do("EXISTS", "go_key"))
	if err != nil {
		fmt.Println("err while checking keys:", err)
	} else {
		fmt.Println(is_key_exit)
	}

	//获取value并转成字符串
	account_balance, err := redis.String(conn.Do("GET", "go_key"))
	if err != nil {
		fmt.Println("err while getting:", err)
	} else {
		fmt.Println(account_balance)
	}

	//删除key
	_, err = conn.Do("DEL", "go_key")
	if err != nil {
		fmt.Println("err while deleting:", err)
	}

	//设置key过期时间
	_, err = conn.Do("SET", "mykey", "superWang", "EX", "5")
	if err != nil {
		fmt.Println("err while setting:", err)
	}

	//对已有key设置5s过期时间
	n, err := conn.Do("EXPIRE", "go_key", 5)
	if err != nil {
		fmt.Println("err while expiring:", err)
	} else if n != int64(1) {
		fmt.Println("failed")
	}

	/*================================================================================================================*/

	//使用 redis.Conn.Do() 方法执行命令
	conn.Do("set", "k1", "redis")

	//redis.String()
	v1, err := redis.String(conn.Do("get", "k1"))
	if err != nil {
		fmt.Printf("err is %v\n", err)
	} else {
		fmt.Printf("v1 is %s\n", v1)
	}

	//添加 list
	conn.Do("lpush", "l1", 1)
	conn.Do("lpush", "l1", 2)
	conn.Do("lpush", "l1", 3)
	conn.Do("lpush", "l1", 4)

	//获取 list 长度
	len, _ := redis.Int64(conn.Do("llen", "l1"))
	fmt.Printf("l1 len is %d\n", len)

	//取第一个值
	l1, _ := redis.Int64(conn.Do("lpop", "l1"))
	fmt.Printf("l1 pop %d\n", l1)

	//遍历
	values, _ := redis.Int64s(conn.Do("lrange", "l1", 0, -1))
	for i, v := range values {
		fmt.Printf("l1[%d] is %d\n", i, v)
	}

	//hash
	conn.Do("hset", "h1", "id", 1)
	conn.Do("hset", "h1", "name", "tom")
	conn.Do("hset", "h1", "age", 18)
	conn.Do("hmset", "h2", "id", 2, "name", "jinx", "age", 17)

	hv, err := redis.Values(conn.Do("hmget", "h1", "id", "name", "age"))
	if err != nil {
		fmt.Printf("err is %v\n", err)
	}
	for _, v := range hv {
		fmt.Printf("%s\n", v.([]byte))
	}

	//
	h2v, err := redis.Values(conn.Do("hgetall", "h2"))
	if err != nil {
		fmt.Printf("err is %v\n", err)
	}
	for _, v := range h2v {
		fmt.Printf("%s\t", v.([]byte))
	}
	fmt.Println()

	//管道批量操作
	conn.Send("HSET", "user", "name", "sun", "age", "30")
	conn.Send("HSET", "user", "sex", "female")
	conn.Send("HGET", "user", "age")
	conn.Flush()
	res1, err := conn.Receive()
	fmt.Printf("Receive res:%v\n", res1)

	res2, err := conn.Receive()
	fmt.Printf("Receive res:%v\n", res2)

	res3, err := conn.Receive()
	fmt.Printf("Receive res:%v\n", res3)
}
