package main

import "fmt"

func main() {

	var mSlice = []int{1, 2, 3, 4, 5}
	modify(mSlice)
	fmt.Printf("values : %v \n", mSlice)
	add(mSlice)
	fmt.Printf("values : %v \n", mSlice)
	add2(&mSlice)
	fmt.Printf("values : %v \n", mSlice)

}

func modify(mSlice []int) {
	mSlice[2] = 10
}

func add(mSlice []int) {
	mSlice = append(mSlice, 6)
}

func add2(mSlice *[]int) {
	*mSlice = append(*mSlice, 6)
}
