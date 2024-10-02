package main

import (
	"fmt"
	"time"
)

func main() {

	channel1 := make(chan int)
	channel2 := make(chan int)

	var res int

	go sendToChannel1(channel1)
	go sendToChannel2(channel2)

	time.Sleep(2*time.Second)
	
	select {
	case res = <-channel1:
		fmt.Printf("revieve %d \n", res)
	case res = <-channel2:
		fmt.Printf("recieve %d \n", res)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout .... \n")
	default:
		fmt.Println("default .... \n")
	}

}

func sendToChannel1(channel chan<- int) {
	time.Sleep(2 * time.Second)
	fmt.Printf("before send to channel 1 -> %s \n", time.Now().String())
	channel <- 1
	fmt.Printf("after send to channel 1 -> %s \n", time.Now().String())
}

func sendToChannel2(channel chan<- int) {
	time.Sleep(1 * time.Second)
	fmt.Printf("before send to channel 2 -> %s \n", time.Now().String())
	channel <- 2
	fmt.Printf("after send to channel 2 -> %s \n", time.Now().String())
}
