package main

import "fmt"

func main() {

	personAge := make(map[string]int)
	personAge["pp"] = 11
	personAge["pk"] = 21
	personAge["jjj"] = 21
	fmt.Println(personAge["jj"])

	rating := map[string]float32{"c": 5, "go": 4.5, "java": 4.5}
	subjectRating, ok := rating["c#"]
	if ok {
		fmt.Println("c# 存在", subjectRating)
	} else {
		println("c# 不存在")
	}

}
