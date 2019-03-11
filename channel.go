package main

import (
	"fmt"
	"strconv"
	"time"
)

var pizzaNum = 0
var pizzaName = ""

func makeDough(stringChan chan string) {
	pizzaNum++
	pizzaName = "Pizza #" + strconv.Itoa(pizzaNum)
	fmt.Println("Make dough for", pizzaName, "and send for sauce")
	stringChan <- pizzaName
	time.Sleep(time.Millisecond * 10)
}

func addSauce(stringChan chan string) {
	pizza := <-stringChan
	fmt.Println("Add sauce and send", pizza, "for toppings")
	stringChan <- pizzaName
	time.Sleep(time.Millisecond * 10)
}

func addTopping(stringChan chan string) {
	pizza := <-stringChan
	fmt.Println("Add toppings to", pizza, "and ship")
	time.Sleep(time.Millisecond * 10)
}
