package main

import (
	"fmt"
	"math"
	"errors"
	"net/http"
	"time"
	"strconv"
)

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
	
	// RegisterRoutes()
	// goRoutine()

	stringChannel := make(chan string)
	for i:= 0; i  < 3; i ++ {
		go makeDough(stringChannel)
		go addSauce(stringChannel)
		go addTopping(stringChannel)
		time.Sleep(time.Millisecond * 5000)
	}
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
	Age int
}

func inc(x int) {
	x++
}

func incWithPointer(x *int) {
	*x++
}

func RegisterRoutes() {
	http.HandleFunc("/", hi)
	http.HandleFunc("/ping", ping)

	port := ":8080"
	fmt.Println("Server is listening at localhost" + port)
	http.ListenAndServe(port, nil)
}

func hi(writer http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(writer, "Hello, World!")
}

func ping(writer http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(writer, "pong")
}

func goRoutine() {
	for i := 0; i < 10; i++ {
		go count(i)
	}

	time.Sleep(time.Millisecond * 11000)
}

func count(id int) {
	for i := 0; i < 10; i++ {
		fmt.Println(fmt.Sprintf("%v: %v", id, i))
		time.Sleep(time.Millisecond * 1000)
	}
}

var pizzaNum = 0
var pizzaName = ""

func makeDough(stringChan chan string) {
	pizzaNum++
	pizzaName = "Pizza #" + strconv.Itoa(pizzaNum)
	fmt.Println("Make dough for", pizzaName, "and send for sauce");
	stringChan <- pizzaName
	time.Sleep(time.Millisecond * 10)
}

func addSauce(stringChan chan string) {
	pizza := <- stringChan
	fmt.Println("Add sauce and send", pizza, "for toppings")
	stringChan <- pizzaName
	time.Sleep(time.Millisecond * 10)
}

func addTopping(stringChan chan string) {
	pizza := <- stringChan
	fmt.Println("Add toppings to", pizza, "and ship")
	time.Sleep(time.Millisecond * 10)
}