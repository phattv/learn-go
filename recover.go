package main

import "fmt"

func safeDivision(num1, num2 int) int {
	defer func() {
		fmt.Println(recover())
	}()

	result := num1 / num2
	fmt.Println(result)
	return result
}

func panicDevision(num1, num2 int) int {
	if num2 == 0 {
		panic("Don't divide by ZERO!")
	}

	result := num1 / num2
	fmt.Println(result)
	return result
}
