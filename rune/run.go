package main

import "fmt"

func main() {

	str := "سلام دنیا !"
	fmt.Printf("print %v , %d \n", str, len(str))

	str2 := []rune(str)
	fmt.Printf("print %v , %d \n", str2, len(str2))

}
