package main

import (
	"fmt"
	"redis_test/redisclient"
)

func main() {
	redisclient.Set(fmt.Sprintf("user_online_state_%d", 100), 0, []byte{1})
	redisclient.Set(fmt.Sprintf("user_online_state_%d", 200), 0, []byte{1})
	redisclient.Set(fmt.Sprintf("user_online_state_%d", 300), 0, []byte{1})
	redisclient.Set(fmt.Sprintf("user_online_state_%d", 400), 0, []byte{1})
	redisclient.Set(fmt.Sprintf("user_online_state_%d", 500), 0, []byte{1})

	keys := make([]interface{}, 0, 6)
	for i := 3; i <= 8; i++ {
		keys = append(keys, fmt.Sprintf("user_online_state_%d00", i))
	}

	hv, err := redisclient.MGet(keys...)
	if err != nil {
		fmt.Printf("err is %v\n", err)
	}

	for _, v := range hv {
		if v == nil {
			fmt.Println(v, 0)
		} else {
			fmt.Println(v, 1)
		}
	}

}
