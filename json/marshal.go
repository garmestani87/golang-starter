package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	Id       int    `json:"-"`
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
	Age      int    `json:"age,omitempty"`
}

func main() {
	p1 := Person{Id: 1, Name: "jafar", Lastname: "hervai"}
	p2 := Person{Id: 2, Name: "naghi", Lastname: "taghavi"}

	persons := []Person{p1, p2}

	res, err := json.Marshal(persons)
	if err != nil {
		log.Fatal("unexpected erro occured !!!")
	}

	fmt.Println(string(res))

}
