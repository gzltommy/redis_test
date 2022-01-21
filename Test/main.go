package main

import (
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func testJsonConvert() {
	c, err := redis.Dial("tcp", "192.168.24.147:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	defer c.Close()

	key := "profile"
	imap := map[string]string{"name": "duncanwang", "sex": "male", "mobile": "13671927788"}
	value, _ := json.Marshal(imap)

	n, err := c.Do("SETNX", key, value)
	if err != nil {
		fmt.Println(err)
	}
	if n == int64(1) {
		fmt.Println("set Json key success。")
	}

	var imapGet map[string]string

	valueGet, err := redis.Bytes(c.Do("GET", key))
	if err != nil {
		fmt.Println(err)
	}

	errShal := json.Unmarshal(valueGet, &imapGet)
	if errShal != nil {
		fmt.Println(err)
	}
	fmt.Println(imapGet["name"])
	fmt.Println(imapGet["sex"])
	fmt.Println(imapGet["mobile"])
}
