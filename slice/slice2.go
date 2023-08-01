package main

import "fmt"

func main() {

	var mSlice = []int{1, 2, 3, 4, 5, 6}
	fmt.Printf("len : %d , cap : %d \n", len(mSlice), cap(mSlice))

	mSlice = mSlice[:0]
	fmt.Printf("len : %d , cap : %d \n", len(mSlice), cap(mSlice))

	mSlice = nil
	fmt.Printf("len : %d , cap : %d \n", len(mSlice), cap(mSlice))

}
