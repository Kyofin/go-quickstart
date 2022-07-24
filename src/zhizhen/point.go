package main

import "fmt"

type User struct {
	name string
	id   int
}

func changeName(user User) {
	user.name = "pk"
	fmt.Println("已经将传入user name属性修改为", user.name)
}
func changeName2(user *User) {
	user.name = "pk"
	fmt.Println("已经将传入user name属性修改为", user.name)

}
func main() {
	user := User{
		name: "jzong",
		id:   12,
	}
	changeName(user)
	fmt.Println(user.name)

	user2 := User{
		name: "jzong",
		id:   12,
	}
	changeName2(&user2)
	fmt.Println(user2.name)

}
