package main

import (
	"fmt"
	"sync"
	"time"
)

var count int
var wg1 sync.WaitGroup
var mutex sync.Mutex

func goroutine1() {
	mutex.Lock()
	len := 1000000
	for i := 0; i < len; i++ {
		count++
	}
	mutex.Unlock()
	wg1.Done()
}

func main() {
	now := time.Now().UnixNano()
	wg1.Add(3)
	go goroutine1()
	go goroutine1()
	go goroutine1()
	wg1.Wait()

	fmt.Println(time.Now().UnixNano() - now)
	fmt.Println(count)
}
