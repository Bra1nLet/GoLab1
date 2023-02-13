package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, world !")
	fmt.Print("2 * 2 is: ")
	fmt.Println(count(2, 2, "*"))
}

func count(a int, b int, command string) int {
	if command == "-" {
		return a - b
	}
	if command == "+" {
		return a + b
	}
	if command == "/" {
		return a / b
	}
	if command == "*" {
		return a * b
	}
	return 0
}
