package main

import "fmt"

type Person struct {
	name     string
	lastname string
	age      int
}

func main() {

	var persons = map[string]Person{}
	persons["1"] = Person{name: "ali", lastname: "rezai", age: 10}
	persons["2"] = Person{name: "jafar", lastname: "gholi", age: 13}
	persons["3"] = Person{name: "mamad", lastname: "jj", age: 25}

	fmt.Printf("values : %v \n", persons)

	persons["2"] = Person{name: "ali", lastname: "kaveh", age: 45}

	fmt.Printf("values : %v \n", persons)

	person, exist := persons["2"]
	if exist {
		fmt.Printf("person 2 is : %v \n", person)
	} else {
		fmt.Print("dose not exist !")
	}

	for i, item := range persons {
		fmt.Printf("index : %s items : %v \n", i, item)
	}

}
