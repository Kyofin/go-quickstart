package main

import (
	"fmt"

	"sync"
	"time"
)

var count int
var wg1 sync.WaitGroup
var wg2 sync.WaitGroup
var wg3 sync.WaitGroup
var wg4 sync.WaitGroup

func goroutine1() {
	wg1.Wait()
	len := 1000000
	for i := 0; i < len; i++ {
		count++
	}
	wg2.Done()
}

func goroutine2() {
	wg2.Wait()
	len := 1000000
	for i := 0; i < len; i++ {
		count++
	}
	wg3.Done()
}

func goroutine3() {
	wg3.Wait()
	len := 1000000
	for i := 0; i < len; i++ {
		count++
	}
	wg4.Done()
}

func main() {
	now := time.Now().UnixNano()
	wg1.Add(1)
	wg2.Add(1)
	wg3.Add(1)
	wg4.Add(1)
	go goroutine1()

	go goroutine2()
	go goroutine3()
	wg1.Done()
	wg4.Wait()

	fmt.Println(time.Now().UnixNano() - now)
	fmt.Println(count)
}
