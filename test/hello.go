package main

import "fmt"

func add(a int, b int) int {
	var sum int = a + b
	return sum
}

func main() {
	fmt.Print("hello", add(100, 222))
}
