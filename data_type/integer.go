package main

import (
	"fmt"
	"math/bits"
	"unsafe"
)

func main() {

	var num int
	var num8 int8
	var num16 int16
	var num32 int32
	var num64 int64

	fmt.Printf("num bytes => %d \n", unsafe.Sizeof(num))
	fmt.Printf("num8 bytes => %d \n", unsafe.Sizeof(num8))
	fmt.Printf("num16 bytes => %d \n", unsafe.Sizeof(num16))
	fmt.Printf("num32 bytes => %d \n", unsafe.Sizeof(num32))
	fmt.Printf("num64 bytes => %d \n", unsafe.Sizeof(num64))

	fmt.Printf("SIZE %d \n", bits.UintSize)

	var fnum32 float32
	var fnum64 float64

	fmt.Printf("fnum32 %d \n", unsafe.Sizeof(fnum32))
	fmt.Printf("fnum32 %d \n", unsafe.Sizeof(fnum64))

}
