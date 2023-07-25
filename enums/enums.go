package main

import "encoding/json"

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

	orderJson1, _ := json.Marshal(order1)
	orderJson2, _ := json.Marshal(order2)

	println(string(orderJson1))
	println(string(orderJson2))

}
