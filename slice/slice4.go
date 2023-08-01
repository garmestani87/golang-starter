package main

import "fmt"

func main() {

	var mSlice = []int{1, 2, 3, 4}
	var mSlice1 = []int{10, 20, 30, 40}

	copy(mSlice, mSlice1)
	fmt.Printf("values : %v", mSlice)

	mSlice = append(mSlice, mSlice1...)
	fmt.Printf("values : %v \n", mSlice)

	put(mSlice1)
	delete(mSlice1)

}

func put(mSlice []int) {

	var slc = append(mSlice, 0)
	copy(slc[3:], mSlice[2:])
	slc[2] = 50
	fmt.Printf("values : %v \n", slc)
}

func delete(mSlice []int) {
	var deleted = append(mSlice[:2], mSlice[3:]...)
	fmt.Printf("values : %v \n", deleted)
}
