package main

import "fmt"

type Human struct {
	name  string
	age   int
	phone string
}

type Student struct {
	Human  // 匿名字段
	school string
	loan   float32 //贷款
}

type Empolyee struct {
	Human   // 匿名字段
	company string
	money   float32
}

func (h Human) SayHi() {
	fmt.Println("hi ,我是 %s,我的手机是%s", h.name, h.phone)
}
func (h Human) Sing(musicName string) {
	fmt.Println("正在唱", musicName)
}

// 重载Human的SayHi
func (e Empolyee) SayHi() {
	fmt.Println("我叫%s,我工作在%s公司")
}

type Men interface {
	SayHi()
	Sing(musicName string)
}

func main() {

	// 案例2
	mike := Student{Human{"Mike", 24, "222-222-xxx"}, "河南大学", 0.0}
	paul := Student{Human{"Paul", 23, "211-222-xxx"}, "北京大学", 100.0}
	sam := Empolyee{Human{"Sam", 33, "444-222-xxx"}, "联通", 9000.0}
	tom := Empolyee{Human{"Tom", 36, "434-222-xxx"}, "移动", 8000.0}
	// 定义Men类型变量i
	var i Men
	// i能存储Student
	i = mike
	fmt.Println("这是mike，是一个学生")
	i.SayHi()
	i.Sing("中国之光")
	// i也能存储Employee
	i = tom
	fmt.Println("这是tom，是一个员工")
	i.SayHi()
	i.Sing("中国之光")
	// 定义了slice Men
	x := make([]Men, 3)
	// 三个不同类型的元素，都实现一个接口Men
	x[0], x[1], x[2] = paul, sam, mike
	for _, men := range x {
		men.SayHi()
	}

}
