package main

import "fmt"

func main() {

	fmt.Println(sumAndProduct(2, 4))
	fmt.Println(sumAndProduct2(2, 4))
	myfun(1, 2, 34)
}

func sumAndProduct(a int, b int) (int, int) {
	return a * b, b + a
}

func sumAndProduct2(a int, b int) (jia int, cheng int) {
	jia = a + b
	cheng = a * b
	return
}

func myfun(arg ...int) {
	for index, i := range arg {
		fmt.Println("索引：", index, ",值：", i)
	}
}
