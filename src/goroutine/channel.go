package main

import (
	"fmt"
	"time"
)

func sum(a []int, c chan int, isSleep bool) {
	sum := 0

	for _, i := range a {
		sum += i
	}
	if isSleep {
		time.Sleep((time.Duration(10) * time.Second))
	}
	c <- sum // 将结果发生给channel c
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8}
	// 创建一个channel，接收类型int
	c := make(chan int)

	ints1 := a[:len(a)/2]
	fmt.Println(ints1)
	// 创建一个goroutine
	go sum(ints1, c, true)

	ints2 := a[len(a)/2:]
	fmt.Println(ints2)
	// 创建一个goroutine
	go sum(ints2, c, false)

	// 阻塞等接收channel结果，ints2会先返回
	x := <-c
	fmt.Println(x)
	y := <-c
	fmt.Println(y)
	// 这里会报错，因为已经不会有数据发送给channel了
	z := <-c
	fmt.Println(z)

}
