package main

import (
	"fmt"
	"redis_test/redisclient"
)

func main() {
	redisclient.ZAdd("fans_list", 1, 100, 2, 200, 3, 300)

	hv, err := redisclient.ZRange("fans_list", 0, 5)
	if err != nil {
		fmt.Printf("err is %v\n", err)
	}

	for _, v := range hv {
		fmt.Println("+++", v)
		if v == nil {
			fmt.Println("=== nil ====", v)
		} else {
			vv := v.([]byte)
			fmt.Println("********", string(vv))
		}
	}

	redisclient.ZRem("fans_list", 200)
}
