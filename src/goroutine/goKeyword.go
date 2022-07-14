package main

import (
	"fmt"
	"runtime"
)

func say(w string) {
	for i := 0; i < 5; i++ {
		runtime.Gosched() // 将cpu让给别人
		fmt.Println(w)
	}
}
func main() {
	go say("hello") // 开启一个新的Goroutines
	say("world")
}
