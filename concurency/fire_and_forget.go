package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	num_cpu := runtime.NumCPU()
	fmt.Println(num_cpu)

	runtime.GOMAXPROCS(num_cpu)

	value := 10
	go func() {
		value += 1
	}()

	go func() {
		value -= 3
	}()

	go func() {
		value *= 5
	}()

	time.Sleep(1 * time.Millisecond)
	fmt.Println(value)
}
