package main

import (
	"fmt"
)

func main() {
	// helloWorld()
	// sum(3, 4)
	// array()
	// slice()
	// maps()
	loops()
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
	fmt.Println(a)
}

func slice() {
	s := []string{"yo", "yay"}
	fmt.Println(s)
}

func maps() {
	m := make(map[string]int)
	m["triange"] = 2
	m["square"] = 3
	m["circle"] = 4
	fmt.Println(m)
	delete(m, "square")
	fmt.Println(m)
}

func loops() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// while loop
	j := 0
	for j < 5 {
		fmt.Println(j)
		j++
	}
}