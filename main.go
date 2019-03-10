package main

import (
	"fmt"
)

func main() {
	// helloWorld()
	// sum(3, 4)
	array()
}

func helloWorld() {
	fmt.Println("Hello, World!")
}

func sum(a, b int) {
	sum := a + b
	fmt.Println(sum)
}

func array() {
	a := [5]int{5, 4, 3, 2, 1}
	fmt.Println(a)
	a[2] = 22
	fmt.Print(a)
}