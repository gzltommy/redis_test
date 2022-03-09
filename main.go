package main

import (
	"fmt"
	redis "redis_test/redisclient"
	"time"
)

func main() {
	id1 := "adfasfasfasf"
	err := redis.Lock(id1, "AAA", 30)
	fmt.Println("=======", err)

	go func() {

		id2 := "adfasfasf666asf"
		err = redis.Unlock(id2, "AAA")
		fmt.Println("-----------", err)

		err = redis.Lock(id2, "AAA", 30)
		fmt.Println("----222-------", err)
	}()

	time.Sleep(time.Second * 1888888)
	return
}
