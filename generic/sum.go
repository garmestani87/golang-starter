package main

import (
	"fmt"
)

type Decimal interface {
	int | float64
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Printf("%+v \n", sum(numbers))
}

func sum[T Decimal](numbers []T) (sum T) {
	for _, item := range numbers {
		sum += item
	}
	return
}
