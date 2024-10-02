package main

import (
	"fmt"
	"sync"
	"time"
)

const NUM int = 10

func main() {

	wg := sync.WaitGroup{}
	channel := make(chan int, NUM)
	now := time.Now()

	func() {
		for i := 1; i <= NUM; i++ {
			wg.Add(1)
			go sendToChannel(channel, i, &wg)
		}
	}()

	wg.Wait()
	close(channel)

	recieveFromChannel(channel)
	
	fmt.Printf("stop watch : %s \n", time.Since(now).String())
}

func sendToChannel(channel chan<- int, idx int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("before sending to channel index = %d \n", idx)
	channel <- idx
	fmt.Printf("after sent to channel index = %d \n", idx)
}

func recieveFromChannel(channel <-chan int) {
	for idx := range channel {
		fmt.Printf("recieved from channel with index = %d \n", idx)
	}
}
