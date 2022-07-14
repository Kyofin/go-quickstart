package main

import "fmt"

type testFunType func(int) bool // 声明一个函数类型

func isOuShu(num int) bool {
	return num%2 == 0
}
func isJiShu(num int) bool {
	return num%2 != 0
}

func filter(slice []int, f testFunType) []int {
	var result []int
	for i := range slice {
		if f(i) {
			result = append(result, i)
		}
	}
	return result
}
func main() {
	num := []int{1, 2, 3, 4, 5, 6, 7}
	zhengshuSlice := filter(num, isOuShu)
	jishuSlice := filter(num, isJiShu)
	fmt.Println("偶数：", zhengshuSlice)
	fmt.Println("奇数：", jishuSlice)
}
