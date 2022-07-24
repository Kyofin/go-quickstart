package main

import (
	"fmt"
	"strconv"
)

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

func (h *Human) SayHi() {
	fmt.Println("hi ,我是 %s,我的手机是%s", h.name, h.phone)
}
func (h *Human) Sing(musicName string) {
	fmt.Println("正在唱", musicName)
}

// 重载Human的SayHi
func (e *Empolyee) SayHi() {
	fmt.Println("我叫%s,我工作在%s公司")
}

func (s *Student) BorrowMoney(amount float32) {
	s.loan += amount //贷款不断增加
}

func (e *Empolyee) SpendSalary(amount float32) {
	e.money -= amount // 花费收入
}

// 实现 fmt String 接口
func (h Human) String() string {
	return "?" + h.name + "-" + strconv.Itoa(h.age) + "years-" + h.phone + "?"
}

type Men interface {
	SayHi()
	Sing(musicName string)
}

type YongPerson interface {
	SayHi()
	Sing(musicName string)
	BorrowMoney(amount float32)
}

type DaGongRen interface {
	SayHi()
	Sing(musicName string)
	SpendSalary(amount float32)
}

func main() {

	empolyee := Empolyee{
		Human:   Human{"j总", 32, "18822698990"},
		company: "RzData",
		money:   10000,
	}
	empolyee.SpendSalary(200)
	fmt.Println("消费完荷包还有钱", empolyee.money)

	fmt.Println(empolyee)

}
