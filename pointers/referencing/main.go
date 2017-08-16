package main

import "fmt"

func main() {
	a := 43
	fmt.Println(a)
	fmt.Println(&a)

	var b *int = &a
	fmt.Println(b)
	//Dereferencing
	//  At this memory address show me the value that is stored there
	fmt.Println(*b)
}
