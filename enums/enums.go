package main

import (
	"fmt"
)

const url string = "http://jj.com"

type status int

const (
	reject  status = 0
	approve status = 1
)

type Order struct {
	id          int
	price       float64
	count       int
	orderStatus status
}

func main() {

	println(url)

	order1 := Order{id: 1, price: 12.3, count: 2, orderStatus: approve}
	order2 := Order{id: 1, price: 12.3, count: 2, orderStatus: reject}

	fmt.Printf("%+v\n", order1)
	fmt.Printf("%+v\n", order2)

}
