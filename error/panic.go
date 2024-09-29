package main

import (
	"fmt"
	"runtime/debug"
)

type ArrayList[T comparable] struct {
	item []T
}

func main() {
	list := ArrayList[int]{}
	list.item = append(list.item, 1)
	list.item = append(list.item, 2)
	list.item = append(list.item, 3)
	list.item = append(list.item, 4)
	list.item = append(list.item, 5)

	defer func(){
		if err := recover(); err != nil{
			fmt.Printf("%s \n", err)
			debug.PrintStack()
		}
	}()
	fmt.Println(list.getValue(10))
	fmt.Printf("after get value by index %d \n")
}

func (arr *ArrayList[T]) getValue(idx int) T {
	// log.Fatal("defer not work after log.fatal !")
	defer func() {
		fmt.Printf("index is out of range %d \n", idx)
	}()
	fmt.Printf("before get value by index %d \n", idx)
	return arr.item[idx]
}
