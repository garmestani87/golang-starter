package main

import (
	"fmt"
	"sync"
)

func main() {
	var res int
	wg := sync.WaitGroup{}
	mx := sync.Mutex{}
	
	sum(&res, &wg, &mx)

	wg.Wait()

	fmt.Println(res)
}

func sum(res *int, wg *sync.WaitGroup, mx *sync.Mutex) {
	for i := 0; i < 1_000_000; i++ {
		wg.Add(1)
		go func(item int) {
			defer wg.Done()
			mx.Lock()
			*res += item
			mx.Unlock()
		}(i)
	}
}
