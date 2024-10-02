package main

import (
	"fmt"
	"time"
)

func main() {

	channel := make(chan int)
	defer close(channel)

	go send(channel)

	recieve(channel)
	recieve(channel)
	recieve(channel)

}

func recieve(channel <-chan int) {
	rec := <-channel
	fmt.Printf("recieved %d time = %s \n", rec, time.Now().String())
}

func send(channel chan<- int) {

	fmt.Printf("before %d time = %s \n", 1, time.Now().String())
	channel <- 1
	fmt.Printf("after %d time = %s \n", 1, time.Now().String())

	fmt.Printf("before %d time = %s \n", 2, time.Now().String())
	channel <- 2
	fmt.Printf("after %d time = %s \n", 2, time.Now().String())

	fmt.Printf("before %d time = %s \n", 3, time.Now().String())
	channel <- 3
	fmt.Printf("after %d time = %s \n", 3, time.Now().String())
}
