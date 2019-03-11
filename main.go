package main

import (
	"errors"
	"fmt"
	"math"
)

// Learning go from the basics, uncomment line by line
func main() {
	// helloWorld()
	// sum(3, 4)
	// array()
	// slice()
	// maps()
	// loops()
	// ranges()
	// result, err := sqrt(-64)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(result)
	// }

	// person := Person{
	// 	Name: "Phat",
	// 	Age: 26,
	// }
	// fmt.Println(person.Name)

	// i := 1
	// inc(i)
	// fmt.Println(i)
	// incWithPointer(&i)
	// fmt.Println(i)

	// safeDivision(3, 0)
	// safeDivision(3, 3)

	// panicDevision(3, 3)
	// panicDevision(3, 0)

	// goRoutine()

	// stringChannel := make(chan string)
	// for i := 0; i < 3; i++ {
	// 	go makeDough(stringChannel)
	// 	go addSauce(stringChannel)
	// 	go addTopping(stringChannel)
	// 	time.Sleep(time.Millisecond * 5000)
	// }

	registerRoutes()
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

func ranges() {
	// slice
	arr := []string{"one", "two", "three"}

	for index, value := range arr {
		fmt.Println("index: ", index, "value: ", value)
	}

	// map
	map1 := make(map[string]string)
	map1["one"] = "first"
	map1["two"] = "second"
	map1["three"] = "third"

	for key, value := range map1 {
		fmt.Println("key: ", key, "value: ", value)
	}
}

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Undefined for negative numbers")
	}

	return math.Sqrt(x), nil
}

type Person struct {
	Name string
	Age  int
}

func inc(x int) {
	x++
}

func incWithPointer(x *int) {
	*x++
}
