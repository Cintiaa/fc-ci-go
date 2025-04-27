package main

import "fmt"

func main() {
	fmt.Println(sum(10, 10))
}

func sum(a int, b int) int {
	return a + b
}

func subtract(a int, b int) int {
	return a - b
}

func multiply(a int, b int) int {
	return a * b
}

func divide(a int, b int) int {
	if b == 0 {
		return 0
	}
	return a / b
}

func modulo(a int, b int) int {
	if b == 0 {
		return 0
	}
	return a % b
}
