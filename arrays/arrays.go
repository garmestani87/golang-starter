package main

import "fmt"

func main() {

	var arr = [4]int{1, 2, 3, 4}

	changeArrayValue(arr)
	fmt.Printf("values after change :  %v \n", arr)

	arr2 := arr
	changeArrayValue(arr2)
	fmt.Printf("values after change :  %v \n", arr2)

	arr3 := &arr
	changeArrayValue2(arr3)
	fmt.Printf("values after change :  %v \n", arr3)

	changeArrayValue2(&arr)
	fmt.Printf("values after change :  %v \n", arr)

}
func changeArrayValue(arr [4]int) {
	arr[1] = 10
	arr[2] = 20
}

func changeArrayValue2(arr *[4]int) {
	arr[1] = 10
	arr[2] = 20
}
