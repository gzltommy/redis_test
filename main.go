package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Second * 1)

	for {
		select {
		case t := <-ticker.C:
			fmt.Println("=====+++=========", t.String())
		}
	}
}
