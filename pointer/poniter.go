package main

import (
	"fmt"
)

func main() {

	var a int = 20
	fmt.Printf("address of a is = %d \n", &a)
	callByValue(a)
	fmt.Printf("a is = %d \n", a)
	callByReference(&a)
	fmt.Printf("a is = %d \n", a)

}

func callByValue(value int) {
	fmt.Printf("address of value is = %d \n", &value)
	value = value + 2
}

func callByReference(value *int) {
	fmt.Printf("address of value is = %d \n", value)
	*value = *value + 2
}
