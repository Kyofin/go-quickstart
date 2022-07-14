package main

import "fmt"

func main() {
	i := 10
	switch i {
	case 10:
		fmt.Println("i是", 10)
		// 强制执行后续
		fallthrough
	case 1:
		fmt.Println("i是", 1)
		// 强制执行后续
		fallthrough
	case 2 | 3 | 4:
		fmt.Println("i 是 2或3或4")
	default:
		fmt.Println("默认输出")
	}
}
