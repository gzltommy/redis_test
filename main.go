package main

import (
	"fmt"
	"redis_test/redisclient"
)

func main() {
	fmt.Println("----", redisclient.HmSet("user_online_state", 100, 1, 200, 1, 300, 1))

	hv, err := redisclient.HmGet("user_online_state", 100, 200, 300, 400, 500)
	if err != nil {
		fmt.Printf("err is %v\n", err)
	}

	fmt.Println("=============", hv)
	for _, v := range hv {
		fmt.Printf("%v\n", v)
	}
}
