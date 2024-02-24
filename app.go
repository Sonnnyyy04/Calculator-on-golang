package main

import (
	"fmt"
)

func main() {
	var a, b int
	var operator string
	fmt.Print("Input 1st nubmer")
	fmt.Scan(&a)
	fmt.Print("Input 2nd nubmer")
	fmt.Scan(&b)
	fmt.Print("Input operation")
	fmt.Scan(&operator)
	result := make(chan int)
	switch operator {
	case "+":
		go add(a, b, result)
	case "-":
		go subtract(a, b, result)
	case "/":
		go divide(a, b, result)
	case "*":
		go multiply(a, b, result)
	default:
		fmt.Println("you wrote wrong operator")
		return
	}
	fmt.Println(<-result)
}

func add(a, b int, result chan int) {
	result <- a + b
}
func subtract(a, b int, result chan int) {
	result <- a - b
}
func multiply(a, b int, result chan int) {
	result <- a * b
}
func divide(a, b int, result chan int) {
	if b == 0 {
		result <- 0
		fmt.Print("Can't divide on zero")
	} else {
		result <- a / b
	}
}
