package main

import "fmt"

func main() {
	vv := []int{1, 2, 3}

	for i := 0; i < len(vv); {
		if vv[i] == 2 {

			vv = append(vv[:i], vv[i+1:]...)
		} else {
			i++
		}

	}

	//for i, _ := range vv {
	//	vv = append(vv[:i], vv[i+1:]...)
	//
	//}
	fmt.Println("----", vv)
}
