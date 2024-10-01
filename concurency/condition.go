package main

import (
	"fmt"
	"sync"
)

type Event struct {
	id     int
	notify bool
}

func main() {

	l := sync.Mutex{}
	cond := sync.NewCond(&l)
	wg := sync.WaitGroup{}
	event := &Event{}

	for i := 0; i < 3; i++ {
		fmt.Printf("establish new request id : %d \n", i)
		wg.Add(2)
		go event.pub(i, cond, &wg)
		go event.sub(i, cond, &wg)
	}
	wg.Wait()
}

func (e *Event) pub(requestId int, cond *sync.Cond, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("send request to quiue id : %d \n", requestId)
	cond.L.Lock()
	defer cond.L.Unlock()
	if !e.notify {
		fmt.Printf("wait for resonse : %d \n", requestId)
		cond.Wait()
	}
	fmt.Printf("send resposne to request id : %d \n", requestId)
}

func (e *Event) sub(requestId int, cond *sync.Cond, wg *sync.WaitGroup) {
	defer wg.Done()
	// time.Sleep(1 * time.Second)
	fmt.Printf("consume request form quiue id : %d \n", requestId)
	e.notify = true
	cond.Broadcast()
}
