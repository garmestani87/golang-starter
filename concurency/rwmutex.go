package main

import (
	"fmt"
	"sync"
)

func main() {
	var res int
	wg := sync.WaitGroup{}
	rwmx := sync.RWMutex{}
	
	sum(&res, &wg, &rwmx)

	wg.Wait()

	fmt.Println(res)
}

func sum(res *int, wg *sync.WaitGroup, mx *sync.RWMutex) {
	for i := 0; i < 1_000_000; i++ {
		wg.Add(1)
		go func(item int) {
			defer wg.Done()
			// I am reading this critical section no one can not write here when i am reading !
			mx.RLock()
			*res += item
			mx.RUnlock()
		}(i)
	}
}
