package main

import (
	"fmt"
	"sync"
)

type Person struct {
	name string
	age  int
}

var person *Person

func main() {

	wg := sync.WaitGroup{}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(item int) {
			person = getPerson(i, &wg)
		}(i)
	}
	wg.Wait()
}

func getPerson(age int, wg *sync.WaitGroup) *Person {
	defer wg.Done()
	once := sync.Once{}
	once.Do(func() {
		if person == nil {
			person = &Person{name: "jafar", age: age}
		}
		fmt.Printf("%+v \n", *person)
	})
	return person
}
