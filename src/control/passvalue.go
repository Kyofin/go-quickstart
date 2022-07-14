package main

import "fmt"

func main() {
	x := 3
	fmt.Println("x=", x)
	// 这里都copy里原值到函数里
	x1 := add1(x)
	fmt.Println("x+1的值是", x1)
	fmt.Println("x的值是", x)
	p := `
		佛祖保佑
          ===
			保佑！
`
	fmt.Println(p)
	j := 3
	// 获取x的内存地址传入函数
	j1 := add2(&j)
	fmt.Println("j+1的值是", j1)
	fmt.Println("j的值是", j)

}

func add1(a int) int {
	a = a + 1
	return a
}

/**
函数入参变成指针类型，即会改变入参原来的内存数据
*/
func add2(a *int) int {
	*a = *a + 1
	return *a
}
