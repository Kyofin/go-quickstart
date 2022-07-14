package main

import "fmt"

func main() {

	salary := make(map[string]int)
	salary["jj"] = 199999
	salary["pk"] = 199999
	for k, v := range salary {
		fmt.Println("map key is ", k)
		fmt.Println("map value is ", v)
	}
}
