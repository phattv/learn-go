package main

import (
	"fmt"
)

func main() {
	// helloWorld()
	sum(3, 4)
}

func helloWorld() {
	fmt.Println("Hello, World!")
}

func sum(a, b int) {
	sum := a + b
	fmt.Println(sum)
}