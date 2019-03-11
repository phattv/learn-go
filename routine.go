package main

import (
	"fmt"
	"time"
)

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
