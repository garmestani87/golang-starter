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
	mx := sync.Mutex{}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(age int) {
			person = getOncePerson(age, &wg, &mx)
		}(i)
	}

	wg.Wait()
}

func getOncePerson(age int, wg *sync.WaitGroup, mx *sync.Mutex) *Person {
	defer wg.Done()
	mx.Lock()
	defer mx.Unlock()
	if person == nil {
		person = &Person{name: "jafar", age: age}
	}
	fmt.Printf("%+v \n", *person)
	return person
}
