package main

import "fmt"

const (
	A = iota
	B = iota
	C
)

// iota value resets once initialized again
const (
	D = iota
	E
	F
)

func main() {
	fmt.Println("a - ", A)
	fmt.Println("b - ", B)
	fmt.Println("c - ", C)
	fmt.Println("d - ", D)
	fmt.Println("e - ", E)
	fmt.Println("f - ", F)
}