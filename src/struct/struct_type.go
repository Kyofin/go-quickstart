package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	p1 := person{name: "pp", age: 13}
	p2 := person{"jj", 14}
	fmt.Println(p1)
	fmt.Println(p2)
}
