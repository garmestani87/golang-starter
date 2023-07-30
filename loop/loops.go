package main

import (
	"fmt"
)

func main() {

	var i int = 0
	for {
		i++
		println("hello loop")
		if i >= 10 {
			break
		}
	}

	for j := 0; j <= 10; j++ {
		println("hello loop")
	}

	var lst = []int{1, 5, 8, 9}
	for index, item := range lst {
		fmt.Printf("index = %d , item = %d \n", index, item)
	}

	for _, item := range lst {
		fmt.Printf("item = %d \n", item)
	}

}
