package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var res int64
	wg := sync.WaitGroup{}

	sum(&res, &wg)

	wg.Wait()

	fmt.Println(res)
}

func sum(res *int64, wg *sync.WaitGroup) {
	var i int64
	for i = 0; i < 1_000_000; i++ {
		wg.Add(1)
		go func(item int64) {
			defer wg.Done()
			atomic.AddInt64(res, int64(item))
		}(i)
	}
}
